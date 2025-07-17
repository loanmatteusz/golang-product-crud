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

### Produtos

- **Criar Produto**

POST /products

Body exemplo:
```json
{
  "name": "Product Name",
  "price": 1050,
	"category_id": "bcc02642-64c4-4bb6-89cc-bb9aa0d9d6c4"
}
```

- **Listar todos os Produtos**

GET /products


- **Obter Produto por ID**

GET /products/:id


- **Atualizar Produto**

PUT /products/:id
> Mesma estrutura do POST, com campos opcionais.

- **Excluir Produto**

DELETE /products/:id
> Não é permitido excluir Produto que tenha lojas vinculadas.


### Categorias

- **Criar Categoria**

POST /categories

Body exemplo:

```json
{
  "name": "Category Name",
}
```

- **Listar todas as categorias**

GET /categories

- **Obter categoria por ID**

GET /categories/:id

- **Atualizar categoria**

PUT /categories/:id
> Mesma estrutura do POST, campos opcionais.

- **Excluit categoria**

DELETE /categories/:id
