package input

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	//"reflect"
)

// ReadInput читает строку ввода от пользователя.
func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// GetInt получает целое число от пользователя.
func GetInt(prompt string) (int64, error) {
	fmt.Printf("%s", prompt)
	input, err := ReadInput()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %w,", err)
	}

	return val, nil
}

// GetFloat получает число с плавающей точкой от пользователя.
func GetFloat(prompt string) (float64, error) {
	fmt.Printf("%s", prompt)
	input, err := ReadInput()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float: %w", err)
	}

	return val, nil
}

// GetBool получает булевое значение от пользователя.
func GetBool(prompt string) (bool, error) {
	fmt.Printf("%s", prompt)
	input, err := ReadInput()
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(input)
	if err != nil {
		return false, fmt.Errorf("invalid bool: %w", err)
	}

	return val, nil
}

// GetString получает строку от пользователя.
func GetString(prompt string) (string, error) {
	fmt.Printf("%s", prompt)
	return ReadInput()
}


// PrintError выводит ошибку в лог.
func LogErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

// GetName запрашивает имя у пользователя и проверяет, чтобы оно содержало только буквы.
func GetName(prompt string) (string, error) {
	fmt.Printf("%s", prompt)
	name, err := GetString("")
	if err != nil {
		return "", err
	}

	for _, char := range name {
		if !unicode.IsLetter(char) {
			return "", errors.New("name must contain only letters")
		}
	}

	return name, nil
}
//12321
