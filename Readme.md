Cohesity App SDK
=================

# Overview

The *Cohesity App SDK*  provides an easy-to-use language binding to 
harness the power of *Cohesity App APIs* in your Golang applications.These
 APIs are available for apps running on Cohesity Apps Infrastructure.


# Getting Started

```bash
go get github.com/cohesity/app-sdk-go
```
Note: To update the package do:

```
 go get -u github.com/cohesity/app-sdk-go
```

## Container Environment Parameters

The App Environment Container has the following parameters initialized by 
Athena.
```
HOST_IP
APPS_API_ENDPOINT_IP
POD_UID
APPS_API_ENDPOINT_PORT
APP_AUTHENTICATION_TOKEN
```
We use the above variables in various use cases to initialize and make call 
to Athena App server.

### How to use
In order to setup authentication and initialization of the API client, you need the following information.

```
AppAuthToken := os.Getenv("APP_AUTHENTICATION_TOKEN")
AppEndpointIp := os.Getenv("APPS_API_ENDPOINT_IP")
AppEndpointPort := os.Getenv("APPS_API_ENDPOINT_PORT")

//Initialize the App Client.
client := CohesityAppSdk.NewAppSdkClient(AppAuthToken, AppEndpointIp, AppEndpointPort)
settings := client.Settings()

//Get App Settings.
var err error
var result *models.AppSettings
result, err = settings.GetAppSettings()
```

