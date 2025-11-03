<script setup lang="ts">
import type { UserType } from "~/type";

const props = defineProps<{
  followingUsers: UserType[];
}>();

const emits = defineEmits<{
  user: [userId: number];
}>();

const toUserPage = (userId: number) => {
  emits("user", userId);
};
</script>

<template>
    <v-container class="h-100">
        <div v-for="(user,index) in props.followingUsers" class="flex justify-center h-100" :key="index">
          <ui-following-record class="m-4" :user-id="user.id" :user-name="user.name" @user="toUserPage"/>
        </div>
    </v-container>
</template>
