package logger

import (
	"encoding/json"
	"fmt"
	"log"
)

func Show(subject interface{}) {
	b, err := json.MarshalIndent(subject, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
	log.Println("")
}
