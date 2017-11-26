package flgs

import (
	"fmt"
	"genSpeech"
	"io/ioutil"
	"playSound"
)

func ShowHelp() {
	fmt.Println(`

Usage: CLI Template [OPTIONS]
Options:
	-h, --help      Print the help log.
	
`)
}

func InitSpeak() {
	polly := genSpeech.New("AKIAJNYLICWX7Z3YN2JA", "KHJ6lsSMknSknpI3qu+6klgK59GmNZxE+eS8URgW")

	polly.Format(genSpeech.OGGMP3)
	polly.Voice(genSpeech.Brian)

	text := "Hi, are you having a good weekend?"
	bytes, err := polly.Speech(text)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./result.ogg", bytes, 0644)
	playSound.Run()
}
