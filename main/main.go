package main

import "KubeAPIExp"

func main() {
	KubeAPIExp.GetLogsFromPod("localhost:8001", "default", "nginx")
}