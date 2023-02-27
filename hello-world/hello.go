package main

import "fmt"

const languageFrench = "French"
const languageSpanish = "Spanish"
const languageLatin = "Latin"
const languageEnglish = "English"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const latinHelloPrefix = "Salve, "

func helloWord(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

// "prefix string" no retorno, automaticamente cria a variável dentro da função.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case languageFrench:
		prefix = frenchHelloPrefix
	case languageSpanish:
		prefix = spanishHelloPrefix
	case languageLatin:
		prefix = latinHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(helloWord("Dyonatha", "Spanish"))
}
