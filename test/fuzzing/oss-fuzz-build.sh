#/bin/bash -eu

compile_go_fuzzer github.com/tdrozdovsky/edge-orchestration/test/fuzzing Fuzz fuzz gofuzz
