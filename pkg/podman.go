package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Podman struct {
	Name string
	Dir  string
}

func NewPodmanImage(funcData FuncData) Podman {

	dir, err := GenerateTempFolder()
	if err != nil {
		log.Panicf("%s: %v", "Failed to create Temp dir", err)
	}
	log.Println("Context Dir", dir)

	templateDir := filepath.Join("templates", funcData.Language)
	err = CopyFolder(templateDir, dir)
	if err != nil {
		fmt.Println("Error copying dockerfile:", err)
	}
	err = writeFuctionToFile(funcData, dir)
	if err != nil {
		fmt.Println("Error copying dockerfile:", err)
	}
	p := Podman{
		Name: funcData.Name,
		Dir:  dir,
	}
	return p
}

func (p Podman) build() error {
	fmt.Println("podman build")
	// Replace this command with your actual Podman build command
	lowerName := strings.ToLower(p.Name)
	cmd := exec.Command("podman", "build", "-t", lowerName, p.Dir)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func (p Podman) push() error {
	// Replace this command with your actual Podman build command
	fmt.Println("podman build")
	// cmd := exec.Command("podman", "push", p.Name)
	// cmd.Dir = p.Dir
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// return cmd.Run()
	return nil
}

func (p Podman) remove() error {
	fmt.Println("podman remove")

	// cmd := exec.Command("podman", "image", "remove", p.Name)
	// cmd.Dir = p.Dir
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// return cmd.Run()
	return nil
}
