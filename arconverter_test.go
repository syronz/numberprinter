package numtoword

import "testing"

func TestArConverter(t *testing.T) {

	samples := []struct {
		in  float64
		out string
	}{
		// {0, ""},
		// {1.266, ""},
		// {6.11, ""},
		// {9.01, ""},
		// {15, "خمسة عشر"},
		// {22, "إثنان و عشرون"},
		// {140, "أربعون"},
		// {240, "أربعون"},

		// {2361, "ستة و ثلاثون"},
		{12222220, "ستة و ثلاثون"},

		// {100,000,000,000, "one hundred billion"},
	}

	for _, v := range samples {
		r := ArConverter(v.in)
		t.Log(r)
		if r != v.out {
			t.Errorf("for %v it should be %q but it is %q", v.in, v.out, r)
		}
	}

}
