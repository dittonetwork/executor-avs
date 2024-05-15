package main

import "github.com/dittonetwork/executor-avs/cmd/operator/app"

func main() {
	app.InitCommands()
	app.Execute()
}
