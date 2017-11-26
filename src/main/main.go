package main

import (
	"flag"
	"flgs"
	"header"
	log "logger"
	"userInput"
)

// If user makes wrong usage of flags, we trigger ShowHelp()
func validateFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		flgs.ShowHelp()
	}
}

func main() {

	// Simple name prompt
	question := "What's your name? \n"
	response := "Your name is: "
	required := false
	userInput.QAR(question, response, required)

	// Define speak
	var speak bool
	flag.BoolVar(&speak, "s", false, "")
	flag.BoolVar(&speak, "speak", false, "")

	// Check if the current flag is valid
	validateFlag(flag.CommandLine)
	// Show the fancy ASCII header
	header.Show()
	// Initialize the logger (to get date)
	log.InitLogger()

	// Parse the command-line flags
	flag.Parse()

	// If bool status is true, execute git status
	if speak {
		flgs.InitSpeak()
	}

}
