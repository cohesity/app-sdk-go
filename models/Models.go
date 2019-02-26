package models



/*
 * Structure for the custom type AppSettings
 */
type AppSettings struct {
    User                User            `json:"user,omitempty" form:"user,omitempty"` //Specifies user information who launched the given app instance.
    AppInstanceSettings AppInstanceSettings `json:"appInstanceSettings,omitempty" form:"appInstanceSettings,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type MountOptions
 */
type MountOptions struct {
    DirName         string          `json:"dirName" form:"dirName"` //Directory where view/namespace is mounted.
    ViewName        string          `json:"viewName" form:"viewName"` //The name of the view that is to be mounted.
    MountOptions    *string         `json:"mountOptions,omitempty" form:"mountOptions,omitempty"` //Additional options for mount. All options from mount command are supported e.g. rw, ro.
    MountProtocol   MountProtocolEnum `json:"mountProtocol,omitempty" form:"mountProtocol,omitempty"` //Type of the mount.
    UserName        *string         `json:"userName,omitempty" form:"userName,omitempty"` //Username if the mount type is smb.
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //Password if the mount type is smb.
    NamespaceName   *string         `json:"namespaceName,omitempty" form:"namespaceName,omitempty"` //The namespace within the view that is to be mounted.
}

/*
 * Structure for the custom type AppInstanceSettings
 */
type AppInstanceSettings struct {
    QosTier                 QosTierEnum     `json:"qosTier,omitempty" form:"qosTier,omitempty"` //Specifies QoS Tier for an app instance. App instances are allocated resources such as memory, CPU and IO based on their QoS Tier.
    ReadViewPrivileges      ViewPrivileges  `json:"readViewPrivileges,omitempty" form:"readViewPrivileges,omitempty"` //Specifies privileges that are required for this app.
    ReadWriteViewPrivileges ViewPrivileges  `json:"readWriteViewPrivileges,omitempty" form:"readWriteViewPrivileges,omitempty"` //Specifies privileges that are required for this app.
}

/*
 * Structure for the custom type ViewPrivileges
 */
type ViewPrivileges struct {
    PrivilegesType  PrivilegesTypeEnum `json:"privilegesType,omitempty" form:"privilegesType,omitempty"` //Specifies if all, none or specific views are allowed to be accessed.
    ViewIds         *[]int64        `json:"viewIds,omitempty" form:"viewIds,omitempty"` //Specifies the ids of the views which are allowed to be accessed in case the privilege type is kSpecific.
}

/*
 * Structure for the custom type User
 */
type User struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Domain of the user who launched the app instance.
    UserName        *string         `json:"userName,omitempty" form:"userName,omitempty"` //Username of the user who launched the app instance.
}

/*
 * Structure for the custom type ManagementAccessToken
 */
type ManagementAccessToken struct {
    AccessToken     *string         `json:"accessToken,omitempty" form:"accessToken,omitempty"` //Token which can be used to call Iris Management apis.
    TokenType       *string         `json:"tokenType,omitempty" form:"tokenType,omitempty"` //Type of the token returned. E.g. Bearer.
}
