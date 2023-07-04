package external

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	TestValueEnvUSpeedoPublicKey            = "f05816ca02feec1b3fc38b80a1c450cc"
	TestValueEnvUSpeedoPrivateKey           = "c45f9bec5fa4c6c47fd871fadd97dd2e"
	TestValueEnvUSpeedoProfile              = "default"
	TestValueEnvUSpeedoTimeout              = time.Duration(30) * time.Second
	TestValueEnvUSpeedoBaseUrl              = "https://api.uspeedo.com/api"
	TestValueEnvUSpeedoSharedConfigFile     = filepath.Join("test-fixtures", "config.json")
	TestValueEnvUSpeedoSharedCredentialFile = filepath.Join("test-fixtures", "credential.json")

	TestValueFileUSpeedoPublicKey  = "18ce27e79596c5adc986f66707b3fd55"
	TestValueFileUSpeedoPrivateKey = "2c125558f9004b73751e3a4658e2f52b"
	TestValueFileUSpeedoTimeout    = time.Duration(30) * time.Second
	TestValueFileUSpeedoBaseUrl    = "https://api.uspeedo.com/api"
)

func TestLoadConfig(t *testing.T) {
	setTestEnv()

	c, err := LoadDefaultUSpeedoConfig()
	assert.NoError(t, err)

	checkTestConfig(t, c)
}

func TestLoadEnvConfig(t *testing.T) {
	setTestEnv()

	c := &config{}
	err := c.loadEnv()
	assert.NoError(t, err)

	checkTestEnvConfig(t, c)
}

func TestLoadSharedFile(t *testing.T) {
	c := &config{
		SharedConfigFile:     TestValueEnvUSpeedoSharedConfigFile,
		SharedCredentialFile: TestValueEnvUSpeedoSharedCredentialFile,
	}
	err := c.loadFileIfExist()
	assert.NoError(t, err)

	checkTestConfig(t, c)
}

func checkTestEnvConfig(t *testing.T, c *config) {
	assert.Equal(t, TestValueEnvUSpeedoPublicKey, c.PublicKey)
	assert.Equal(t, TestValueEnvUSpeedoPrivateKey, c.PrivateKey)
	assert.Equal(t, TestValueEnvUSpeedoBaseUrl, c.BaseUrl)
	assert.Equal(t, TestValueEnvUSpeedoTimeout, c.Timeout)
	assert.Equal(t, TestValueEnvUSpeedoProfile, c.Profile)
	assert.Equal(t, TestValueEnvUSpeedoSharedConfigFile, c.SharedConfigFile)
	assert.Equal(t, TestValueEnvUSpeedoSharedCredentialFile, c.SharedCredentialFile)
}

func checkTestConfig(t *testing.T, c ConfigProvider) {
	cfg, cred := c.Config(), c.Credential()

	checkTestCredential(t, cred)
	checkTestClientConfig(t, cfg)
}

func setTestEnv() {
	_ = os.Setenv(USpeedoPublicKeyEnvVar, TestValueEnvUSpeedoPublicKey)
	_ = os.Setenv(USpeedoPrivateKeyEnvVar, TestValueEnvUSpeedoPrivateKey)
	_ = os.Setenv(USpeedoAPIBaseURLEnvVar, TestValueEnvUSpeedoBaseUrl)
	_ = os.Setenv(USpeedoSharedProfileEnvVar, TestValueEnvUSpeedoProfile)
	_ = os.Setenv(USpeedoSharedConfigFileEnvVar, TestValueEnvUSpeedoSharedConfigFile)
	_ = os.Setenv(USpeedoSharedCredentialFileEnvVar, TestValueEnvUSpeedoSharedCredentialFile)
	durationStr := strings.TrimSuffix(TestValueEnvUSpeedoTimeout.String(), "s")
	_ = os.Setenv(USpeedoTimeoutSecondEnvVar, durationStr)
}
