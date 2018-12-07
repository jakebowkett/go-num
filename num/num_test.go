package num

import (
	"testing"
)

func TestRoman(t *testing.T) {

	cases := []struct {
		n       int
		want    string
		wantErr bool
	}{
		{-1, "", true},
		{0, "", true},
		{1, "I", false},
		{3, "III", false},
		{4, "IV", false},
		{5, "V", false},
		{9, "IX", false},
		{11, "XI", false},
		{294, "CCXCIV", false},
		{442, "CDXLII", false},
		{900, "CM", false},
		{1989, "MCMLXXXIX", false},
		{4859, "MMMMDCCCLIX", false},
	}

	for _, c := range cases {
		if got, err := Roman(c.n); got != c.want || err == nil && c.wantErr {

			errStr := "nil"
			if c.wantErr {
				errStr = "error"
			}

			t.Errorf("Roman(%d)\n"+
				"    return %q, %v\n"+
				"    wanted %q, %s\n",
				c.n, got, err, c.want, errStr)
		}
	}
}

func TestAlpha(t *testing.T) {

	cases := []struct {
		n       int
		want    string
		wantErr bool
	}{
		{-1, "", true},
		{0, "A", false},
		{3, "D", false},
		{4, "E", false},
		{52, "BA", false},
	}

	for _, c := range cases {
		if got, err := Alpha(c.n); got != c.want || err == nil && c.wantErr {

			errStr := "nil"
			if c.wantErr {
				errStr = "error"
			}

			t.Errorf("Alpha(%d)\n"+
				"    return %q, %v\n"+
				"    wanted %q, %s\n",
				c.n, got, err, c.want, errStr)
		}
	}
}

func TestEncode(t *testing.T) {

	cases := []struct {
		n       int
		enc     string
		want    string
		wantErr bool
	}{

		{0, "世界世", "", true}, // Duplicate character in encoding.

		{5, "A", "", true},

		{0, "世界", "世", false},
		{1, "世界", "界", false},
		{2, "世界", "界世", false},
		{3, "世界", "界界", false},
		{4, "世界", "界世世", false},

		{-1, "世界地球風火災水稲妻太陽", "", true},
		{0, "世界地球風火災水稲妻太陽", "世", false},
		{4, "世界地球風火災水稲妻太陽", "風", false},
		{13, "世界地球風火災水稲妻太陽", "界界", false},

		{-1, "0123456789", "", true},
		{0, "0123456789", "0", false},
		{1, "0123456789", "1", false},
		{10, "0123456789", "10", false},
		{11, "0123456789", "11", false},
		{100, "0123456789", "100", false},
		{298648, "0123456789", "298648", false},

		{2, "!@#$%^&*()", "#", false},
		{11, "!@#$%^&*()", "@@", false},
		{99, "!@#$%^&*()", "))", false},
		{67427, "!@#$%^&*()", "&*%#*", false},

		// Emojis.
		{2, "😀😁😂🤣😄😅", "😂", false},
		{6, "😀😁😂🤣😄😅", "😁😀", false},
	}

	for _, c := range cases {
		if got, err := Encode(c.n, c.enc); got != c.want || err == nil && c.wantErr {

			errStr := "nil"
			if c.wantErr {
				errStr = "error"
			}

			t.Errorf("Encode(%d, %q)\n"+
				"    return %q, %v\n"+
				"    wanted %q, %s\n",
				c.n, c.enc, got, err, c.want, errStr)
		}
	}
}
