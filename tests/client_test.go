package api

import (
	"context"
	"testing"
	"time"
	"net/http"
	"net/http/httptest"
	"github.com/HTM1000/goexpert-busca-cep/models"
	"github.com/HTM1000/goexpert-busca-cep/api"
)

func TestGetFastestResponse(t *testing.T) {
	mockAPI1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep":"01153000","logradouro":"Rua Teste","bairro":"Bairro Teste","localidade":"Cidade Teste","uf":"SP"}`))
	}))
	defer mockAPI1.Close()

	mockAPI2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond) 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cep":"01153000","logradouro":"Rua Lenta","bairro":"Bairro Lento","localidade":"Cidade Lenta","uf":"RJ"}`))
	}))
	defer mockAPI2.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := api.GetFastestResponse(ctx, mockAPI1.URL, mockAPI2.URL)
	if err != nil {
		t.Fatalf("esperado sem erros, mas recebeu: %v", err)
	}

	if result.Cep != "01153000" || result.Localidade != "Cidade Teste" {
		t.Errorf("resultado esperado: %+v, mas recebeu: %+v", models.Address{
			Cep: "01153000", Localidade: "Cidade Teste",
		}, result)
	}
}

func TestGetFastestResponseTimeout(t *testing.T) {
	mockAPI1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) 
	}))
	defer mockAPI1.Close()

	mockAPI2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) 
	}))
	defer mockAPI2.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := api.GetFastestResponse(ctx, mockAPI1.URL, mockAPI2.URL)
	if err == nil {
		t.Fatal("esperado erro de timeout, mas nenhum erro ocorreu")
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Fatalf("esperado erro de timeout (context deadline exceeded), mas recebeu: %v", err)
	}
}
