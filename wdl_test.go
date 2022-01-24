package main

import (
	"testing"

	"github.com/go-test/deep"
)

func TestDistinctLetters(t *testing.T) {
	type testcase struct {
		words []string
		exp   bool
	}
	tcases := []testcase{
		{[]string{"abc"}, true},
		{[]string{"abc", "xyz"}, true},
		{[]string{"abbca"}, false},
		{[]string{"abc", "xyzb"}, false},
	}
	for _, tc := range tcases {
		got := (distinctLetters(tc.words...))
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
}

func TestAlmostDistinctLetters(t *testing.T) {
	type testcase struct {
		ndups int
		words []string
		exp   bool
	}
	tcases := []testcase{
		{1, []string{"abcde", "afghi"}, true},
		{2, []string{"abcde", "afghi"}, true},
		{1, []string{"abcde", "afghe"}, false},
		{1, []string{"abcde", "afgha"}, false},
		{2, []string{"abcde", "afghe"}, true},
	}
	for _, tc := range tcases {
		got := (almostDistinctLetters(tc.ndups, tc.words...))
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
}

func TestHasOneVowel(t *testing.T) {
	type testcase struct {
		word string
		exp  bool
	}
	tcases := []testcase{
		{"phlox", true},
		{"cases", false},
	}
	for _, tc := range tcases {
		got, _ := hasOneVowel(tc.word)
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
}
func TestFindCandidates(t *testing.T) {
	type testcase struct {
		wordlist []string
		exp      []string
	}
	tcases := []testcase{
		{[]string{"slept"}, []string{"slept"}},
		{[]string{"slept", "sleep", "nod"}, []string{"slept"}},
		{[]string{"slept", "sleep", "nod", "munch"}, []string{"slept", "munch"}},
	}
	for _, tc := range tcases {
		got := findCandidates(tc.wordlist)
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
}
