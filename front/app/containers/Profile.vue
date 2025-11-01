<script setup lang="ts">
import { useUserStore } from "~/stores/userStore";
const props = defineProps<{
  userId: string;
}>();

const router = useRouter();
const user = useUserStore();
const auth = useAuthStore();
const TOKEN = auth.idToken;
const uid = Number(props.userId);
const data = await $fetch(
  `/api/user_profiles/${uid}`, {
    headers: {
      "Authorization": "Bearer " + TOKEN,
      "Content-Type": "application/json",
    },
  },
);

const isFollow = await $fetch(
  `/api/follows/${uid}`, {
    headers: {
      "Authorization": "Bearer " + TOKEN,
      "Content-Type": "application/json",
    },
  },
);

const follow = async (id: number) => {
  await $fetch("/api/follows", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + TOKEN,
      "Content-Type": "application/json",
    },
    body: {
      follower_id: user.userId,
      followed_id: id,
    },
  });
};

const unfollow = async (id: number) => {
  await $fetch("/api/unfollows", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + TOKEN,
      "Content-Type": "application/json",
    },
    body: {
      follower_id: user.userId,
      followed_id: id,
    },
  });
};

const toBack = () => {
  router.back();
};

</script>

<template>
    <ui-profile
                :id="data.id"
                :is-follow="isFollow.followed"
                :name="data.name"
                :profile-image="data.profile_image"
                :follow-count="data.follow_count"
                :follower-count="data.follower_count"
                @follow="follow"
                @unfollow="unfollow"
                @back="toBack"
                />
</template>
