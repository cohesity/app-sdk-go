// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for VolumeTypeEnum enum
 */
type VolumeTypeEnum int

/**
 * Value collection for VolumeTypeEnum enum
 */
const (
    VolumeType_KNFS VolumeTypeEnum = 1 + iota
    VolumeType_KISCSI
)

func (r VolumeTypeEnum) MarshalJSON() ([]byte, error) {
    s := VolumeTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *VolumeTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  VolumeTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts VolumeTypeEnum to its string representation
 */
func VolumeTypeEnumToValue(volumeTypeEnum VolumeTypeEnum) string {
    switch volumeTypeEnum {
        case VolumeType_KNFS:
    		return "kNfs"
        case VolumeType_KISCSI:
    		return "kIscsi"
        default:
        	return "kNfs"
    }
}

/**
 * Converts VolumeTypeEnum Array to its string Array representation
*/
func VolumeTypeEnumArrayToValue(volumeTypeEnum []VolumeTypeEnum) []string {
    convArray := make([]string,len( volumeTypeEnum))
    for i:=0; i<len(volumeTypeEnum);i++ {
        convArray[i] = VolumeTypeEnumToValue(volumeTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func VolumeTypeEnumFromValue(value string) VolumeTypeEnum {
    switch value {
        case "kNfs":
            return VolumeType_KNFS
        case "kIscsi":
            return VolumeType_KISCSI
        default:
            return VolumeType_KNFS
    }
}
