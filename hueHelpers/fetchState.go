package hueHelpers

import (
	"fmt"
	"go-hue-controls/httpRequestClient"
	marshalData "go-hue-controls/marshal-data"
	"net/http"
	"os"
)

type LightState struct {
	On bool `json:"on"`
}

type Response struct {
	State LightState `json:"state"`
}

func RequestHueState(lightNo string) Response {
	hueIpAddress := os.Getenv("HUE_IP_ADDRESS")
	hueUsername := os.Getenv("HUE_USERNAME")

	hueStateApiUrl := fmt.Sprintf("http://%s/api/%s/lights/%s", hueIpAddress, hueUsername, lightNo)
	lightStatusResponse := httpRequestClient.Request(http.MethodGet, hueStateApiUrl, nil)

	var lightStatus Response

	marshalData.UnmarshalJson(lightStatusResponse, &lightStatus)

	return lightStatus
}
