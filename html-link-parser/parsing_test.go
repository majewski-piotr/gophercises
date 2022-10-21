package hlp

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var ex1 string = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`

func TestIsLink(t *testing.T) {
	doc, _ := html.Parse(strings.NewReader(ex1))
	actualLink := doc.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling
	fakeLink := doc.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling

	if !isLink(actualLink) {
		t.Error("Link not detected: ", actualLink)
	}
	if isLink(fakeLink) {
		t.Error("This should not be marked as link: ", fakeLink)
	}
	if isLink(nil) {
		t.Error("nil accepted as a link1")
	}
}

var ex2 string = `<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>`

func TestGetNestedText(t *testing.T) {
	doc, _ := html.Parse(strings.NewReader(ex2))
	expected := "Gophercises is onGithub!"
	actual := getNestedText(doc.FirstChild.FirstChild.NextSibling.NextSibling.
		FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.
		NextSibling.NextSibling.FirstChild)

	if actual != expected {
		t.Errorf("Strings does't match got: \n%s \nEXPECTED\n%s", actual, expected)
	}
	if getNestedText(nil) != "" {
		t.Errorf("Incorrect result after nil value")
	}

}

func TestGetLinks(t *testing.T) {
	doc, _ := html.Parse(strings.NewReader(ex2))
	links := GetLinks(doc)

	if len(links) != 2 {
		t.Error("Wrong size of links")
	}
	if links[0].Text != "Check me out on twitter" {
		t.Error("Wrong link text, should be Check me out on twitter, got", links[0].Text)
	}
	if links[0].Url != "https://www.twitter.com/joncalhoun" {
		t.Error("Wrong link text, should be got https://www.twitter.com/joncalhoun , got", links[0].Url)
	}
}
