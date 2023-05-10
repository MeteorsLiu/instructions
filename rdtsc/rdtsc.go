package rdtsc

func rdtsc() (n uint64)
func RDTSCP() uint64 {
	return rdtsc()
}
