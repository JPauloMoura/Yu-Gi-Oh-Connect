## API Yu-Gi-Oh! Connect
### É hora do duelo!
![Yu-Gi-Oh-Trading-Card-Games-2237302-wallhere com](https://github.com/JPauloMoura/Yu-Gi-Oh-Connect/assets/62079201/3acd52d7-32f7-4aa8-bafa-ab580fbe6d0e)

### Visão Geral

A API Yu-Gi-Oh! Connect permite criar, atualizar, listar, recuperar e excluir informações de duelistas.
Duelistas são como são chamados os jogadores de Yu-Gi-oh TCG. E para ajudar a conectar esses jogadores novas batalhas esse projeto foi criado!

### Setup
1. Intale o Go na sua máquina
2. Instale o Docker e Docker-compose 
3. Instale o go-migrations para que seja possível realizara criação automatizada das tabelas:
```bash
$ brew install golang-migrate
```

### Rodando o projeto

Antes de executar a API, certifique-se de ter feito o setup comentado acima.

Para instalar as dependências e executar a API, execute os seguintes comandos:

```bash
# suba a infaestrutura com docker-compose (o processo segura o seu terminal)
# obs: o docker deve está sendo rodando
$ make run

# -> Em um novo terminal:
# instale os pacotes de dependencias do projeto
$ make tidy

# crie as tabelas do banco de dados
$ make migrations-up

# suba a api da aplicação
$ make run-api
```

## Endpoints

Você pode instalar a extensão `REST Client` no VScode para fazer as requisições diretamente do arquivo `Request.http` que está na raiz do projeto ;)

## Criar Duelista

Este endpoint é usado para criar um novo Duelista.
Utilizamos o cep informado para consultar os dados de endenreço do duelista e salvar.

- **URL**
  
  `/duelist`

- **Método**

  `POST`

- **Parâmetros do corpo**
  
  | Nome          | Tipo   | Descrição                                     |
  |---------------|--------|-----------------------------------------------|
  | name          | string | O nome do Duelista                           |
  | presentation  | string | Uma breve apresentação do Duelista            |
  | birthDate     | string | A data de nascimento do Duelista (DD/MM/AAAA) |
  | cep           | string | O CEP do endereço do Duelista                |
  | email         | string | O email do Duelista                          |
  | phone         | string | O número de telefone do Duelista             |

<br/>
<details>
<summary><b>Exemplo de Requisição</b></summary>

  ```http
  POST http://localhost:3001/duelist
  Content-Type: application/json

  {
    "name": "JP",
    "presentation": "Duelista experiente com paixão por jogabilidade estratégica!",
    "birthDate": "24/06/1998",
    "cep": "72007040",
    "email": "duelist@example.com",
    "phone": "9999999999"
  }
  ```
</details>

<details>
<summary><b>Exemplo de Response</b></summary>
  
  ```http
  HTTP/1.1 201 Created
  Content-Type: application/json
  Date: Fri, 09 Feb 2024 14:38:40 GMT
  Content-Length: 410
  Connection: close

  {
    "data": {
      "id": "51ee3bc7-8ecb-467c-833c-da1db0314fc4",
      "name": "JP",
      "presentation": "Experienced duelist with a passion for strategic gameplay!",
      "birthDate": "1998-06-24T00:00:00Z",
      "address": {
        "state": "DF",
        "city": "Brasília",
        "street": "Rua Rua 8 Chácara 220",
        "district": "Setor Habitacional Vicente Pires",
        "cep": "72007040"
      },
      "contact": {
        "email": "duelist@example.com",
        "phone": "9999999999"
      }
    },
    "error": "",
    "statusCode": 201
  }
  ```
</details>

## Atualizar Duelista

Este endpoint é usado para atualizar as informações de um Duelista existente.

- **URL**
  
  `/duelist/{id}`

- **Método**

  `PUT`

- **Parâmetros do corpo**
  
  | Nome          | Tipo   | Descrição                                     |
  |---------------|--------|-----------------------------------------------|
  | name          | string | O nome do Duelista                           |
  | presentation  | string | Uma breve apresentação do Duelista            |
  | birthDate     | string | A data de nascimento do Duelista (DD/MM/AAAA) |
  | cep           | string | O CEP do endereço do Duelista                |
  | email         | string | O email do Duelista                          |
  | phone         | string | O número de telefone do Duelista             |

