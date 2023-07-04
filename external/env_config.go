package external

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	USpeedoPublicKeyEnvVar = "USPEEDO_PUBLIC_KEY"

	USpeedoPrivateKeyEnvVar = "USPEEDO_PRIVATE_KEY"

	USpeedoAPIBaseURLEnvVar = "USPEEDO_API_BASE_URL"

	USpeedoTimeoutSecondEnvVar = "USPEEDO_TIMEOUT_SECOND"

	USpeedoSharedProfileEnvVar = "USPEEDO_PROFILE"

	USpeedoSharedConfigFileEnvVar = "USPEEDO_SHARED_CONFIG_FILE"

	USpeedoSharedCredentialFileEnvVar = "USPEEDO_SHARED_CREDENTIAL_FILE"
)

func loadEnvConfig() (*config, error) {
	cfg := &config{
		PublicKey:            os.Getenv(USpeedoPublicKeyEnvVar),
		PrivateKey:           os.Getenv(USpeedoPrivateKeyEnvVar),
		BaseUrl:              os.Getenv(USpeedoAPIBaseURLEnvVar),
		SharedConfigFile:     os.Getenv(USpeedoSharedConfigFileEnvVar),
		SharedCredentialFile: os.Getenv(USpeedoSharedCredentialFileEnvVar),
		Profile:              os.Getenv(USpeedoSharedProfileEnvVar),
	}

	durstr, ok := os.LookupEnv(USpeedoTimeoutSecondEnvVar)
	if ok {
		durnum, err := strconv.Atoi(durstr)
		if err != nil {
			return nil, fmt.Errorf("parse environment variable USPEEDO_TIMEOUT_SECOND [%s] error : %v", durstr, err)
		}
		cfg.Timeout = time.Second * time.Duration(durnum)
	}
	return cfg, nil
}
