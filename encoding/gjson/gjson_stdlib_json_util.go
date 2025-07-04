// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson

import (
	"bytes"
	"log"

	"github.com/ximplez-go/gf/errors/gerror"
	"github.com/ximplez-go/gf/internal/json"
	"github.com/ximplez-go/gf/util/gconv"
)

// Valid checks whether `data` is a valid JSON data type.
// The parameter `data` specifies the json format data, which can be either
// bytes or string type.
func Valid(data interface{}) bool {
	return json.Valid(gconv.Bytes(data))
}

// Marshal is alias of Encode in order to fit the habit of json.Marshal/Unmarshal functions.
func Marshal(v interface{}) (marshaledBytes []byte, err error) {
	return Encode(v)
}

// MarshalIndent is alias of json.MarshalIndent in order to fit the habit of json.MarshalIndent function.
func MarshalIndent(v interface{}, prefix, indent string) (marshaledBytes []byte, err error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal is alias of DecodeTo in order to fit the habit of json.Marshal/Unmarshal functions.
func Unmarshal(data []byte, v interface{}) (err error) {
	return DecodeTo(data, v)
}

// Encode encodes any golang variable `value` to JSON bytes.
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// MustEncode performs as Encode, but it panics if any error occurs.
func MustEncode(value interface{}) []byte {
	b, err := Encode(value)
	if err != nil {
		panic(err)
	}
	return b
}

// EncodeString encodes any golang variable `value` to JSON string.
func EncodeString(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	return string(b), err
}

// MustEncodeString encodes any golang variable `value` to JSON string.
// It panics if any error occurs.
func MustEncodeString(value interface{}) string {
	return string(MustEncode(value))
}

// Decode decodes json format `data` to golang variable.
// The parameter `data` can be either bytes or string type.
func Decode(data interface{}, options ...Options) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(gconv.Bytes(data), &value, options...); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// DecodeTo decodes json format `data` to specified golang variable `v`.
// The parameter `data` can be either bytes or string type.
// The parameter `v` should be a pointer type.
func DecodeTo(data interface{}, v interface{}, options ...Options) (err error) {
	decoder := json.NewDecoder(bytes.NewReader(gconv.Bytes(data)))
	if len(options) > 0 {
		// The StrNumber option is for certain situations, not for all.
		// For example, it causes converting issue for other data formats, for example: yaml.
		if options[0].StrNumber {
			decoder.UseNumber()
		}
	}
	if err = decoder.Decode(v); err != nil {
		err = gerror.Wrap(err, `json Decode failed`)
	}
	return
}

// DecodeToJson codes json format `data` to a Json object.
// The parameter `data` can be either bytes or string type.
func DecodeToJson(data interface{}, options ...Options) (*Json, error) {
	if v, err := Decode(gconv.Bytes(data), options...); err != nil {
		return nil, err
	} else {
		if len(options) > 0 {
			return New(v, options[0].Safe), nil
		}
		return New(v), nil
	}
}

func ToJsonSilent(data interface{}) string {
	return toJsonSilent(data, false)
}

func ToJsonSilentPretty(data interface{}) string {
	return toJsonSilent(data, true)
}

func toJsonSilent(data interface{}, pretty bool) string {
	if pretty {
		jsonByte, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Printf("[ToJson] error. %s", err.Error())
		}
		return string(jsonByte)
	} else {
		jsonByte, _ := json.Marshal(data)
		return string(jsonByte)
	}
}

func PhaseJsonSilent[T any](data []byte) *T {
	v := new(T)
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Printf("[PhaseJson] error. %s", err.Error())
	}
	return v
}

func PhaseJsonFromString[T any](data string) *T {
	return PhaseJsonSilent[T]([]byte(data))
}
