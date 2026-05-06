package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 88, Capacity: 87, Latency: 27, Risk: 14, Weight: 12}, wantScore: 185, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 89, Capacity: 106, Latency: 14, Risk: 5, Weight: 9}, wantScore: 271, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 99, Capacity: 101, Latency: 18, Risk: 18, Weight: 12}, wantScore: 215, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
