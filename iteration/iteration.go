package iteration

// refractor: Added const inside function
const repeatCount = 5

func Repeat(char string) string {
	var iteration string
	for i := 0; i < repeatCount; i++ {
		iteration += char
	}
	return iteration
}
