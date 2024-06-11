package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testValue struct {
	arg []string
	ans string
}

func getTests() []testValue {
	testValues := []testValue{
		{
			[]string{
				"9 45000",
			},
			"4 0 5"},
		{
			[]string{
				"20 196000",
			},
			"-1 -1 -1"},
		{
			[]string{
				"1000 1234000",
			},
			// "14 27 959"},
			"26 0 974"},
		{
			[]string{
				"2000 20000000",
			},
			"2000 0 0"},
	}

	return testValues
}

func stubio(inbuf string, f func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()

	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr

	inw.Write([]byte(inbuf))
	inw.Close()

	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw

	outC := make(chan string)
	errC := make(chan string)
	defer close(outC)
	defer close(errC)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, outr)
		outr.Close()
		outC <- buf.String()
	}()

	go func() {
		var buf strings.Builder
		io.Copy(&buf, errr)
		errr.Close()
		errC <- buf.String()
	}()

	f()

	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	errw.Close()

	return <-outC, <-errC
}

func Test_main(t *testing.T) {
	tests := getTests()
	for i, tt := range tests {
		si := strconv.Itoa(i)
		t.Run("Case "+si, func(t *testing.T) {
			ret, err := stubio(strings.Join(tt.arg, " "), main)
			if err != "" {
				t.Errorf("error func: %s ", err)
			}
			ans := fmt.Sprintf("%s\n", tt.ans)
			if ret != ans {
				t.Errorf("Unexpected output: '%s' Need: '%s'", ret, ans)
			}
		})
	}
}
