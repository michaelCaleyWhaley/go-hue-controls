package main

import (
	"encoding/json"
	"fmt"
	hueRequest "go-hue-controls/hue-request"
	"go-hue-controls/initialise"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initialise.LoadEnvVariables()
}

func rootHandler(c *gin.Context) {

	data := map[string]interface{}{
		"on": false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	hueIpAddress := os.Getenv("HUE_IP_ADDRESS")
	hueUsername := os.Getenv("HUE_USERNAME")

	hueApiUrl := fmt.Sprintf("http://%s/api/%s/lights/2/state", hueIpAddress, hueUsername)

	// lights to be toggled
	// 1,2,3,4

	resLightTwo := hueRequest.Request(http.MethodPut, hueApiUrl+"/lights/2/state", jsonData)

	if resLightTwo == nil {
		fmt.Println("Failed to turn on light")
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})

}

func main() {
	r := gin.Default()
	r.GET("/", rootHandler)
	r.Run()
}