</br>
<details>
<summary><b>Exemplo de Requisição</b></summary>

  ```http
  PUT http://localhost:3001/duelist/51ee3bc7-8ecb-467c-833c-da1db0314fc4
  Content-Type: application/json

  {
    "cep": "64290000"
  }
  ```
</details>

<details>
<summary><b>Exemplo de Response</b></summary>
  
  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Fri, 09 Feb 2024 14:40:50 GMT
  Content-Length: 47
  Connection: close

  {
    "data": "updated",
    "error": "",
    "statusCode": 200
  }
  ```
</details>

## Listar Duelistas

Este endpoint é usado para listar todos os Duelistas.
Nele tamnbém está implementado a páginação e ordenação dos resultados.

- **URL**
  
  `/duelist`

- **Método**

  `GET`

- **Parâmetros da consulta**
  
  | Nome    | Tipo   | Descrição                                                                |
  |---------|--------|--------------------------------------------------------------------------|
  | sort    | string | A ordem de classificação dos Duelistas (asc ou desc)                     |
  | field   | string | O campo pelo qual os Duelistas devem ser classificados (name, birthDate) |
  | limit   | int    | O número máximo de Duelistas a serem retornados                          |
  | page    | int    | O número da página de resultados                                         |

</br>
<details>
<summary><b>Exemplo de Requisição</b></summary>

  ```http
  GET http://localhost:3001/duelist?sort=asc&field=birthDate&limit=10&page=1
  Content-Type: application/json
  ```
</details>

<details>
<summary><b>Exemplo de Response</b></summary>
  
  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Fri, 09 Feb 2024 14:42:44 GMT
  Content-Length: 354
  Connection: close

  {
    "data": [
      {
        "id": "51ee3bc7-8ecb-467c-833c-da1db0314fc4",
        "name": "JP",
        "presentation": "Experienced duelist with a passion for strategic gameplay!",
        "birthDate": "1998-06-24T00:00:00Z",
        "address": {
          "state": "PI",
          "city": "Altos",
          "street": "",
          "district": "",
          "cep": "64290000"
        },
        "contact": {
          "email": "duelist@example.com",
          "phone": "9999999999"
        }
      }
    ],
    "error": "",
    "statusCode": 200
  }
  ```
</details>

## Excluir Duelista

Este endpoint é usado para excluir um Duelista existente.

- **URL**
  
  `/duelist/{id}`

- **Método**

  `DELETE`

</br>
<details>
<summary><b>Exemplo de Requisição</b></summary>

  ```http
  DELETE http://localhost:3001/duelist/51ee3bc7-8ecb-467c-833c-da1db0314fc4
  Content-Type: application/json
  ```
</details>

<details>
<summary><b>Exemplo de Response</b></summary>
  
  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Fri, 09 Feb 2024 14:45:50 GMT
  Content-Length: 47
  Connection: close

  {
    "data": "deleted",
    "error": "",
    "statusCode": 200
  }
  ```
</details>

## Obter Duelista

Este endpoint é usado para obter os detalhes de um Duelista específico.

- **URL**
  
  `/duelist/{id}`

- **Método**

  `GET`

</br>
<details>
<summary><b>Exemplo de Requisição</b></summary>

  ```http
  GET http://localhost:3001/duelist/ec6c6447-0848-4354-b08c-19ab940edddb
  Content-Type: application/json
  ```
</details>

<details>
<summary><b>Exemplo de Response</b></summary>
  
  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Fri, 09 Feb 2024 14:44:08 GMT
  Content-Length: 352
  Connection: close

  {
    "data": {
      "id": "51ee3bc7-8ecb-467c-833c-da1db0314fc4",
      "name": "JP",
      "presentation": "Experienced duelist with a passion for strategic gameplay!",
      "birthDate": "1998-06-24T00:00:00Z",
      "address": {
        "state": "PI",
        "city": "Altos",
        "street": "",
        "district": "",
        "cep": "64290000"
      },
      "contact": {
        "email": "duelist@example.com",
        "phone": "9999999999"
      }
    },
    "error": "",
    "statusCode": 200
  }
  ```
</details>

### Tratamento de Erros

A API pode retornar os seguintes erros:

- **400 Bad Request**: Indica que a solicitação era inválida ou malformada.
- **404 Not Found**: Indica que o recurso solicitado não foi encontrado.
- **500 Internal Server Error**: Indica que ocorreu um erro inesperado no servidor.

---

Divirta-se!
