// Code generated from Pkl module `chaeyeonswav.cosmo_accounts.PostgresConfig`. DO NOT EDIT.
package postgresconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type PostgresConfig struct {
	Enabled bool `pkl:"enabled"`

	Host string `pkl:"host"`

	Port uint16 `pkl:"port"`

	Database string `pkl:"database"`

	Auth *Auth `pkl:"auth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a PostgresConfig
func LoadFromPath(ctx context.Context, path string) (ret *PostgresConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a PostgresConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*PostgresConfig, error) {
	var ret PostgresConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
