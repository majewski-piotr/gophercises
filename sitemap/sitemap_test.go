package sitemap

import "testing"

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
