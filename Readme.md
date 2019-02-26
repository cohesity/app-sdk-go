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
