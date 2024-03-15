package check

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/lavanet/lava/app"
	"github.com/lavanet/lava/utils"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lavaprovider-monitor/pkg/alert"
	"lavaprovider-monitor/pkg/log"

	"time"
)

func exponentialBackoff(retry int) {
	time.Sleep(time.Duration((1<<retry)*1000) * time.Millisecond)
}

func LavaProviderChecker(address string, lavaGrpc string, LavaChainid string, maxRetries int) (types.QueryAccountInfoResponse, error) {
	grpcConn, err := grpc.Dial(lavaGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)))
	if err != nil {
		// Just log because grpc redials
		log.Log.Error("Failed to dial grpc", zap.Error(err))
	}
	defer grpcConn.Close()
	clientCtx := client.Context{
		GRPCClient:        grpcConn,
		ChainID:           LavaChainid,
		InterfaceRegistry: app.MakeEncodingConfig().InterfaceRegistry,
		TxConfig:          app.MakeEncodingConfig().TxConfig,
	}
	specQuerier := spectypes.NewQueryClient(clientCtx)
	ctx := context.Background()
	var allChains *spectypes.QueryShowAllChainsResponse
	for retry := 0; retry < maxRetries; retry++ {
		allChains, err = specQuerier.ShowAllChains(ctx, &spectypes.QueryShowAllChainsRequest{})
		if err == nil {
			break
		}
		if retry < maxRetries-1 {
			exponentialBackoff(retry)
		} else {
			return types.QueryAccountInfoResponse{}, utils.LavaFormatError("failed to get lava all chains", err)
		}
	}

	pairingQuerier := types.NewQueryClient(clientCtx)
	epochStorageQuerier := epochstoragetypes.NewQueryClient(clientCtx)
	serviceClient := tmservice.NewServiceClient(grpcConn)

	latestBlockRes, err := serviceClient.GetLatestBlock(
		ctx,
		&tmservice.GetLatestBlockRequest{},
	)
	if err != nil {
		return types.QueryAccountInfoResponse{}, utils.LavaFormatError("Failed to get latest block", err)
	}

	header := latestBlockRes.GetBlock().GetHeader()
	currentBlock := header.GetHeight()
	// gather information
	var info types.QueryAccountInfoResponse

	// fill the objects
	for _, chainStructInfo := range allChains.ChainInfoList {
		chainID := chainStructInfo.ChainID
		var response *types.QueryProvidersResponse
		for retry := 0; retry < maxRetries; retry++ {
			response, err = pairingQuerier.Providers(ctx, &types.QueryProvidersRequest{
				ChainID:    chainID,
				ShowFrozen: true,
			})
			if err == nil {
				break
			}

			if retry < maxRetries-1 {
				exponentialBackoff(retry)
			} else {
				return types.QueryAccountInfoResponse{}, utils.LavaFormatError("failed to get all providers after retries", err)
			}
		}
		if len(response.StakeEntry) > 0 {
			for _, provider := range response.StakeEntry {
				if provider.Address == address {
					if provider.StakeAppliedBlock > uint64(currentBlock) {
						info.Frozen = append(info.Frozen, provider)
					} else {
						info.Provider = append(info.Provider, provider)
					}
					break
				}
			}
		}
	}
	for retry := 0; retry < maxRetries; retry++ {
		unstakeEntriesAllChains, err := epochStorageQuerier.StakeStorage(ctx, &epochstoragetypes.QueryGetStakeStorageRequest{
			Index: epochstoragetypes.StakeStorageKeyUnstakeConst,
		})
		if err == nil {
			if len(unstakeEntriesAllChains.StakeStorage.StakeEntries) > 0 {
				for _, unstakingProvider := range unstakeEntriesAllChains.StakeStorage.StakeEntries {
					if unstakingProvider.Address == address {
						info.Unstaked = append(info.Unstaked, unstakingProvider)
					}
				}
			}
			break
		}
		if retry < maxRetries-1 {
			exponentialBackoff(retry)
		} else {
			return types.QueryAccountInfoResponse{}, utils.LavaFormatError("failed to get all unstake entries", err)
		}
	}
	return info, nil
}

// CheckSpecificChainProvider check specific chain provider, ChainsToCheck example ["LAV1","EVMOST"]
func CheckSpecificChainProvider(address string, lavaGrpc string, LavaChainid string, ChainsToCheck []string) []alert.Alert {
	info, err := LavaProviderChecker(address, lavaGrpc, LavaChainid, 3)
	var alerts []alert.Alert
	if err != nil {
		log.Log.Error("Failed to check lava provider info", zap.Error(err))
		alerts = append(alerts, alert.Alert{
			Type:      "warning",
			Title:     "Failed to check lava provider info",
			Message:   "Failed to check lava provider info",
			Details:   fmt.Sprintf("Failed to check lava provider %s info", address),
			Timestamp: time.Now(),
			Origin:    "Lava Provider Monitor",
		})
		return alerts
	}
	//check chains to check is staked or not
	for _, chain := range ChainsToCheck {
		var staked bool
		for _, provider := range info.Provider {
			if provider.Chain == chain {
				log.Log.Info("Lava provider is working fine", zap.String("chain", chain))
				staked = true
				break
			}
		}
		if !staked {
			log.Log.Info("Your lava provider is jailed or frozen", zap.String("chain", chain))
			alerts = append(alerts, alert.Alert{
				Type:      "Error",
				Title:     "Provider Jailed(or Frozen)",
				Message:   "Your lava provider is jailed or frozen",
				Details:   fmt.Sprintf("Your lava provider on %s is not staked", chain),
				Timestamp: time.Now(),
				Origin:    "Lava Provider Monitor",
			})
		}
	}
	return alerts
}
