package numtoword

import "testing"

func TestArConverter(t *testing.T) {

	samples := []struct {
		in  uint
		out string
	}{
		{0, ""},
		{1, ""},
		{6, ""},
		{9, ""},
		{15, "خمسة عشر"},
		{22, "إثنان و عشرون"},
		{140, "أربعون"},
		{240, "أربعون"},

		{2361, "ستة و ثلاثون"},
		{1000000, "ستة و ثلاثون"},

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
