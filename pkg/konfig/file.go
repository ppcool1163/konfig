package konfig

import (
	"fmt"
	"konfig/model"
	"log"
	"os"
)

func populate(variables []model.Variable, path string) {
	if len(variables) == 0 {
		fmt.Println("no config variable found")
		return
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, v := range variables {
		_, err = f.WriteString(fmt.Sprintf("%s=%s\n", v.Name, v.Value))
	}
	fmt.Println(f.Name())
}
