// Package contracts does not need to be used directly; this file is just for `go:generate` commands.
package contracts

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=./abi/DittoEntryPoint.json --pkg=dittoentrypoint --out=./gen/dittoentrypoint/dittoentrypoint.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=./abi/IServiceManager.json --pkg=iservicemanager --out=./gen/iservicemanager/iservicemanager.go
