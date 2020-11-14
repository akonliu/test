package test

import "testing"

func Test_TestHelper(t *testing.T) {
	TestHelper()
	t.Logf("pass")
}

func Test_TestSayHello(t *testing.T) {
	SayHello("ak")
	t.Logf("pass")
}
