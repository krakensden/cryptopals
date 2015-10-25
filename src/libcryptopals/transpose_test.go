package libcryptopals

import "testing"

func TestTransposeBlocks(t *testing.T) {
	output := TransposeBlocks([]byte{0, 1, 2, 3, 4, 5}, 2)

	for i, ar := range [][]byte{[]byte{0, 2, 4}, []byte{1, 3, 5}} {
		for j, v := range ar {
			if output[i][j] != v {
				t.Errorf("%d should equal %d\n", output[i][j], v)
			}
		}
	}
}
