// Copyright 2019 Cohesity Inc.
package models

type DeleteUnmountParams struct {
	DirName string
}

type CreateMountParams struct {
	MountOptions *MountOptions
}

type CreateVolumeParams struct {
	VolumeName string
	VolumeSpec *VolumeSpec
}

type GetVolumeParams struct {
	VolumeName string
}

type DeleteVolumeParams struct {
	VolumeName string
}

type GetSourceVolumeInfoParams struct {
	SourceId int64
}

/*
 Functions taking no params, Need to see if this should be done for empty params
type GetAppSettingsParams struct {
}

type CreateManagementAccessTokenParams struct {
}

*/
