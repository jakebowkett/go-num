/*
Package num provides functions that create different
representations of integers (such as Roman numerals).
*/
package num

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	kiB float64 = 1024
	miB float64 = kiB * 1024
	giB float64 = miB * 1024
)

/*
Bytes takes n representing a number of bytes and produces
a human-friendly string that chooses the most compact form
possible. If there is a fractional unit it will be represented
with one position after the decimal point.

Since the output of Bytes is intended to be read by non-programmers
it notates its measurements as GB, MB, and KB (e.g. "50KB") rather
than GiB, MiB, and KiB respectively. However, when calculating
the output Bytes does interpret 1024 bytes as a kilobyte and so on.

In short, Bytes is not intended to be highly precise. Its output
is intended to be read by users performing tasks such as uploading
images.

	println(num.Bytes(70000000000)) // 65.2GB
	println(num.Bytes(6000000000))  // 5.6GB
	println(num.Bytes(500000000))   // 476.8MB
	println(num.Bytes(30000000))    // 28.6MB
	println(num.Bytes(2000000))     // 2MB
	println(num.Bytes(300000))      // 293KB
	println(num.Bytes(20000))       // 19.5KB
	println(num.Bytes(1000))        // 1KB
	println(num.Bytes(600))         // 600B
	println(num.Bytes(10))          // 10B
*/
func Bytes(n int) string {
	var unit string
	threshold := 0.95
	f := float64(n)
	switch {
	case f/giB > threshold:
		unit = "GB"
		f /= giB
	case f/miB > threshold:
		unit = "MB"
		f /= miB
	case f/kiB > threshold:
		unit = "KB"
		f /= kiB
	default:
		return fmt.Sprintf("%dB", n)
	}
	fractional := f - float64(int(f))
	if fractional < 0.1 || fractional > 0.9 {
		return fmt.Sprintf("%.0f%s", math.Round(f), unit)
	}
	return fmt.Sprintf("%.1f%s", f, unit)
}

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

func WordFloat(f float64, precision int) string {

	// Get words for whole number then remove
	// that part of the number.
	s := Word(int(f))
	f -= float64(int(f))

	if precision > 0 {
		s += " point"
	}

	for i := 0; i < precision; i++ {

		// Multiply fractional part til first number is a whole number.
		n := int(f * 10)

		// Remove the first number from fractional part.
		f = f*10 - float64(n)

		// Add word for first number in fractional part.
		s += " " + Word(n)
	}

	return s
}

/*
Word takes n and returns an English language rendition of it.

	Word(7232) // "seven thousand two hundred and thirty-two"
	Word(-5)   // "negative five"

*/
func Word(n int) string {

	// Zero screws with our logic so we handle it here.
	if n == 0 {
		return "zero"
	}

	/*
		We have to record n's sign here because we
		modify n below. We also force it to be a
		positive number so we can use the same logic
		for negative or positive.
	*/
	var s string
	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}

	type unit struct {
		number int
		word   string
	}
	units := []unit{
		{1000000000, "billion"},
		{1000000, "million"},
		{1000, "thousand"},
		{100, "hundred"},
		{90, "ninety"},
		{80, "eighty"},
		{70, "seventy"},
		{60, "sixty"},
		{50, "fifty"},
		{40, "fourty"},
		{30, "thirty"},
		{20, "twenty"},
		{19, "nineteen"},
		{18, "eighteen"},
		{17, "seventeen"},
		{16, "sixteen"},
		{15, "fifteen"},
		{14, "fourteen"},
		{13, "thirteen"},
		{12, "twelve"},
		{11, "eleven"},
		{10, "ten"},
		{9, "nine"},
		{8, "eight"},
		{7, "seven"},
		{6, "six"},
		{5, "five"},
		{4, "four"},
		{3, "three"},
		{2, "two"},
		{1, "one"},
	}

	for _, u := range units {

		instances := n / u.number
		n %= u.number

		if instances == 0 {
			continue
		}

		/*
			If we've already got preceding words and there's
			no trailing hyphen we should add "and" before
			numbers less than 100 - e.g. two hundred and five,
			six thousand and eighty-four, etc. Regardless, we
			always add a trailing space.
		*/
		if len(s) > 0 && !strings.HasSuffix(s, "-") {
			if u.number < 100 {
				s += " and"
			}
			s += " "
		}

		if instances == 1 {

			/*
				Single instances of "hundred" and greater units
				("thousand", etc) need to be prefixed with "one"
				- e.g. one hundred, one thousand, etc.
			*/
			if u.number >= 100 {
				s += "one "
			}

			// Add the actual word.
			s += u.word

			/*
				If there's still more of n left and the number
				we're currently dealing with is less than 100
				we need a hyphen - e.g. sixty-nine.
			*/
			if u.number < 100 && n > 0 {
				s += "-"
			}

			continue
		}

		/*
			If there are multiple instances of the unit number -
			e.g. in 2,400,000 there are two instances of the unit
			"million" - we recurse to get the word for the number
			of instances.
		*/
		s += Word(instances) + " " + u.word
	}

	/*
		We prefix "negative" right before returning
		otherwise it messes with the conditionals
		that decide when to add "and" between words.
	*/
	if negative {
		s = "negative " + s
	}

	return s
}

/*
Alpha converts n to a base 52 string where each numeral
is represented by an upper or lower case alphabet character.
Returns an error if n is negative.

	s, _ := Alpha(-1) // Error; n is negative.
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
	s, _ = Encode(2, "")             // Error; encoding is an empty string.
	s, _ = Encode(5, "A")            // Error; encoding contains < 2 characters.
	s, _ = Encode(-1, "01123")       // Error; encoding contains duplicates.

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

	enc := strings.Split(encoding, "")
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
