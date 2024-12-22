package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HTM1000/goexpert-busca-cep/api"
)

func main() {
	cep := "01153000"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	brasilAPI := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	viaCEP := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	result, err := api.GetFastestResponse(ctx, brasilAPI, viaCEP)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Printf("Resultado mais rápido:\nAPI: %s\nEndereço: %+v\n", result.SourceAPI, result)
}
