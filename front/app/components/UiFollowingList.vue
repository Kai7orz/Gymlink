<script setup lang="ts">
const props = defineProps<{
  followingUsers: number[];
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
        <div v-for="(userId,index) in props.followingUsers" class="flex justify-center h-100" :key="index">
          <ui-following-record :user-id="userId" @user="toUserPage"/>
        </div>
    </v-container>
</template>
