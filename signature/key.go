package signature

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

func GenerateRandomKeyPair() (string, string) {
	privKeyBytes := make([]byte, 32)
	if _, err := rand.Read(privKeyBytes); err != nil {
		panic(err)
	}
	privKeyObj, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	pubkey := fmt.Sprintf("%x", privKeyObj.PubKey().SerializeCompressed())
	privKeyStr := fmt.Sprintf("%x", privKeyBytes)
	return pubkey, privKeyStr
}

func Hash256(s string) string {
	hash_result := sha256.Sum256([]byte(s))
	hash_string := string(hash_result[:])
	return hash_string
}

func SignMessage(message string, privKey string) string {
	privKeyBytes, _ := hex.DecodeString(privKey)
	privKeyObj, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	sig, _ := privKeyObj.Sign([]byte(Hash256(Hash256(message))))
	return fmt.Sprintf("%x", sig.Serialize())
}

func VerifySignature(message string, signature string, pubKey string) bool {
	pubKeyBytes, _ := hex.DecodeString(pubKey)
	pubKeyObj, _ := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256(Hash256(message))), pubKeyObj)
	return verified
}

