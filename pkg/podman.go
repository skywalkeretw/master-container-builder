package pkg

import (
	"fmt"
	"log"
)

type Podman struct {
	Name       string
	Dir        string
	DockerFile string
}

func NewPodmanImage(name, language string) Podman {
	var dockerfile string
	switch language {
	case "python":
		dockerfile = "Dockerfile.python"
	case "javascript":
		dockerfile = "Dockerfile.node"
	case "golang":
		dockerfile = "Dockerfile.golang"
	}
	fmt.Println("dockerfile: ", dockerfile)
	dir, err := GenerateTempFolder()
	if err != nil {
		log.Panicf("%s: %v", "Failed to create Temp dir", err)
	}
	fmt.Println("dockerfilepath: ", "/dockerfiles/"+dockerfile)
	fmt.Println("tempdir: ", dir)
	err = CopyFile("/dockerfiles/"+dockerfile, dir)
	if err != nil {
		fmt.Println("Error copying dockerfile:", err)
	}
	p := Podman{
		Name:       name,
		Dir:        dir,
		DockerFile: dockerfile,
	}
	return p
}

func (p Podman) build() error {
	fmt.Println("podman build")
	// Replace this command with your actual Podman build command
	// cmd := exec.Command("podman", "build", "-f", p.DockerFile, "-t", p.Name, ".")
	// cmd.Dir = p.Dir
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// return cmd.Run()
	return nil
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
