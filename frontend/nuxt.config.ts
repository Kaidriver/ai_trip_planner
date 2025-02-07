// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-01-13',
  devtools: { enabled: true },
  css: ['~/assets/main.scss'],
  ssr: false,
  modules: ['@nuxt/ui'],
  runtimeConfig: {
    public: {
      secret: process.env.API_KEY,
      prod_endpoint: process.env.API_ENDPOINT,
    }
  },
  nitro: {
    preset: 'static'
  }
})