package executor_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/models"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor/mocks"
)

func TestExecutor_Handle(t *testing.T) {
	type ethereumClientMockBuilder func(*testing.T) *mocks.EthereumClient
	type dittoEntryPointMockBuilder func(*testing.T) *mocks.DittoEntryPoint

	// Test values
	ctx := context.Background()
	nonce := uint64(1)
	blockNum := big.NewInt(1)
	block := types.NewBlock(&types.Header{Number: blockNum}, nil, nil, nil)
	blockHash := common.HexToHash("0x1")
	contractAddr := common.HexToAddress("0x123")
	data := []byte("data")
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(0), 0, big.NewInt(0), data)
	result := "0x1"
	activeWorkflows := []models.Workflow{
		{WorkflowID: big.NewInt(1), VaultAddress: common.HexToAddress("0x111")},
		{WorkflowID: big.NewInt(2), VaultAddress: common.HexToAddress("0x222")},
	}
	wfSimulatedTx1 := types.NewTransaction(nonce, activeWorkflows[0].VaultAddress, big.NewInt(0), 0, big.NewInt(0), data)
	wfSimulatedTx2 := types.NewTransaction(nonce, activeWorkflows[1].VaultAddress, big.NewInt(0), 0, big.NewInt(0), data)

	type fields struct {
		client     ethereumClientMockBuilder
		entryPoint dittoEntryPointMockBuilder
	}

	tests := []struct {
		name        string
		fields      fields
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Success flow",
			fields: fields{
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(ctx).Return(true, nil)
					m.EXPECT().IsValidExecutor(ctx, blockNum).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(ctx).Return(activeWorkflows, nil)
					m.EXPECT().RunMultipleWorkflows(ctx, activeWorkflows).Return(tx, nil)
					m.EXPECT().GetRunWorkflowTx(ctx, activeWorkflows[0].VaultAddress, big.NewInt(1)).
						Return(wfSimulatedTx1, nil)
					m.EXPECT().GetRunWorkflowTx(ctx, activeWorkflows[1].VaultAddress, big.NewInt(2)).
						Return(wfSimulatedTx2, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(ctx, blockHash).Return(block, nil)
					m.EXPECT().SimulateTransaction(ctx, wfSimulatedTx1, blockNum).Return(result, nil)
					m.EXPECT().SimulateTransaction(ctx, wfSimulatedTx2, blockNum).Return(result, nil)
					m.EXPECT().SendTransaction(ctx, tx).Return(nil)
					return m
				},
			},
			wantErr: false,
		},
		{
			name: "Error on IsExecutor=false",
			fields: fields{
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(ctx).Return(false, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(ctx, blockHash).Return(block, nil)
					return m
				},
			},
			wantErr:     true,
			expectedErr: executor.ErrUnregisteredExecutor,
		},
		{
			name: "Return on IsValidExecutor=false",
			fields: fields{
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(ctx).Return(true, nil)
					m.EXPECT().IsValidExecutor(ctx, blockNum).Return(false, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(ctx, blockHash).Return(block, nil)
					return m
				},
			},
			wantErr: false,
		},
		{
			name: "only one (out of 2) workflow could be simulated",
			fields: fields{
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(ctx).Return(true, nil)
					m.EXPECT().IsValidExecutor(ctx, blockNum).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(ctx).Return(activeWorkflows, nil)
					m.EXPECT().RunMultipleWorkflows(ctx, activeWorkflows[:1]).Return(tx, nil)
					m.EXPECT().GetRunWorkflowTx(ctx, activeWorkflows[0].VaultAddress, big.NewInt(1)).
						Return(wfSimulatedTx1, nil)
					m.EXPECT().GetRunWorkflowTx(ctx, activeWorkflows[1].VaultAddress, big.NewInt(2)).
						Return(wfSimulatedTx2, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(ctx, blockHash).Return(block, nil)
					m.EXPECT().SimulateTransaction(ctx, wfSimulatedTx1, blockNum).Return(result, nil)
					m.EXPECT().SimulateTransaction(ctx, wfSimulatedTx2, blockNum).Return("0x0", nil)
					m.EXPECT().SendTransaction(ctx, tx).Return(nil)
					return m
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &executor.Executor{
				Client:     tt.fields.client(t),
				EntryPoint: tt.fields.entryPoint(t),
			}
			if err := r.Handle(ctx, blockHash); (err != nil) != tt.wantErr {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
