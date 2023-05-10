package rdrand

func rdrand64() (ok bool, ret uint64)
func rdseed64() (ok bool, ret uint64)

func RDRand() (ok bool, ret uint64) {
	return rdrand64()
}

func RDSeed() (ok bool, ret uint64) {
	return rdseed64()
}
