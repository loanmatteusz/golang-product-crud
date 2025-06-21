# Admin Dashboard - Estabelecimentos e Lojas

Painel de administração feito em **Nuxt 3** utilizando **Nuxt UI**, com gerenciamento de **estabelecimentos** e **lojas** conectados a uma **API REST** local.

## ✨ Funcionalidades

- 📋 Listagem de Estabelecimentos e Lojas
- ➕ Criação de Lojas com associação a Estabelecimentos
- ✏️ Edição de dados de Lojas
- 🗑️ Remoção de Lojas
- 🔍 Paginação e ordenação nas tabelas
- 📦 Integração com API backend

## 🧰 Tecnologias

- [Nuxt 3](https://nuxt.com/)
- [Nuxt UI](https://ui.nuxt.com/)
- [VueUse](https://vueuse.org/)
- [Zod](https://zod.dev/) para validação de formulários
- [@tanstack/vue-table](https://tanstack.com/table/v8/docs/guide/vue/overview) para gerenciamento de tabelas



## 🚀 Start do projeto

1. Instale as dependências:

```bash
npm install
```

2. Rode o projeto:

```bash
npm run dev
```

3. Acesse via navegador:

```
http://localhost:3000
```

## 🔧 Variáveis de Ambiente

Você pode configurar o endpoint da API em `composables/useStore.ts` e `useEstablishment.ts`:

```ts
const API_URL = 'http://localhost:3333/stores'
```

Para um projeto mais flexível, mova isso para `.env`:

```ts
const API_URL = useRuntimeConfig().public.apiBase + '/stores';
```

E no `nuxt.config.ts`:

```ts
runtimeConfig: {
  public: {
    apiBase: process.env.API_BASE || 'http://localhost:3333'
  }
}
```

## 🧪 TODO / Melhorias Futuras

- [ ] Autenticação com token JWT
- [ ] Filtro e busca por nome/estabelecimento
- [ ] Paginação real via backend
---
