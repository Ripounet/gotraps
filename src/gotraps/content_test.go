package gotraps

import (
	"os/exec"
	"strings"
	"testing"
)

var traps = []struct {
	name              string
	error             bool
	expectedPhrases   []string
	unexpectedPhrases []string
}{
	{"slice-traversal", false, []string{"Hello 0", "Hello 1", "Hello 2"}, nil},
	{"string-traversal", false, []string{"Character #3 is é", "Character #5 is -"}, nil},
	{"map-traversal", false, nil, nil},
	{"boosted-iteration", false, nil, []string{"Louie"}},
	{"rendezvous-iteration", false, []string{"Louie"}, []string{"Huey", "Dewey"}},
	{"yet-another-iteration", true, []string{"deadlock!"}, nil},
	{"pipeline-traversal", false, []string{"Huey", "Dewey", "Louie"}, []string{"0", "1", "2"}},
	// Don't run "collatz-conjecture"! It's an infinite loop.
	// {"collatz-conjecture", false, []string{}, nil},
	{"collatz-2", false, []string{"(1+3i)"}, nil},
	{"watchman", false, []string{"error calling secureGreet: <nil>"}, nil},
	{"watchman-2", false, []string{"Here is my greeting", "No greeting because of"}, nil},
	{"watchman-3", false, []string{"Huey , you shall not pass!", "Dewey , you shall not pass!"}, nil},
	{"append", false, []string{"[9 1 2 88 99]"}, nil},
	{"just-saying", false, []string{"speaks"}, []string{"Rosie", "barks"}},
	{"map-indexing", false, []string{"0"}, []string{"panic"}},
	{"addition", false, []string{"367"}, []string{"369"}},
	{"equation", false, []string{"Found 0 solutions", "Solutions ="}, nil},
	{"mean", false, nil, []string{"lower = 55"}},
	{"pretty", false, []string{"155"}, []string{"4", "+", "3", "i"}},
	{"symmetry", false, []string{"{1 2}", "{3 4}", "{5 6}"}, []string{"{2 1}", "{4 3}", "{6 5}"}},
	{"symmetry-2", false, []string{"{1 2}", "{3 4}", "{5 6}"}, []string{"{2 1}", "{4 3}", "{6 5}"}},
	// Don't run "event-loop"! It hangs.
	// {"event-loop", true, nil, nil},
	{"date-format", false, []string{"2006-02-01"}, []string{"2006-01-01", "2006-01-02"}},
}

func TestContent(t *testing.T) {
	for _, trap := range traps {
		if trap.name == "" {
			continue
		}
		b, err := exec.Command("go", "run", "content/"+trap.name+"/code.go").CombinedOutput()
		if err != nil && !trap.error {
			t.Errorf("[%s] Should not have produced error: %v", trap.name, err)
			continue
		}
		if err == nil && trap.error {
			t.Errorf("[%s] Should have produced an error, but didn't", trap.name)
		}

		output := string(b)

		for _, fragment := range trap.expectedPhrases {
			if !strings.Contains(output, fragment) {
				t.Errorf("[%s] Expected %q not found in output\n %s", trap.name, fragment, output)
			}
		}

		for _, fragment := range trap.unexpectedPhrases {
			if strings.Contains(output, fragment) {
				t.Errorf("[%s] Unxpected %q found in output\n %s", trap.name, fragment, output)
			}
		}
	}
}
