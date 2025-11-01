<script setup lang="ts">
import { useUserStore } from "~/stores/userStore";
const user = useUserStore();
const auth = useAuthStore();
const TOKEN = auth.idToken;
const uid = Number(user.userId);
const followingUsers = await $fetch(`/api/following/${uid}`, {
  headers: {
    "Authorization": "Bearer " + TOKEN,
    "Content-Type": "application/json",
  },
});
console.log("followingUsers:", followingUsers);
</script>

<template>
  <div>test</div>
</template>
