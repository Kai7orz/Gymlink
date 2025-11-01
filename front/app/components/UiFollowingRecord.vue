<script setup lang="ts">
// props で user id 受け取る
// button click時に user id を emits させる
const props = defineProps<{
  userId: number;
  userName: string;
}>();

const emits = defineEmits<{
  user: [userId: number];
}>();

const btnHeight = "100px";

const toUserPage = (userId: number) => {
  emits("user", userId);
};
</script>

<template>
  <div class="flex justify-center record-design">
    <v-btn :height="btnHeight" class="button-design" @click="()=> toUserPage(props.userId)" />
  </div>
</template>

<style>
  .button-design {
    width: 600px;
    background: rgba(255, 255, 255, 0.4);
    @media (max-width: 700px){
        width: 300
    }
  }

  .record-design {
    width: 80%;
    height: 80%;
  }
</style>
