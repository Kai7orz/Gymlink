<script setup lang="ts">
const uploadLabel = "upload";

const toUpload = () => {
  navigateTo("/upload");
};

</script>

<template>
      <div class="upload-pos">
          <v-btn color="light-blue-accent-4" @click="toUpload">{{ uploadLabel }}</v-btn>
      </div>
</template>
