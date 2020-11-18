package main

import (
	"KubeAPIExp/Utils"
)

func main() {
	Utils.GetLogsFromPod("localhost:8001", "default", "nginx")
}