// Copyright 2019 Cohesity Inc.
package volume

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
type VOLUME_IMPL struct {
	config configuration.CONFIGURATION
}

/**
 * Gets the status of persistent volume owned by this app.
 * @param    string        volumeName     parameter: Required
 * @return	Returns the *models.VolumeInfo response from the API call
 */
func (me *VOLUME_IMPL) GetVolume(
	GetVolume *models.GetVolumeParams) (*models.VolumeInfo, error) {
	//the endpoint path uri
	_pathUrl := "/volumes/{volumeName}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"volumeName": GetVolume.VolumeName,
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
		err = apihelper.NewAPIError("Volume doesn't exist.", _response.Code, _response.RawBody)
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
	var retVal *models.VolumeInfo = &models.VolumeInfo{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Delete a previously created persistent volume owned by this app.
 * @param    string        volumeName     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *VOLUME_IMPL) DeleteVolume(
	DeleteVolume *models.DeleteVolumeParams) error {
	//the endpoint path uri
	_pathUrl := "/volumes/{volumeName}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"volumeName": DeleteVolume.VolumeName,
	})
	if err != nil {
		//error in template param handling
		return err
	}

	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.ENUM_DEFAULT, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "app-Go-sdk-1.1.1",
		"Authorization": fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
	}

	//prepare API request
	_request := unirest.Delete(_queryBuilder, headers, nil)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 400 {
		err = apihelper.NewAPIError("Invalid parameters.", _response.Code, _response.RawBody)
	} else if _response.Code == 401 {
		err = apihelper.NewAPIError("Unauthorized.", _response.Code, _response.RawBody)
	} else if _response.Code == 404 {
		err = apihelper.NewAPIError("Volume doesn't exist.", _response.Code, _response.RawBody)
	} else if _response.Code == 500 {
		err = apihelper.NewAPIError("Unexpected error.", _response.Code, _response.RawBody)
	} else if _response.Code == 502 {
		err = apihelper.NewAPIError("Bad Gateway.", _response.Code, _response.RawBody)
	} else if _response.Code == 504 {
		err = apihelper.NewAPIError("Gateway Timeout.", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil

}

/**
 * Use this API to create a new kubernetes persistent volume backed up by cohesity view.
 * @param    string                        volumeName     parameter: Required
 * @param    *models.VolumeSpec        volumeSpec     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *VOLUME_IMPL) CreateVolume(
	CreateVolume *models.CreateVolumeParams) error {
	//the endpoint path uri
	_pathUrl := "/volumes/{volumeName}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"volumeName": CreateVolume.VolumeName,
	})
	if err != nil {
		//error in template param handling
		return err
	}

	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.ENUM_DEFAULT, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "app-Go-sdk-1.1.1",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("Bearer %s", me.config.OAuthAccessToken()),
	}

	//prepare API request
	_request := unirest.Put(_queryBuilder, headers, CreateVolume.VolumeSpec)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 401 {
		err = apihelper.NewAPIError("Unauthorized.", _response.Code, _response.RawBody)
	} else if _response.Code == 409 {
		err = apihelper.NewAPIError("Volume already exists with different parameters.", _response.Code, _response.RawBody)
	} else if _response.Code == 500 {
		err = apihelper.NewAPIError("Unexpected error.", _response.Code, _response.RawBody)
	} else if _response.Code == 502 {
		err = apihelper.NewAPIError("Bad Gateway.", _response.Code, _response.RawBody)
	} else if _response.Code == 504 {
		err = apihelper.NewAPIError("Gateway Timeout.", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil

}
