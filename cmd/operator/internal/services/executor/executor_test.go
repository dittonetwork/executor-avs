package executor_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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
			// None of the input args are checked, under the hood testify/mock uses assert.ObjectsAreEqual(expected, Anything),
			//   we could just replace it with mock.Anything. Better to pass AnythingOfType("TypeName"), though.
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
			r := executor.NewExecutor(tt.fields.client(t), tt.fields.entryPoint(t))
			if err := r.Handle(ctx, blockHash); (err != nil) != tt.wantErr {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func workflowsAreTheSame(a, b []models.Workflow) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, elA := range a {
		elB := b[idx]

		if elA.VaultAddress != elB.VaultAddress || elA.WorkflowID != elB.WorkflowID {
			return false
		}
	}
	return true
}

func TestExecutor_Cache(t *testing.T) {
	type ethereumClientMockBuilder func(*testing.T) *mocks.EthereumClient
	type dittoEntryPointMockBuilder func(*testing.T) *mocks.DittoEntryPoint

	// Test values
	tx := types.NewTx(&types.DynamicFeeTx{})
	wf11 := models.Workflow{WorkflowID: big.NewInt(1), VaultAddress: common.HexToAddress("0x111")}
	wf12 := models.Workflow{WorkflowID: big.NewInt(2), VaultAddress: common.HexToAddress("0x111")}
	wf2 := models.Workflow{WorkflowID: big.NewInt(1), VaultAddress: common.HexToAddress("0x222")}
	wf3 := models.Workflow{WorkflowID: big.NewInt(1), VaultAddress: common.HexToAddress("0x333")}
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
					m.EXPECT().IsExecutor(mock.Anything).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, mock.Anything).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(mock.Anything).Return([]models.Workflow{wf11, wf12, wf2}, nil).Twice()
					m.EXPECT().GetAllActiveWorkflows(mock.Anything).Return([]models.Workflow{wf11, wf12, wf2, wf3}, nil)
					// Unfortunatelly, order does not matter. Declaring the same
					//   MatchedBy with Once() 2 times == declaring with Twice() 1 time
					m.EXPECT().RunMultipleWorkflows(mock.Anything, mock.MatchedBy(func(wfs []models.Workflow) bool {
						return workflowsAreTheSame(wfs, []models.Workflow{wf11, wf12, wf2})
					})).Return(tx, nil).Once()
					m.EXPECT().RunMultipleWorkflows(mock.Anything, mock.MatchedBy(
						func(wfs []models.Workflow) bool { return workflowsAreTheSame(wfs, []models.Workflow{wf3}) }),
					).Return(tx, nil).Once()
					m.EXPECT().RunMultipleWorkflows(mock.Anything, mock.MatchedBy(
						func(wfs []models.Workflow) bool { return workflowsAreTheSame(wfs, []models.Workflow{wf11, wf12, wf2}) }),
					).Return(tx, nil).Once()
					m.EXPECT().RunMultipleWorkflows(mock.Anything, mock.MatchedBy(
						func(wfs []models.Workflow) bool {
							return workflowsAreTheSame(wfs, []models.Workflow{wf11, wf12, wf2, wf3})
						}),
					).Return(tx, nil).Once()
					m.EXPECT().GetRunWorkflowTx(
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).Return(types.NewTransaction(
						uint64(1), common.HexToAddress("0x123"), big.NewInt(0), 0, big.NewInt(0), []byte("data"),
					), nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					for _, blockNum := range []int64{0, 1, 4, 5, 10} {
						m.EXPECT().BlockByHash(mock.Anything, mock.Anything).Return(
							types.NewBlock(&types.Header{Number: big.NewInt(blockNum)}, nil, nil, nil), nil).Once()
					}
					m.EXPECT().SimulateTransaction(mock.Anything, mock.Anything, mock.Anything).Return("0x1", nil)
					m.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil)
					return m
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := executor.NewExecutor(tt.fields.client(t), tt.fields.entryPoint(t))
			for i := 0; i < 5; i++ {
				err := r.Handle(context.Background(), common.HexToHash("0x1"))
				if err != nil {
					t.Error(err)
				}
			}
		})
	}
}
