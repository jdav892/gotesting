package dependancyInjection

import(
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Justin")

	got := buffer.String()
	want := "Hello, Justin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
