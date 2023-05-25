package adapters

import "github.com/dgraph-io/ristretto"

func CreateLocalCacheWithConfig(config *ristretto.Config) (*ristretto.Cache, error) {
	return ristretto.NewCache(config)
}
