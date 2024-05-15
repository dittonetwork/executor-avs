package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Workflow struct {
	VaultAddress    common.Address
	WorkflowID      *big.Int
	CouldBeExecuted bool
}
