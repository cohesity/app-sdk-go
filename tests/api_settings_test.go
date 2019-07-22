package test

import (
	"testing"

	appsdk "github.com/cohesity/app-sdk-go/appsdk"
	"github.com/cohesity/app-sdk-go/models"
	"github.com/stretchr/testify/suite"
)

type Settings struct {
	suite.Suite
	athena AthenaStruct
}

func (settings *Settings) SetupTest() {
	settings.athena.setmeUp()
}
func (setting *Settings) Test01GetAppSettings() {
	result, _ := setting.athena.appClient.Settings().GetAppSettings()
	setting.Equal(*result.User.UserName, "admin")
	setting.Equal(*result.User.Domain, "LOCAL")
	setting.Equal(models.QosTierEnum(1), result.AppInstanceSettings.QosTier)
	//setting.Equal(models.PrivilegesTypeEnum(0), result.AppInstanceSettings.ReadViewPrivileges.PrivilegesType)
	//var ViewIdReturnType *[]int64
	//setting.IsType(ViewIdReturnType, result.AppInstanceSettings.ReadViewPrivileges.ViewIds)
}

func (setting *Settings) Test02NegativeBogusToken() {
	AppEndpointIp := getString(setting.athena.App["AppEndpointIp"])
	AppEndpointPort := getString(setting.athena.App["AppEndpointPort"])
	setting.athena.appClient = appsdk.NewAppSdkClient("bogusToken", AppEndpointIp, AppEndpointPort)
	_, err := setting.athena.appClient.Settings().GetAppSettings()
	setting.NotEqual(err, nil)
	print(err)
}

func TestSettings(t *testing.T) {
	// This is what actually runs our suite
	a := new(Settings)
	suite.Run(t, a)
}
