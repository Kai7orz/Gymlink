<script setup lang="ts">
import Profile from "~/containers/Profile.vue"
const route = useRoute();
console.log("route id:::",route.value)
</script>

<template>
    <profile 
              :userId="route"/>
</template>
