package password

import (
	"math"
	"strings"
	"unicode"
)

func (p *Password) CreateAdjacencyMatrix(passwords []string) map[rune]map[rune]float64 {
	adjacencyMatrix := make(map[rune]map[rune]float64)
	totalTransitions := make(map[rune]int)

	for _, password := range passwords {
		// Приводим пароль к нижнему регистру для учета регистра символов
		password = strings.ToLower(password)

		// Проходим по каждой паре символов в пароле
		for i := 0; i < len(password)-1; i++ {
			char1 := rune(password[i])
			char2 := rune(password[i+1])

			// Создаем вложенную карту, если еще не создана
			if adjacencyMatrix[char1] == nil {
				adjacencyMatrix[char1] = make(map[rune]float64)
			}

			// Увеличиваем счетчик перехода от char1 к char2
			adjacencyMatrix[char1][char2]++

			// Увеличиваем общий счетчик переходов для char1
			totalTransitions[char1]++
		}
	}

	// Преобразуем частоты в вероятности
	for char1 := range adjacencyMatrix {
		for char2 := range adjacencyMatrix[char1] {
			adjacencyMatrix[char1][char2] /= float64(totalTransitions[char1])
		}
	}

	return adjacencyMatrix
}

func (p *Password) CalculateCombinationScore(password string) int {
	adjacencyMatrix := p.CreateAdjacencyMatrix(p.commonPatterns)
	score := 0

	// Проходим по всем парам символов в пароле
	for i := 0; i < len(password)-1; i++ {
		currentChar := rune(password[i])
		nextChar := rune(password[i+1])
		if adjacencies, exists := adjacencyMatrix[currentChar]; exists {
			if probability, exists := adjacencies[nextChar]; exists {
				if probability < 0.1 { // Пороговая вероятность для редких комбинаций
					score += 2
				}
			}
		}
	}
	return score
}

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

func (p *Password) CalculatePositionalDistributionScore(password string) int {
	// Инициализируем счетчики для биграмм и триграмм
	bigrams := make(map[string]int)
	trigrams := make(map[string]int)

	// Приводим пароль к нижнему регистру для учета регистра символов
	password = strings.ToLower(password)

	// Считаем биграммы и триграммы в пароле
	for i := 0; i < len(password)-1; i++ {
		bigram := password[i : i+2]
		bigrams[bigram]++

		if i < len(password)-2 {
			trigram := password[i : i+3]
			trigrams[trigram]++
		}
	}

	// Вычисляем суммарную оценку на основе биграмм и триграмм
	score := 0
	for _, count := range bigrams {
		score += count
	}
	for _, count := range trigrams {
		score += count
	}

	return score
}