package pkg

import (
	"os"
	"os/exec"
)

type Podman struct {
	Name       string
	Dir        string
	DockerFile string
}

func NewPodmanImage(name, language string) Podman {

	p := Podman{
		Name:       "",
		Dir:        "",
		DockerFile: "",
	}
	return p
}

func (p Podman) build() error {
	// Replace this command with your actual Podman build command
	cmd := exec.Command("podman", "build", "-f", p.DockerFile, "-t", p.Name, ".")
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p Podman) push() error {
	// Replace this command with your actual Podman build command
	cmd := exec.Command("podman", "push", p.Name)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (p Podman) remove() error {

	cmd := exec.Command("podman", "image", "remove", p.Name)
	cmd.Dir = p.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
