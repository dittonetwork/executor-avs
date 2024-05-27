// package pusher

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/trustwallet/backend-blacklist/pkg/blacklist"
// 	"github.com/trustwallet/backend-txhub-events/config"
// 	"github.com/trustwallet/backend-txhub-events/internal/app/pusher/mocks"
// 	"github.com/trustwallet/backend-txhub-events/internal/models"
// 	"github.com/trustwallet/backend-txhub-events/test/pkg"
// 	txhub "github.com/trustwallet/backend-txhub/pkg/models"
// 	"github.com/trustwallet/go-primitives/coin"
// )

// func TestRunMain(t *testing.T) {
// 	type customersUCMockBuilder func(*testing.T) *mocks.CustomersUsecase
// 	type subscriptionsUCMockBuilder func(*testing.T) *mocks.SubscriptionsUsecase
// 	type eventsCacheMockBuilder func(*testing.T) *mocks.EventsCache
// 	type blacklistServiceMockBuilder func(*testing.T) *mocks.BlacklistClient

// 	chain := coin.Ethereum()

// 	webhookServerPort := "5555"
// 	webhookURLPath := "/test/webhook"
// 	wst := pkg.NewWebhookServerTest(webhookServerPort, webhookURLPath)

// 	blockCreatedAt := pkg.ConvertStrToTime(time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
// 	inputTxs := []txhub.Transaction{
// 		{
// 			Hash:        "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 			Chain:       60,
// 			BlockNumber: 17369185,
// 			BlockTs:     blockCreatedAt,
// 			CreatedTs:   blockCreatedAt,
// 			Status:      "completed",
// 			Fees: []txhub.AssetAmount{
// 				{
// 					Asset: "c60",
// 					Value: "1869000000000000",
// 				},
// 			},
// 			Error: "",
// 			Nonce: 22119,
// 			Memo:  "",
// 			Events: []txhub.Event{
// 				{
// 					Type: txhub.EventTransfer,
// 					Data: txhub.EventDataTransfer{
// 						From:  "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 						To:    "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 						Asset: "c60",
// 						Value: "133000000000000000",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	testCases := map[string]struct {
// 		customersUCMock      customersUCMockBuilder
// 		subscriptionsUCMock  subscriptionsUCMockBuilder
// 		eventsCacheMock      eventsCacheMockBuilder
// 		blacklistServiceMock blacklistServiceMockBuilder

// 		txs        []txhub.Transaction
// 		discardTTL time.Duration

