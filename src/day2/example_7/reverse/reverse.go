package reverse

import "fmt"

//反转字符串
func Reverse(str string) string {
	strlen := len(str)
	var result string
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen-i-1])
	}

	return result
}

func Reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	strLen := len(tmp)
	for i, _ := range tmp {
		result = append(result, tmp[strLen-i-1])
	}
	return string(result)
}
