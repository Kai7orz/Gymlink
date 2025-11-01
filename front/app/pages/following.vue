<script setup lang="ts">
import FollowingList from "~/containers/FollowingList.vue";
const pageLabel = "Following Users";
</script>

<template>
<div class="text-center">
  <h1 class="text-4xl m-10">{{ pageLabel }}</h1>
  <following-list />
</div>
</template>
