import { defineNuxtPlugin } from 'nuxt/app'
import vuetify from '~/utils/vuetify'
import { VDateInput } from 'vuetify/labs/VDateInput'

export default defineNuxtPlugin((app) => {
  app.vueApp.use(vuetify)
})