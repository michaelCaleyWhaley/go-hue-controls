package marshalData

import (
	"encoding/json"
	"fmt"
)

func MarshalJson(data map[string]interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil
	}

	return jsonData
}

func UnmarshalJson(data []byte, parsedJson any) {
	json.Unmarshal(data, &parsedJson)
}
