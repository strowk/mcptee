package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// mcptee is a very simple CLI tool that runs a command and pipes its output to the standard output,
// while also piping the standard input line by line to the command, in that it proxies stdin and stdout.
// Additionally it also terminates the command when the standard input is closed and exits with the same exit code as the command.
// It would every input line to provided file following prefix in: and every output line to provided file following prefix out:
// Before every input line it would also output multi-document YAML separator --- and endline.

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: mcptee <out-file> <command> [args...]")
		os.Exit(1)
	}

	outFile := os.Args[1]
	command := os.Args[2]
	restArgs := os.Args[3:]

	// Open the output file
	logOut, err := os.Create(outFile)
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer logOut.Close()

	// Run the command
	cmd := exec.Command(command, restArgs...)

	cmdin, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Pipe the standard input line by line to the command
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error reading standard input: %s\n", err)
				break
			}
			fmt.Fprintf(logOut, "---\n")
			fmt.Fprintf(logOut, "in: %s", line)
			cmdin.Write([]byte(line))
		}
	}()

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Pipe the command standard output line by line to the standard output and file
	go func() {
		reader := bufio.NewReader(cmdOut)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Printf("Error reading command output: %s\n", err)
				break
			}
			fmt.Fprintf(logOut, "out: %s", line)
			fmt.Print(line)
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	log.Printf("Command %s %v started", command, restArgs)
	err = cmd.Wait()

	// Check for errors
	if err != nil {
		log.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Get the exit code
	exitCode := cmd.ProcessState.ExitCode()

	// Exit with the same exit code
	os.Exit(exitCode)
}

// examples:
// go run main.go test.yaml go run -C C:/work/mcp-k8s-go main.go
// go run main.go test.yaml mcp-k8s-go
// mcptee test.yaml mcp-k8s-go
