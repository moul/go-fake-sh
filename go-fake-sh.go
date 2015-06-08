package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var verbose bool   // -x flag
var command string // -c flag

func init() {
	flag.StringVar(&command, "c", "", "Is $0 /bin/sh ?")
	flag.BoolVar(&verbose, "x", false, "Increase verbosity")

	flag.Parse()

	// logging
	log.SetOutput(os.Stderr)
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.WithFields(log.Fields{
		"verbose": verbose,
		"command": command,
		"tail":    flag.Args(),
	}).Debug("Args")

}

// Usage in a Dockerfile:
//   RUN {COMMAND}
func main() {
	execCommand := strings.Split(command, " ")
	spawn := exec.Command(execCommand[0], execCommand[1:]...)
	spawn.Stdout = os.Stdout
	spawn.Stdin = os.Stdin
	spawn.Stderr = os.Stderr
	err := spawn.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
