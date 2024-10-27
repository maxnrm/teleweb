package envs

import (
	"fmt"
	"os"
	"strconv"
)

var BOT_TOKEN = getEnvStr("BOT_TOKEN")
var LOG_LEVEL = getEnvInt("LOG_LEVEL")

func getEnvStr(key string) string {
	v, exists := os.LookupEnv(key)
	if !exists || v == "" {
		fmt.Fprintf(os.Stderr, "error: %v\n", ErrNoEnv{key})
		os.Exit(1)
	}
	return v
}

func getEnvInt(key string) int {
	strVal := getEnvStr(key)

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", ErrNoEnv{key})
		os.Exit(1)
	}

	return intVal
}

type ErrNoEnv struct {
	key string
}

func (e ErrNoEnv) Error() string {
	return fmt.Sprintf("env var with name: %s not found or empty string", e.key)
}

type ErrParseIntFailed struct {
	key string
}

func (e ErrParseIntFailed) Error() string {
	return fmt.Sprintf("env %s can not be parsed as int", e.key)
}
