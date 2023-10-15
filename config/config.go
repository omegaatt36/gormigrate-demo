package config

import (
	"os"
)

// GetString returns the value of the environment variable named by the key.
func GetString(key string) string {
	return os.Getenv(key)
}

// // GetBool returns the value of the environment variable named by the key.
// func GetBool(key string) bool {
// 	val, err := strconv.ParseBool(os.Getenv(key))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return val
// }
