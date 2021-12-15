package answer

func Swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
