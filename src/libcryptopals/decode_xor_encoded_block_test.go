package libcryptopals

import (
	"bytes"
	"testing"
)

func TestDecodeBoyz(t *testing.T) {
	text := "To be, or not to be, that is the question: Whether 'tis Nobler in the mind to suffer The Slings and Arrows of outrageous Fortune, Or to take Arms against a Sea of troubles, And by opposing end them: to die, to sleep No more; and by a sleep, to say we end The Heart-ache, and the thousand Natural shocks That Flesh is heir to? 'Tis a consummation Devoutly to be wished. To die, to sleep, To sleep, perchance to Dream; aye, there's the rub, For in that sleep of death, what dreams may come,"

	for _, key := range []string{"a", "XXX", "_-'----", "BOYZINTH", "BOYZINTHa", "fiddledeedee-fiddledeeday"} {

		outhex, err := StringRepeatingXor(text, key)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		outbytes, err := Hex2Byte(outhex)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		decoded := DecodeXorEncodedBlock(outbytes)
		if comp := bytes.Compare(decoded, []byte(text)); comp != 0 {
			t.Errorf("comp came back %d with key %s from \n%s\nvs\n%s\n", comp, key, string(decoded), text)
		}
	}
}
