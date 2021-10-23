package main

import (
	"fmt"
	"os"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func GetApis(version string) []string {
	apis := make([]string, 0)
	goPath := os.Getenv("GOPATH")
	fmt.Println(goPath)
	result, err := gojsonq.New().File(goPath + "/pkg/mod/github.com/!zgxh/grpc-gen@" + version + "/proto/gatewayApis.json").FindR("apis")
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
	GetApis("v1.0.4")
}
