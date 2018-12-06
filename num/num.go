/*
Package num converts base-10 integers to other bases.
*/
package num

import (
	"errors"
	"fmt"
	"golang.org/x/text/unicode/norm"
)

/*
Roman converts n to a roman numeral of type string.

Keep in mind there is no roman numeral for zero and that
the largest value a single roman numeral can represent is
M (1000). Therefore an n over a few thousand will produce
results like "MMMMMMMCDXII" (7412) which may be undesired.

	s1, _ := Roman(0)    // This will return an error.
	s2, _ := Roman(1)    // s2 is "I"
	s3, _ := Roman(4)    // s3 is "IV"
	s4, _ := Roman(467)  // s4 is "CDLXVII"
	s5, _ := Roman(1991) // s5 is "MCMXCI"

Returns an error if n is less than one.
*/
func Roman(n int) (string, error) {
	if n == 0 {
		return "", errors.New("Input cannot be zero.")
	}
	if n < 0 {
		return "", errors.New(
			fmt.Sprintf("Input cannot be a negative number. Got %d.", n))
	}
	type multiple struct {
		number int
		letter string
	}
	multiples := []multiple{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}
	s := ""
	for _, m := range multiples {
		s += repeat(m.letter, n/m.number)
		n %= m.number
	}
	return s, nil
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

/*
Alpha converts n to a base 52 string where each numeral
is represented by an upper or lower case alphabet character.

	s1, _ := Alpha(0)  // s1 is "A"
	s2, _ := Alpha(25) // s2 is "Z"
	s3, _ := Alpha(52) // s3 is "AA"

Returns an error if n is negative.
*/
func Alpha(n int) (string, error) {
	encoding := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return Encode(n, encoding)
}

/*
Encode converts n to a string that uses the characters
in encoding as its numerals. Multi-byte characters such
as kanji, emojis, and so on will be treated as a single
character. The base of the result will be determined by
the number of characters (not bytes) in encoding.

	s1, _ := Encode(0, "世界") // s1 is "世"
	s2, _ := Encode(1, "世界") // s2 is "界"
	s3, _ := Encode(2, "世界") // s3 is "世世"

Returns an error if n is negative.
*/
func Encode(n int, encoding string) (string, error) {

	if n < 0 {
		return "", errors.New(
			fmt.Sprintf("Input number cannot be negative. Got %d.", n))
	}

	n++ // Our loop below won't begin if n was 0.
	result := ""
	quotient := n
	remainder := 0

	// Get a slice of the characters (rather than the bytes) in encoding.
	enc := chars(encoding)
	length := len(enc)

	for quotient != 0 {
		decremented := quotient - 1 // Compensate for zero-based indexing.
		quotient = decremented / length
		remainder = decremented % length
		result = enc[remainder] + result
	}

	return result, nil
}

func chars(s string) []string {
	var it norm.Iter
	it.InitString(norm.NFC, s)
	cc := []string{}
	for !it.Done() {
		cc = append(cc, string(it.Next()))
	}
	return cc
}
