# GoExpert Busca CEP

Este projeto é uma aplicação Go que realiza a busca de um endereço baseado no CEP fornecido, utilizando duas APIs diferentes para realizar a consulta. O objetivo principal do sistema é obter o endereço mais rápido entre as duas APIs e retornar o resultado para o usuário.

## Funcionalidades

  - Faz duas requisições simultâneas para duas APIs diferentes:
    1. BrasilAPI
    2. ViaCEP
  - O sistema escolhe a resposta mais rápida entre as duas e retorna o endereço correspondente.
  - Se a resposta demorar mais de 1 segundo, o sistema retorna um erro de "timeout".
  - Caso uma das APIs falhe ou tenha erro, o sistema irá tratar a falha e devolver a mensagem de erro apropriada.

  ## Como Rodar

  ### Pré-requisitos

  Go instalado em sua máquina.

  ### Passos para Execução
  1. Clone o repositório:
  `git clone https://github.com/HTM1000/goexpert-busca-cep.git`
  
  2. Navegue até o diretório do projeto:
  `cd goexpert-busca-cep`

  3. Baixe as dependências do projeto (se houver):
  `go mod tidy`

  4. Execute o programa:
  `go run main.go`

  5. O programa irá solicitar o CEP que você deseja consultar e, em seguida, exibirá o endereço retornado pela API mais rápida. Se o tempo limite for excedido ou ocorrer algum erro, ele exibirá uma mensagem de erro.

  ### Exemplo de Execução

  `go run main.go`

  Saída esperada (se a API responder corretamente):
  ```plaintext
  Resultado mais rápido:
  API: BrasilAPI
  Endereço: {Logradouro:Rua Exemplo Bairro:Centro Cidade:São Paulo UF:SP}
  ```

  Saída em caso de erro (tempo limite ou falha na API):
  `Erro: tempo limite excedido`

  ## Estrutura do Projeto

  O projeto é dividido da seguinte forma:
  ```plaintext
  goexpert-busca-cep/
  ├── api/
  │   └── client.go   # Contém a lógica para fazer requisições para as APIs e retornar o endereço
  ├── models/
  │   └── address.go  # Estrutura que representa o endereço retornado pela API
  ├── main.go         # Arquivo principal que executa a aplicação
  └── README.md       # Este arquivo
  ```

  - api/client.go: Implementa a função GetFastestResponse que faz as requisições paralelamente para as duas APIs e retorna o endereço da API mais rápida.
  - models/address.go: Define o modelo de dados Address, que é usado para representar o endereço.
  - main.go: Responsável pela execução da aplicação e interação com o usuário.

  ## Testes

  ### Como rodar os testes

  Para garantir que o código funcione corretamente, você pode rodar os testes unitários do Go:
  `go test ./...`
  
  O código inclui testes para verificar a funcionalidade das requisições e o tratamento de erros.