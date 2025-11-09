package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// TemplateOptions contiene toda la información para crear y configurar la plantilla.
type Script struct {
	Name string
	Url  string
}

func (j Script) Execute() error {

	fmt.Printf("clone repository %s in %s...\n", j.Url, j.Name)

	cmd := exec.Command("git", "clone", j.Url, j.Name)

	if output, err := cmd.CombinedOutput(); err != nil {

		return fmt.Errorf("error : %v\nOutput Git: %s", err, string(output))
	}

	gitDir := filepath.Join(j.Name, ".git")

	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("\n--- Process completed!\nconfigure your environment: ---")
	return nil
}

func InitialScript(name string, url string) IScript {
	return Script{
		Name: name,
		Url:  url,
	}
}
