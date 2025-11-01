<script setup lang="ts">
import { useUserStore } from "~/stores/userStore";
const props = defineProps<{
  id: number;
  isFollow: boolean;
  name: string;
  profileImage: string;
  followCount: number;
  followerCount: number;
}>();

const emits = defineEmits<{
  follow: [id: number];
  unfollow: [id: number];
  back: [];
  following: [id: number];
  followed: [id: number];
}>();

const isError = ref(false);
const user = useUserStore();
const state = ref(props.isFollow);
const toBackLabel = "戻る";
const followingLabel = "フォロー中: ";
const followedLabel = "フォロワー: ";
const followWarning = "cannot follow yourself";
const closeLabel = "close";

const toggleFollow = () => {
  if (props.id == user.userId) {
    isError.value = true;
    return;
  }
  state.value = !state.value;
};

const toFollowing = () => {
  if (props.id == user.userId) navigateTo("/following"); // 現状は自身のフォロー一覧しか見られない仕様
  else return;
};

const toBack = () => {
  emits("back");
};

onUnmounted(() => {
  if (props.id == user.userId || state.value == props.isFollow) {
    return;
  }
  else if (state.value) {
    emits("follow", props.id);
  }
  else {
    emits("unfollow", props.id);
  }
});

</script>

<template>
    <v-container class="d-flex flex-column justify-center items-center">
        <div class="flex items-center gap-3 m-4 text-white/90 mr-auto">
            <v-btn
            density="comfortable"
            variant="text"
            class="rounded-full hover:opacity-90"
            icon="mdi-arrow-left"
            @click="toBack"
            />
            <span class="text-sm opacity-80">{{ toBackLabel }}</span>
        </div>
        <v-avatar :image="props.profileImage" size="200" class="mx-auto my-3" />
        <div class="name-size">{{ props.name }}</div>
        <v-container class="d-flex mx-auto my-10 gap-10">
            <v-btn  class="mx-auto p-4 bg-grey-darken-3 rounded-lg" @click="toFollowing">{{ followingLabel }} {{ props.followCount }}</v-btn>
            <v-btn  class="mx-auto p-4 bg-grey-darken-3 rounded-lg" @click="emits('followed',props.id)">{{ followedLabel }} {{ props.followerCount }}</v-btn>
        </v-container>
        <v-container class="d-flex mx-auto my-10 gap-10">
            <v-btn class="bg-red mx-auto p-5 px-10 rounded-lg" color="primary" :class="{'bg-red': state===true, 'bg-blue': state=== false}" @click="toggleFollow">
                {{ state ? "フォロー解除" : "フォロー" }}
            </v-btn>
            <v-snackbar
v-model="isError" class="mb-20"
                    multi-line>
                    {{ followWarning }}
                <template #actions>
                    <v-btn
                        color="red"
                        variant="text"
                        @click="isError = false">
                        {{ closeLabel }}
                    </v-btn>
                </template>
            </v-snackbar>
        </v-container>
    </v-container>
</template>

<style>
 .name-size {
    font-size: 60px;
    @media (max-width:700px) {
        font-size: 40px;
    }
 }
</style>
