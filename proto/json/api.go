package main

import (
	"fmt"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func GetApis() []string {
	apis := make([]string, 0)
	result, err := gojsonq.New().File("../gatewayApis.json").FindR("apis")
	if err != nil {
		fmt.Println("failed to get apis fields")
	}
	fmt.Println(result)
	err = result.As(&apis)
	if err != nil {
		fmt.Println("failed to get apis")
	}

	return apis
}

func main() {
	apis := GetApis()
	fmt.Println(apis)
}
