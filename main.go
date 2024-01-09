package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func pullImage(imageName string) error {
	cmd := exec.Command("podman", "pull", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to pull image: %v", err)
	}

	return nil
}

func main() {
	imageName := "docker.io/library/nginx:latest"

	err := pullImage(imageName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Image %s successfully pulled!\n", imageName)

	// Sleep for 10 minutes
	fmt.Println("Sleeping for 10 minutes...")
	time.Sleep(10 * time.Minute)

	fmt.Println("Wake up after 10 minutes!")
}
