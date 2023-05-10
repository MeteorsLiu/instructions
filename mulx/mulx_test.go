package mulx

import (
	"math/bits"
	"testing"
)

func TestMulx(t *testing.T) {
	t.Log(Mulx(5, 6))
}

func BenchmarkMulxSmall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mulx(5, 6)
	}
}

func BenchmarkGoSmall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bits.Mul64(5, 6)
	}
}
func BenchmarkMulxLarge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mulx(0xa0761d6478bd642f, 0xe7037ed1a0b428db)
	}
}

func BenchmarkGoLarge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bits.Mul64(0xa0761d6478bd642f, 0xe7037ed1a0b428db)
	}
}
