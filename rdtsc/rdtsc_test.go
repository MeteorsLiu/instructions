package rdtsc

import (
	"testing"
	"time"

	_ "unsafe"
)

//go:linkname nanotime runtime.nanotime
func nanotime() int64

func TestRdtsc(t *testing.T) {
	t1 := RDTSCP()
	//t.Log(t1)
	time.Sleep(time.Second)
	t2 := RDTSCP()

	t.Log(time.Duration(t2 - t1).Seconds())
}

func TestNanotime(t *testing.T) {
	t1 := nanotime()
	//t.Log(t1)
	time.Sleep(time.Second)
	t2 := nanotime()

	t.Log(time.Duration(t2 - t1).Seconds())
}

func BenchmarkRDTSCP(b *testing.B) {
	b.Run("rdtscp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RDTSCP()
		}
	})

	b.Run("nanotime", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nanotime()
		}
	})
}
