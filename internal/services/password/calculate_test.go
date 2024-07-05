package password

import (
	"testing"
)

func TestCalculateVariety(t *testing.T) {
	tests := []struct {
		password string
		expected int
	}{
		{"", 0},             // Пустой пароль
		{"abc", 2},          // Только нижний регистр
		{"ABC", 2},          // Только верхний регистр
		{"123", 2},          // Только цифры
		{"!@#", 2},          // Только специальные символы
		{"abcABC", 4},       // Верхний и нижний регистр
		{"abc123", 4},       // Нижний регистр и цифры
		{"abcABC123", 6},    // Верхний, нижний регистр и цифры
		{"abcABC123!@#", 8}, // Все типы символов
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			p := Password{}
			score := p.CalculateVariety(tt.password)
			if score != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, score)
			}
		})
	}
}

func TestCalculateEntropy(t *testing.T) {
	tests := []struct {
		password string
		expected int
	}{
		{"", 0},              // Пустой пароль
		{"a", 4},             // Один символ
		{"abc", 13},          // Три символа (нижний регистр)
		{"abcABC", 25},       // Шесть символов (верхний и нижний регистр)
		{"abcABC123", 35},    // Девять символов (верхний, нижний регистр и цифры)
		{"abcABC123!@#", 44}, // Двенадцать символов (все типы символов)
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			p := Password{}
			entropy := p.CalculateEntropy(tt.password)
			if entropy != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, entropy)
			}
		})
	}
}

func TestCalculatePoolSize(t *testing.T) {
	tests := []struct {
		password string
		expected int
	}{
		{"", 0},      // Пустой пароль
		{"a", 26},    // Нижний регистр
		{"A", 26},    // Верхний регистр
		{"1", 10},    // Цифры
		{"!", 32},    // Специальные символы
		{"aA1!", 94}, // Все типы символов
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			p := Password{}
			poolSize := p.CalculatePoolSize(tt.password)
			if poolSize != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, poolSize)
			}
		})
	}
}

func TestCheckCommonPatterns(t *testing.T) {
	p := Password{
		commonPatterns: []string{"password", "123456", "qwerty"},
	}

	tests := []struct {
		password string
		expected int
	}{
		{"", 0},                // Пустой пароль
		{"password", 10},       // Распространенный шаблон
		{"123456", 10},         // Распространенный шаблон
		{"qwerty", 10},         // Распространенный шаблон
		{"uniquePass123", 0},   // Уникальный пароль
		{"passWord123456", 10}, // Содержит распространенный шаблон
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			score := p.CheckCommonPatterns(tt.password)
			if score != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, score)
			}
		})
	}
}
