package set1

// Xor takes two equal-length buffers and produces their XOR combination.
func Xor(input, xoree []byte) []byte {
	res := make([]byte, len(input))

	for i := range input {
		res[i] = input[i] ^ xoree[i]
	}

	return res
}
