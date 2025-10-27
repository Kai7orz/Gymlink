<script setup lang="ts">
import TheHeader from "~/containers/TheHeader.vue";
</script>

<template>
  <NuxtLayout>
    <v-app>
      <the-header/>
      <v-main class="min-h-screen custom-bg" >
        <NuxtPage />
      </v-main>
    </v-app>
  </NuxtLayout>
</template>

<style>
main {
  min-height: 100vh;
  margin: 0;
}

@keyframes gradientMove {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.custom-bg{
  background: linear-gradient(135deg, #85a1f0, #e6e6e6);
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
