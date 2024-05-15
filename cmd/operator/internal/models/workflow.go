package models

import "github.com/ethereum/go-ethereum/core/types"

type Workflow struct {
	VaultAddress types.Account
	WorkflowID   uint64
}
