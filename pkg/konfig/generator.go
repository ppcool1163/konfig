package konfig

import (
	"fmt"
)

func Generate(endpoint string, id int, path string) {
	config := download(endpoint, id)
	parsedConfig, err := Parse(config)
	if err != nil {
		fmt.Println("error occurred while parsing")
		return
	}
	println(parsedConfig.Environment)
	if parsedConfig.Environment != "prod" {
		populate(parsedConfig.Variables, path)
		return
	}
	fmt.Println("This environment config is forbidden")
}
