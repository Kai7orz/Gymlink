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

<style>
main {
  background: linear-gradient(135deg, #0044ff, #ff5e5e);
  height: 100vh;
  margin: 0;
}

@keyframes gradientMove {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.page-enter-active,
.page-leave-active {
  transition: all 0.2s;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(1rem);
}
</style>
