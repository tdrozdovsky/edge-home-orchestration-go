//go:build gofuzz
// +build gofuzz

package fuzz

import (
	"github.com/lf-edge/edge-home-orchestration-go/internal/common/commandvalidator/injectionchecker"
)

func Fuzz(data []byte) int {
	str := "12345"
	if HasInjectionOperator(str) != false {
		return 0
	}
	return 1
}
