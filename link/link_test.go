package link

import (
	"io"
	"testing"
)

type parseTest struct {
	arg      io.Reader
	expected []Link
	err error
}

var parseTests = []parseTest{
	{arg: nil, expected: nil, err: nil},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		output, err := parse(test.arg)

		if len(output) != len(test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

		for i, link := range output {
			if link.Text != test.expected[i].Text || link.Href != test.expected[i].Href {
				t.Errorf("Output %q not equal to expected %q", link, test.expected[i])
			}
		}

		if err != test.err {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
