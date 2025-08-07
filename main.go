package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Println("Starting Sentinel Auth Server...")
	
	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}