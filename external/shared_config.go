package external

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/log"
)

// DefaultProfile is the default named profile for uspeedo sdk
const DefaultProfile = "default"

// DefaultSharedConfigFile will return the default shared config filename
func DefaultSharedConfigFile() string {
	return filepath.Join(userHomeDir(), ".uspeedo", "config.json")
}

// DefaultSharedCredentialsFile will return the default shared credential filename
func DefaultSharedCredentialsFile() string {
	return filepath.Join(userHomeDir(), ".uspeedo", "credential.json")
}

// LoadUSpeedoConfigFile will load uspeedo client config from config file
func LoadUSpeedoConfigFile(cfgFile, profile string) (*uspeedo.Config, error) {
	if len(profile) == 0 {
		return nil, fmt.Errorf("expected uspeedo named profile is not empty")
	}

	cfgMaps, err := loadConfigFile(cfgFile)
	if err != nil {
		return nil, err
	}

	c := getSharedConfig(cfgMaps, profile)
	return c.Config(), nil
}

// LoadUSpeedoCredentialFile will load uspeedo credential config from config file
func LoadUSpeedoCredentialFile(credFile, profile string) (*auth.Credential, error) {
	if len(profile) == 0 {
		return nil, fmt.Errorf("expected uspeedo named profile is not empty")
	}

	credMaps, err := loadCredFile(credFile)
	if err != nil {
		return nil, err
	}

	c := getSharedCredential(credMaps, profile)
	return c.Credential(), nil
}

type sharedConfig struct {
	BaseURL string `json:"base_url"`
	Timeout int    `json:"timeout_sec"`
	Profile string `json:"profile"`
	Active  bool   `json:"active"`
}

type sharedCredential struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Profile    string `json:"profile"`
}

func loadConfigFile(cfgFile string) ([]sharedConfig, error) {
	realCfgFile := cfgFile
	cfgs := make([]sharedConfig, 0)

	// try to load default config
	if len(realCfgFile) == 0 {
		realCfgFile = DefaultSharedConfigFile()
	}

	// load config file
	err := loadJSONFile(realCfgFile, &cfgs)

	if err != nil {
		// skip error for loading default config
		if len(cfgFile) == 0 && os.IsNotExist(err) {
			log.Debugf("config file is empty")
		} else {
			return nil, err
		}
	}

	return cfgs, nil
}

func getCredFilePath(credFile string) string {
	realCredFile := credFile
	homePath := fmt.Sprintf("~%s", string(os.PathSeparator))
	if strings.HasPrefix(credFile, homePath) {
		realCredFile = strings.Replace(credFile, "~", userHomeDir(), 1)
	}
	// try to load default credential
	if len(credFile) == 0 {
		realCredFile = DefaultSharedCredentialsFile()
	}
	return realCredFile
}

func loadCredFile(credFile string) ([]sharedCredential, error) {
	realCredFile := getCredFilePath(credFile)
	creds := make([]sharedCredential, 0)

	// load credential file
	err := loadJSONFile(realCredFile, &creds)

	if err != nil {
		// skip error for loading default credential
		if len(credFile) == 0 && os.IsNotExist(err) {
			log.Debugf("credential file is empty")
		} else {
			return nil, err
		}
	}

	return creds, nil
}

func loadSharedConfigFile(cfgFile, credFile, profile string) (*config, error) {
	cfgs, err := loadConfigFile(cfgFile)
	if err != nil {
		return nil, err
	}

	creds, err := loadCredFile(credFile)
	if err != nil {
		return nil, err
	}

	c := &config{
		Profile:              profile,
		SharedConfigFile:     cfgFile,
		SharedCredentialFile: credFile,
	}
	_ = c.merge(getSharedConfig(cfgs, profile))
	_ = c.merge(getSharedCredential(creds, c.Profile))

	return c, nil
}

func getSharedConfig(cfgs []sharedConfig, profile string) *config {
	cfg := &sharedConfig{}

	if profile != "" {
		for i := 0; i < len(cfgs); i++ {
			if cfgs[i].Profile == profile {
				cfg = &cfgs[i]
			}
		}
	} else {
		for i := 0; i < len(cfgs); i++ {
			if cfgs[i].Active {
				cfg = &cfgs[i]
			}
		}
	}

	return &config{
		Profile: cfg.Profile,
		BaseUrl: cfg.BaseURL,
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}
}

func getSharedCredential(creds []sharedCredential, profile string) *config {
	cred := &sharedCredential{}

	for i := 0; i < len(creds); i++ {
		if creds[i].Profile == profile {
			cred = &creds[i]
		}
	}

	return &config{
		PublicKey:  cred.PublicKey,
		PrivateKey: cred.PrivateKey,
	}
}
