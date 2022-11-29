package popcount

func PopCount3(x uint64) int {
	var sum int

	for i:=0; i<64; i++ {
		x = x >> 1
		if (x & 1) == 1 {
			sum++
		}
	}

	return sum
}
