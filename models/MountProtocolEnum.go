// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for MountProtocolEnum enum
 */
type MountProtocolEnum int

/**
 * Value collection for MountProtocolEnum enum
 */
const (
    MountProtocol_KNFS MountProtocolEnum = 1 + iota
    MountProtocol_KSMB
    MountProtocol_KVOLUME
)

func (r MountProtocolEnum) MarshalJSON() ([]byte, error) {
    s := MountProtocolEnumToValue(r)
    return json.Marshal(s)
}

func (r *MountProtocolEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  MountProtocolEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts MountProtocolEnum to its string representation
 */
func MountProtocolEnumToValue(mountProtocolEnum MountProtocolEnum) string {
    switch mountProtocolEnum {
        case MountProtocol_KNFS:
    		return "kNfs"
        case MountProtocol_KSMB:
    		return "kSmb"
        case MountProtocol_KVOLUME:
    		return "kVolume"
        default:
        	return "kNfs"
    }
}

/**
 * Converts MountProtocolEnum Array to its string Array representation
*/
func MountProtocolEnumArrayToValue(mountProtocolEnum []MountProtocolEnum) []string {
    convArray := make([]string,len( mountProtocolEnum))
    for i:=0; i<len(mountProtocolEnum);i++ {
        convArray[i] = MountProtocolEnumToValue(mountProtocolEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func MountProtocolEnumFromValue(value string) MountProtocolEnum {
    switch value {
        case "kNfs":
            return MountProtocol_KNFS
        case "kSmb":
            return MountProtocol_KSMB
        case "kVolume":
            return MountProtocol_KVOLUME
        default:
            return MountProtocol_KNFS
    }
}
