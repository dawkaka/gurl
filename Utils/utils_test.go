package utils
import "testing"


func TestParseCookies(t *testing.T) {
  got := ParseCookies("foo=bar;mo=on;big = small")
  want := map[string]string{foo:"bar",mo:"on",big:"small"}
  if got != want {
    t.ErrorF("ParseCookies returns %q, want %q",got, want)
  }
}

func TestParseHeaders(t *testing.T) {
  got,_ := ParseHeaders("foo:bar,mo:on,big : small")
  want,_ := map[string]string{foo:"bar",mo:"on",big:"small"}
  if got != want {
    t.ErrorF("ParseCookies returns %q, want %q",got, want)
  }
}
