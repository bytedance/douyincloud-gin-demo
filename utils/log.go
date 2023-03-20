package utils

import "encoding/json"

func ToJsonString(i interface{}) string {
	if i == nil {
		return ""
	}

	data, err := json.Marshal(i)
	if err != nil {
		return ""
	}

	return string(data)
}
