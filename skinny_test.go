package skinny

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var tests = []struct {
	key    string
	plain  string
	cipher string
}{
	{
		/* Skinny -64 -64 */
		`f5269826fc681238`,
		`06034f957724d19d`,
		`bb39dfb2429b8ac7`,
	},
	{
		// Skinny -64 -128
		`9eb93640d088da6376a39d1c8bea71e1`,
		`cf16cfe8fd0f98aa`,
		`6ceda1f43de92b9e`,
	},

	{
		// Skinny -64 -192
		`ed00c85b120d68618753e24bfd908f60b2dbb41b422dfcd0`,
		`530c61d35e8663c3`,
		`dd2cf1a8f330303c`,
	},
	{
		// Skinny -128 -128
		`4f55cfb0520cac52fd92c15f37073e93`,
		`f20adb0eb08b648a3b2eeed1f0adda14`,
		`22ff30d498ea62d7e45b476e33675b74`,
	},
	{
		// Skinny -128 -256
		`009cec81605d4ac1d2ae9e3085d7a1f31ac123ebfc00fddcf01046ceeddfcab3`,
		`3a0c47767a26a68dd382a695e7022e25`,
		`b731d98a4bde147a7ed4a6f16b9b587f`,
	},
	{
		// Skinny -128 -384
		`df889548cfc7ea52d296339301797449ab588a34a47f1ab2dfe9c8293fbea9a5ab1afac2611012cd8cef952618c3ebe8`,
		`a3994b66ad85a3459f44e92b08f550cb`,
		`94ecf589e2017c601b38c6346a10dcfa`,
	},
}

func TestSKINNY(t *testing.T) {

	for i, tt := range tests {
		k, _ := hex.DecodeString(tt.key)
		p, _ := hex.DecodeString(tt.plain)
		c, _ := hex.DecodeString(tt.cipher)

		t.Log(p, k, i)
		Encrypt(p, k, i)

		if !bytes.Equal(p, c) {
			t.Errorf("enc(%v,%v,%v)=%v, want %v", tt.plain, tt.key, i, p, tt.cipher)
		}

		p, _ = hex.DecodeString(tt.plain)

		t.Log(c, k, i)
		Decrypt(c, k, i)

		if !bytes.Equal(p, c) {
			t.Errorf("dec(%v,%v,%v)=%v, want %v", tt.key, tt.cipher, i, p, tt.plain)
		}

	}
}
