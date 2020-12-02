package main

import "testing"

type Fixture struct {
	policyVal1 int
	policyVal2 int
	c          byte
	password   []byte
	expected   bool
}

func TestIsPasswordValidPolicy1(t *testing.T) {
	fixtures := []Fixture{
		{1, 3, 'a', []byte("abcde"), true},
		{1, 3, 'b', []byte("cdefg"), false},
		{2, 9, 'c', []byte("ccccccccc"), true},
	}

	for _, fixture := range fixtures {
		if isPasswordValid(fixture.policyVal1, fixture.policyVal2, fixture.c, fixture.password) != fixture.expected {
			t.Error(fixture)
		}
	}
}

func TestIsPasswordValidPolicy2(t *testing.T) {
	fixtures := []Fixture{
		{1, 3, 'a', []byte("abcde"), true},
		{1, 3, 'b', []byte("cdefg"), false},
		{2, 9, 'c', []byte("ccccccccc"), false},
	}

	for _, fixture := range fixtures {
		if isPasswordValid2(fixture.policyVal1-1, fixture.policyVal2-1, fixture.c, fixture.password) != fixture.expected {
			t.Error(fixture)
		}
	}
}
