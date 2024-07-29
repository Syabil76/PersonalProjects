package fileencryption

import (
	"io"
	"os"
)

func Encrypt(source string, password []byte) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}
	srcFile, err := os.Open(source)
	if er != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	plaintext, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
}

func Decrypt() {

}
