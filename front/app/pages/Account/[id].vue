<script setup lang="ts">
import Profile from "~/containers/Profile.vue"
const route = useRoute();
console.log("route::",route.params.id)
</script>

<template>
    <profile 
              :userId="route.params.id"/>
</template>
