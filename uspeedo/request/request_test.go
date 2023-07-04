package request

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

func TestRequestAccessor(t *testing.T) {
	var err error
	req := &CommonBase{}

	s := "foo"
	assert.Equal(t, "", req.GetAction())
	assert.Nil(t, req.GetActionRef())
	err = req.SetActionRef(&s)
	assert.NoError(t, err)
	assert.NotNil(t, req.GetActionRef())

	req.SetRetryCount(42)
	assert.Equal(t, 42, req.GetRetryCount())

	req.SetRetryable(true)
	assert.Equal(t, true, req.GetRetryable())
	req.SetRetryable(false)

	req.WithRetry(42)
	assert.Equal(t, 42, req.GetMaxretries())
	assert.Equal(t, true, req.GetRetryable())

	now := time.Now()
	req.SetRequestTime(now)
	assert.Equal(t, now, req.GetRequestTime())

	req.WithTimeout(1 * time.Second)
	assert.Equal(t, 1*time.Second, req.GetTimeout())

	err = req.SetAction("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.GetAction())
	assert.Equal(t, String("Zm9v"), ToBase64Query(String(req.GetAction())))
}

func TestRequestAccessorForGeneric(t *testing.T) {
	var err error
	req := &BaseGenericRequest{}

	assert.Equal(t, "", req.GetAction())
	err = req.CommonBase.SetAction("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.GetAction())

	assert.Equal(t, map[string]interface{}{
		"Action": "foo",
	}, req.GetPayload())
	assert.NoError(t, req.SetPayload(map[string]interface{}{
		"Action": "bar",
	}))
	assert.Equal(t, "bar", req.GetAction())

	assert.Error(t, req.SetPayload(map[string]interface{}{
		"Action": 1,
	}))
}

func TestEncodeOne(t *testing.T) {
	i := 42

	tests := []struct {
		name        string
		inputVector interface{}
		golden      string
		wantErr     bool
	}{
		{"int", 42, "42", false},
		{"uint", uint(42), "42", false},
		{"float", 42.0, "42", false},
		{"float", 42.1, "42.1", false},
		{"bool", true, "true", false},
		{"string", "foo", "foo", false},
		{"pointer", &i, "42", false},
		{"error", struct{}{}, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rv := reflect.ValueOf(test.inputVector)
			got, err := encodeOne(&rv)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.golden, got)
				assert.NoError(t, err)
			}
		})
	}
}

func TestToQueryMap(t *testing.T) {
	type Composite struct {
	}

	type compositedRequest struct {
		CommonBase
	}

	type args struct {
		req Common
	}

	tests := []struct {
		name    string
		args    args
		json    map[string]interface{}
		wantErr bool
	}{
		{
			"Ok",
			args{
				req: &struct {
					CommonBase
					Id      int
					Name    string
					IsValid bool
					Ips     []string
				}{
					Id:      1,
					Name:    "liLei",
					IsValid: true,
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]interface{}{
				"Id":      1,
				"Name":    "liLei",
				"IsValid": true,
				"Ips":     []interface{}{"127.0.0.1", "192.168.1.1"},
			},
			false,
		},
		{
			"Partial",
			args{
				req: &struct {
					CommonBase
					Id      *int
					Name    *string
					IsValid *bool
					Ips     []string
				}{
					Id:      Int(1),
					IsValid: Bool(true),
					Ips:     []string{"127.0.0.1", "192.168.1.1"},
				},
			},
			map[string]interface{}{
				"Id":      1,
				"IsValid": true,
				"Ips":     []interface{}{"127.0.0.1", "192.168.1.1"},
			},
			false,
		},
		{
			"IsComposited",
			args{
				req: &struct {
					CommonBase
					Composite
				}{
					CommonBase{},
					Composite{},
				},
			},
			map[string]interface{}{
				"Composite": map[string]interface{}{},
			},
			false,
		},
		{
			"Common",
			args{
				req: func() Common {
					req := &compositedRequest{}
					return req
				}(),
			},
			map[string]interface{}{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonEncoded, err := EncodeJSON(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.json, jsonEncoded) {
				t.Errorf("EncodeJSON() = %v, json %v", jsonEncoded, tt.json)
			}
		})
	}
}

func TestToQueryMapForGeneric(t *testing.T) {
	var typedNil *string
	tests := []struct {
		name    string
		payload map[string]interface{}
		wantErr bool
	}{
		{
			"Ok",
			map[string]interface{}{
				"Id":      1,
				"Name":    "liLei",
				"IsValid": true,
				"Ips":     []string{"127.0.0.1", "192.168.1.1"},
			},
			false,
		},

		{
			"Partial",
			map[string]interface{}{
				"Id":      Int(1),
				"Name":    typedNil,
				"IsValid": Bool(true),
				"Ips":     []string{"127.0.0.1", "192.168.1.1"},
			},
			false,
		},
		{
			"Ok_nest_map",
			map[string]interface{}{
				"Str":  "str",
				"Int":  1,
				"Bool": true,
				"Nest": []map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": 1,
							"Str": "str",
						},
					},
				},
			},
			false,
		},
		{
			"Ok_ptr_map",
			map[string]interface{}{
				"Nest": []*map[string]interface{}{
					{
						"Nest2": map[string]interface{}{
							"Int": 1,
							"Str": "str",
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &BaseGenericRequest{}
			assert.NoError(t, req.SetPayload(tt.payload))

			jsonEncoded, err := EncodeJSON(req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.payload, jsonEncoded) {
				t.Errorf("EncodeJSON() = %v, json %v", jsonEncoded, tt.payload)
			}

		})
	}
}

func TestEncoderError(t *testing.T) {
	var err error
	tests := []struct {
		name        string
		req         map[string]interface{}
		wantFormErr bool
		wantJSONErr bool
	}{
		{
			"Invalid_type",
			map[string]interface{}{
				"foo": auth.NewCredential(),
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		req := &BaseGenericRequest{}
		assert.NoError(t, req.SetPayload(tt.req))

		_, err = EncodeJSON(req)
		if (err != nil) != tt.wantJSONErr {
			t.Errorf("EncodeJSON() error = %v, wantErr %v", err, tt.wantJSONErr)
			return
		}

	}
}
