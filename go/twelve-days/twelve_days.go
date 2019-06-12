package twelve

import (
	"fmt"
	"strings"
)

var verses = map[int][]string{
	1:  {"first", "and a Partridge in a Pear Tree."},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

//Verse recites a particular line verse in the song "Twelve Days of Christmas"
func Verse(verseNumber int) string {
	suffix := []string{}
	for i := verseNumber; i > 0; i-- {
		suffix = append(suffix, verses[i][1])
	}
	if verseNumber == 1 {
		suffix = []string{"a Partridge in a Pear Tree."}
	}
	completeSuffix := strings.Join(suffix, ", ")
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", verses[verseNumber][0], completeSuffix)
}

//Song recites the song "Twelve Days of Christmas" in it's entirity
func Song() string {
	song := []string{}
	for i := 1; i < 13; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n") + "\n"
}
