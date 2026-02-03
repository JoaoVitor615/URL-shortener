package encoder

func Encode(num int) string {
	if num == 0 {
		return string(alphabet[0])
	}

	encoded := ""
	for num > 0 {
		encoded = string(alphabet[num%Base]) + encoded
		num = num / Base
	}
	return encoded
}
