# Transaction System API

Essa API permite realizar as operações de rotina de transações financeiras.

Cada cliente possui uma conta com seus dados e a cada operação realizada pelo cliente uma transação é criada e associada a sua respectiva conta.

As transações possuem tipos e dependendo do tipo de operação o valor será registrado como negativo. Ex: compras e saques são registradas com valor negativo e pagamento é registrado com valor positivo 

## Estrutura dos dados da aplicação

- Accounts:
  - ``id`` - id da conta 
  - ``document_number`` indicando o documento do cliente.
  

- Operations Types:
  - ``id`` - id do tipo de operação
  - ``description`` - Descrição do tipo de operação . (Ex: PAGAMENTO)
  - ``negative_amount`` - Indica se para esse tipo de transação o valor deve ser registrado como negativo

    
- Transactions
  - ``id`` - id da transação
  - ``account_id`` - indicando a conta vinculada a essa transação
  - ``operation_type_id`` - indicando qual o tipo de operação vinculada a essa transação
  - ``amount`` - Valor da transação, onde dependendo do tipo de operação (_operation_type_id_) pode ser negativo ou positivo
  - ``event_date`` - data de criação da transação

## Arquitetura da aplicação

A aplicação foi desenvolvida utilizando os princípios da [Arquitetura Limpa](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), visando desacoplar detalhes de aplicação e domínio de detalhes de implementação de ferramentas, além de essa abordagem facilitar a testabilidade da aplicação.



## Como executar a aplicação

### Docker

Para executar a aplicação via Docker, na pasta raíz do projeto execute:

```
docker-compose up
```

Isso irá subir os containers do banco de dados e da API, no caso da API além de subir a API também executará a ``migration`` e `seeder` conforme será melhor explicado abaixo.

Caso deseje executar os serviços de maneira isolada (um em cada terminal por exemplo), execute:

```
docker-compose up <nome_do_servico>

Ex: docker-compose up db_transaction_system
```

### Local

Caso queira executar a aplicação de maneira local, tenha a versão 1.22 do Golang e caso tenha o ``make`` instalado, também é possível executar via ``make``

__OBS__: Antes de subir a aplicação pela primeira vez, deve-ser executar a migration antes para realizar a criaçao das tabelas e a inserção de informações que já devem constar no banco.

Etapas:

__1 - Crie um arquivo `.env` com as variáveis de ambiente:__

Pode ser feito copiando o arquivo ``.env.example``:

````
cp .env.example .env
````

__OBS__: Modifique o `DB_HOST` para ``localhost`` no `.env`

__2 - Executar a Migration (caso seja a primeira vez):__

Via Go:
````
go run cmd/migration/main.go
````

Via Make:
```
make migration
```

__3 - Executar a aplicação:__

Via Go:
````
go run cmd/transaction-system/main.go
````

Via Make:
```
make run
```



## Migration

A migration foi criada utilizando a lib [migrate](https://github.com/golang-migrate/migrate), sendo responsável pela criação das tabelas e o seeder (inserção) dos dados relacionados a ``operation_types``

A príncipio não foi identificada a necessidade de utilizar um ORM (Ex: Gorm) mas pode ser uma evolução futura caso seja necessário.

## Endpoints

A aplicação atualmente é composta por 3 endpoints:
````
POST /accounts (criação de uma conta)

GET /accounts/:accountId (consulta informações de uma conta)

POST /transactions (criação de uma transação)
````

Abaixo segue alguns exemplos de curl dessas requisições:

__POST /accounts__
```
curl --request POST \
  --url http://localhost:8080/accounts \
  --header 'Content-Type: application/json' \
  --data '{
	"document_number": "123457890"
}'
```
__GET /accounts/:accountId__

```
curl --request GET \
  --url http://localhost:8080/accounts/1
```

__POST /transactions__

```
curl --request POST \
  --url http://localhost:8080/transactions \
  --header 'Content-Type: application/json' \
  --data '{
	"account_id" : 1,
	"operation_type_id": 1,
	"amount": 100.40
}'
```
__OBS__: O valor de amount é registrado como negativo somente dentro da aplicação, no caso da request e response seu valor é absoluto.


### Collection
Dentro da pasta ``assets`` também contém um JSON de collection para importar no Imnsonia com esses endpoints.

### Swagger

Uma documentação dos endpoints pode ser vista via Swagger, acessando http://localhost:8080/swagger/index.html ao executar a aplicação.

Foi utilizada a lib [swag](https://github.com/swaggo/swag) para gerar a documentação.

#### Atualizar a doc

Caso seja necessário atualizar a doc do swagger, instale CLI da swag conforme mostra a documentação acima, feito isso após realizar a atualização, execute o comando para atualizar os arquivos de documentação no  ``/docs``:

```
swag init -g cmd/transaction-system/main.go  
```

## Testes

A API conta com testes unitários e de integração que podem ser executados diretamente via ``go`` ou via ``make`` para isso execute os comandos abaixo na raíz do projeto:

### unit tests

Via Go:
```
go test -v ./internal/tests/unit_tests/*
```

Via Make:
```
make unit_tests
```

### integration tests

No caso dos testes de integração é necessário que o banco de dados esteja em execução, feito isso, os comandos para executar são:

Via Go:
```
go test -v ./internal/tests/integration_tests/*
```

Via Make:
```
make integration_tests
```

## Stack

- Golang 1.22
- Postgres 16.2