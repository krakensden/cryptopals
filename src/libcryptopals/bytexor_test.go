package libcryptopals

import "testing"

func TestByteXor(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	abyte, _ := Hex2Byte(a)
	bbyte, _ := Hex2Byte(b)
	final, _ := SliceXor(abyte, bbyte)
	desired, _ := Hex2Byte("746865206b696420646f6e277420706c6179")

	if len(final) != len(desired) {
		t.Errorf("!!! lengths don't match")
	}
}
