//go:build gofuzz
// +build gofuzz

package fuzz

func Fuzz(data []byte) int {
	str := "12345"
	if HasInjectionOperator(str) != false {
		return 0
	}
	return 1
}
