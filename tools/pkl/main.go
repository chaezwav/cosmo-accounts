package pkl

import (
	"context"

	"github.com/chaeyeonswav/cosmo-accounts/gen/appconfig"
)

func LoadFromPath(ctx context.Context, path string) (*appconfig.AppConfig, error) {
	cfg, err := appconfig.LoadFromPath(ctx, path)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
