package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestBucket(t *testing.T) {
	repeated := Repeater(1)
	expected := 5

	if repeated != expected {
		t.Errorf("expected %d and got %d", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	//setup
	for b.Loop() {
		//code to measure
		Repeat("a")
	}
	//cleanup
}
