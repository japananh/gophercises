package link

import (
	"os"
	"testing"
)

type parseTest struct {
	filename string
	expected []Link
	err      error
}

var parseTests = []parseTest{
	{filename: "ex1.html", expected: []Link{
		{Href: "/other-page", Text: "A link to another page"},
	}, err: nil},
	{filename: "ex2.html", expected: []Link{
		{Href: "https://www.twitter.com/joncalhoun", Text: "        Check me out on twitter\n            "},
		{Href: "https://github.com/gophercises", Text: "        Gophercises is on Github!\n    "},
	}, err: nil},
	{filename: "ex3.html", expected: []Link{
		{Href: "#", Text: "Login "},
		{Href: "/lost", Text: "Lost? Need help?"},
		{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
	}, err: nil},
	{filename: "ex4.html", expected: []Link{
		{Href: "/dog-cat", Text: "dog cat "},
	}, err: nil},
}

func TestParse(t *testing.T) {
	for i, test := range parseTests {
		s, err := os.Open(test.filename)
		if err != nil {
			t.Errorf("Error when open file %s", test.filename)
			continue
		}

		output, err := Parse(s)

		if len(output) != len(test.expected) {
			t.Errorf("Testcase %d: The number of output %d not equal to expected %d", i, len(output), len(test.expected))
			continue
		}

		for i, link := range output {
			if link.Href != test.expected[i].Href {
				t.Errorf("Output `Href` %q not equal to expected %q", link.Href, test.expected[i].Href)
			}
			if link.Text != test.expected[i].Text {
				t.Errorf("Output `Text` %q not equal to expected %q", link.Text, test.expected[i].Text)
			}
		}

		if err != nil && err != test.err {
			t.Errorf("Output %q not equal to expected %q", err, test.err)
		}
	}
}
