<script setup lang="ts">
import animationData from '~/assets/animations/dolphin.json'

</script>

<template>
      <ui-lottie-container
        :animation-data="animationData"
        :loop="true"
        :autoplay="true"
        width="320px"
        height="320px"
      />
</template>
