package utils

import "os"

// EnvOrDefault get env by key, return default value if not set
func EnvOrDefault(k, dv string) string {
	vv := os.Getenv(k)
	if vv == "" {
		return dv
	}
	return vv
}
