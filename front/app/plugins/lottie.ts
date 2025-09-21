import { defineNuxtPlugin } from 'nuxt/app'
import lottie from 'lottie-web'

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.provide('lottie', lottie)
})