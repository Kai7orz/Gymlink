<script setup lang="ts">
import LottieContainer from "~/containers/LottieContainer.vue";
const comment = "頑張りましょう";
</script>
<template>
  <div>
    <lottie-container />
    <h1 class="text-h4 text-center mt-10">
      {{ comment }}
    </h1>
  </div>
</template>
