#/bin/bash -eu

# go mod init github.com/tdrozdovsky/edge-home-orchestration-go/test/fuzzing
#compile_go_fuzzer github.com/tdrozdovsky/edge-home-orchestration-go/internal/common/commandvalidator/injectionchecker Fuzz fuzz gofuzz
compile_go_fuzzer github.com/tdrozdovsky/edge-home-orchestration-go/test/fuzzing Fuzz fuzz gofuzz
