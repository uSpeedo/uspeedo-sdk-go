package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
QueryUSMSTemplateResp -
*/
type QueryUSMSTemplateResp struct {
	response.CommonBase

	/**
	 *
	 */
	Data []models.OutTemplate `required:""`
}
