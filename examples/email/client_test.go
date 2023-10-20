package email

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/email"
	"github.com/uspeedo/uspeedo-sdk-go/services/email/apis"
	"github.com/uspeedo/uspeedo-sdk-go/services/email/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"testing"
)

func TestEmailClient_SendEmailTemplate(t *testing.T) {
	cfg := uspeedo.NewConfig()

	credential := auth.NewCredential()
	credential.PublicKey = "PublicKey"
	credential.PrivateKey = "PrivateKey"

	smsClient := email.NewClient(&cfg, &credential)

	// send request
	resp, err := smsClient.SendEmailTemplate(&apis.SendEmailTemplateReq{
		EmailContent: []models.TargetEmail{
			{
				EmailAddress:           uspeedo.String("example@example.com"),
				TemplateVariableParams: []string{"variableName{##}variableValue", "variableName{##}variableValue"},
			},
		},
		SendEmail:  uspeedo.String("example@example.com"),
		TemplateId: uspeedo.String("UETXXXXXXXXXXX"),
		AccountId:  uspeedo.Int(0),
	})
	if err != nil {
		t.Logf("something bad happened: %s\n", err)
		return
	} else {
		t.Logf("response of the request: %+v\n", resp)
	}

	t.Logf("%+v \n", resp)

	t.Log("Message = ", resp.GetMessage())
	t.Log("RetCode = ", resp.GetRetCode())
	for _, e := range resp.FailContent {
		t.Logf("FailContent FailureReason = %+v \n", *e.FailureReason)
	}

	t.Logf("SessionNo = %+v \n", *resp.SessionNo)

	t.Logf("SuccessCount = %+v \n", *resp.SuccessCount)
}
