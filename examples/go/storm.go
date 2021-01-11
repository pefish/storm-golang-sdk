package main

import (
	"fmt"
	. "github.com/pefish/storm-golang-sdk/remote"
	. "github.com/pefish/storm-golang-sdk/signature"
)

func main() {

	//生成密钥对
	pubKey, privKey := GenerateRandomKeyPair()
	fmt.Printf("public key is: %s\n private key is: %s\n", pubKey, privKey)

	// 签名消息
	sigMsg := SignMessage("haha", privKey)
	fmt.Println(`签名后的消息为：`, sigMsg)

	// 校验签名
	fmt.Println(`校验签名：`, VerifySignature(`haha`, sigMsg, pubKey))

	// 初始化
	var remote = NewRemote(`http://localhost:8000`, WithKey(`021fb22fede82c1180f46b7d1a6efc1f0d7ecadc5ce5c830e3b9adf980affd6bc7`, `3611036fd51a5d57b2421c9c1c10ef666dafadd5931728c1b29dc2837d36859b`, `0342c859a3dafd76c7949862b2265bcf71cb53afd976800c5e7f93472bccb790df`))

	// 获取充值地址
	result, err := remote.GetNewDepositAddress(GetNewDepositAddressParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Index:    1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf(`address: %s`, result.Address)

	// 校验地址格式
	isValidAddr, err := remote.ValidateAddress(ValidateAddressParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Address:  `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
	})
	if err != nil {
		panic(err)
	}
	if isValidAddr != true {
		fmt.Println(`address invalid`)
		return
	}

	//# 判断是否平台地址
	isPlatformAddress, err := remote.IsPlatformAddress(IsPlatformParam{
		Currency: `BTC`,
		Chain:    `Btc`,
		Address:  `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
	})
	if err != nil {
		panic(err)
	}
	if isPlatformAddress != false {
		fmt.Println(`not storm address`)
	}

	//# 获取充值交易详情
	results, err := remote.ListDepositTransaction(ListDepositTransactionParam{
		TxId: `67e885c11f0dacf982dd2d1e10a7a62c37e454d9ead827eab7a96124fc628628`,
	})
	if err != nil {
		panic(err)
	}
	if results[0].UserId <= 0 {
		fmt.Println(`error`)
	}

	//  获取账户余额

	balance, err := remote.ListBalance()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, balance)

	//  发起提现
	errMsg := remote.Withdraw(WithdrawParam{
		Currency:  `BTC`,
		Chain:     `Btc`,
		RequestId: `185-aa`,
		Address:   `3HqH1qGAqNWPpbrvyGjnRxNEjcUKD4e6ea`,
		Amount:    `0.01`,
	})
	fmt.Printf(`%#v`, errMsg)

	//  获取提现交易详情
	txs, err := remote.ListWithdrawTransaction(ListWithdrawTransactionParam{
		TxId: `67e885c11f0dacf982dd2d1e10a7a62c37e454d9ead827eab7a96124fc628629`,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("%#v\n", txs)
}
