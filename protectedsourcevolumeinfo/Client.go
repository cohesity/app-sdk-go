// Copyright 2019 Cohesity Inc.
package protectedsourcevolumeinfo

import (
	"encoding/json"
	"fmt"

	"github.com/cohesity/app-sdk-go/apihelper"
	"github.com/cohesity/app-sdk-go/configuration"
	"github.com/cohesity/app-sdk-go/models"
	unirest "github.com/cohesity/management-sdk-go/unirest-go"
)

/*
 * Client structure as interface implementation
 */
type PROTECTEDSOURCEVOLUMEINFO_IMPL struct {
	config configuration.CONFIGURATION
}

/**
 * Gets the list of volumes for a snapshot of a protected source.
 * @param    int64        sourceId     parameter: Required
 * @return	Returns the *models.ProtectedSourceVolumeInfo response from the API call
 */
func (me *PROTECTEDSOURCEVOLUMEINFO_IMPL) GetProtectedSourceVolumeInfo(
	SourceInfo *models.GetSourceVolumeInfoParams) (*models.ProtectedSourceVolumeInfo, error) {
	//the endpoint path uri
	_pathUrl := "/protectedSourceVolumeInfo/{sourceId}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"sourceId": SourceInfo.SourceId,
	})
	if err != nil {
		//error in template param handling
		return nil, err
	}

	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.ENUM_DEFAULT, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "app-Go-sdk-1.1.1",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 401 {
		err = apihelper.NewAPIError("Unauthorized", _response.Code, _response.RawBody)
	} else if _response.Code == 404 {
		err = apihelper.NewAPIError("Snapshot does not exist.", _response.Code, _response.RawBody)
	} else if _response.Code == 500 {
		err = apihelper.NewAPIError("Unexpected error", _response.Code, _response.RawBody)
	} else if _response.Code == 502 {
		err = apihelper.NewAPIError("Bad Gateway.", _response.Code, _response.RawBody)
	} else if _response.Code == 504 {
		err = apihelper.NewAPIError("Gateway Timeout.", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ProtectedSourceVolumeInfo = &models.ProtectedSourceVolumeInfo{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}
