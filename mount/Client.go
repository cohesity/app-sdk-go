// Copyright 2019 Cohesity Inc.
package mount

import (
	"fmt"

	"github.com/cohesity/app-sdk-go/apihelper"
	"github.com/cohesity/app-sdk-go/configuration"
	"github.com/cohesity/app-sdk-go/models"
	unirest "github.com/cohesity/management-sdk-go/unirest-go"
)

/*
 * Client structure as interface implementation
 */
type MOUNT_IMPL struct {
	config configuration.CONFIGURATION
}

/**
 * Unmount previously mounted view/namespace or volume of a protected entity.
 * @param    string        dirName     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *MOUNT_IMPL) DeleteUnmount(
	DeleteUnmount *models.DeleteUnmountParams) error {
	//the endpoint path uri
	_pathUrl := "/mounts/{dirName}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"dirName": DeleteUnmount.DirName,
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
		err = apihelper.NewAPIError("Invalid token.", _response.Code, _response.RawBody)
	} else if _response.Code == 404 {
		err = apihelper.NewAPIError("Directory doesn't exist.", _response.Code, _response.RawBody)
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
 * Allows you to mount a cohesity external view or snapshots of a protected object (VM volumes or NAS).
 * @param    *models.MountOptions        mountOptions     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *MOUNT_IMPL) CreateMount(
	CreateMount *models.CreateMountParams) error {
	//the endpoint path uri
	_pathUrl := "/mounts"

	//variable to hold errors
	var err error = nil
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
	_request := unirest.Post(_queryBuilder, headers, CreateMount.MountOptions)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 400 {
		err = apihelper.NewAPIError("Validation errors.", _response.Code, _response.RawBody)
	} else if _response.Code == 401 {
		err = apihelper.NewAPIError("Invalid token.", _response.Code, _response.RawBody)
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
