package utils

import (
	"gmcm/static"
	"os"
	"os/exec"
)

const DefaultEditor = "vi"

func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	if IsNotExist(filename) {
		content := static.DefaultConfigContent
		_ = os.WriteFile(filename, content, 0644)
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
