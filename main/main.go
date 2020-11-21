package main

import (
	"KubeAPIExp/Utils"
	"log"
)

func main() {
	err := Utils.ExecFunction()
	if err != nil {
		log.Fatal(err.Error())
	}
}