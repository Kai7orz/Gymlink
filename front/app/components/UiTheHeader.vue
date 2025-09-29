<script setup lang="ts">
  const emits = defineEmits<{
    home: [],
    account: [],
    monster: [],
    team: [],
    share:[],
  }>();
</script>

<template>
      <v-app-bar scroll-behavior="hide" color="black" class="mb-4 ">
        <v-app-bar-title>
          Gymlink
        </v-app-bar-title>
        <v-icon class="mx-5" icon="mdi-book-open-page-variant" @click="emits('share')"></v-icon>
        <v-icon class="mx-5" icon="mdi-home" @click="emits('home')"></v-icon>
        <v-icon class="mx-5" icon="mdi-sheep" @click="emits('monster')"></v-icon>
        <v-icon class="mx-5" icon="mdi-account" @click="emits('account')"></v-icon>
        <v-icon class="mx-5" icon="mdi-account-group" @click="emits('team')"></v-icon>
      </v-app-bar>
</template>
