package mms

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/uspeedo/uspeedo-sdk-go/services/mms"
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/log"
)

const (
	AccountId  = 1
	PublicKey  = "PublicKey"
	PrivateKey = "PrivateKey"
)

const OK = 0

func getMMSClient() *mms.MMSClient {
	cfg := uspeedo.NewConfig()
	cfg.LogLevel = log.InfoLevel
	cfg.Timeout = 5 * time.Second

	credential := auth.NewCredential()
	credential.PublicKey = PublicKey
	credential.PrivateKey = PrivateKey

	return mms.NewClient(&cfg, &credential)
}

func ExampleMMSTemplate() {
	client := getMMSClient()

	// Create MMS Template
	createTemplateReq := client.NewCreateMMSTemplateReq()
	createTemplateReq.Action = uspeedo.String("CreateMMSTemplate")
	createTemplateReq.AccountId = uspeedo.Int(AccountId)
	createTemplateReq.TemplateName = uspeedo.String("MyTemplate")
	createTemplateReq.Text = uspeedo.String("This is Text")
	createTemplateReq.MediaType = uspeedo.String("image/png")
	createTemplateReq.Media = uspeedo.String("iVBORw0KGgoAAAANSUhEUgAAAOYAAACKCAMAAACJgsaoAAAAGFBMVEVHcEwAm9YA56cATP8A56cAS/8A56cAS/9WLrTHAAAABnRSTlMAQICAv78Cuu4JAAAEpUlEQVR42tWdWZLcOAxEkVjI+994HP0xCi9tkEwVQL8DuPSYCYdaCyUcAOwLAPIhAEB/AACQWmAW81ciDPIiUPXxC+4KFCnG/J4IvGM4vsf906ZmMycAznHkOPC5HOcq55nqWOUzmeaOvKn62EJRLsmL6tjHUSCZYYRkuajNc4yQLBW1yWG0ZI4KCWySgLCsCtQmSxCSVYHGpAFhWROoTR4sS/Jom6Xlls2eMXkit+z1xEzhK+vjbbS6sHll4eN9vNYyryzGR3AUW04cjCWPF1saYcngpZZBNJZDCy0nEsteT8yCMH10e86XAGHJg8QyCsLU8XlqLCcyy87a2vx8mBij2ROzIMxRBFNZPkwfReiLlY0vkjDpyroTnlxlzSAPQMRSmGMPVTw/sWtLV9ZMfgcRaZhO3jqAs3Ea4fiY/j1M5TsncCZOsJKPKLjKPpKsqB5X1iQDloRJSO6KHoZpwoB9S3LF9ChMEwrfleQDPQjThAMvX8vR/SUL2jLH96PkPTfD5C1xYMl76l6YJixOSBKe8oACSxxIPk9axQ8MJ57YCNOExk8kLbKHNnRj8QosJTkYTU4/vzVdv2oLypLX1L1zbOzFudxZeQOoe664tPSx5amLYULeA1D1L1QVOL0ohY3a6lqYIfXYTLCNONc0IeXEfOA9daWzJtVg99ohVjTtssraXGM5Tl9oCPot85rpwnDeFWYc3SbONXFTmJg7rMaJdPkaC5uDxTg104ybLUMeNNO8p7MxtzBZ1fRsGkrHkrAUJP8H4ZLOGmGZDmfyz+NaS2z91Z4MxPVj+aDnmnHtWG5rMqPZP5YPIDQvHct9TSTLdv9Y8pq4sbANmv2WD+eaF45lhWb/WHZp9heWn83rLe/XtMqXsMQIzdKxzHBC86rC9mv2W4qTpwf9lmBvnzpxTls5ljmgNfsLy3ZWkwt7V1jymiA0C8cyB4PRxA1jyWtmv2wXFJbvrKe//Y9YYnCa1mYZkIydG2LINPvHMoe9izvvLeyr9+TtYkv+CYuH7rHkH3rEyjFY71jyYY61xy97C8uH6WuLbZdbahImFg/EGseSr6wvH0nfWPKWA8Sz0VWFzfF1TSN+vtlSR4LvHI5dMZbkewsyU9Aylrylb647iCj7LAd2N+jAdZY+Unz/qOyuscRYAAdPrETBWAJ8YR/8aPkDHy6srr6+qXq2zQxeGyELwvIL5aJ88CSC80CN2HnZF9/HVWbPoLlIWCJ5GCV+ykGR1JXa3IGJFHld88I+OJA4Jji/uVUE8JPikyNt+eAO4P83y9Rf2ebKTjaXMbM8xLwUWrjZnk0KYiy90FLQZKnFO7NZk+X9G5Z1jCW/C2Z5lOINloJiS/RsmShWW9gCyyZPkwLLnPjXLVWk2zPkwQssSU8+SowCyx5Pq8mS9+QlC05+Oj3DCnaphUivpxXsxuvkDRAe6/t6RuF5Hwo+uqByBuJjUT54yVj2f0VDC6LMwJuSnyuuQ0iMlSz5hhgPgrXMgbOSnaImy8AZyU5Rky3ghGSXqJlsAyckO0xNDlFfd2QkedEwEwK40448QESiSAP81dQVUgI+/8lqwN3/YHiq+B+zd7pmidG7ZwAAAABJRU5ErkJggg==")
	createTemplateReq.Subject = uspeedo.String("MySubject")

	createTemplateRes, err := client.CreateMMSTemplate(createTemplateReq)
	if err != nil {
		panic(err)
	}
	if createTemplateRes.GetRetCode() != OK {
		panic(createTemplateRes.GetMessage())
	}

	templateId := *createTemplateRes.TemplateId
	fmt.Println("Successfully created MMS template, template id:", templateId)

	// Query MMS Template
	queryTemplateReq := client.NewQueryMMSTemplateReq()
	queryTemplateReq.Action = uspeedo.String("QueryMMSTemplate")
	queryTemplateReq.AccountId = uspeedo.Int(AccountId)
	queryTemplateReq.TemplateIds = []string{templateId}

	queryTemplateRes, err := client.QueryMMSTemplate(queryTemplateReq)
	if err != nil {
		panic(err)
	}
	if queryTemplateRes.GetRetCode() != OK {
		panic(queryTemplateRes.GetMessage())
	}

	for _, template := range queryTemplateRes.Data {
		mediaFileAddr, _ := base64.StdEncoding.DecodeString(*template.MediaFile)
		fmt.Printf("Successfully queried MMS template, template id: %v, text: %v, media file addr: %v\n", template.TemplateId, template.Text, mediaFileAddr)
	}

	// Update MMS Template
	updateTemplateReq := client.NewUpdateMMSTemplateReq()
	updateTemplateReq.Action = uspeedo.String("UpdateMMSTemplate")
	updateTemplateReq.AccountId = uspeedo.Int(AccountId)
	updateTemplateReq.TemplateId = uspeedo.String(templateId)
	updateTemplateReq.TemplateName = uspeedo.String("MyTemplate")
	updateTemplateReq.Text = uspeedo.String("This is Text")
	updateTemplateReq.MediaType = uspeedo.String("image/png")
	updateTemplateReq.Media = uspeedo.String("iVBORw0KGgoAAAANSUhEUgAAAOYAAACKCAMAAACJgsaoAAAAGFBMVEVHcEwAm9YA56cATP8A56cAS/8A56cAS/9WLrTHAAAABnRSTlMAQICAv78Cuu4JAAAEpUlEQVR42tWdWZLcOAxEkVjI+994HP0xCi9tkEwVQL8DuPSYCYdaCyUcAOwLAPIhAEB/AACQWmAW81ciDPIiUPXxC+4KFCnG/J4IvGM4vsf906ZmMycAznHkOPC5HOcq55nqWOUzmeaOvKn62EJRLsmL6tjHUSCZYYRkuajNc4yQLBW1yWG0ZI4KCWySgLCsCtQmSxCSVYHGpAFhWROoTR4sS/Jom6Xlls2eMXkit+z1xEzhK+vjbbS6sHll4eN9vNYyryzGR3AUW04cjCWPF1saYcngpZZBNJZDCy0nEsteT8yCMH10e86XAGHJg8QyCsLU8XlqLCcyy87a2vx8mBij2ROzIMxRBFNZPkwfReiLlY0vkjDpyroTnlxlzSAPQMRSmGMPVTw/sWtLV9ZMfgcRaZhO3jqAs3Ea4fiY/j1M5TsncCZOsJKPKLjKPpKsqB5X1iQDloRJSO6KHoZpwoB9S3LF9ChMEwrfleQDPQjThAMvX8vR/SUL2jLH96PkPTfD5C1xYMl76l6YJixOSBKe8oACSxxIPk9axQ8MJ57YCNOExk8kLbKHNnRj8QosJTkYTU4/vzVdv2oLypLX1L1zbOzFudxZeQOoe664tPSx5amLYULeA1D1L1QVOL0ohY3a6lqYIfXYTLCNONc0IeXEfOA9daWzJtVg99ohVjTtssraXGM5Tl9oCPot85rpwnDeFWYc3SbONXFTmJg7rMaJdPkaC5uDxTg104ybLUMeNNO8p7MxtzBZ1fRsGkrHkrAUJP8H4ZLOGmGZDmfyz+NaS2z91Z4MxPVj+aDnmnHtWG5rMqPZP5YPIDQvHct9TSTLdv9Y8pq4sbANmv2WD+eaF45lhWb/WHZp9heWn83rLe/XtMqXsMQIzdKxzHBC86rC9mv2W4qTpwf9lmBvnzpxTls5ljmgNfsLy3ZWkwt7V1jymiA0C8cyB4PRxA1jyWtmv2wXFJbvrKe//Y9YYnCa1mYZkIydG2LINPvHMoe9izvvLeyr9+TtYkv+CYuH7rHkH3rEyjFY71jyYY61xy97C8uH6WuLbZdbahImFg/EGseSr6wvH0nfWPKWA8Sz0VWFzfF1TSN+vtlSR4LvHI5dMZbkewsyU9Aylrylb647iCj7LAd2N+jAdZY+Unz/qOyuscRYAAdPrETBWAJ8YR/8aPkDHy6srr6+qXq2zQxeGyELwvIL5aJ88CSC80CN2HnZF9/HVWbPoLlIWCJ5GCV+ykGR1JXa3IGJFHld88I+OJA4Jji/uVUE8JPikyNt+eAO4P83y9Rf2ebKTjaXMbM8xLwUWrjZnk0KYiy90FLQZKnFO7NZk+X9G5Z1jCW/C2Z5lOINloJiS/RsmShWW9gCyyZPkwLLnPjXLVWk2zPkwQssSU8+SowCyx5Pq8mS9+QlC05+Oj3DCnaphUivpxXsxuvkDRAe6/t6RuF5Hwo+uqByBuJjUT54yVj2f0VDC6LMwJuSnyuuQ0iMlSz5hhgPgrXMgbOSnaImy8AZyU5Rky3ghGSXqJlsAyckO0xNDlFfd2QkedEwEwK40448QESiSAP81dQVUgI+/8lqwN3/YHiq+B+zd7pmidG7ZwAAAABJRU5ErkJggg==")
	updateTemplateReq.Subject = uspeedo.String("MySubject")

	updateTemplateRes, err := client.UpdateMMSTemplate(updateTemplateReq)
	if err != nil {
		panic(err)
	}
	if updateTemplateRes.GetRetCode() != OK {
		panic(updateTemplateRes.GetMessage())
	}

	fmt.Println("Successfully updated MMS template, template id:", templateId)

	// Delete MMS Template
	deleteTemplateReq := client.NewDeleteMMSTemplateReq()
	deleteTemplateReq.Action = uspeedo.String("QueryMMSTemplate")
	deleteTemplateReq.AccountId = uspeedo.Int(AccountId)
	deleteTemplateReq.TemplateIds = []string{templateId}

	deleteTemplateRes, err := client.DeleteMMSTemplate(deleteTemplateReq)
	if err != nil {
		panic(err)
	}
	if deleteTemplateRes.GetRetCode() != OK {
		panic(deleteTemplateRes.GetMessage())
	}

	fmt.Println("Successfully deleted MMS template, template id:", templateId)
}

