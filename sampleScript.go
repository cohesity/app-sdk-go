// Copyright 2019 Cohesity Inc.
//
// Script to demonstrate usage of Cohesity App SDK in Golang.
//
// Usage: go run sampleSript.go

package main

import (
	"fmt"
	"os"

	CohesityAppSdk "github.com/cohesity/app-sdk-go/appsdk"
	"github.com/cohesity/app-sdk-go/models"
)

func main() {

	// Get Environment variables.
	AppAuthToken := os.Getenv("APP_AUTHENTICATION_TOKEN")
	AppEndpointIp := os.Getenv("APPS_API_ENDPOINT_IP")
	AppEndpointPort := os.Getenv("APPS_API_ENDPOINT_PORT")

	// Initialize the App Client.
	client := CohesityAppSdk.NewAppSdkClient(AppAuthToken, AppEndpointIp, AppEndpointPort)
	settings := client.Settings()

	var err error
	var result *models.AppSettings
	result, err = settings.GetAppSettings()

	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println("App Settings User: ", *result.User.UserName)
	}

}
