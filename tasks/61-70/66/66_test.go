package main

import (
	"math/rand"
	"testing"
)

func TestGeneratePasswordDeterministic(t *testing.T) {
	rand.Seed(42) // фиксируем сид
	pass := generatePassword(5)
	expected := "aB7cZ" // ← один раз запускаем, смотрим, что получилось
	if pass != expected {
		t.Errorf("ожидали %q, получили %q", expected, pass)
	}
}
