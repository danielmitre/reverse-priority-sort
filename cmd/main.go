package main

import (
	"fmt"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"os"
	"reverse-scheduler/pkg/reverse"
)

func main() {
	fmt.Println("Hello World")

	command := app.NewSchedulerCommand(
		app.WithPlugin(reverse.Name, reverse.New))

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
