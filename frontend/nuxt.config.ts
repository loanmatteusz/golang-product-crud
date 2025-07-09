// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite"

export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  srcDir: 'src/',
  modules: [
    "@nuxt/ui",
    'nuxt-toast',
    'shadcn-nuxt',
    '@nuxt/icon',
  ],
  ui: {
    prefix: 'Nuxt'
  },
  icon: {
    customCollections: [
      {
        prefix: 'svg',
        dir: './assets/svg',
      },
    ],
  },
  shadcn: {
    prefix: '',
    /**
     * Directory that the component lives in.
     * @default "./components/ui"
     */
    componentDir: './components/ui'
  },
  css: ['./assets/css/main.css'],
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
  runtimeConfig: {
    public: {
      apiUrl: process.env.API_URL || 'http://localhost:3333',
    }
  },
  vue: {
    compilerOptions: {
      isCustomElement: (tag) => ['lite-youtube'].includes(tag),
    },
  },
});