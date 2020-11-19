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

func Test_TestSayHelloBack(t *testing.T) {
	s := SayHelloBack("akonliu")
	t.Logf("pass %v", s)
}
