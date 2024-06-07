package httpRequestClient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func RequestGet(endpoint string) []byte {
	return Request(http.MethodGet, endpoint, nil)
}

func Request(method string, endpoint string, marshalledData []byte) []byte {
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

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	return body

}
