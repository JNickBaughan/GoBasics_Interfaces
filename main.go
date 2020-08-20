package main

import "fmt"

type speakingBot interface {
	getGreeting() string
	getFarewell() string
}

type englishSpeakingBot struct {
	greeting string
	farewell string
}

type spanishSpeakingBot struct {
	saludo    string
	despedida string
}

func main() {
	esb := englishSpeakingBot{"hello", "bye"}
	ssb := spanishSpeakingBot{"hola", "adios"}

	printGreeting(esb)
	printFarewell(esb)

	printGreeting(ssb)
	printFarewell(ssb)

}

func (esb englishSpeakingBot) getGreeting() string {
	return esb.greeting
}

func (esb englishSpeakingBot) getFarewell() string {
	return esb.farewell
}

func (ssb spanishSpeakingBot) getGreeting() string {
	return ssb.saludo
}

func (ssb spanishSpeakingBot) getFarewell() string {
	return ssb.despedida
}

func printGreeting(sb speakingBot) {
	fmt.Println(sb.getGreeting())
}

func printFarewell(sb speakingBot) {
	fmt.Println(sb.getFarewell())
}
