import type { RouterConfig } from '@nuxt/schema'

export default {
  // https://router.vuejs.org/api/interfaces/routeroptions.html#routes
  routes: (_routes) => [
    {
      path: '/',
      component: () => import('~/pages/index.vue')
    },
    {
      name: 'login',
      path: '/login',
      component: () => import('~/pages/login.vue')
    },
    {
      name: 'register',
      path: '/register',
      component: () => import('~/pages/register.vue')
    },
    {
      name: 'establishment',
      path: '/establishment',
      component: () => import('~/pages/main/establishment.vue')
    },
    {
      name: 'store',
      path: '/store',
      component: () => import('~/pages/main/store.vue')
    },
  ],
} satisfies RouterConfig
