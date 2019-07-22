// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for DiskTypeEnum enum
 */
type DiskTypeEnum int

/**
 * Value collection for DiskTypeEnum enum
 */
const (
    DiskType_KBOOT DiskTypeEnum = 1 + iota
    DiskType_KDATA
)

func (r DiskTypeEnum) MarshalJSON() ([]byte, error) {
    s := DiskTypeEnumToValue(r)
    return json.Marshal(s)
}

func (r *DiskTypeEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  DiskTypeEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts DiskTypeEnum to its string representation
 */
func DiskTypeEnumToValue(diskTypeEnum DiskTypeEnum) string {
    switch diskTypeEnum {
        case DiskType_KBOOT:
    		return "kBoot"
        case DiskType_KDATA:
    		return "kData"
        default:
        	return "kBoot"
    }
}

/**
 * Converts DiskTypeEnum Array to its string Array representation
*/
func DiskTypeEnumArrayToValue(diskTypeEnum []DiskTypeEnum) []string {
    convArray := make([]string,len( diskTypeEnum))
    for i:=0; i<len(diskTypeEnum);i++ {
        convArray[i] = DiskTypeEnumToValue(diskTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func DiskTypeEnumFromValue(value string) DiskTypeEnum {
    switch value {
        case "kBoot":
            return DiskType_KBOOT
        case "kData":
            return DiskType_KDATA
        default:
            return DiskType_KBOOT
    }
}
