package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int64
	var m int64
	var ns string
	var ms string
	var steps bool
	var prompt bool
	var preview bool

	// Define flags
	flag.Int64Var(&n, "n", 0, "the n value as an integer")
	flag.Int64Var(&m, "m", 0, "the m value as an integer")
	flag.StringVar(&ns, "ns", "", "the n as ns value as an integer if you want prompting for the input")
	flag.StringVar(&ms, "ms", "", "the m as ms value as an integer if you want prompting for the input")
	flag.BoolVar(&steps, "steps", false, "if true will show you each step of the calc")
	flag.BoolVar(&prompt, "prompt", false, "get asked for the values instead of entering them in the command")
	flag.BoolVar(&preview, "preview", false, "you enter two numbers n and m. the program then "+
		"calculates the lowest common denominater for those two numbers")

	// Parse flags
	flag.Parse()

	// Optionally print flags and exit if DEBUG is set
	if os.Getenv("DEBUG") != "" {
		fmt.Println("n:", n)
		fmt.Println("m:", m)
		fmt.Println("ns:", ns)
		fmt.Println("ms:", ms)
		fmt.Println("steps:", steps)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)
		os.Exit(0)
	}

	// Conditionally read from stdin
	if prompt {
		ms, ns = renderPrompt()
		fmt.Println("ms:", ms)
		fmt.Println("ns:", ns)
		n, _ := strconv.Atoi(ns)
		m, _ := strconv.Atoi(ms)
		fmt.Println("n:", n)
		fmt.Println("m:", m)
	}
	remainder := calcLcd(n, m)
	for remainder != 0 {
		n = m
		m = remainder
		remainder = calcLcd(n, m)
	}
	fmt.Print("lcd is: ", m)
}

func calcLcd(n, m int64) int64 {
	r := n % m

	fmt.Println("r:", r)
	return r
}

func renderPrompt() (ns, ms string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Integer n: ")
	ns, _ = reader.ReadString('\n')
	ns = strings.TrimSpace(ns)

	fmt.Print("Integer m: ")
	ms, _ = reader.ReadString('\n')
	ms = strings.TrimSpace(ms)
	//n, err := strconv.Atoi(ns)
	//m, err := strconv.Atoi(ms)
	return
}
