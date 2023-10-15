package main

import (
	_ "embed"
	"fmt"
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

func PrintSliceItems[T any](s []T, separtor string) {
	sb := strings.Builder{}
	elementsCount := len(s)
	sb.Grow(elementsCount + (elementsCount - 1))
	for i := range s {
		sb.WriteString(fmt.Sprint(s[i]))
		if i < len(s)-1 {
			sb.WriteString(separtor)
		}
	}
	fmt.Println(sb.String())
}

func main() {
	diacritics = strings.ReplaceAll(diacritics, " ", "")
	diacritics_slice := strings.Split(diacritics, "\n")
	// fmt.Println(diacritics_slice)
	diacritics_map := make(map[rune]string)
	for _, str := range diacritics_slice {
		strs := strings.Split(str, "=")
		for _, r := range strs[1] {
			diacritics_map[r] = strs[0]
		}
	}
	replacer := strings.NewReplacer(" ", "")
	sentence := replacer.Replace(text)
	builder := strings.Builder{}
	var nextRune rune
	words := []string{}
	sRunes := []rune(sentence)
	for i, r := range sRunes {
		if r == MEDIAL_HA {
			if sRunes[i+1] == 'င' {
				fmt.Println(string(sRunes[i+2]))
			}
		}
	}

	for i := 0; i < len(sRunes); i++ {
		r := sRunes[i]

		//checking whether index out of bounds
		//for nextRune.
		//if out of bound current rune
		//and next is the same
		//this is current rune is the last one
		if i != len(sRunes)-1 {
			nextRune = sRunes[i+1]
		} else {
			nextRune = r
		}
		if _, ok := diacritics_map[r]; ok {
			//we skipping checking if next rune is diacritic
			//if currnent rune is ္
			if r != VIRAMA {
				if _, ok = diacritics_map[nextRune]; !ok {
					builder.WriteRune(r)
					//we check if the next rune is
					//something like တ်
					if i+2 <= len(sRunes)-1 {
						//if it is, then current word in the buffer is
						//something like နတ်
						n2 := sRunes[i+2]
						if n2 == ASAT || n2 == DOT_BELOW {
							continue
						}
						words = append(words, builder.String())
						builder.Reset()
						continue
					}
				}
			}
		}
		// if all above procedures isn't executed
		//we can safe to assume that current rune
		//is part of a word
		builder.WriteRune(r)

		//if currnent rune is not ္
		//and the next rune is not a diacritics
		//or if the current rune is the last one
		//we do the following
		if _, ok := diacritics_map[nextRune]; !ok && r != VIRAMA || nextRune == r {
			//again checking for something like နတ်
			if i+2 <= len(sRunes)-1 {
				if sRunes[i+2] == ASAT || sRunes[i+2] == DOT_BELOW {
					continue
				}
			}
			words = append(words, builder.String())
			builder.Reset()
		}
	}
	PrintSliceItems(words, ",")
}
