package iteration

func Repeat(inputcharacter string, times int) string {
	ans := ""
	for i := 0; i < times; i++ {
		ans += inputcharacter
	}
	return ans
}
