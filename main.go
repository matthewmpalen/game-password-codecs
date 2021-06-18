package main

import (
	"fmt"
	"game-password-codecs/metroid"
	"io/ioutil"
	"log"
)

var (
	m = metroid.Metroid{}
)

func decodeExamples() {
	invalidPasswords := []string{
		"",
		"ABCDEFGHIJKLMNOPQRSTUVWX",
		"cnaibiadsjSKDJLND?djkdfX",
		"NewContextRetro000000000",
	}

	for _, password := range invalidPasswords {
		_, err := m.Decode(password)
		log.Printf("ERROR: password: '%s'\n%v\n", password, err)
	}

	validPasswords := []string{
		"JUSTINBAILEY------------",
		"NARPASSWORD0000000COPADO",
		"NewContext00Retro000001V",
		"000000000000000E0000000E",
	}

	for _, password := range validPasswords {
		state, _ := m.Decode(password)
		fmt.Printf("password: '%s'\n%s\n", password, state.ToJSON(true))

		encodedPassword, _ := m.Encode(*state)
		fmt.Printf("encodedPassword: %s\n", encodedPassword)
	}

	brokenPasswords := []string{
		"ENGAGERIDLEYMUFFINKICKER",
	}

	for _, password := range brokenPasswords {
		state, _ := m.Decode(password)
		fmt.Printf("password: '%s'\n%s\n", password, state.ToJSON(true))
		if state.Reset {
			log.Print("WARNING: reset bit is set! Will crash the game!")
		}
	}
}

func encodeExamples() {
	json1, _ := ioutil.ReadFile("./metroid/examples/2_reset.json")
	jsonState, _ := metroid.NewGameStateFromJSON(string(json1))
	password, _ := m.Encode(*jsonState)
	fmt.Println(metroid.PrettyPassword(password))
}

func main() {
	decodeExamples()
	encodeExamples()
}
