import type { RouterConfig } from '@nuxt/schema';

export default {
  // https://router.vuejs.org/api/interfaces/routeroptions.html#routes
  routes: (_routes) => [
    {
      path: '/',
      component: () => import('~/pages/index.vue'),
    },
    {
      name: 'login',
      path: '/login',
      component: () => import('~/pages/login.vue'),
    },
    {
      name: 'register',
      path: '/register',
      component: () => import('~/pages/register.vue'),
    },
    {
      name: 'products',
      path: '/products',
      component: () => import('~/pages/main/products.vue'),
    },
    {
      name: 'categories',
      path: '/categories',
      component: () => import('~/pages/main/categories.vue'),
    },
  ],
} satisfies RouterConfig
