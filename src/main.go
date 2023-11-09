package tops

func Tops(msg string) string {
	for tops, n, i := []uint8{}, 2, 1; ; n, i = n+1, 2*pow(n, 2)-n {
		if i >= len(msg) {
			reverse(tops)
			return string(tops)
		}
		tops = append(tops, msg[i])
	}
}

func reverse(str []uint8) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func pow(x int, y int) int {
	result := x
	for i := 0; i < y-1; i++ {
		result *= x
	}
	return result
}
