package rdtsc

import "testing"

func TestRdtsc(t *testing.T) {
	t.Log(RDTSCP())
}