## Questions or Feedback :
We would love to hear from you. Please send your questions and feedback to: *developer@cohesity.com*
# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [protectedsourcevolumeinfo](#protectedsourcevolumeinfo)
* [volume](#volume)
* [tokenmanagement](#tokenmanagement)
* [settings](#settings)
* [mount](#mount)

## <a name="protectedsourcevolumeinfo"></a>![Class: ](https://apidocs.io/img/class.png ".protectedsourcevolumeinfo") protectedsourcevolumeinfo

### Get instance

Factory for the ``` PROTECTEDSOURCEVOLUMEINFO ``` interface can be accessed from the package protectedsourcevolumeinfo.

```go
protectedSourceVolumeInfo := protectedsourcevolumeinfo.NewPROTECTEDSOURCEVOLUMEINFO()
```

### <a name="get_protected_source_volume_info"></a>![Method: ](https://apidocs.io/img/method.png ".protectedsourcevolumeinfo.GetProtectedSourceVolumeInfo") GetProtectedSourceVolumeInfo

> Gets the list of volumes for a snapshot of a protected source.


```go
func (me *PROTECTEDSOURCEVOLUMEINFO_IMPL) GetProtectedSourceVolumeInfo(*models.GetSourceVolumeInfoParams)(*models.ProtectedSourceVolumeInfo,error)
```

#### Parameters

| Parameter Type | Tags | Description |
|-----------|------|-------------|
| *models.GetSourceVolumeInfoParams |  ``` Required ```  | Unique ID of the protected source. |


#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Unauthorized |
| 404 | Snapshot does not exist. |
| 500 | Unexpected error |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



[Back to List of Controllers](#list_of_controllers)

## <a name="volume"></a>![Class: ](https://apidocs.io/img/class.png ".volume") volume

### Get instance

Factory for the ``` VOLUME ``` interface can be accessed from the package volume.

```go
volume := volume.NewVOLUME()
```

### <a name="get_volume"></a>![Method: ](https://apidocs.io/img/method.png ".volume.GetVolume") GetVolume

> Gets the status of persistent volume owned by this app.


```go
func (me *VOLUME_IMPL) GetVolume(VolumeParams *models.GetVolumeParams)
(*models.VolumeInfo, error)
```

#### Parameters

| Parameter Type | Tags | Description |
|-----------|------|-------------|
| *models.GetVolumeParams |  ``` Required ```  | Name of the volume unique within the app instance. |




#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Unauthorized |
| 404 | Volume doesn't exist. |
| 500 | Unexpected error |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



### <a name="delete_volume"></a>![Method: ](https://apidocs.io/img/method.png ".volume.DeleteVolume") DeleteVolume

> Delete a previously created persistent volume owned by this app.


```go
func (me *VOLUME_IMPL) DeleteVolume(DeleteVolume *models.DeleteVolumeParams)(,error)
```

#### Parameters

| Parameter Type | Tags | Description |
|-----------|------|-------------|
| *models.DeleteVolumeParams |  ``` Required ```  | Name of the volume unique 
within the app instance. |



#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid parameters. |
| 401 | Unauthorized. |
| 404 | Volume doesn't exist. |
| 500 | Unexpected error. |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



### <a name="create_volume"></a>![Method: ](https://apidocs.io/img/method.png ".volume.CreateVolume") CreateVolume

> Use this API to create a new kubernetes persistent volume backed up by cohesity view.


```go
func (me *VOLUME_IMPL) CreateVolume( VolumeParam *models.CreateVolumeParams)(,error)
```

#### Parameters

| Parameter Type | Tags | Description |
|-----------|------|-------------|
| *models.CreateVolumeParams |  ``` Required ```  | Name of the volume unique within the app instance. |


#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Unauthorized. |
| 409 | Volume already exists with different parameters. |
| 500 | Unexpected error. |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



[Back to List of Controllers](#list_of_controllers)

## <a name="tokenmanagement"></a>![Class: ](https://apidocs.io/img/class.png ".tokenmanagement") tokenmanagement

### Get instance

Factory for the ``` TOKENMANAGEMENT ``` interface can be accessed from the package tokenmanagement.

```go
tokenManagement := tokenmanagement.NewTOKENMANAGEMENT()
```

### <a name="create_management_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokenmanagement.CreateManagementAccessToken") CreateManagementAccessToken

> Use this api to get a new management api token.


```go
func (me *TOKENMANAGEMENT_IMPL) CreateManagementAccessToken()(*models.ManagementAccessToken,error)
```

#### Example Usage

```go

var result *models.ManagementAccessToken
result,_ = tokenManagement.CreateManagementAccessToken()

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Invalid token. |
| 500 | Unexpected error. |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



[Back to List of Controllers](#list_of_controllers)

## <a name="settings"></a>![Class: ](https://apidocs.io/img/class.png ".settings") settings

### Get instance

Factory for the ``` SETTINGS ``` interface can be accessed from the package settings.

```go
settings := settings.NewSETTINGS()
```

### <a name="get_app_settings"></a>![Method: ](https://apidocs.io/img/method.png ".settings.GetAppSettings") GetAppSettings

> Returns app settings object.


```go
func (me *SETTINGS_IMPL) GetAppSettings()(*models.AppSettings,error)
```

#### Example Usage

```go

var result *models.AppSettings
result,_ = settings.GetAppSettings()

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 401 | Invalid token |
| 500 | Unexpected error |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



[Back to List of Controllers](#list_of_controllers)

## <a name="mount"></a>![Class: ](https://apidocs.io/img/class.png ".mount") mount

### Get instance

Factory for the ``` MOUNT ``` interface can be accessed from the package mount.

```go
mount := mount.NewMOUNT()
```

### <a name="delete_unmount"></a>![Method: ](https://apidocs.io/img/method.png ".mount.DeleteUnmount") DeleteUnmount

> Unmount previously mounted view/namespace or volume of a protected entity.


```go
func (me *MOUNT_IMPL) DeleteUnmount(MountParams *models.DeleteUnmountParams)(,
error)
```

#### Parameters

| Parameter Type | Tags | Description |
|-----------|------|-------------|
| *models.DeleteUnmountParams |  ``` Required ```  | Name of the mount directory. |



#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid parameters. |
| 401 | Invalid token. |
| 404 | Directory doesn't exist. |
| 500 | Unexpected error. |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |



### <a name="create_mount"></a>![Method: ](https://apidocs.io/img/method.png ".mount.CreateMount") CreateMount

> Allows you to mount a cohesity external view or snapshots of a protected object (VM volumes or NAS).


```go
func (me *MOUNT_IMPL) CreateMount(MountParam *models.CreateMountParams)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| *models.CreateMountParams |  ``` Required ```  | TODO: Add a parameter description |



#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Validation errors. |
| 401 | Invalid token. |
| 500 | Unexpected error. |
| 502 | Bad Gateway. |
| 504 | Gateway Timeout. |

