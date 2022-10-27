package sitemap

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	d := Domain{Name: "elderscrolls.org", Sites: nil}

	if !d.contains("http://www.elderscrolls.org/oblivion") {
		t.Error("http://www.elderscrolls.org/oblivion is from domain elderscrolls")
	}

	if !d.contains("https://www.elderscrolls.org/oblivion") {
		t.Error("https://www.elderscrolls.org/oblivion is from domain elderscrolls")
	}

	if !d.contains("/oblivion") {
		t.Error("/oblivion is from domain elderscrolls")
	}

	if d.contains("http://www.elderscrolls.oblivion/badfaces") {
		t.Error("http://www.elderscrolls.oblivion/badfaces isn't from domain elderscrolls")
	}
}

func TestContainsError(t *testing.T) {
	//given
	d := Domain{Name: "elderscrolls.org", Sites: nil}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	//when
	d.contains("HTTascxsvcvsdd?#$%!aax")
	actual := buf.String()[20:]

	//then
	expected := `during parsing url of HTTascxsvcvsdd?#$%!aax occured error:parse "HTTascxsvcvsdd?#$%!aax": invalid URL escape "%!a"`
	if strings.EqualFold(expected, actual) {
		t.Errorf("Wrong err log, should be \n%v\ngot\n%v", expected, actual)
	}
}
