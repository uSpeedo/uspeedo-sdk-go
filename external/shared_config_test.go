package external

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

func TestLoadSharedConfig(t *testing.T) {
	cfg, err := LoadUSpeedoConfigFile(
		TestValueEnvUSpeedoSharedConfigFile,
		TestValueEnvUSpeedoProfile,
	)
	assert.NoError(t, err)
	checkTestClientConfig(t, cfg)

	cred, err := LoadUSpeedoCredentialFile(
		TestValueEnvUSpeedoSharedCredentialFile,
		TestValueEnvUSpeedoProfile,
	)
	assert.NoError(t, err)
	checkTestCredential(t, cred)
}

func checkTestCredential(t *testing.T, cred *auth.Credential) {
	assert.Equal(t, TestValueFileUSpeedoPublicKey, cred.PublicKey)
	assert.Equal(t, TestValueFileUSpeedoPrivateKey, cred.PrivateKey)
}

func checkTestClientConfig(t *testing.T, cfg *uspeedo.Config) {
	assert.Equal(t, TestValueFileUSpeedoTimeout, cfg.Timeout)
	assert.Equal(t, TestValueFileUSpeedoBaseUrl, cfg.BaseUrl)
}
