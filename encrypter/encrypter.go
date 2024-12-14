package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{Key: key}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	// Мы будем шифровать с помощью aes - это библиотека aescrypt, которая встроена в Go
	// Она позволяет создать новый block, block - это по сути объект, представляющий симметричный блочный шифр
	// Важно понимать, что это симметричное шифрование
	block, err := aes.NewCipher([]byte(enc.Key)) // этот шифр создаётся с нашим исходным ключом, который позволит его зашифровать
	if err != nil {
		panic(err.Error())
	}
	// Теперь нам необходимо из этого блока создать некоторый GSM - это счётчик с идентификацией Галуа, который широко применяется в различных симметричных шифрованиях и с помощью него, добавляя туда некоторую случайную будем считать числом (NNOU) - nouns number once used.
	// Мы будем его (это число) шифровать, чтобы мы не могли предсказать, как будет происходить каждый раз это шифрование и соответсвенно мы могли чётко понимать, что наши данные действительно в безопасности.
	// Для того, чтобы создать этот GSM (Galois/Counter Mode) [Счётчик с идентификацией Галуа], необходимо использовать cipher, вызываем NewGCM, внутрь принимаем блок шифрования, который мы создали ранее.
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// Теперь создаём случайное уникальное значение, которое мы будем использовать для безопасности наших криптографических операций
	nonce := make([]byte, aesGSM.NonceSize())
	// Теперь этот nonce, который необходимо каким-то образом случайным заполнить
	// Для этого мы будем использовать некоторый ioReadFull, вместе с rand.Reader, который позволяет нам создать некоторое случайное число
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	// И наконец последним шагом, нам надо зашифровать наш текст и его вернуть
	return aesGSM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // этот шифр создаётся с нашим исходным ключом, который позволит его зашифровать
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
