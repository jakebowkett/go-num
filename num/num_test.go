package num

import (
	"testing"
)

func TestRoman(t *testing.T) {

	cases := []struct {
		n        int
		want     string
		want_err bool
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
		if got, err := Roman(c.n); got != c.want || err == nil && c.want_err {
			err_str := "nil"
			if c.want_err {
				err_str = "error"
			}
			t.Errorf("Roman(\"%d\")\n"+
				"    return \"%s\", %s\n"+
				"    wanted \"%s\", %s\n",
				c.n, got, err, c.want, err_str)
		}
	}
}

func TestAlpha(t *testing.T) {
	cases := []struct {
		n        int
		want     string
		want_err bool
	}{
		{-1, "", true},
		{0, "A", false},
		{3, "D", false},
		{4, "E", false},
		{52, "AA", false},
	}
	for _, c := range cases {
		if got, err := Alpha(c.n); got != c.want || err == nil && c.want_err {
			err_str := "nil"
			if c.want_err {
				err_str = "error"
			}
			t.Errorf("Alpha(\"%d\")\n"+
				"    return \"%s\", %s\n"+
				"    wanted \"%s\", %s\n",
				c.n, got, err, c.want, err_str)
		}
	}
}

func TestEncode(t *testing.T) {
	cases := []struct {
		n        int
		enc      string
		want     string
		want_err bool
	}{
		{-1, "世界地球風火災水稲妻太陽", "", true},
		{0, "世界地球風火災水稲妻太陽", "世", false},
		{4, "世界地球風火災水稲妻太陽", "風", false},
		{12, "世界地球風火災水稲妻太陽", "世世", false},
	}
	for _, c := range cases {
		if got, err := Encode(c.n, c.enc); got != c.want || err == nil && c.want_err {
			err_str := "nil"
			if c.want_err {
				err_str = "error"
			}
			t.Errorf("Encode(\"%d\")\n"+
				"    return \"%s\", %s\n"+
				"    wanted \"%s\", %s\n",
				c.n, got, err, c.want, err_str)
		}
	}
}
