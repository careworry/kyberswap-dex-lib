package dvm

import (
	"context"

	"github.com/KyberNetwork/ethrpc"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/shared"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	pooltrack "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool/tracker"
)

type PoolTracker struct {
	sharedPoolTracker *shared.PoolTracker
}

var _ = pooltrack.RegisterFactoryCE(PoolType, NewPoolTracker)

func NewPoolTracker(
	cfg *shared.Config,
	ethrpcClient *ethrpc.Client,
) (*PoolTracker, error) {
	sharedPoolTracker, err := shared.NewPoolTracker(cfg, ethrpcClient)
	if err != nil {
		return nil, err
	}

	return &PoolTracker{
		sharedPoolTracker: sharedPoolTracker,
	}, nil
}

func (d *PoolTracker) GetNewPoolState(
	ctx context.Context,
	p entity.Pool,
	params pool.GetNewPoolStateParams,
) (entity.Pool, error) {
	return d.sharedPoolTracker.GetNewPoolState(ctx, p, params)
}
