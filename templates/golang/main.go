package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func getEnvBool(key string, defaultValue bool) bool {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(envValue)
	if err != nil {
		return defaultValue
	}

	return boolValue
}

func main() {
	var wg sync.WaitGroup

	if getEnvBool("MESSAGING", false) {
		wg.Add(1)
		go HandleMessaging(&wg)
	}
	if getEnvBool("HTTP", true) {
		wg.Add(1)
		go HandleHttp(&wg)
	}
	wg.Wait()

	fmt.Println("All handlers are running.")
}
