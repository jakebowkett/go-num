// Package num provides some simple functions for converting
// base-10 integers to other bases.
package num

import (
	"errors"
	"fmt"
	"golang.org/x/text/unicode/norm"
)

// Roman converts n to a roman numeral of type string.
// An error will be returned if n is less than one.
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

// Alpha converts n to a base 52 string where each numeral
// is represented by an upper or lower case alphabet character.
// An error will be returned if n is negative.
func Alpha(n int) (string, error) {
	encoding := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return Encode(n, encoding)
}

// Encode converts n to a string that uses the characters
// in encoding as its basis. The base of the result will
// be determined by the number of characters (not bytes)
// in encoding.

// Adapted from: http://bideowego.com/base-26-conversion
func Encode(n int, encoding string) (string, error) {

	if n < 0 {
		return "", errors.New(
			fmt.Sprintf("Input number cannot be negative. Got %d.", n))
	}

	result := ""
	n++

	quotient := n
	var remainder int

	// Get number of characters in encoding rather than
	// the number of bytes.
	enc := chars(encoding)
	length := len(enc)

	// Until quotient of 0.
	for quotient != 0 {

		// Compensate for zero-based indexing.
		decremented := quotient - 1

		// Divide by our encoding length.
		quotient = decremented / length

		// Get remainder.
		remainder = decremented % length

		// Prepend letter at index of remainder.
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
