package popcount

var pc2 [256]byte

func init() {
	for i := range pc2 {
		pc2[i] = pc2[i/2] + byte(i&1)
	}
}

func PopCount2(x uint64) int {
	var sum byte
	for i:=0; i<8; i++ {
		sum += pc2[byte(x>>(i*8))]
	}

	return int(sum)
}
