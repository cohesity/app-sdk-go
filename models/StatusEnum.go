// Copyright 2019 Cohesity Inc.
package models

import(
    "encoding/json"
)

/**
 * Type definition for StatusEnum enum
 */
type StatusEnum int

/**
 * Value collection for StatusEnum enum
 */
const (
    Status_KINITIALIZING StatusEnum = 1 + iota
    Status_KAVAILABLE
    Status_KBOUND
    Status_KFAILED
)

func (r StatusEnum) MarshalJSON() ([]byte, error) {
    s := StatusEnumToValue(r)
    return json.Marshal(s)
}

func (r *StatusEnum) UnmarshalJSON(data []byte) error {
    var s string
    json.Unmarshal(data, &s)
    v :=  StatusEnumFromValue(s)
    *r = v
    return nil
 }


/**
 * Converts StatusEnum to its string representation
 */
func StatusEnumToValue(statusEnum StatusEnum) string {
    switch statusEnum {
        case Status_KINITIALIZING:
    		return "kInitializing"
        case Status_KAVAILABLE:
    		return "kAvailable"
        case Status_KBOUND:
    		return "kBound"
        case Status_KFAILED:
    		return "kFailed"
        default:
        	return "kInitializing"
    }
}

/**
 * Converts StatusEnum Array to its string Array representation
*/
func StatusEnumArrayToValue(statusEnum []StatusEnum) []string {
    convArray := make([]string,len( statusEnum))
    for i:=0; i<len(statusEnum);i++ {
        convArray[i] = StatusEnumToValue(statusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusEnumFromValue(value string) StatusEnum {
    switch value {
        case "kInitializing":
            return Status_KINITIALIZING
        case "kAvailable":
            return Status_KAVAILABLE
        case "kBound":
            return Status_KBOUND
        case "kFailed":
            return Status_KFAILED
        default:
            return Status_KINITIALIZING
    }
}
