package response

import (
	"encoding/json"
	"fmt"
)

type Responses []ResponseElement

func UnmarshalResponses(data []byte) (Responses, error) {
	var r Responses
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalResponseElement(data []byte) (ResponseElement, error) {
	var r ResponseElement
	err := json.Unmarshal(data, &r)
	return r, err
}

type ResponseElement struct {
	ID        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	Exception string `json:"exception,omitempty"`
	Errorcode string `json:"errorcode,omitempty"`
	Message   string `json:"message,omitempty"`
	Debuginfo string `json:"debuginfo,omitempty"`
}

func ParseResponseToStruct(data []byte) (Responses, error) {
	if r, err := UnmarshalResponseElement(data); err == nil {
		return Responses{r}, err
	}
	if re, err := UnmarshalResponses(data); err == nil {
		return re, err
	}
	return nil, fmt.Errorf("não foi possivel realizar o parse: %s", string(data))
}
