package test

import (
	"encoding/json"
	"fmt"
	"os"

	appsdk "github.com/cohesity/app-sdk-go/appsdk"
)

type AthenaStruct struct {
	appClient appsdk.COHESITYAPPSDK
	App       map[string]interface{}
}

func LoadConfiguration() map[string]interface{} {
	var App map[string]interface{}
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&App)
	return App
}

func (athena *AthenaStruct) setmeUp() {
	athena.App = LoadConfiguration()
	AppAuthToken := getString(athena.App["AppAuthToken"])
	AppEndpointIp := getString(athena.App["AppEndpointIp"])
	AppEndpointPort := getString(athena.App["AppEndpointPort"])
	athena.appClient = appsdk.NewAppSdkClient(AppAuthToken, AppEndpointIp, AppEndpointPort)
}

func getString(str interface{}) string {
	return fmt.Sprintf("%v", str)
}
