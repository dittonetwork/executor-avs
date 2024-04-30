package main

import "github.com/dittonetwork/executor-avs/cmd/operator/app"

func main() {
	wg := app.Run()

	wg.Wait()
}