// 		expectedTxEvents []models.TransactionEvent
// 		expectedErrMsg   string
// 	}{
// 		"success": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, nil)
// 				m.EXPECT().MarkAsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f").
// 					Return(nil)
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:        inputTxs,
// 			discardTTL: 2 * time.Second,
// 			expectedTxEvents: []models.TransactionEvent{
// 				{
// 					Hash:      "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 					Type:      txhub.EventTransfer,
// 					Direction: txhub.DirectionIncoming,
// 					Chain:     60,
// 					Status:    "completed",
// 					From:      "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					To:        "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					Values: []txhub.AssetAmount{
// 						{Asset: "c60", Value: "133000000000000000"},
// 					},
// 					BlockCreatedAt: blockCreatedAt,
// 				},
// 			},
// 			expectedErrMsg: "",
// 		},
// 		"discarding old events": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				time.Sleep(2 * time.Microsecond)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache { return nil },
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:              inputTxs,
// 			discardTTL:       1 * time.Microsecond,
// 			expectedTxEvents: []models.TransactionEvent{},
// 			expectedErrMsg:   "",
// 		},
// 		"event was pushed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(true, nil)
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:              inputTxs,
// 			discardTTL:       2 * time.Second,
// 			expectedTxEvents: []models.TransactionEvent{},
// 			expectedErrMsg:   "",
// 		},
// 		"same tx: 1 customer received event, 1 customer not": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  "", // no need to have a webhook because this customer already got an event
// 							Status:      models.CustomerStatusActive,
// 						},
// 						{
// 							CustomerID:  "customer_binance",
// 							CompanyName: "Binance DeFi",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet", "customer_binance"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet", "customer_binance"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(true, nil)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_binance",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, nil)
// 				m.EXPECT().MarkAsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_binance",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f").
// 					Return(nil)
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:        inputTxs,
// 			discardTTL: 2 * time.Second,
// 			expectedTxEvents: []models.TransactionEvent{
// 				{
// 					Hash:      "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 					Type:      txhub.EventTransfer,
// 					Direction: txhub.DirectionIncoming,
// 					Chain:     60,
// 					Status:    "completed",
// 					From:      "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					To:        "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					Values: []txhub.AssetAmount{
// 						{Asset: "c60", Value: "133000000000000000"},
// 					},
// 					BlockCreatedAt: blockCreatedAt,
// 				},
// 			},
// 			expectedErrMsg: "",
// 		},
// 		"customer does not exist": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet_new_id",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet_old_id"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet_old_id"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache { return nil },
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:              inputTxs,
// 			discardTTL:       2 * time.Second,
// 			expectedTxEvents: nil,
// 			expectedErrMsg:   "",
// 		},
// 		"mark as pushed: failed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, nil)
// 				m.EXPECT().MarkAsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f").
// 					Return(errors.New("some error"))
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:        inputTxs,
// 			discardTTL: 2 * time.Second,
// 			expectedTxEvents: []models.TransactionEvent{
// 				{
// 					Hash:      "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 					Type:      txhub.EventTransfer,
// 					Direction: txhub.DirectionIncoming,
// 					Chain:     60,
// 					Status:    "completed",
// 					From:      "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					To:        "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					Values: []txhub.AssetAmount{
// 						{Asset: "c60", Value: "133000000000000000"},
// 					},
// 					BlockCreatedAt: blockCreatedAt,
// 				},
// 			},
// 			expectedErrMsg: "",
// 		},
// 		"is pushed: failed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, errors.New("some error"))
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:              inputTxs,
// 			discardTTL:       2 * time.Second,
// 			expectedTxEvents: nil,
// 			expectedErrMsg:   "checking if event is pushed: some error",
// 		},
// 		"pushing to webhook: failed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  "invalid webhook",
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 						[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, nil)
// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:              inputTxs,
// 			discardTTL:       2 * time.Second,
// 			expectedTxEvents: nil,
// 			expectedErrMsg:   "pushing to webhook",
// 		},
// 		"get customer: failed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return(nil, errors.New("some error"))
// 				return m
// 			},
// 			subscriptionsUCMock:  func(*testing.T) *mocks.SubscriptionsUsecase { return nil },
// 			eventsCacheMock:      func(*testing.T) *mocks.EventsCache { return nil },
// 			blacklistServiceMock: func(*testing.T) *mocks.BlacklistClient { return nil },
// 			txs:                  inputTxs,
// 			expectedTxEvents:     nil,
// 			expectedErrMsg:       "get customers: some error",
// 		},
// 		"get customers ids by addresses: failed": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(nil, errors.New("some error")).Maybe()
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(nil, errors.New("some error")).Maybe()

// 				return m
// 			},
// 			eventsCacheMock:      func(*testing.T) *mocks.EventsCache { return nil },
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient { return nil },
// 			txs:                  inputTxs,
// 			expectedTxEvents:     nil,
// 			expectedErrMsg:       "getting customers by address: some error",
// 		},
// 		"blacklist returned blocked address which was not marked as blocked": {
// 			customersUCMock: func(*testing.T) *mocks.CustomersUsecase {
// 				m := mocks.NewCustomersUsecase(t)
// 				m.EXPECT().
// 					GetCustomers(context.Background()).
// 					Return([]models.Customer{
// 						{
// 							CustomerID:  "customer_trustwallet",
// 							CompanyName: "Trust Wallet",
// 							WebhookURL:  fmt.Sprintf("http://localhost:%s%s", webhookServerPort, webhookURLPath),
// 							Status:      models.CustomerStatusActive,
// 						},
// 					}, nil)
// 				return m
// 			},
// 			subscriptionsUCMock: func(*testing.T) *mocks.SubscriptionsUsecase {
// 				m := mocks.NewSubscriptionsUsecase(t)
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c", "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 						"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().GetCustomersIDsByAddresses(context.Background(), chain.Handle,
// 					[]string{"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A", "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(map[string][]string{
// 						"0x1AB426c5725b7B69de74C2768ee16D3304E8d06A": {"customer_trustwallet"},
// 						"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c": {"customer_trustwallet"},
// 					}, nil).Maybe()
// 				m.EXPECT().
// 					BlockSubscriptionsByAddress(context.Background(), chain.Handle, "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A").
// 					Return(nil)
// 				return m
// 			},
// 			eventsCacheMock: func(*testing.T) *mocks.EventsCache {
// 				m := mocks.NewEventsCache(t)
// 				m.EXPECT().IsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 				).Return(false, nil)
// 				m.EXPECT().MarkAsPushed(
// 					context.Background(),
// 					chain.Handle,
// 					"0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					"customer_trustwallet",
// 					"0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f").
// 					Return(nil)

