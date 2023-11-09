package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TALL_AA   rune = 'ါ'
	AA             = 'ာ'
	I              = 'ိ'
	II             = 'ီ'
	U              = 'ု'
	UU             = 'ူ'
	E              = 'ေ'
	AI             = 'ဲ'
	ANUSVARA       = 'ံ'
	DOT_BELOW      = '့'
	VISARGA        = 'း'
	VIRAMA         = '္'
	ASAT           = '်'
	MEDIAL_YA      = 'ျ'
	MEDIAL_RA      = 'ြ'
	MEDIAL_WA      = 'ွ'
	MEDIAL_HA      = 'ှ'
)

//go:embed diacritics.txt
var diacritics string

//go:embed text.txt
var text string

func main() {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatalln(err)
	}
	m, max := Splitter(strings.ReplaceAll(string(data), " ", ""))
	i := make([]string, max+1, max+1)
	for k, v := range m {
		for _, j := range v {
			i[j] = k
		}
	}
	fmt.Println(i)
}
