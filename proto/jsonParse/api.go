package jsonParse

import (
	"fmt"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func getApis() []string {
	apis := make([]string, 0)
	result, err := gojsonq.New().File("../gatewayApis.json").FindR("apis")
	if err != nil {
		fmt.Println("failed to get apis fields")
	}
	err = result.As(&apis)
	if err != nil {
		fmt.Println("failed to get apis")
	}

	return apis
}
