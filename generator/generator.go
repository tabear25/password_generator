package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

// パスワードで使用する文字のリスト
const (
	LetterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberBytes    = "0123456789"
	SymbolBytes    = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	AlphaOnly      = 1
	AlphaNum       = 2
	NumOnly        = 3
	AlphaNumSymbol = 4
)

// 指定されたパターンに基づいてランダムな文字を選択する
func GetRandomChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}

// 指定された文字数とパターンに基づいてパスワードを作成する
func GeneratePassword(length int, charset string) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}
	var password strings.Builder
	for i := 0; i < length; i++ {
		char, err := GetRandomChar(charset)
		if err != nil {
			return "", err
		}
		password.WriteByte(char)
	}
	return password.String(), nil
}

// パターンが合っているのかの確認
func GetCharset(choice int) (string, error) {
	switch choice {
	case AlphaOnly:
		return LetterBytes, nil
	case AlphaNum:
		return LetterBytes + NumberBytes, nil
	case NumOnly:
		return NumberBytes, nil
	case AlphaNumSymbol:
		return LetterBytes + NumberBytes + SymbolBytes, nil
	default:
		return "", errors.New("invalid charset choice")
	}
}
