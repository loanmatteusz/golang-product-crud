# Desafio Fullstack SNet: Backend Go + Frontend Nuxt

Aplicação fullstack com backend desenvolvido em Go (Echo framework) e frontend em Nuxt 3 (Vue 3). O backend expõe uma API REST para gerenciar usuários, estabelecimentos e lojas, e o frontend consome essa API para apresentar a interface ao usuário.

---

## Tecnologias utilizadas

- Backend:
  - Go 1.24+
  - Echo framework
  - PostgreSQL
- Frontend:
  - Nuxt 3
  - Vue 3
  - TailwindCSS
  - Zod para validação
- Docker e Docker Compose para containerização e orquestração

---

## Configuração local

### Pré-requisitos

- Docker e Docker Compose instalados
- (Opcional) Go instalado para rodar backend localmente
- (Opcional) Node.js e npm/yarn para rodar frontend localmente

---

### Variáveis de ambiente

O backend e o frontend usam variáveis de ambiente para configuração, principalmente:

- **Backend:**

```env
DATABASE_URL=postgres://user:password@host:port/dbname
PORT=3333
```

- **Frontend:**

```env
API_URL=http://localhost:3333
```

---

## Rodando a aplicação com Docker Compose

Na raiz do projeto, execute:

```bash
docker-compose up --build
```

Isso vai:

- Construir e subir o container do backend na porta `3333`
- Construir e subir o container do frontend na porta `3000`

Acesse a aplicação no navegador via:

```
http://localhost:3000
```

---

## Rodando localmente sem Docker

### Backend

No diretório `backend/`:

```bash
go run ./cmd/api
```

Certifique-se de que o banco de dados está configurado e rodando.

### Frontend

No diretório `frontend/`:

```bash
npm install
npm run dev
```
---
