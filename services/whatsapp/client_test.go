package whatsapp

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"os"
	"testing"
	"time"
)

func getClient() *WhatsAppClient {
	cfg := uspeedo.NewConfig()
	credential := auth.NewCredential()
	credential.PublicKey = "PublicKey"
	credential.PrivateKey = "PrivateKey"
	return NewClient(&cfg, &credential)
}

func TestGet(t *testing.T) {
	ass := assert.New(t)
	c := getClient()

	phones, err := c.GetAccountPhoneList(&apis.GetAccountPhoneListRequest{})
	ass.Nil(err)
	ass.NotNil(phones)
	ass.Equal(phones.RetCode, 0)
	ass.NotNil(phones.Data)
	ass.NotEmpty(phones.Data.PhoneList)

	time.Sleep(100 * time.Millisecond)
	phone := phones.Data.PhoneList[0]
	templates, err := c.GetTemplates(&apis.Console__getTemplateRequest{
		BusinessPhone: phone.Number,
	})
	ass.Nil(err)
	ass.NotNil(templates)
	ass.Equal(templates.RetCode, 0)
	ass.NotNil(templates.Data)
	ass.NotEmpty(templates.Data.Data)

	sum, err := c.GetMessageSummary(&apis.Client__GetMessageSummaryRequest{})
	ass.Nil(err)
	ass.NotNil(sum)
	ass.Equal(sum.RetCode, 0)
	ass.NotNil(sum.Data)
}

func TestMedia(t *testing.T) {
	ass := assert.New(t)
	c := getClient()

	phones, err := c.GetAccountPhoneList(&apis.GetAccountPhoneListRequest{})
	ass.Nil(err)
	ass.NotNil(phones)
	ass.Equal(phones.RetCode, 0)
	ass.NotNil(phones.Data)
	ass.NotEmpty(phones.Data.PhoneList)
	phone := phones.Data.PhoneList[0]

	time.Sleep(100 * time.Millisecond)
	b, err := os.ReadFile("./example.jpg")
	ass.Nil(err)
	fileString := base64.StdEncoding.EncodeToString(b)
	umRes, err := c.UploadMedia(&apis.Console__uploadMediaRequest{
		BusinessPhone: phone.Number,
		File:          &fileString,
	})
	ass.Nil(err)
	ass.NotNil(umRes)
	ass.Equal(umRes.RetCode, 0)
	ass.NotNil(umRes.Data)
	ass.NotEmpty(umRes.Data.Id)
	mediaId := umRes.Data.Id

	time.Sleep(100 * time.Millisecond)
	md, err := c.GetMedia(&apis.Console__getMediaRequest{
		BusinessPhone: phone.Number,
		MediaId:       mediaId,
	})
	ass.Nil(err)
	ass.NotNil(md)
	ass.Equal(md.RetCode, 0)
	ass.NotNil(md.Data)
	ass.NotNil(md.Data.URL)

	time.Sleep(100 * time.Millisecond)
	dr, err := c.DeleteMedia(&apis.Console__deleteMediaRequest{
		BusinessPhone: phone.Number,
		MediaId:       mediaId,
	})
	ass.Nil(err)
	ass.NotNil(dr)
	ass.Equal(dr.RetCode, 0)
}

func TestMessage(t *testing.T) {
	ass := assert.New(t)
	c := getClient()

	phones, err := c.GetAccountPhoneList(&apis.GetAccountPhoneListRequest{})
	ass.Nil(err)
	ass.NotNil(phones)
	ass.Equal(phones.RetCode, 0)
	ass.NotNil(phones.Data)
	ass.NotEmpty(phones.Data.PhoneList)

	time.Sleep(100 * time.Millisecond)
	phone := phones.Data.PhoneList[0]
	templates, err := c.GetTemplates(&apis.Console__getTemplateRequest{
		BusinessPhone: phone.Number,
	})
	ass.Nil(err)
	ass.NotNil(templates)
	ass.Equal(templates.RetCode, 0)
	ass.NotNil(templates.Data)
	ass.NotEmpty(templates.Data.Data)

	to := "WhatsappNumber"
	msgType := "text"
	content := "none"
	msg, err := c.SendWhatsappMessage(&apis.Console__sendMessageRequest{
		BusinessPhone: phone.Number,
		To:            &to,
		Type:          &msgType,
		Content:       &content,
	})
	// no money in this account
	ass.Error(err)
	ass.NotNil(msg)
	ass.Equal(msg.RetCode, 17125)
}
