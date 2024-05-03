package hueRequest

import (
	"bytes"
	"fmt"
	"net/http"
)

func Request(method string, endpoint string, marshalledData []byte) *http.Response {
	request, newReqerr := http.NewRequest(method, endpoint, bytes.NewBuffer(marshalledData))
	request.Header.Set("Content-Type", "application/json")

	if newReqerr != nil {
		fmt.Println("New request error:", newReqerr)
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", response.StatusCode)
		return nil
	}

	return response

}
