package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
	"github.com/holmen1/holmcoin/utils"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func hashRIPEMD160(data []byte) []byte {
	hash := ripemd160.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func addVersion(data []byte) []byte {
	result := make([]byte, 21)
	result[0] = 0x00
	copy(result[1:], data)
	return result
}

func addChecksum(data []byte) []byte {
	firstSHA := hashSHA256(data)
	secondSHA := hashSHA256(firstSHA)
	checksum := secondSHA[:4]

	result := make([]byte, 25)
	copy(result[:21], data)
	copy(result[21:], checksum)
	return result
}

func (w *Wallet) GenerateAddress() string {
	publicKeyBytes := append(w.publicKey.X.Bytes(), w.publicKey.Y.Bytes()...)
	// 1. Perform SHA-256 hashing on the public key (32 bytes)
	hashedPublicKey := hashSHA256(publicKeyBytes)
	// 2. Perform RIPEMD-160 hashing on the result of SHA-256 (20 bytes)
	ripeHash := hashRIPEMD160(hashedPublicKey)
	// 3. Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
	versionedPayload := addVersion(ripeHash)
	// 4. Perform SHA-256 hash on the extended RIPEMD-160 result
	// 5. Perform SHA-256 hash on the result of the previous SHA-256 hash
	// 6. Take the first 4 bytes of the second SHA-256 hash for checksum
	// 7. Add the 4 checksum bytes from 7 at the end of extended RIPEMD-160 hash from 4 (25 bytes)
	fullPayload := addChecksum(versionedPayload)
	// 8. Convert the result from a byte string into base58
	address := base58.Encode(fullPayload)
	return address
}

func NewWallet() *Wallet {
	// Creating ECDSA private key (32 bytes) public key (64 bytes)
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	w.blockchainAddress = w.GenerateAddress()
	log.Printf("action=new wallet, address=%s\n", w.blockchainAddress)
	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string, value float32) *Transaction {
	return &Transaction{privateKey, publicKey, sender, recipient, value}
}

func (t *Transaction) GenerateSignature() (*utils.Signature, error) {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, err := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	return &utils.Signature{r, s}, nil
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
