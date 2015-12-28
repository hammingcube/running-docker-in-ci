package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestAverage(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed because of %v", err)
	}
	dir, err := filepath.Abs(filepath.Join(cwd, "../"))
	if err != nil {
		t.Errorf("Failed because of %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "main.cpp")); err != nil {
		t.Errorf("Failed because of %v", err)
	}
	fmt.Printf("Got working directory as %s\n", cwd)
	cmd := []string{"docker", "run", "--rm", "-v", fmt.Sprintf("%s:/app", dir), "-w", "/app", "glot/clang", "sh", "-c", "g++ -std=c++11 main.cpp -o a.out && ./a.out"}
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		t.Errorf("Failed because of %v", err)
	}
	/*fmt.Printf("Got output %s\n", output)
	if string(output) != "Hello" {
		t.Error("Tests failed because did not get 'Hello'")
	}
	*/
}
