<script setup lang="ts">
import Profile from "~/containers/Profile.vue"
const route = useRoute();
</script>

<template>
    <profile 
              :userId="route.params.id"/>
</template>
