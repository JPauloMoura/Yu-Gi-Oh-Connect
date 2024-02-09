## API Yu-Gi-Oh! Connect
### É hora do duelo!

### Visão Geral

A API Yu-Gi-Oh! Connect permite criar, atualizar, listar, recuperar e excluir informações de duelistas.
Duelistas são como são chamados os jogadores de Yu-Gi-oh TCG. E para ajudar a conectar esses jogadores novas batalhas esse projeto foi criado!

### Setup
1. Intale o Go instalado na sua máquina
2. Instale o Docker e Docker-compose 
3. Instale o go-migrations para que seja possivel realizara criação automatizado das tabelas:
```bash
$ brew install golang-migrate
```

### Rodando o projeto

Antes de executar a API, certifique-se de ter feito o setup comentado acima.

Para instalar as dependências e executar a API, execute os seguintes comandos:

```bash
# sobe a infaestrutura com docker-compose (o processo segura o seu terminal)
# obs: o docker deve está sendo executado
$ make run

# -> Em um novo terminal:
# instala os pacotes de dependencias do projeto
$ make tidy

# use para criar as tabelas do banco de dados
$ make migrations-up

# sobe a aplicação
$ make run-api
```

### Endpoints

#### Criar Duelista

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

#### Atualizar Duelista

```http
PUT http://localhost:3001/duelist/92530245-b02a-46e5-9cc3-c7710800a5b8
Content-Type: application/json

{
  "cep": "64290000"
}
```

#### Listar Duelistas

```http
GET http://localhost:3001/duelist?sort=asc&field=birthDate&limit=10&page=1
Content-Type: application/json
```

#### Excluir Duelista

```http
DELETE http://localhost:3001/duelist/ec6c6447-0848-4354-b08c-19ab940edddb
Content-Type: application/json
```

#### Obter Duelista

```http
GET http://localhost:3001/duelist/ec6c6447-0848-4354-b08c-19ab940edddb
Content-Type: application/json
```

### Tratamento de Erros

A API pode retornar os seguintes erros:

- **400 Bad Request**: Indica que a solicitação era inválida ou malformada.
- **404 Not Found**: Indica que o recurso solicitado não foi encontrado.
- **500 Internal Server Error**: Indica que ocorreu um erro inesperado no servidor.

---

Divirta-se!