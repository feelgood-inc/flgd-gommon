package adapters

import "github.com/dgraph-io/ristretto"

func CreateLocalCache(config *ristretto.Config) (*ristretto.Cache, error) {
	return ristretto.NewCache(config)
}
