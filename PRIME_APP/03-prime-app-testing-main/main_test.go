package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
	"bytes"
)

func Test_isPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 9, false, "9 is not a prime number because it is divisible by 3"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"negative number", -5, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {

		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

}

func Test_prompt(t *testing.T) {

	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}

}

func Test_intro(t *testing.T) {

	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("incorrect prompt: expected Enter a whole number but got %s", string(out))
	}

}

func Test_checkNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "negative number", input: "-5", expected: "Negative numbers are not prime, by definition!"},
		{name: "prime number", input: "7", expected: "7 is a prime number!"},
		{name: "non prime number", input: "8", expected: "8 is not a prime number because it is divisible by 2"},
		{name: "decimal", input:"0.4", expected:"Please enter a whole number"},
		{name:"quit", input:"q", expected: ""},
		{name:"QUIT",input:"Q",expected:""},
	}

	for _, e := range tests {

		// here we are creating a reader which reads from usually a file. but in this case string
		input := strings.NewReader(e.input)

		// the NewScanner returns a scanner which can read from any type which implements a reader interface
		// io.Stdout is one of them and also here it will read from the string rather than from command line which
		// is done by the actual program
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s",e.name,e.expected,res)
		}
	}

}


func Test_readUserInput(t *testing.T) {

	// to test this channel we need a channel, and an instance of io.Reader
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin,doneChan)
	<- doneChan
	close(doneChan)


}
