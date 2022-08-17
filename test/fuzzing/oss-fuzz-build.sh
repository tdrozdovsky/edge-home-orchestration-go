#/bin/bash -eu

compile_go_fuzzer github.com/tdrozdovsky/edge-home-orchestration-go/test/fuzzing Fuzz fuzz gofuzz
