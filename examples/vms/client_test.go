package vms

import (
	"testing"

	"github.com/uspeedo/uspeedo-sdk-go/services/vms"
	"github.com/uspeedo/uspeedo-sdk-go/services/vms/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

func TestVMSClient_SendVMSMessage(t *testing.T) {
	cfg := uspeedo.NewConfig()

	credential := auth.NewCredential()
	credential.PublicKey = "PublicKey"
	credential.PrivateKey = "PrivateKey"

	smsClient := vms.NewClient(&cfg, &credential)

	// send request
	resp, err := smsClient.SendVMSMessage(&apis.SendMessageRequest{
		AccountId:    uspeedo.Int(0),
		CalledNumber: uspeedo.String("+1111111111"),
		TemplateId:   uspeedo.String("UVVTXXXXXXXXXXX"),
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

	t.Logf("SessionNo = %+v \n", *resp.SessionNo)
}
