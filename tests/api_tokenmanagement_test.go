package test

import (
	"fmt"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	mgmtModel "github.com/cohesity/management-sdk-go/models"
	"github.com/stretchr/testify/suite"
)

type Token struct {
	suite.Suite
	athena AthenaStruct
}

func (token *Token) SetupTest() {
	token.athena.setmeUp()
}
func (token *Token) Test01CreateManagementAccessToken() {
	generatedToken, err := token.athena.appClient.TokenManagement().CreateManagementAccessToken()
	if err != nil {
		panic(err)
	}
	irisToken := new(mgmtModel.AccessToken)
	irisToken.AccessToken = generatedToken.AccessToken
	irisToken.TokenType = generatedToken.TokenType
	token.Equal(fmt.Sprintf("%v", token.athena.App["TokenType"]), *generatedToken.TokenType)

	var fetchStats *bool
	client := CohesityManagementSdk.NewCohesitySdkClientWithToken(fmt.Sprintf("%v", token.athena.App["AppEndpointIp"]), irisToken)
	clusterDetails, err := client.Cluster().GetCluster(fetchStats)
	if err != nil {
		println(err.Error())
	}
	token.Equal(*clusterDetails.Name, fmt.Sprintf("%v", token.athena.App["ClusterName"]))

	//CohesityManagementSdk.NewCohesitySdkClientWithToken()
}

func TestToken(t *testing.T) {
	// This is what actually runs our suite
	a := new(Token)
	suite.Run(t, a)
}
