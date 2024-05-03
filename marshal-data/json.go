package marshalData

import (
	"encoding/json"
	"fmt"
)

func Json(data map[string]interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil
	}

	return jsonData
}
