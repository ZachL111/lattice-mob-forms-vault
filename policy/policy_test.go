package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 88, Capacity: 87, Latency: 27, Risk: 14, Weight: 12}
	if got := Score(signal); got != 185 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 89, Capacity: 106, Latency: 14, Risk: 5, Weight: 9}
	if got := Score(signal); got != 271 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 99, Capacity: 101, Latency: 18, Risk: 18, Weight: 12}
	if got := Score(signal); got != 215 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
