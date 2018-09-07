package utils

import (
	"fmt"
	"os"
)

// EnvOrDefault get env by key, return default value if not set
func EnvOrDefault(k, dv string) string {
	vv := os.Getenv(k)
	if vv == "" {
		return dv
	}
	return vv
}

// RequiredEnv ensure the env must set
func RequiredEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("env %s is required. ", k))
	}
	return v
}
