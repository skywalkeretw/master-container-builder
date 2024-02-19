package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type Podman struct {
	Name              string
	ImgName           string
	Dir               string
	DockerhubUsername string
	DockerhubPassword string
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
		fmt.Println("Error writting function to file:", err)
	}
	err = writeSpecToFile(funcData, dir)
	if err != nil {
		fmt.Println("Error writting specs:", err)
	}

	lowerName := strings.ToLower(funcData.Name)
	dockerhub_username := GetEnvSting("DOCKERHUB_USERNAME", "username")
	dockerhub_password := GetEnvSting("DOCKERHUB_PASSWORD", "password")
	re := regexp.MustCompile(`[\s\n\r\t]+`)
	cleanedUsername := re.ReplaceAllString(dockerhub_username, "")
	iamgeName := fmt.Sprintf("%s/master-imgs:%s", cleanedUsername, lowerName)
	fmt.Println("imagenaem", iamgeName)
	p := Podman{
		Name:              funcData.Name,
		ImgName:           iamgeName,
		Dir:               dir,
		DockerhubUsername: cleanedUsername,
		DockerhubPassword: dockerhub_password,
	}
	fmt.Println("podman data", p)
	return p
}

func (p Podman) build() error {
	fmt.Println("podman build")

	cmd := exec.Command("podman", "build", "-t", p.ImgName, p.Dir)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func (p Podman) login() error {

	fmt.Println("login to dockerhub")

	cmd := exec.Command("podman", "login", "--username", p.DockerhubUsername, "--password", p.DockerhubPassword, "docker.io")
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p Podman) push() (string, error) {
	// Replace this command with your actual Podman build command
	err := p.login()
	if err != nil {
		return p.ImgName, fmt.Errorf("failed to login to dockerhub not pushing image: %v", err)
	}

	fmt.Println("podman push")
	cmd := exec.Command("podman", "push", p.ImgName)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return p.ImgName, cmd.Run()
}

func (p Podman) remove() error {
	fmt.Println("podman remove")

	cmd := exec.Command("podman", "image", "remove", p.ImgName)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
