package main

import (
	"KubeAPIExp/Utils"
)

func main() {
	Utils.GetLogs("default", "nginx-d46f5678b-7vd9j", []string{"nginx"})
}