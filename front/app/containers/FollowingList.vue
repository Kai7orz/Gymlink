<script setup lang="ts">
import { useUserStore } from "~/stores/userStore";

const user = useUserStore();
const auth = useAuthStore();
const TOKEN = auth.idToken;
const uid = Number(user.userId);
const followingUsers = await $fetch<number[]>(`/api/following/${uid}`, {
  headers: {
    "Authorization": "Bearer " + TOKEN,
    "Content-Type": "application/json",
  },
});
// user id に基づいたレコード情報を取得する（とりあえずユーザー名とユーザーレコードページへの遷移ボタン）
const toFollowingPage = (userId: number) => {
  navigateTo(`/users/${userId}`);
};
</script>

<template>
   <ui-following-list :following-users="followingUsers" @user="toFollowingPage"/>
</template>
