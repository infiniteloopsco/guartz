package utils

import "os"

func IsProd() bool {
	return isEnv("prod")
}

func IsTest() bool {
	return isEnv("test")
}

func IsDev() bool {
	return isEnv("dev")
}

func isEnv(env string) bool {
	return os.Getenv("GIN_MODE") == env
}
