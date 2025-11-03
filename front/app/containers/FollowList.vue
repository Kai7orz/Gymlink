<script setup lang="ts">
import { useUserStore } from "~/stores/userStore";
import type { UserType } from "~/type";
// followinglist コンポーネントに  :isFollowing をつけて props 渡す．
// true: フォローユーザーリストを返す
// false: フォロワーユーザーリストを返す
const props = defineProps<{
  isFollowing: boolean;
}>();
const user = useUserStore();
const auth = useAuthStore();
const TOKEN = auth.idToken;
const uid = Number(user.userId);
const url = props.isFollowing ? `/api/following/${uid}` : `/api/followed/${uid}`;

const users = await $fetch<UserType[]>(url, {
  headers: {
    "Authorization": "Bearer " + TOKEN,
    "Content-Type": "application/json",
  },
});
// user id に基づいたレコード情報を取得する（とりあえずユーザー名とユーザーレコードページへの遷移ボタン）
const toFollowPage = (userId: number) => {
  navigateTo(`/users/${userId}`);
};
</script>

<template>
   <ui-follow-list :following-users="users" @user="toFollowPage"/>
</template>
