package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hue-controls/initialise"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initialise.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

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

		request, err := http.NewRequest(http.MethodPut, hueApiUrl+"/lights/2/state", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("ERROR:", err)
		}

		request.Header.Set("Content-Type", "application/json")

		// Create a new HTTP client
		client := &http.Client{}

		// Send the request
		response, err := client.Do(request)

		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer response.Body.Close()

		// Check the response status code
		if response.StatusCode != http.StatusOK {
			fmt.Println("Unexpected status code:", response.StatusCode)
			return
		}

		// Print the response status
		fmt.Println("Response:", response.Status)

		if err != nil {
			fmt.Print(err)
		}

		c.JSON(200, gin.H{
			"message": "Success",
		})

	})
	r.Run()
}
