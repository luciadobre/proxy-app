// core/json.go
package core

import (
	"encoding/json"
)

func ModifyJSONResponse(responseBody []byte) ([]byte, error) {
	var data map[string]interface{}
	err := json.Unmarshal(responseBody, &data)
	if err != nil {
		return nil, err
	}

	data["foo"] = "bar"

	modifiedBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return modifiedBody, nil
}
