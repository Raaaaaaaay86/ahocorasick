package ahocorasick_test

import (
	"strings"
	"testing"

	"github.com/raaaaaaaay86/ahocorasick"
)

var dictionary = []string{
	"apple",
	"banana",
	"applepie",
	"pie",
	"orange",
	"grape",
	"grapefruit",
	"grape fruit",
	"fruit",
	"a",
	"an",
	"the",
	"hello",
	"world",
	"word",
	"coding",
	"code",
	"program",
}

func TestMatching(t *testing.T) {
	sentences := []string{
		"apple",
		"orange banana grape",
		"applepie",
		"grapefruit",
		"grape fruit",
		"I love applepie and grape juice, learning to code.",
		"cherry melon kiwi",
		"",
		"Apple Orange",
		"banana banana pie",
		"an apple and an orange",
		"hello world, this is a programming example, where coding is important for the program.",
		"hello world hell",
		"hellow",
	}

	acTrie := ahocorasick.NewTrie(dictionary)

	for _, sentence := range sentences {
		t.Run(sentence, func(t *testing.T) {
			t.Parallel()

			matches := make(map[string]struct{})

			for _, match := range acTrie.FindAllMatches(sentence) {
				start, end := match.StartIndex, match.EndIndex
				matches[sentence[start:end]] = struct{}{}
			}

			for _, text := range dictionary {
				contains := strings.Contains(sentence, text)

				if _, matched := matches[text]; matched != contains {
					t.Errorf("Expected %s to be %v in sentence '%s', but got %v", text, contains, sentence, matched)
				}
			}
		})
	}
}
