// spb facilitates solving NYT Spelling Bee games
package main

import (
	"fmt"
	"strings"
)

// distinctLetters accepts one or more words and returns true if there are no
// duplicate letters in the words (taken together)
func distinctLetters(words ...string) (yes bool) {
	var lmap = map[rune]int{}
	for _, word := range words {
		for _, r := range word {
			_, oldrune := lmap[r]
			if !oldrune {
				lmap[r] = 1
			} else {
				return // found first duplicate, no need to keep checking
			}
		}
	}
	for r := range lmap {
		if lmap[r] > 1 {
			return
		}
	}
	yes = true
	return
}

// almostDistinctLetters accepts one or more words and returns true if there is
// at most one duplicate letter.  in the words (taken together)
func almostDistinctLetters(ndups int, words ...string) (yes bool) {
	var lmap = map[rune]int{}
	for _, word := range words {
		for _, r := range word {
			_, oldrune := lmap[r]
			if !oldrune {
				lmap[r] = 1
			} else {
				lmap[r] += 1
			}
		}
		sumdups := 0
		for _, v := range lmap {
			if v > 1 {
				sumdups += v - 1
			}
		}
		if sumdups > ndups {
			return
		}
	}
	yes = true
	return
}

// findCandidates returns a slice of string containing all 5 letter words
// with no duplicate letters found in wordlist
func findCandidates(wordlist []string) []string {
	words := []string{}
	for _, word := range wordlist {
		if len(word) != 5 {
			continue
		}
		if !distinctLetters(word) {
			continue
		}
		// add it to the list of candidates
		words = append(words, word)
	}
	return words
}

// hasOneVowel returns true if word has exactly one vowel. The vowel is also
// returned.
func hasOneVowel(word string) (bool, rune) {
	wlower := strings.ToLower(word)
	var nvowels = 0
	var v0 rune // first vowel found
	for _, r := range "aeiou" {
		if strings.Contains(wlower, string(r)) {
			nvowels++
			v0 = r
		}
		if nvowels > 1 {
			return false, v0
		}
	}
	return true, v0
}

// OneVowelWords is a convenience struct type for holding lists of one-vowel
// words in separate slices for each vowel
type OneVowelWords struct {
	a, e, i, o, u []string
}

// Main finds and prints sets of five five-letter words that cover at least 24
// letters using a depth-first search.  Note that breadth-first brute force
// searching is impractical since that would result in more than 10 quintillion,
// (10e16) comparisons.
func main() {
	candidates := findCandidates(AllWords)
	screened := OneVowelWords{}
	for _, w := range candidates {
		if ok, r := hasOneVowel(w); ok {
			switch r {
			case 'a':
				screened.a = append(screened.a, w)
			case 'e':
				screened.e = append(screened.e, w)
			case 'i':
				screened.i = append(screened.i, w)
			case 'o':
				screened.o = append(screened.o, w)
			case 'u':
				screened.u = append(screened.u, w)
			}
		}
	}
	fmt.Printf("%d 'a' candidates found\n", len(screened.a))
	fmt.Printf("%d 'e' candidates found\n", len(screened.e))
	fmt.Printf("%d 'i' candidates found\n", len(screened.i))
	fmt.Printf("%d 'o' candidates found\n", len(screened.o))
	fmt.Printf("%d 'u' candidates found\n", len(screened.u))

	// Search order is (u, e, o, i, a). I chose this order because 'u' has the
	// fewest candidates and 'a' has the most.

	nfound := 0
	for _, wu := range screened.u {
		for _, we := range screened.e {
			if distinctLetters(wu, we) {
				for _, wo := range screened.o {
					if distinctLetters(wu, we, wo) {
						for _, wi := range screened.i {
							if distinctLetters(wu, we, wo, wi) {
								for _, wa := range screened.a {
									// Prior exploration has shown that there are no 5 word combinations that
									// cover 25 letters. However, there are over 600 that cover 24 letters.
									if almostDistinctLetters(1, wu, we, wo, wi, wa) {
										nfound++
										fmt.Printf("%d: %s, %s, %s, %s, %s\n", nfound, wu, we, wo, wi, wa)

									}
								}
							}
						}
					}
				}
			}
		}
	}
}
