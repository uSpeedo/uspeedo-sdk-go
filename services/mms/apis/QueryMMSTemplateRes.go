package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
QueryMMSTemplateRes -
*/
type QueryMMSTemplateRes struct {
	response.CommonBase

	/**
	 *
	 */
	Data []models.OutTemplate `required:""`
}
