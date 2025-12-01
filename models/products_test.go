package models

import "testing"

func TestNewChatMessage(t *testing.T) {
	p := NewChatMessage()
	if p.CodProduct != "5" {
		t.Error("Esperado 5, obtido ", p.CodProduct)
	}
}

func TestNewLoverR(t *testing.T) {
	p := NewLover("r")
	if p.CodProduct != "13" {
		t.Error("Esperado 13, obtido ", p.CodProduct)
	}
}

func TestNewLoverAirflow(t *testing.T) {
	p := NewLover("airflow")
	if p.CodProduct != "1" {
		t.Error("Esperado 1, obtido ", p.CodProduct)
	}

}

func TestNewPresent(t *testing.T) {
	p := NewPresent()
	if p.CodProduct != "11" {
		t.Error("Esperado 11, obtido ", p.CodProduct)
	}
}

func TestNewStreakPresent(t *testing.T) {
	p := NewStreakPresent()
	if p.CodProduct != "12" {
		t.Error("Esperado 12, obtido ", p.CodProduct)
	}
}

func TestNewChurn(t *testing.T) {

	type testPair struct {
		churn       float64
		codExpected string
	}

	pairs := []testPair{
		{0.01, "7"},
		{0.025, "7"},
		{0.045, "8"},
		{0.050, "8"},
		{0.075, "6"},
		{0.100, "6"},
		{0.110, "38"},
		{0.500, "38"},
		{0.999, "38"},
	}

	for _, pair := range pairs {
		p := NewChurn(pair.churn)
		if p.CodProduct != pair.codExpected {
			t.Error("Esperado ", pair.codExpected, " obtido ", p.CodProduct)
		}
	}

}
