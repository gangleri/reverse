package reverse

import "testing"

func TestReverse(t *testing.T) {
	expected := "olleH"
	actual := Reverse("Hello")
	if actual != expected {
		t.Fatal("Failed to reverse the string")
	}
}
