# Backend API - Snet Project

API RESTful desenvolvida em Go com Echo, seguindo arquitetura em camadas (handler > service > repository). Utiliza PostgreSQL para persistência de dados e suporta operações CRUD para estabelecimentos e lojas, além de autenticação básica de usuários.

---

## Tecnologias

- Go 1.24+
- Echo Framework
- GORM ORM
- PostgreSQL
- Docker + Docker Compose

---

## Arquitetura

- **Handlers**: Responsáveis por receber as requisições HTTP e enviar as respostas.
- **Services**: Contém a lógica de negócio e regras de validação.
- **Repositories**: Interface com o banco de dados usando GORM.
- **Models**: Representação das entidades do sistema.
- **DTOs**: Objetos de transferência para entrada/saída de dados.

---

## Endpoints

Veja todos os endpoints [aqui](routes.md)!

---

## Variáveis de Ambiente
`.env`

```bash
APP_PORT=3333
DB_HOST=db
DB_USER=docker
DB_PASSWORD=123
DB_NAME=snet_db
DB_PORT=5432
```

---

## Rodar o projeto Localmente (Docker)

```bash
$ docker-compose up --build
```
