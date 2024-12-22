package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/HTM1000/goexpert-busca-cep/models"
)

func fetch(ctx context.Context, url, source string, resultChan chan<- models.Address, errorChan chan<- error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errorChan <- fmt.Errorf("erro ao criar requisição para %s: %v", source, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errorChan <- fmt.Errorf("erro ao fazer requisição para %s: %v", source, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("API %s retornou status inválido: %s", source, resp.Status)
		return
	}

	var addr models.Address
	if err := json.NewDecoder(resp.Body).Decode(&addr); err != nil {
		errorChan <- fmt.Errorf("erro ao decodificar resposta da API %s: %v", source, err)
		return
	}

	addr.SourceAPI = source
	select {
	case resultChan <- addr:
	default:
	}
}

func GetFastestResponse(ctx context.Context, url1, url2 string) (models.Address, error) {
	var wg sync.WaitGroup
	resultChan := make(chan models.Address, 1)
	errorChan := make(chan error, 1)

	wg.Add(2)
	go func() {
		defer wg.Done()
		fetch(ctx, url1, "BrasilAPI", resultChan, errorChan)
	}()
	go func() {
		defer wg.Done()
		fetch(ctx, url2, "ViaCEP", resultChan, errorChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return models.Address{}, err
	case <-ctx.Done():
		return models.Address{}, errors.New("tempo limite excedido")
	}
}
