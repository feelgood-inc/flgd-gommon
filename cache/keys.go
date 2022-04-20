package cache

import "os"

func TransformToScopedKey(key string) string {
	scope := os.Getenv("CACHE_SCOPE")
	if scope == "" {
		return key
	}

	return scope + key
}
