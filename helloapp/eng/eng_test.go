package eng

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello"
	if got := Greet(); got != want {
		t.Errorf("Greet() got %q. want : %q", got, want)
	}
}

func TestGreet2(t *testing.T) {
	want := "Hello, Hello"
	if got := Greet2(); got != want {
		t.Errorf("Greet2() = %q, want : %q", got, want)
	}
}

func TestGreet3(t *testing.T) {
	want := "Hello, Gopher!"
	if got := Greet3(); got != want {
		t.Errorf("Greet3() = %q, want : %q", got, want)
	}
}
