package request

import (
	"fmt"
	"reflect"
)

type GenericRequest interface {
	Common

	SetPayload(m map[string]interface{}) error
	GetPayload() map[string]interface{}
}

type BaseGenericRequest struct {
	CommonBase

	payload map[string]interface{}
}

func (r *BaseGenericRequest) SetPayload(m map[string]interface{}) error {
	if m["Action"] != nil && reflect.ValueOf(m["Action"]).Type().Kind() != reflect.String {
		return fmt.Errorf("request SetPayload error, the Action must set a String value")
	}
	r.payload = m
	return nil
}

func (r *BaseGenericRequest) GetPayload() map[string]interface{} {
	m := make(map[string]interface{})
	if len(r.CommonBase.GetAction()) != 0 {
		m["Action"] = r.CommonBase.GetAction()
	}

	for k, v := range r.payload {
		m[k] = v
	}
	return m
}

func (r *BaseGenericRequest) GetAction() string {
	if r.payload["Action"] != nil {
		return r.payload["Action"].(string)
	}

	return r.CommonBase.GetAction()
}
