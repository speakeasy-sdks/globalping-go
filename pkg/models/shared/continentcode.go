// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ContinentCode string

const (
	ContinentCodeAf ContinentCode = "AF"
	ContinentCodeAn ContinentCode = "AN"
	ContinentCodeAs ContinentCode = "AS"
	ContinentCodeEu ContinentCode = "EU"
	ContinentCodeNa ContinentCode = "NA"
	ContinentCodeOc ContinentCode = "OC"
	ContinentCodeSa ContinentCode = "SA"
)

func (e ContinentCode) ToPointer() *ContinentCode {
	return &e
}

func (e *ContinentCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AF":
		fallthrough
	case "AN":
		fallthrough
	case "AS":
		fallthrough
	case "EU":
		fallthrough
	case "NA":
		fallthrough
	case "OC":
		fallthrough
	case "SA":
		*e = ContinentCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContinentCode: %v", v)
	}
}
