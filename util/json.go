package util

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

func MarshalToStringNoErr(v interface{}) string {
	if v == nil {
		return ""
	}
	toString, _ := jsoniter.MarshalToString(v)
	return toString
}

func UnMarshalToStringNoErr(s string, v interface{}) {
	_ = json.Unmarshal([]byte(s), v)
}
