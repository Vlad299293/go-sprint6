package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	trimmedData := strings.TrimSpace(data)
	if trimmedData == "" {
		return "", errors.New("пустой файл")
	}

	isMorseCode := !strings.ContainsFunc(trimmedData, func(r rune) bool {
		return r != '.' && r != '-' && r != ' ' && r != '\n' && r != '\r' && r != '\t'
	})

	if isMorseCode {
		return morse.ToText(data), nil
	}

	return morse.ToMorse(data), nil
}
