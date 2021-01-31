package numtoword

import "testing"

func TestEnConverter(t *testing.T) {

	samples := []struct {
		in  float64
		out string
	}{
		//{1052, "thousand and fifty two"},
		{1.546, "one"},
		{6, "six"},
		{9, "nine"},
		{15, "fifteen"},
		{25, "twenty and five"},
		{22, "twenty and two"},
		{40, "forty"},
		{36, "thirty and six"},
		{71, "seventy and one"},
		{15, "fifteen"},
		{106, "one hundred and six"},
		{100, "one hundred"},
		{223, "two hundred and twenty and three"},
		{100000000000, "one hundred billion"},
		{100000000001, "one hundred billion and one"},
		{3522, "three thousand and five hundred and twenty and two"},
		{6372428, "six million and three hundred and seventy and two thousand and four hundred and twenty and eight"},
		{1002003004, "one billion and two million and three thousand and four"},
		// {100,000,000,000, "one hundred billion"},
	}

	for _, v := range samples {
		r := EnConverter(v.in)
		t.Log(r)
		if r != v.out {
			t.Errorf("for %v it should be %q but it is %q", v.in, v.out, r)
		}
	}

}
