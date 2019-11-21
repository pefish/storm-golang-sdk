package signature

import (
	"fmt"
	"testing"
)

func TestGenerateRandomKeyPair(t *testing.T) {
	fmt.Println(GenerateRandomKeyPair())
}

func TestSignMessage(t *testing.T) {
	//pubKey := `0302e2b8b44fe202ac68f2f74a8a503496fbec3d9c3b686b9f06812e08d9fbbd78`
	privKey := `1b75f9d2651786bbc3da124f2f52fb853ad7f4abadf1c2fac545f0b68d4d29d9`
	fmt.Println(SignMessage(`haha`, privKey))
}

func TestVerifySignature(t *testing.T) {
	pubKey := `0302e2b8b44fe202ac68f2f74a8a503496fbec3d9c3b686b9f06812e08d9fbbd78`
	sig := `3045022100f063bba67553d77f486eac8d89867c5a6d7e61a11e432bc1c71a4f114daabeb902206a88287f43cca1092fffbc2970b0ba15695d672e678dc66eae90c49c9c8c70b5`
	fmt.Println(VerifySignature(`haha`, sig, pubKey))
}
