package main

import "testing"

func Test_isPrime(t *testing.T) {

	primeTests := []struct{
		name string
		testNum int
		expected bool
		msg string
	}{
		{"prime",7,true,"7 is a prime number!"},
	}

	for _,e := range primeTests {

		result,msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s: expected true but got false",e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true",e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s",e.name,e.msg,msg)
		}
	}
	
}