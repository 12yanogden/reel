package reel

import (
	"testing"

	"github.com/12yanogden/slices"
)

func TestReel(t *testing.T) {
	var reel Reel
	frames := []string{
		"[        ]",
		"[       ]",
		"[      ]",
		"[     ]",
		"[    ]",
		"[   ]",
		"[  ]",
		"[ ]",
		"[]",
		"[ ]",
		"[ P]",
		"[ PA]",
		"[ PAS]",
		"[ PASS]",
		"[ PASSE]",
		"[ PASSED]",
		"[ PASSED ]",
	}
	expected := "[        ][       ] [      ]  [     ]   [    ]    [   ]     [  ]      [ ]       []        [ ]       [ P]      [ PA]     [ PAS]    [ PASS]   [ PASSE]  [ PASSED] [ PASSED ]"
	actual := ""

	reel.Init(frames, 10)

	for range slices.Indexes(frames) {
		actual += reel.Play()
	}

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s\n", expected, actual)
	}
}
