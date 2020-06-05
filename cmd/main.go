package main

import (
	"fmt"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"os"
<<<<<<< HEAD
	reverse "reverse-scheduler/pkg/reverse"
=======
	"reverse-scheduler/pkg/reverse"
>>>>>>> 882d9b4... initial commit
)

func main() {
	fmt.Println("Hello World")

	command := app.NewSchedulerCommand(
		app.WithPlugin(reverse.Name, reverse.New))

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
