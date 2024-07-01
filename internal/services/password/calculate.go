package password

import (
	"math"
	"strings"
	"unicode"
)

func (p *Password) CalculateVariety(password string) int {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	score := 0
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSpecial {
		score++
	}

	return score * 2 // Умножаем на 2 для увеличения вклада в общую оценку
}

func (p *Password) CalculateEntropy(password string) int {
	poolSize := p.CalculatePoolSize(password)
	entropy := float64(len(password)) * math.Log2(float64(poolSize))
	return int(entropy / 10) // Приводим к удобной шкале
}

func (p *Password) CalculatePoolSize(password string) int {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	poolSize := 0
	if hasUpper {
		poolSize += 26
	}
	if hasLower {
		poolSize += 26
	}
	if hasNumber {
		poolSize += 10
	}
	if hasSpecial {
		poolSize += 32 // Примерное количество специальных символов
	}

	return poolSize
}

func (p *Password) CheckCommonPatterns(password string) int {
	for _, pattern := range p.commonPatterns {
		if strings.Contains(strings.ToLower(password), pattern) {
			return 10 // Штрафуем за использование распространённых шаблонов
		}
	}
	return 0
}