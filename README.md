# lcdflags

this version uses flags for the input data and command line feedback
there will be an lcdcobra version that does the same functions but use cobra instead of flags
input two integers and it will calculate the lcd


DEBUG=true go run main.go
DEBUG=true go run main.go -n 500 -m 60
DEBUG=true go run main.go -n 10 -m 100 -preview -prompt


DEBUG=true go run main.go -greeting "Hello" -name "Keith" -preview -prompt

go mod init github.com/zanzeeba/lcdflags