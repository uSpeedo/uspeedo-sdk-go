package main

import (
	"fmt"
	"time"

	"github.com/uspeedo/uspeedo-sdk-go/services/asms"
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/log"
)

func TestTemplateApi() {
	asmsClient := getAsmsClient()

	// Create Template

	createTemplateReq := asmsClient.NewCreateUSMSTemplateReq()
	createTemplateReq.Template = uspeedo.String("Your verification code is {1}")
	createTemplateReq.Purpose = uspeedo.Int(1)
	createTemplateReq.Remark = uspeedo.String("this is a test template")
	createTemplateReq.TemplateName = uspeedo.String("test template")

	createTemplateResp, err := asmsClient.CreateUSMSTemplate(createTemplateReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", createTemplateResp)
	}

	// Query Template

	queryTemplateReq := asmsClient.NewQueryUSMSTemplateReq()
	queryTemplateReq.TemplateId = []string{"UTAXXXXXXXXXXX"}
	queryTemplateResp, err := asmsClient.QueryUSMSTemplate(queryTemplateReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", queryTemplateResp)
	}

	// Update Template

	updateTemplateReq := asmsClient.NewUpdateUSMSTemplateReq()
	updateTemplateReq.TemplateId = uspeedo.String("UTAXXXXXXXXXXX")
	updateTemplateReq.Template = uspeedo.String("Your verification code is {1},thanks")
	updateTemplateResp, err := asmsClient.UpdateUSMSTemplate(updateTemplateReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", updateTemplateResp)
	}

	// Delete Template

	deleteTemplateReq := asmsClient.NewDeleteUSMSTemplateReq()
	deleteTemplateReq.TemplateIds = []string{"UTAXXXXXXXXXXX"}
	deleteTemplateResp, err := asmsClient.DeleteUSMSTemplate(deleteTemplateReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", deleteTemplateResp)
	}

}

func getAsmsClient() *asms.AsmsClient {
	cfg := uspeedo.NewConfig()
	cfg.LogLevel = log.DebugLevel
	cfg.Timeout = 5 * time.Second
	cfg.MaxRetries = 1

	credential := auth.NewCredential()

	credential.PublicKey = "Get the access key id from console.uspeedo.com"
	credential.PrivateKey = "Get the access key secret from console.uspeedo.com"

	asmsClient := asms.NewClient(&cfg, &credential)

	return asmsClient
}

func TestSendMessage() {
	asmsClient := getAsmsClient()
	req := asmsClient.NewSendBatchUSMSMessageReq()
	req.TaskContent = []models.SendInfo{
		{
			TemplateId: uspeedo.String("UTAXXXXXXXXXXX"),
			SenderId:   uspeedo.String("uspeedo"),
			Target: []models.TargetPhone{
				{
					Phone:          uspeedo.String("(11)11111111"),
					TemplateParams: []string{"111111"},
				},
			},
		},
	}

	// send request
	resp, err := asmsClient.SendBatchUSMSMessage(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	} else {
		fmt.Printf("response of the request: %v\n", resp)
	}
}

func main() {
	TestTemplateApi()
	TestSendMessage()
}
