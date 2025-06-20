## Endpoints

### Health Check

GET /health
Retorna status do servidor.

---

### Autenticação (Simulada)

- **Registrar usuário**

POST /auth/register

Body:
```json
{
  "name": "User Name",
  "email": "user@example.com",
  "password": "123456"
}
```

- **Login**

POST /auth/login

Body:
```json
{
  "email": "user@example.com",
  "password": "123456"
}
```

### Estabelecimentos

- **Criar Estabelecimento**

POST /establishments

Body exemplo:
```json
{
  "number": "100",
  "name": "Establishment Name",
  "legal_name": "Establishment Legal Name",
  "address": {
    "street": "Rua A",
    "number": "123",
    "city": "Cidade X",
    "state": "UF",
    "zip_code": "000000"
  }
}
```

- **Listar todos os estabelecimentos**

GET /establishments


- **Obter estabelecimento por ID**

GET /establishments/:id


- **Atualizar estabelecimento**

PUT /establishments/:id
> Mesma estrutura do POST, com campos opcionais.

- **Excluir estabelecimento**

DELETE /establishments/:id
> Não é permitido excluir estabelecimento que tenha lojas vinculadas.


### Lojas

- **Criar Loja**

POST /stores

Body exemplo:

```json
{
  "number": "177",
  "name": "Store Name",
  "legal_name": "Store Legal Name",
  "address": {
    "street": "Rua B",
    "number": "242",
    "city": "Cidade Y",
    "state": "UF",
    "zip_code": "111111"
  },
  "establishment_id": "uuid-do-estabelecimento"
}
```

- **Listar todas as lojas**

GET /stores

- **Obter loja por ID**

GET /stores/:id

- **Listar lojas pelo ID do estabelecimento**

GET /stores/establishment/:id

- **Atualizar Loja**

PUT /stores/:id
> Mesma estrutura do POST, campos opcionais.

- **Excluit Loja**

DELETE /stores/:id
