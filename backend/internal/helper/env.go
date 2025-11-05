package helper

import (
	"fmt"
	"os"
	"strconv"
)

func GetStrEnvOrPanic(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable [%s] not found", key))
	}

	return val
}

func GetIntEnvOrPanic(key string) int {
	val := GetStrEnvOrPanic(key)

	intEnv, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("environment variable [%s] is not an integer [%s]", key, val))
	}

	return intEnv
}
