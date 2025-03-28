package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	templates := []string{
		"https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04/deployment.yaml",
		"https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04/service.yaml",
		"https://raw.githubusercontent.com/FelipeSaijo/desafio-k8s/refs/heads/main/iniciante/04/netwokpolicies.yaml",
	}

	fmt.Println("Applying Templates...")

	// Open /dev/null to suppress output
	devNull, err := os.OpenFile(os.DevNull, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening /dev/null. Exiting...")
		os.Exit(1)
	}
	defer devNull.Close()

	for _, template := range templates {
		// Run kubectl apply command and suppress output
		cmd := exec.Command("kubectl", "apply", "-f", template)
		cmd.Stdout = devNull
		cmd.Stderr = devNull

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error applying templates. Exiting...")
			os.Exit(1)
		}
	}

	fmt.Println("All templates applied successfully!")
}
