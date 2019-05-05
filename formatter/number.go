package formatter

import "strconv"

// NumberFormat returns String with commas as thousands separators
func NumberFormat(n int64) string {
	signPadding := func(sign byte) int {
		return map[bool]int{true: 1, false: 0}[sign == '-']
	}

	in := strconv.FormatInt(n, 10)
	commaCount := (len(in) - 1 - signPadding(in[0])) / 3
	out := make([]byte, len(in)+commaCount)

	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, digit := 0, signPadding(out[0]), len(in); i < len(in); i, j, digit = i+1, j+1, digit-1 {
		out[j] = in[i]

		if digit > 3 && digit%3 == 1 {
			j++
			out[j] = ','
		}
	}
	return string(out)
}
