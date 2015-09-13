package libcryptopals

import (
	"errors"
	"fmt"
)

var hex_char_map map[byte]byte = map[byte]byte{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
	// user friendly!
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
}

var Base64Alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var base64_char_map map[byte]byte = map[byte]byte{}

func init() {
	for i, v := range Base64Alphabet {
		base64_char_map[byte(v)] = byte(i)
	}
	base64_char_map['='] = 0
}

func Byte2Hex(input []byte) string {
	// Each byte becomes a pair of runes
	output := make([]byte, 0, len(input)*2)
	for _, val := range input {
		output = append(output, []byte(fmt.Sprintf("%.2x", val))...)
	}
	return string(output)
}

func Hex2Byte(input string) ([]byte, error) {
	// Each pair of runes encodes one byte. If there's an odd number of runes, assume it's 0x4, not 0x40.

	output := make([]byte, 0, len(input)/2+len(input)%2)
	for i := 0; i < len(input); i += 2 {
		byteval0, ok := hex_char_map[input[i]]
		if !ok {
			return output, errors.New(fmt.Sprintf("invalid character detected at offset %d in %s", i, input))
		}
		if i+1 < len(input) {
			byteval1, ok := hex_char_map[input[i+1]]
			if !ok {
				return output, errors.New(fmt.Sprintf("invalid character detected at offset %d in %s", i+1, input))
			}
			output = append(output, (byteval0<<4)+byteval1)
		} else {
			output = append(output, byteval0)
		}
	}
	return output, nil
}

func Byte2Base64(input []byte) (output string) {
	var max_length int = len(input)
	if len(input)%3 != 0 {
		max_length += 3 - len(input)%3
	}
	var output_builder []byte = make([]byte, 0, max_length)
	var remainder byte = 0       // to hold over bits between bytes
	var sixmask byte = 252       // 0b11111100
	var twomask byte = 192       // 0b11000000
	var topfourmask byte = 240   // 0b11110000
	var bottomfourmask byte = 15 // 0b00001111

	var cur byte

	// three bytes, four base64 characters
	for i := 0; i < max_length; i++ {
		if i >= len(input) {
			cur = 0
		} else {
			cur = input[i]
		}

		switch i % 3 {
		case 0:
			output_builder = append(output_builder, Base64Alphabet[(cur&sixmask)>>2])
			remainder = cur &^ sixmask
		case 1:
			var top byte = (remainder << 4)
			var bot byte = ((cur & topfourmask) >> 4)
			output_builder = append(output_builder, Base64Alphabet[top|bot])
			remainder = cur & bottomfourmask
		case 2:
			if i >= len(input) {
				output_builder = append(output_builder, '=')
			} else {
				var top1 byte = (remainder << 2)
				var bot1 byte = ((cur & twomask) >> 6)

				output_builder = append(output_builder, Base64Alphabet[top1|bot1])
			}
			remainder = 0

			if i > len(input) {
				output_builder = append(output_builder, '=')
			} else {
				output_builder = append(output_builder, Base64Alphabet[cur&^twomask])
			}
		}
	}
	return string(output_builder)
}

func Base642Byte(input []byte) (output []byte, err error) {
	// Each base64 character encodes 6 bits, each byte encodes 8 bits, so
	// for each 4 b64 characters there are 3 bytes
	var max_length int = len(input)
	if len(input)%4 != 0 {
		max_length += 4 + len(input)%4
	}
	var output_builder []byte = make([]byte, 0, max_length)
	var toptwomask byte = 192 //    0b11000000

	var cur, rem byte = 0, 0

	// three bytes, four base64 characters
	for i := 0; i < max_length; i++ {
		if i >= len(input) {
			cur = 0
		} else {
			if val, ok := base64_char_map[input[i]]; ok {
				cur = val
				fmt.Println("input ", input[i], " b64 val ", cur)
			} else {
				return nil, errors.New(fmt.Sprintf("input contains %s, which is not a valid Base64 character", input[i]))
			}
		}
		fmt.Println("i is ", i, " rem is ", rem, " output_builder ", output_builder)

		switch i % 4 {
		// T
		case 0:
			rem = cur << 2
		// W
		case 1:
			fmt.Println("appending ", rem|(cur&toptwomask)>>6)
			output_builder = append(output_builder, rem|(cur&toptwomask)>>6)
			// Save the bottom half in the top half
			rem = cur << 4
		// F
		case 2:
			fmt.Println("appending ", (rem)|(cur>>2))
			// bottom four bits of W and top four bits of F
			output_builder = append(output_builder, (rem)|(cur>>2))
			rem = (cur & 3) << 6
		// u
		case 3:
			fmt.Println("appending ", (rem)|(cur))
			output_builder = append(output_builder, (rem)|(cur))
			rem = 0 // not strictly necessary
		}
	}
	return output_builder, nil
}

func Hex2Base64(input string) (output string, err error) {
	bytestr, err := Hex2Byte(input)
	if err != nil {
		return "", err
	}
	return Byte2Base64(bytestr), nil
}
