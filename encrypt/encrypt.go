package encrypt
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// EncryptFile
// Example
//  path, err := os.Getwd()
//	_, err = encrypt.EncryptFile(filepath.Join(path, FILE), SECRET_KEY)
//	if err != nil {
//		panic(err)
//	}
func EncryptFile(file string, k string) (string, error) {
	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	err = ioutil.WriteFile(file+".bin", ciphertext, 0777)
	if err != nil {
		return "", err
	}
	return file + ".bin", err
}

func DecryptFile(file string, k string) (string, error) {
	ciphertext, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	name := strings.TrimSuffix(file, filepath.Ext(file))
	err = ioutil.WriteFile(name, plaintext, 0777)
	if err != nil {
		return "", err
	}
	fmt.Println(name)
	return name, err
}

func GenerateKey() string {
	key := make([]byte, 32)

	_, _ = rand.Read(key)
	encoded := base64.StdEncoding.EncodeToString(key)
	return encoded
}
