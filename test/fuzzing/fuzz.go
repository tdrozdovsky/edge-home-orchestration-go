//go:build gofuzz
// +build gofuzz

package fuzz

import "github.com/tdrozdovsky/edge-home-orchestration-go/internal/common/commandvalidator/injectionchecker"

func Fuzz(data []byte) int {
	str := "12345"
	if injectionchecker.HasInjectionOperator(str) != false {
		return 0
	}
	return 1
}
