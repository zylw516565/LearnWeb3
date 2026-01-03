package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"hash"
	"os"
)

const nickname = "peter"

func main() {
	fmt.Println("Begin to simulate pow!")

	nick := []byte(nickname)
	nonce := make([]byte, 16)

	hash := sha256.New()
	resHash := calcHash(4, hash, nick, nonce)

	// --- 步骤 1: 生成 2048 位的 RSA 密钥对 ---
	fmt.Println("--- Generating RSA Key Pair (2048 bits) ---")
	privateKey, publicKey, err := generateRSAKeyPair(2048)
	if err != nil {
		fmt.Printf("Failed to generate key pair: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Key pair generated successfully.\n")

	signature, err := signData(privateKey, resHash)
	if err != nil {
		fmt.Printf("Failed to sign data: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated Signature (hex): %x\n\n", signature)

	err = verifySignature(publicKey, resHash, signature)
	if err != nil {
		fmt.Printf("Signature verification FAILED: %v\n", err)
	} else {
		fmt.Println("Signature verification SUCCEEDED!")
	}
	fmt.Println()

	// --- 测试：篡改数据后验证 ---
	fmt.Println("--- Testing with Tampered Data ---")
	tamperedData := []byte("This is a message that needs to be signed. (TAMPERED!)")
	err = verifySignature(publicKey, tamperedData, signature)
	if err != nil {
		fmt.Printf("Signature verification FAILED for tampered data: %v\n", err)
	} else {
		fmt.Println("Signature verification SUCCEEDED for tampered data (This should not happen!)")
	}
}

func calcHash(leadZeroBitwise uint, hash hash.Hash, nick, nonce []byte) []byte {
	var result []byte
	for {
		full_name := make([]byte, 0, len(nick)+len(nonce))
		genNonce(nonce)
		full_name = append(full_name, nick...)
		full_name = append(full_name, nonce...)

		hash.Write(full_name)
		result = hash.Sum(nil)

		// fmt.Printf("nonce: %x\t result:%x\t result[0]:%08b\n", nonce, result, result[0])
		if (result[0] >> (8 - leadZeroBitwise)) == 0 {
			break
		}
	}

	return result
}

func genNonce(nonce []byte) {
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}
}

// 1. 生成 RSA 密钥对 (与加密示例相同)
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// 2. 使用私钥对数据进行签名
func signData(privateKey *rsa.PrivateKey, hashData []byte) ([]byte, error) {
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashData)
	if err != nil {
		return nil, err
	}

	return sign, nil
}

// 3. 使用公钥验证签名
func verifySignature(publicKey *rsa.PublicKey, hashData []byte, signature []byte) error {

	// b. 使用公钥验证签名
	// VerifyPKCS1v15 验证给定的签名是否是数据哈希值的有效 RSA 签名
	// 参数与 SignPKCS1v15 类似
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashData, signature)
	return err // 如果验证通过，err 为 nil；否则，err 不为 nil
}