func ExampleSendMMSMessage() {
	client := getMMSClient()

	// Send MMS Message
	sendReq := client.NewSendBatchMMSMessageReq()
	sendReq.Action = uspeedo.String("SendBatchMMSMessage")
	sendReq.AccountId = uspeedo.Int(AccountId)
	sendReq.TaskContent = []models.SendInfo{
		{
			SenderId: uspeedo.String("uSpeedo"),
			Target: []models.TargetPhone{
				{
					Phone:          uspeedo.String("(11)1111111"),
					TemplateParams: []string{"TemplateParam1"},
				},
			},
			TemplateId: uspeedo.String("TemplateId"),
		},
	}

	sendRes, err := client.SendBatchMMSMessage(sendReq)
	if err != nil {
		panic(err)
	}
	if sendRes.GetRetCode() != OK {
		panic(sendRes.GetMessage())
	}

	var sessionNo string
	if sendRes.SessionNo != nil {
		sessionNo = *sendRes.SessionNo
	}
	for _, failContent := range sendRes.FailContent {
		if failContent.FailureDetails != nil && *failContent.FailureDetails != "" {
			fmt.Printf("failed to send mms message, template id: %v, sender id: %v, error message: %v\n", failContent.TemplateId, failContent.SenderId, failContent.FailureDetails)
		}

		for _, target := range failContent.Target {
			fmt.Printf("failed to send mms message, phone: %v, invalid: %v, template params: %v, error message: %v\n", target.Phone, target.Invalid, target.TemplateParams, target.FailureDetails)
		}
	}

	fmt.Printf("session no: %v, success count: %v\n", sessionNo, sendRes.SuccessCount)

	if sessionNo == "" {
		return
	}

	// Get MMS Send Receipt
	getSendReceiptReq := client.NewGetMMSSendReceiptReq()
	getSendReceiptReq.Action = uspeedo.String("GetMMSSendReceipt")
	getSendReceiptReq.AccountId = uspeedo.Int(AccountId)
	getSendReceiptReq.SessionNoSet = []string{sessionNo}

	getSendReceiptRes, err := client.GetMMSSendReceipt(getSendReceiptReq)
	if err != nil {
		panic(err)
	}
	if getSendReceiptRes.GetRetCode() != OK {
		panic(getSendReceiptRes.GetMessage())
	}

	for _, datum := range getSendReceiptRes.Data {
		fmt.Println("session no:", datum.SessionNo)
		for _, target := range datum.ReceiptSet {
			fmt.Printf("phone: %v, receipt code: %v, receipt result: %v, receipt desc: %v\n", target.Phone, target.ReceiptCode, target.ReceiptResult, target.ReceiptDesc)
		}
	}
}

func main() {
	ExampleMMSTemplate()
	ExampleSendMMSMessage()
}
