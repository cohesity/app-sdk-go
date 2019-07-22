# Before running: 

Make sure the config file is populated with all the values 
Make sure the app has access to current and future views in the cluster. 

# Import testify 
```
go get github.com/stretchr/testify
go get github.com/cohesity/app-sdk-go
go get github.com/cohesity/management-sdk-go
```

# Make sure your app is accessible outside the cluster => Disable firewall. 
Login to cluster:
```
sudo service firewalld stop 
```

### Test run commands 
```
go test github.com/cohesity/app-sdk-go/tests -run TestSettings/Test01GetAppSettings -v 
go test github.com/cohesity/app-sdk-go/tests -run TestToken/Test01CreateManagementAccessToken -v 
go test github.com/cohesity/app-sdk-go/tests -run TestMount -v 
go test github.com/cohesity/app-sdk-go/tests -run TestVolume -v 
``` 