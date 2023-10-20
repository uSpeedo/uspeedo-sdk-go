package auth

import (
	"strings"
	"testing"
)

func TestCredential_CreateSign(t *testing.T) {
	type fields struct {
		PublicKey  string
		PrivateKey string
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"standard",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialCreateSignQuery00},
			"d8b2114fc7bf4bd24fd0424326bb9d1ca7bfdda7",
		},
		{
			"unOrder",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialCreateSignQuery01},
			"d8b2114fc7bf4bd24fd0424326bb9d1ca7bfdda7",
		},
		{
			"noPublicKey",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialCreateSignQuery02},
			"11a0430ae6f1a4dc0c656e64dfc1886b0ac4b63a",
		},
		{
			"noParams",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialCreateSignQuery03},
			"ceabd6349a8f229a82271119a8b3af30ba8a2428",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Credential{
				PublicKey:  tt.fields.PublicKey,
				PrivateKey: tt.fields.PrivateKey,
			}
			if got := c.CreateSign(tt.args.query); got != tt.want {
				t.Errorf("Credential.CreateSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredential_BuildCredentialedQuery(t *testing.T) {
	type fields struct {
		PublicKey  string
		PrivateKey string
	}
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			"standard",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialBuildCredentialedQuery01},
			"11a0430ae6f1a4dc0c656e64dfc1886b0ac4b63a",
		},
		{
			"longArray",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialBuildCredentialedQuery02},
			"258dc3d2d7115e9e7c9f4dd4910e671e17b224f1",
		},
		{
			"noParams",
			fields{"0dca1d51a9b3113c6f562acb0f813bce", "NTk1Mzk0MAItOWI0My10MGM4LTg0NmMMNDM0ZGM5Y2ZkMmNk"},
			args{testCredentialBuildCredentialedQuery03},
			"ceabd6349a8f229a82271119a8b3af30ba8a2428",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCredential()
			c.PublicKey = tt.fields.PublicKey
			c.PrivateKey = tt.fields.PrivateKey

			if got := c.BuildCredentialedQuery(tt.args.params); !strings.Contains(got, tt.want) {
				t.Errorf("Credential.BuildCredentialedQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testCredentialCreateSignQuery00 = strings.Join(strings.Split(`AccountId=60000011
&Action=CreateUSMSTemplate
&PublicKey=0dca1d51a9b3113c6f562acb0f813bce
&Purpose=1
&Template=your%20code%20is%20%7B1%7D
&TemplateName=demo`, "\n"), "")

var testCredentialCreateSignQuery01 = strings.Join(strings.Split(`Action=CreateUSMSTemplate
&PublicKey=0dca1d51a9b3113c6f562acb0f813bce
&AccountId=60000011
&Template=your%20code%20is%20%7B1%7D
&Purpose=1
&TemplateName=demo`, "\n"), "")

var testCredentialCreateSignQuery02 = strings.Join(strings.Split(`Action=CreateUSMSTemplate
&AccountId=60000011
&Purpose=1
&Template=your%20code%20is%20%7B1%7D
&TemplateName=demo`, "\n"), "")

var testCredentialCreateSignQuery03 = strings.Join(strings.Split(`Action=GetNSPermission`, "\n"), "")

var testCredentialBuildCredentialedQuery01 = map[string]interface{}{
	"Action":       "CreateUSMSTemplate",
	"AccountId":    "60000011",
	"Purpose":      "1",
	"Template":     "your code is {1}",
	"TemplateName": "demo",
}

var testCredentialBuildCredentialedQuery02 = map[string]interface{}{
	"Action":     "QueryUSMSTemplate",
	"AccountId":  "60000011",
	"TemplateId": []string{"UTA230227EL4IW1", "UTA230227JVIB02"},
}

var testCredentialBuildCredentialedQuery03 = map[string]interface{}{
	"Action": "GetNSPermission",
}
