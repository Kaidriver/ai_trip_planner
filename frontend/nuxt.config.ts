// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-01-13',
  devtools: { enabled: true },
  css: ['~/assets/main.scss'],
  modules: ['@nuxt/ui'],
  ssr: false,
  runtimeConfig: {
    public: {
      secret: process.env.API_KEY,
      prod_endpoint: process.env.API_ENDPOINT,
    }
  },
})