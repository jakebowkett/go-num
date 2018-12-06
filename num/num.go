/*
Package num provides functions for creating different
representations for integers (such as Roman numerals).
*/
package num

import (
	"errors"
	"fmt"
	"golang.org/x/text/unicode/norm"
)

/*
Roman converts n to a Roman numeral of type string.
Returns an error if n is less than one.

Keep in mind there is no Roman numeral for zero. Further,
the largest value a single Roman numeral can represent is
M (1000). Therefore an n above a few thousand such as 7412
will produce the output "MMMMMMMCDXII" which may be
undesirable.

	s1, _ := Roman(0)    // This will return an error.
	s2, _ := Roman(1)    // s2 is "I"
	s3, _ := Roman(4)    // s3 is "IV"
	s4, _ := Roman(467)  // s4 is "CDLXVII"
	s5, _ := Roman(1991) // s5 is "MCMXCI"

*/
func Roman(n int) (string, error) {
	if n == 0 {
		return "", errf("Input cannot be 0.")
	}
	if n < 0 {
		return "", errf("Input cannot be a negative number. Got %d.", n)
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
Returns an error if n is negative.

	s, _ := Alpha(0) // "A"
	s, _ = Alpha(25) // "Z"
	s, _ = Alpha(52) // "AA"

*/
func Alpha(n int) (string, error) {
	encoding := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return Encode(n, encoding)
}

/*
Encode converts n to a string that uses the characters
in encoding as its numerals. Returns an error if n is
negative, encoding is an empty string, or if encoding
contains duplicate characters.

Multi-byte characters such as kanji, emojis, and so on
will be treated as a single character. The base of the
result will be determined by the number of characters
(not bytes) in encoding. The first character of encoding
acts as the zero value.

	s, _ := Encode(0, "0123456789") // "0"
	s, _ = Encode(1, "0123456789")  // "1"
	s, _ = Encode(10, "0123456789") // "10"

	s, _ = Encode(0, "世界") // "世"
	s, _ = Encode(1, "世界") // "界"
	s, _ = Encode(2, "世界") // "界世"
	s, _ = Encode(3, "世界") // "界界"
	s, _ = Encode(4, "世界") // "界世世"

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
		return "", errf("Input number cannot be negative. Got %d", n)
	}
	if encoding == "" {
		return "", errf("Encoding cannot be an empty string.")
	}
	enc := chars(encoding)
	if err := unique_set(enc); err != nil {
		return "", err
	}
	if n == 0 {
		return enc[0], nil
	}
	result := ""
	quotient := n
	remainder := 0
	length := len(enc)
	for quotient != 0 {
		decremented := quotient
		quotient = decremented / length
		remainder = decremented % length
		result = enc[remainder] + result
	}
	return result, nil
}

func unique_set(ss []string) error {
	seen := map[string]bool{}
	for _, s := range ss {
		if seen[s] {
			return errf(`"%s" appears multiple times.`, s)
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

func errf(s string, v ...interface{}) error {
	return errors.New(fmt.Sprintf(s, v...))
}
