package executor_test

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"testing"
	"time"

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
	nonce := uint64(1)
	blockNum := big.NewInt(1)
	ignoreEndBlocks := big.NewInt(3)
	// for check if executor will be valid in N blocks
	futureBlockNum := new(big.Int).Add(blockNum, ignoreEndBlocks)
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
					m.EXPECT().IsExecutor(mock.Anything).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, blockNum).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, futureBlockNum).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(mock.Anything).Return(activeWorkflows, nil)
					m.EXPECT().RunMultipleWorkflows(mock.Anything, activeWorkflows, mock.Anything).Return(tx, nil)
					m.EXPECT().GetRunWorkflowTx(mock.Anything, activeWorkflows[0].VaultAddress, big.NewInt(1)).
						Return(wfSimulatedTx1, nil)
					m.EXPECT().GetRunWorkflowTx(mock.Anything, activeWorkflows[1].VaultAddress, big.NewInt(2)).
						Return(wfSimulatedTx2, nil)
					m.EXPECT().GetSucceededWorkflows(mock.Anything).Return([]models.Workflow{activeWorkflows[0]}, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(mock.Anything, blockHash).Return(block, nil)
					m.EXPECT().SimulateTransaction(mock.Anything, mock.Anything, mock.Anything).Return(result, nil)
					m.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(types.NewReceipt([]byte{}, false, 1337), nil)
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
					m.EXPECT().IsExecutor(mock.Anything).Return(false, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(mock.Anything, blockHash).Return(block, nil)
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
					m.EXPECT().IsExecutor(mock.Anything).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, blockNum).Return(false, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(mock.Anything, blockHash).Return(block, nil)
					return m
				},
			},
			wantErr: false,
		},
		{
			name: "don't execute if end of turn (not valid executor in N blocks)",
			fields: fields{
				entryPoint: func(t *testing.T) *mocks.DittoEntryPoint {
					m := mocks.NewDittoEntryPoint(t)
					m.EXPECT().IsExecutor(mock.Anything).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, blockNum).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, futureBlockNum).Return(false, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(mock.Anything, blockHash).Return(block, nil)
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
					m.EXPECT().IsExecutor(mock.Anything).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, blockNum).Return(true, nil)
					m.EXPECT().IsValidExecutor(mock.Anything, futureBlockNum).Return(true, nil)
					m.EXPECT().GetAllActiveWorkflows(mock.Anything).Return(activeWorkflows, nil)
					m.EXPECT().RunMultipleWorkflows(mock.Anything, activeWorkflows[:1], mock.Anything).Return(tx, nil)
					m.EXPECT().GetRunWorkflowTx(mock.Anything, activeWorkflows[0].VaultAddress, big.NewInt(1)).
						Return(wfSimulatedTx1, nil)
					m.EXPECT().GetRunWorkflowTx(mock.Anything, activeWorkflows[1].VaultAddress, big.NewInt(2)).
						Return(wfSimulatedTx2, nil)
					m.EXPECT().GetSucceededWorkflows(mock.Anything).Return([]models.Workflow{activeWorkflows[0]}, nil)
					return m
				},
				client: func(t *testing.T) *mocks.EthereumClient {
					m := mocks.NewEthereumClient(t)
					m.EXPECT().BlockByHash(mock.Anything, blockHash).Return(block, nil)
					m.EXPECT().SimulateTransaction(mock.Anything, wfSimulatedTx1, blockNum).Return(result, nil)
					m.EXPECT().SimulateTransaction(mock.Anything, wfSimulatedTx2, blockNum).Return("0x0", nil)
					m.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(types.NewReceipt([]byte{}, false, 1337), nil)
					return m
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := executor.NewExecutor(tt.fields.client(t), tt.fields.entryPoint(t), ignoreEndBlocks, 5*time.Second)
			if err := r.Handle(context.Background(), blockHash); (err != nil) != tt.wantErr {
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

// TestConditionalAcquireWorkflowsLock tests adding new workflows to the cache.
func TestCacheConditionalAcquireWorkflowsLock(t *testing.T) {
	cache := executor.NewWorkflowsCache()
	workflows := []models.Workflow{
		{VaultAddress: common.HexToAddress("0x1"), WorkflowID: big.NewInt(1)},
		{VaultAddress: common.HexToAddress("0x1"), WorkflowID: big.NewInt(2)},
		{VaultAddress: common.HexToAddress("0x2"), WorkflowID: big.NewInt(1)},
	}

	inserted := cache.ConditionalAcquireWorkflowsLock(workflows)
	assert.Len(t, inserted, 3)

	insertedAgain := cache.ConditionalAcquireWorkflowsLock(workflows)
	assert.Empty(t, insertedAgain)
}

// TestReleaseWorkflowsLock tests the removal of workflows from the cache.
func TestCacheReleaseWorkflowsLock(t *testing.T) {
	cache := executor.NewWorkflowsCache()
	var wf1 = models.Workflow{VaultAddress: common.HexToAddress("0x1"), WorkflowID: big.NewInt(1)}
	var wf2 = models.Workflow{VaultAddress: common.HexToAddress("0x1"), WorkflowID: big.NewInt(2)}

	acquired := cache.ConditionalAcquireWorkflowsLock([]models.Workflow{wf1, wf2})
	assert.Len(t, acquired, 2)

	cache.ReleaseWorkflowsLock([]models.Workflow{wf1})

	acquired = cache.ConditionalAcquireWorkflowsLock([]models.Workflow{wf1, wf2})
	assert.Len(t, acquired, 1)

	assert.Equal(t, []models.Workflow{wf1}, acquired)
}

func getRandomSubset[T any](items []T, subsetSize int) []T {
	if len(items) == 0 || subsetSize <= 0 {
		return nil
	}
	if subsetSize > len(items) {
		subsetSize = len(items)
	}

	subset := make([]T, subsetSize)
	pickedIndices := make(map[int]bool)

	for i := range subsetSize {
		for {
			idx := rand.Intn(len(items))
			if !pickedIndices[idx] {
				subset[i] = items[idx]
				pickedIndices[idx] = true
				break
			}
		}
	}
	return subset
}

func TestConcurrentAcquisitionWithSets(t *testing.T) {
	cache := executor.NewWorkflowsCache()
	var wg sync.WaitGroup

	totalWorkflows := 200
	workflows := make([]models.Workflow, totalWorkflows)
	for i := range totalWorkflows {
		workflows[i] = models.Workflow{
			VaultAddress: common.HexToAddress("0x" + fmt.Sprintf("%02x", i)),
			WorkflowID:   big.NewInt(int64(i)),
		}
	}

	acquired := make(map[common.Address]map[string]int)
	mu := sync.Mutex{}

	workerCount := 20
	wg.Add(workerCount)
	for range workerCount {
		go func() {
			defer wg.Done()
			for j := range 10 {
				subset := getRandomSubset(workflows, j)
				acquiredWorkflows := cache.ConditionalAcquireWorkflowsLock(subset)
				mu.Lock()
				for _, aWf := range acquiredWorkflows {
					if _, exists := acquired[aWf.VaultAddress]; !exists {
						acquired[aWf.VaultAddress] = make(map[string]int)
					}
					acquired[aWf.VaultAddress][aWf.WorkflowID.String()]++
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	for vaultAddr, wfs := range acquired {
		for wfID, timesAquired := range wfs {
			assert.Equal(t, 1, timesAquired, "Workflow %d of vault %s was acquired more than once", wfID, vaultAddr)
		}
	}
}