// 				return m
// 			},
// 			blacklistServiceMock: func(t *testing.T) *mocks.BlacklistClient {
// 				m := mocks.NewBlacklistClient(t)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A"}).
// 					Return(true, nil)
// 				m.EXPECT().Blocked(context.Background(), blacklist.Address{Chain: chain.ID, Address: "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c"}).
// 					Return(false, nil)

// 				return m
// 			},
// 			txs:        inputTxs,
// 			discardTTL: 2 * time.Second,
// 			expectedTxEvents: []models.TransactionEvent{
// 				{
// 					Hash:      "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 					Type:      txhub.EventTransfer,
// 					Direction: txhub.DirectionOutgoing,
// 					Chain:     60,
// 					Status:    "completed",
// 					From:      "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					To:        "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					Values: []txhub.AssetAmount{
// 						{Asset: "c60", Value: "133000000000000000"},
// 					},
// 					BlockCreatedAt: blockCreatedAt,
// 				},
// 			},
// 			expectedErrMsg: "",
// 		},
// 	}

// 	wst.RunHTTPServerWithWebhook()

// 	var i int
// 	for name, tc := range testCases {
// 		webhookDataMap := wst.GetWebhookDataMap()

// 		i++
// 		t.Run(name, func(t *testing.T) {
// 			processor := NewProcessor(
// 				config.Pusher{
// 					Enabled:                  true,
// 					DiscardEventsIfLaterThan: tc.discardTTL,
// 					SendToWebhookTimeout:     1 * time.Second,
// 				},
// 				chain,
// 				tc.customersUCMock(t),
// 				tc.subscriptionsUCMock(t),
// 				tc.eventsCacheMock(t),
// 				tc.blacklistServiceMock(t),
// 				WithProcessorMetrics(fmt.Sprintf("test_%d", i), prometheus.DefaultRegisterer),
// 			)

// 			ctx := context.Background()
// 			err := processor.pushEvents(ctx, tc.txs)

// 			if tc.expectedErrMsg == "" {
// 				assert.NoError(t, err)
// 			} else {
// 				assert.ErrorContains(t, err, tc.expectedErrMsg)
// 			}

// 			assert.ElementsMatch(t, webhookDataMap[chain.Handle], tc.expectedTxEvents)

// 			wst.CleanWebhookDataMap()
// 		})
// 	}
// }

// func TestPushToWebhook(t *testing.T) {
// 	testCases := map[string]struct {
// 		webhookURL     string
// 		txEvents       []models.TransactionEvent
// 		expectedErrMsg string
// 	}{
// 		"success": {
// 			webhookURL: "http://localhost:5556/test/webhook2",
// 			txEvents: []models.TransactionEvent{
// 				{
// 					Hash:      "0xef689ec8a3dafb0c63d408b7b327bd46b3eaed063ba7fe932d1f2cf65f7efa7f",
// 					Type:      txhub.EventTransfer,
// 					Direction: txhub.DirectionIncoming,
// 					Chain:     60,
// 					Status:    "completed",
// 					From:      "0x28f2E7888bbff2f49e9842E1c6b37fEdAFf7364c",
// 					To:        "0x1AB426c5725b7B69de74C2768ee16D3304E8d06A",
// 					Values: []txhub.AssetAmount{
// 						{Asset: "c60", Value: "133000000000000000"},
// 					},
// 					BlockCreatedAt: pkg.ConvertStrToTime(time.Now().UTC().Format("2006-01-02T15:04:05.000Z")),
// 				},
// 			},
// 			expectedErrMsg: "",
// 		},
// 		"invalid status code": {
// 			webhookURL:     "http://localhost:5556/invalid_webhook",
// 			txEvents:       []models.TransactionEvent{},
// 			expectedErrMsg: "unsuccessful status code",
// 		},
// 	}

// 	webhookServerPort := "5556"
// 	webhookURLPath := "/test/webhook2"
// 	wst := pkg.NewWebhookServerTest(webhookServerPort, webhookURLPath)
// 	wst.RunHTTPServerWithWebhook()

// 	for name, tc := range testCases {
// 		webhookDataMap := wst.GetWebhookDataMap()

// 		t.Run(name, func(t *testing.T) {
// 			ctx := context.Background()
// 			err := pushToWebhook(ctx, tc.webhookURL, tc.txEvents, 1*time.Second)

// 			if tc.expectedErrMsg == "" {
// 				assert.NoError(t, err)
// 				assert.ElementsMatch(t,
// 					webhookDataMap[coin.Coins[uint(tc.txEvents[0].Chain)].Handle], tc.txEvents)
// 			} else {
// 				assert.ErrorContains(t, err, tc.expectedErrMsg)
// 			}

// 			wst.CleanWebhookDataMap()
// 		})
// 	}
// }
