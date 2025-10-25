<script setup lang="ts">
  const emits = defineEmits<{
    home: [],
    account: [],
    monster: [],
    share:[],
    signout:[],
  }>();
</script>

<template>
      <v-app-bar scroll-behavior="hide" color="black" class="flex mb-4 ">
        <v-app-bar-title class="hidden md:inline-block">
          KMonitor
        </v-app-bar-title>
        <v-icon class="mx-5" icon="mdi-book-open-page-variant" @click="emits('share')"></v-icon>
        <v-icon class="mx-5" icon="mdi-home" @click="emits('home')"></v-icon>
        <v-icon class="mx-5" icon="mdi-sheep" @click="emits('monster')"></v-icon>
        <v-icon class="mx-5" icon="mdi-account" @click="emits('account')"></v-icon>
        <v-icon class="mx-5 bg-red" icon="mdi-exit-run" @click="emits('signout')"></v-icon>
      </v-app-bar>
</template>