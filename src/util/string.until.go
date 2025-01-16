package util

import (
	"encoding/json"
	"fmt"
)

func StringIsNil(object string) string  {
	if object == "" {
		fmt.Printf("object is nil %s\n",object)
		return ""
	}
	return object
}
func ConvectData( object interface{})interface{} {



	if list ,ok := object.([]interface{}); ok  {
		return list
	}
	if mapData,err := object.(map[string]interface{}); err {
		return mapData
	}
	return nil

}

func ConvectDataRedis(data string) interface{}  {
	var object interface{}
	err := json.Unmarshal([]byte(data),&object)
	if err != nil {
		fmt.Printf("object is nil %s\n",object)
	}
	return object
}