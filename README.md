# apiGo

API REST em Go para gerenciamento de tarefas (tasks), com operações básicas de CRUD. Projeto desenvolvido para praticar Go com arquitetura limpa, PostgreSQL e sem uso de frameworks externos.

## Tecnologias

- **Go** 1.25
- **PostgreSQL** (via `pgx/v5` com connection pool)
- **godotenv** para variáveis de ambiente

## Estrutura do projeto

```
apiGo/
├── cmd/
│   └── main.go
├── internal/
│   ├── api/
│   │   └── router.go
│   ├── config/
│   │   └── database.go
│   ├── domain/
│   │   ├── task.go
│   │   ├── tasks_errors.go
│   │   └── tasks_repository.go
│   ├── handler/
│   │   └── handler.go
│   ├── repository/
│   │   └── repository.go
│   └── service/
│       └── task_service.go
├── .env
├── go.mod
└── go.sum
```

O projeto segue uma arquitetura em camadas:

- **domain** — entidades, interface do repositório e erros de domínio
- **repository** — implementação do acesso ao banco (PostgreSQL)
- **service** — regras de negócio e validações
- **handler/api** — recebimento das requisições HTTP e roteamento
- **config** — inicialização de dependências externas (banco de dados)

## Endpoints

| Método | Rota         | Descrição                   |
|--------|--------------|-----------------------------|
| POST   | `/task`      | Cria uma nova tarefa        |
| GET    | `/task`      | Retorna todas as tarefas    |
| GET    | `/task/{id}` | Retorna uma tarefa pelo ID  |
| PUT    | `/task/{id}` | Atualiza uma tarefa         |
| DELETE | `/task/{id}` | Remove uma tarefa pelo ID   |

## Modelo de dados

```json
{
  "id": "string",
  "title": "string",
  "done": false,
  "created_at": "2024-01-01T00:00:00Z"
}
```

## Como rodar localmente

### Pré-requisitos

- Go instalado
- PostgreSQL rodando

### Configuração

Crie um arquivo `.env` na raiz do projeto:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5433/SeuBanco
```

### Criando a tabela

```sql
CREATE TABLE tasks (
    id         VARCHAR PRIMARY KEY,
    title      VARCHAR NOT NULL,
    done       BOOLEAN,
    created_at TIMESTAMP
);
```

### Rodando

```bash
go run ./cmd/main.go
```

A API sobe na porta `8080`.

## Exemplos de uso

**Criar tarefa**
```bash
curl -X POST http://localhost:8080/task \
  -H "Content-Type: application/json" \
  -d '{"id": "1", "title": "Estudar Go", "done": false, "created_at": "2024-01-01T00:00:00Z"}'
```

**Listar todas**
```bash
curl http://localhost:8080/task
```

**Buscar por ID**
```bash
curl http://localhost:8080/task/1
```

**Atualizar**
```bash
curl -X PUT http://localhost:8080/task/1 \
  -H "Content-Type: application/json" \
  -d '{"id": "1", "title": "Estudar Go avançado", "done": true, "created_at": "2024-01-01T00:00:00Z"}'
```

**Deletar**
```bash
curl -X DELETE http://localhost:8080/task/1
```

## Códigos de resposta

| Código | Descrição                               |
|--------|-----------------------------------------|
| 200    | OK                                      |
| 201    | Tarefa criada com sucesso               |
| 204    | Tarefa deletada (sem conteúdo)          |
| 400    | Requisição inválida                     |
| 404    | Tarefa não encontrada                   |
| 409    | Conflito (ex: tarefa já concluída)      |
| 500    | Erro interno do servidor                |
