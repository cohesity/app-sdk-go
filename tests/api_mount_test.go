package test

import (
	"fmt"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	mgmtModels "github.com/cohesity/management-sdk-go/models"
	"github.com/stretchr/testify/suite"
)

type Mount struct {
	suite.Suite
	athena     AthenaStruct
	mgmtClient CohesityManagementSdk.COHESITYMANAGEMENTSDK
	viewName   string
}

func (mnt *Mount) SetupTest() {
	mnt.athena.setmeUp()

	// Create View
	var err error
	var Domain string = "LOCAL"
	ClusterVIP := fmt.Sprintf("%v", mnt.athena.App["AppEndpointIp"])
	ClusterUsername := fmt.Sprintf("%v", mnt.athena.App["ClusterUsername"])
	ClusterPassword := fmt.Sprintf("%v", mnt.athena.App["ClusterPassword"])
	mnt.mgmtClient, err = CohesityManagementSdk.NewCohesitySdkClient(ClusterVIP, ClusterUsername, ClusterPassword, Domain)
	if err != nil {
		panic(err)
	}
	storageDomain := mnt.mgmtClient.ViewBoxes()
	var defaultStorageDomain []string
	defaultStorageDomain = append(defaultStorageDomain, getString(mnt.athena.App["DefaultStorageDomain"]))
	result, err := storageDomain.GetViewBoxes(nil, nil, nil, defaultStorageDomain, nil, nil)
	if err != nil {
		panic(err)
	}
	sdID := result[0].Id
	view_controller := mnt.mgmtClient.Views()
	var viewRequest *mgmtModels.CreateViewRequest
	viewRequest.Name = getString(mnt.athena.App["ViewName"])
	viewRequest.ViewBoxId = *sdID
	viewResult, err := view_controller.CreateView(viewRequest)
	if err != nil {
		panic(err)
	}
	mnt.viewName = *viewResult.Name

	//mnt.athena.appClient.Mount().CreateMount()
}

func (mnt *Mount) TearDownSuite() {
	viewController := mnt.mgmtClient.Views()
	err := viewController.DeleteView(mnt.viewName)
	if err != nil {
		panic(err)
	}
}
func (mnt *Mount) Test01Mount() {

}

func (mnt *Mount) Test02Unmount() {

}

func (mnt *Mount) Test03NegativeBogusUnMount() {

}

func (mnt *Mount) Test04NegativeMountingWithWritePermission() {
	//TODO
	// This operation should fail because app has ro access insteda of rw access.
}

func (mnt *Mount) Test05NegativeMountWithNestedDirectories() {

}

func (mnt *Mount) Test04NegativeBogusPayloadMount() {

}

func TestMount(t *testing.T) {
	// This is what actually runs our suite
	a := new(Mount)
	suite.Run(t, a)
}
