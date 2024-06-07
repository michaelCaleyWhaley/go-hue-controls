package main

import (
	"fmt"
	"go-hue-controls/httpRequestClient"
	hueHelpers "go-hue-controls/hue"
	"go-hue-controls/initialise"
	marshalData "go-hue-controls/marshal-data"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initialise.LoadEnvVariables()
}

func rootHandler(c *gin.Context) {
	hueIpAddress := os.Getenv("HUE_IP_ADDRESS")
	hueUsername := os.Getenv("HUE_USERNAME")
	lightNo := c.Param("number")

	lightStatus := hueHelpers.RequestHueState(lightNo)

	jsonData := marshalData.MarshalJson(map[string]interface{}{
		"on": !lightStatus.State.On,
	})
	hueApiUrl := fmt.Sprintf("http://%s/api/%s/lights/%s/state", hueIpAddress, hueUsername, lightNo)

	resLight := httpRequestClient.Request(http.MethodPut, hueApiUrl, jsonData)

	if resLight == nil {
		fmt.Println("Failed to turn on light")
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})

}

func main() {
	r := gin.Default()
	r.GET("/light/:number", rootHandler)
	r.Run()
}
