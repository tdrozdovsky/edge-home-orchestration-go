package fuzz

import (
	// "fmt"
	"github.com/tdrozdovsky/edge-home-orchestration-go/internal/common/commandvalidator/injectionchecker"
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
//     testcases := []struct {
//         in, want string
//     }{
//         {"Hello, world", "dlrow ,olleH"},
//         {" ", " "},
//         {"!12345", "54321!"},
//     }
//     for _, tc := range testcases {
//         rev := Reverse(tc.in)
//         if rev != tc.want {
//                 t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

// func Fuzzer(f *testing.F) {
// 	testcases := []string{"Hello, world", " ", "!12345"}
// 	for _, tc := range testcases {
// 		f.Add(tc) // Use f.Add to provide a seed corpus
// 	}
// 	f.Fuzz(func(t *testing.T, str string) {
// 		// str := string(data)
// 		fmt.Println(str)
// 		if HasInjectionOperator(str) != false {
// 			t.Error("unexpected success")
// 		} else {
// 			return
// 		}
// 		// fmt.Println(data2)
// 		// fmt.Println(data3)
// 	})
// }

func FuzzReverse(f *testing.F) {
	// testcases := []string{"Hello, world", " ", "!12345"}
	// for _, tc := range testcases {
	// 	f.Add(tc) // Use f.Add to provide a seed corpus
	// }
	f.Fuzz(func(t *testing.T, data []byte) int {
		orig := string(data)
		rev, err1 := injectionchecker.Reverse(orig)
		if err1 != nil {
			return 1
		}
		doubleRev, err2 := injectionchecker.Reverse(rev)
		if err2 != nil {
			return 1
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
			return 1
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
			return 1
		}
		return 0
	})
}
