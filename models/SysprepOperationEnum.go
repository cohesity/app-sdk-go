// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for SysprepOperationEnum enum
 */
type SysprepOperationEnum int

/**
 * Value collection for SysprepOperationEnum enum
 */
const (
    SysprepOperation_KADDUSERSSHKEY SysprepOperationEnum = 1 + iota
    SysprepOperation_KADDSWAPMOUNTPOINT
    SysprepOperation_KADDVIRTIODRIVER
)

func (r SysprepOperationEnum) MarshalJSON() ([]byte, error) {
    s := SysprepOperationEnumToValue(r)
    return json.Marshal(s)
}

func (r *SysprepOperationEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  SysprepOperationEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts SysprepOperationEnum to its string representation
 */
func SysprepOperationEnumToValue(sysprepOperationEnum SysprepOperationEnum) string {
    switch sysprepOperationEnum {
        case SysprepOperation_KADDUSERSSHKEY:
    		return "kAddUserSshKey"
        case SysprepOperation_KADDSWAPMOUNTPOINT:
    		return "kAddSwapMountPoint"
        case SysprepOperation_KADDVIRTIODRIVER:
    		return "kAddVirtioDriver"
        default:
        	return "kAddUserSshKey"
    }
}

/**
 * Converts SysprepOperationEnum Array to its string Array representation
*/
func SysprepOperationEnumArrayToValue(sysprepOperationEnum []SysprepOperationEnum) []string {
    convArray := make([]string,len( sysprepOperationEnum))
    for i:=0; i<len(sysprepOperationEnum);i++ {
        convArray[i] = SysprepOperationEnumToValue(sysprepOperationEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func SysprepOperationEnumFromValue(value string) SysprepOperationEnum {
    switch value {
        case "kAddUserSshKey":
            return SysprepOperation_KADDUSERSSHKEY
        case "kAddSwapMountPoint":
            return SysprepOperation_KADDSWAPMOUNTPOINT
        case "kAddVirtioDriver":
            return SysprepOperation_KADDVIRTIODRIVER
        default:
            return SysprepOperation_KADDUSERSSHKEY
    }
}
