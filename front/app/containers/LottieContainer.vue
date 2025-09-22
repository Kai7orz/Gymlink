<script setup lang="ts">
import animationData from '~/assets/animations/huwaData.json'

</script>

<template>
      <ui-lottie-container
        class="mx-auto mt-10"
        :animation-data="animationData"
        :loop="true"
        :autoplay="true"
        width="320px"
        height="320px"
      />
</template>
