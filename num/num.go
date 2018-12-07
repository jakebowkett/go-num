/*
Package num provides functions for creating different
representations for integers (such as Roman numerals).
*/
package num

import (
	"errors"
	"fmt"
	"golang.org/x/text/unicode/norm"
	"strings"
)

/*
Roman converts n to a Roman numeral of type string.
Returns an error if n is less than one.

Keep in mind there is no Roman numeral for zero. Further,
the largest value a single Roman numeral can represent is
M (1000). Therefore an n above a few thousand such as 7412
will produce the output "MMMMMMMCDXII" which may be
undesirable.

	s, _ := Roman(-1)  // Error; n is negative.
	s, _ = Roman(0)    // Error; no Roman numeral for zero.
	s, _ = Roman(1)    // "I"
	s, _ = Roman(4)    // "IV"
	s, _ = Roman(467)  // "CDLXVII"
	s, _ = Roman(1991) // "MCMXCI"

*/
func Roman(n int) (string, error) {

	if n == 0 {
		return "", errors.New("Input cannot be 0.")
	}

	if n < 0 {
		return "", fmt.Errorf("Input cannot be a negative number. Got %d.", n)
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

	var s string
	for _, m := range multiples {
		s += strings.Repeat(m.letter, n/m.number)
		n %= m.number
	}

	return s, nil
}

/*
Alpha converts n to a base 52 string where each numeral
is represented by an upper or lower case alphabet character.
Returns an error if n is negative.

	s, _ := Alpha(-1) // Error; n in negative.
	s, _ = Alpha(0)   // "A"
	s, _ = Alpha(25)  // "Z"
	s, _ = Alpha(52)  // "AA"

*/
func Alpha(n int) (string, error) {
	const encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return Encode(n, encoding)
}

/*
Encode converts n to a string that uses the characters
in encoding as its numerals. Returns an error if: n is
negative, encoding is an empty string, encoding contains
less than two characters, or encoding contains duplicate
characters.

Multi-byte characters such as kanji, emojis, and so on
will be treated as a single character. The base of the
result will be determined by the number of characters
(not bytes) in encoding.

The first character of encoding acts as the zero value.
This means the encoding string must contain at least
two characters.

	s, _ := Encode(-1, "0123456789") // Error; n is negative.
	s, _ = Encode(2, "") 			 // Error; encoding is an empty string.
	s, _ = Encode(5, "A") 			 // Error; encoding contains < 2 characters.
	s, _ = Encode(-1, "01123") 		 // Error; encoding contains duplicates.

	s, _ = Encode(0, "0123456789")  // "0"
	s, _ = Encode(1, "0123456789")  // "1"
	s, _ = Encode(10, "0123456789") // "10"

	s, _ = Encode(0, "世界") // "世"
	s, _ = Encode(1, "世界") // "界"
	s, _ = Encode(2, "世界") // "界世"
	s, _ = Encode(3, "世界") // "界界"
	s, _ = Encode(4, "世界") // "界世世"

	s, _ = Encode(2, "😀😁😂🤣😄😅") // "😂"
	s, _ = Encode(6, "😀😁😂🤣😄😅") // "😁😀"

	s, _ = Encode(2, "!@#$%^&*()")     // "#"
	s, _ = Encode(11, "!@#$%^&*()")    // "@@"
	s, _ = Encode(67427, "!@#$%^&*()") // "&*%#*"

*/
func Encode(n int, encoding string) (string, error) {

	if n < 0 {
		return "", fmt.Errorf("Input number cannot be negative. Got %d", n)
	}

	if encoding == "" {
		return "", errors.New("Encoding cannot be an empty string.")
	}

	enc := chars(encoding)
	if err := uniqueSet(enc); err != nil {
		return "", err
	}
	if n == 0 {
		return enc[0], nil
	}

	length := len(enc)
	if length == 1 {
		return "", errors.New("Encoding must have at least two characters.")
	}

	var result string
	var remainder int
	quotient := n

	for quotient != 0 {
		decremented := quotient
		quotient = decremented / length
		remainder = decremented % length
		result = enc[remainder] + result
	}

	return result, nil
}

func uniqueSet(ss []string) error {
	seen := make(map[string]bool, len(ss))
	for _, s := range ss {
		if seen[s] {
			return fmt.Errorf("%q appears multiple times.", s)
		}
		seen[s] = true
	}
	return nil
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
