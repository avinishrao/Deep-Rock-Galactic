package main

import (
	"fmt"
	"strings"
)
// dwarvenMap is a map that holds the mapping between Dwarven characters and English characters
var dwarvenMap = map[rune]string{
	'ᛃ': "A", 'ᛒ': "B", 'ᛞ': "C", 'ᛑ': "D", 'ᛉ': "E",
	'ᛈ': "F", 'ᛋ': "G", 'ᚿ': "H", 'ᚾ': "I", 'ᚽ': "J",
	'ᚷ': "K", 'ᚶ': "L", 'ᚵ': "M", 'ᚴ': "N", 'ᚱ': "O",
	'ᚰ': "P", 'ᚯ': "Q", 'ᚮ': "R", 'ᚭ': "S", 'ᚬ': "T",
	'ᚫ': "U", 'ᚪ': "V", 'ᚩ': "W", 'ᚨ': "X", 'ᚧ': "Y",
	'ᚦ': "Z",
	'ᚥ': "1", 'ᚤ': "2", 'ᚣ': "3", 'ᚢ': "4",
	'ᚡ': "5", 'ᚠ': "6", '᚟': "7", '᚞': "8",
	'᚝': "9", '᚜': "0",
}

// dwarvenToText converts Dwarven text to English text
func dwarvenToText(dwarvenText string) string {
	var englishText strings.Builder

	// Iterate over each character in the Dwarven text
	for _, char := range dwarvenText {
		if char == ' ' {
			englishText.WriteString(" ")
		} else if englishChar, found := dwarvenMap[char]; found {
			englishText.WriteString(englishChar)
		}
	}

	return englishText.String() // Return the converted English text
}

func main() {
	dwarvenSentence := "ᚬᚿᛉ ᚯᚫᚾᛞᚷ ᛒᚮᚱᚩᚴ ᛈᚱᚨ ᚽᚫᚵᚰᚭ ᚱᚪᛉᚮ ᚥᚣ ᚶᛃᚦᚧ ᛑᚱᛋᚭ"
	fmt.Println("Dwarven language:", dwarvenSentence)

	// Convert Dwarven text to English text
	englishText := dwarvenToText(dwarvenSentence)
	fmt.Println("English text:", englishText)
}
