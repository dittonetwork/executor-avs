package executor

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor/mocks"
)

func TestExecutor_Handle(t *testing.T) {
	type ethereumClientMockBuilder func(*testing.T) *mocks.EthereumClient
	type dittoEntryPointMockBuilder func(*testing.T) *mocks.DittoEntryPoint

	nonce := uint64(1)
	blockNum := big.NewInt(1)
	contractAddr := common.HexToAddress("0x123")
	data := []byte("data")
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), 0, big.NewInt(0), data)
	result := new(interface{})

	type fields struct {
		client     ethereumClientMockBuilder
		entryPoint dittoEntryPointMockBuilder
	}
	type args struct {
		ctx      context.Context
		blockNum *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success flow",
			fields: fields{
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().SimulateTransaction(context.Background(), tx, blockNum, result).Return(nil)
					m.EXPECT().SendTransaction(context.Background(), tx).Return(nil)
					return m
				},
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(context.Background()).Return(true, nil)
					m.EXPECT().IsValidExecutor(context.Background(), blockNum).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(context.Background()).Return(nil, nil)
					m.EXPECT().RunMultipleWorkflows(context.Background(), nil).Return(nil, nil)
					m.EXPECT().GetRunWorkflowTx(context.Background(), contractAddr, big.NewInt(0)).Return(tx, nil)
					return m
				},
			},
			args: args{
				ctx:      context.Background(),
				blockNum: blockNum,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Executor{
				client:     tt.fields.client(t),
				entryPoint: tt.fields.entryPoint(t),
			}
			if err := r.Handle(tt.args.ctx, tt.args.blockNum); (err != nil) != tt.wantErr {
				t.Errorf("Executor.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
