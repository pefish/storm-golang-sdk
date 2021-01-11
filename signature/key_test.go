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
	privKey := `60a10883c8252658bed264951385a62096d69f5ce4fd9e43b690f091e29ab2bb`
	fmt.Println(SignMessage(`{"amount":"14098.3268000000","chain":"Pmeer","confirmation":3,"confirmation_threshold":15,"created_at":"2020-06-09T03:35:21.000Z","currency":"PMEER","decimals":8,"height":null,"network_fee":"0.0010000000","network_fee_currency":"PMEER","network_fee_decimal":8,"request_id":"574-","side":"withdraw","to_address":"TmYXzyTiNKrK6mQG8qkxdJeXiE55xrbhYgu","tx_id":"6e8d378ae7f199247255f28f41531b4fdbae797c1f8eee1935eb1383daf819e0","user_id":3}`, privKey))
}

func TestVerifySignature(t *testing.T) {
	pubKey := `0302e2b8b44fe202ac68f2f74a8a503496fbec3d9c3b686b9f06812e08d9fbbd78`
	sig := `304402202e3e0b1ecf96c3d190c9cebc20af9fcaa3f006da0fde34f64548e78134cb80bd0220027d314829124c0d88d3c3bc60feca357308ae7d566792c3c5ae993e7a4b33c8`
	if VerifySignature(`{"amount":232.95032569,"block_id":"0000040102e563e8878c0f6d9ca228e5a70a7be993b41a00af55abe73530d6e8","chain":"Pmeer","confirmation":100,"confirmation_threshold":10,"created_at":"2020-06-08T08:30:42Z","currency":"PMEER","decimals":8,"from_address":"","height":"463184","network_fee":0,"network_fee_currency":"PMEER","network_fee_decimal":8,"output_index":5,"request_id":"","side":"deposit","tag":"","to_address":"TmgaeDCBR9sHVXYbyTm4HwUX7nyUsNtsDJz","tx_id":"874cfe640ba987b9bcaffcc37e03d243e576feb4fb6397a4644742942c273561","user_id":3}`, sig, pubKey) != true {
		t.Error()
	}
}
