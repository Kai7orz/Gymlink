<script setup lang="ts">
  import TheHeader from '~/containers/TheHeader.vue';
</script>

<template>
  <NuxtLayout>
    <v-app>
      <the-header>
      </the-header>
      <v-main>
        <NuxtPage />
      </v-main>
    </v-app>
  </NuxtLayout>
</template>
