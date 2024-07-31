package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(source string, password []byte) {

	//check for file; panic error if null
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	//open source file
	srcFile, err := os.Open(source)
	if err != nil {
		fmt.Println("source file not found!!!")
		panic(err.Error())
	}
	defer srcFile.Close()

	//assigns 'plaintext' as srcfile
	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	//nonce refers to unique random number generated for authentication
	key := password
	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//derived key from public based key derivation function v.2
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	//advanced encryption standard to cipherise plaontext
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//create ciphertext
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)

	//create destination file for ciphertext from source
	dstFile, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(ciphertext)
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}
	srcFile, err := os.Open(source)
	if err != nil {
		fmt.Println("source file not found!!!")
		panic(err.Error())
	}
	defer srcFile.Close()

	//assigns 'plaintext' as srcfile
	ciphertext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext)-12:]
	str := hex.EncodeToString(salt)
	nonce, err := hex.DecodeString(str)

	dk := pbkdf2.Key(name, nonce, 4096, 32, sha1.New())
	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}
	
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//initialise plaintext as a decrypted ciphertext - 12 char as those are the nonce
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext,[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err := dstFile. Write(plaintext)
	if err != nil {
		panic(err.Error())
	}


}
