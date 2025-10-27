<script setup lang="ts">
import Profile from "~/containers/Profile.vue";
const route = useRoute();
</script>

<template>
    <profile
              :user-id="route.params.id"/>
</template>
