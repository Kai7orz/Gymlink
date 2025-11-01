<script setup lang="ts">
import FollowingList from "~/containers/FollowingList.vue";
const router = useRouter();
const toBack = () => {
  router.back();
};
const pageLabel = "Following Users";
const toBackLabel = "戻る";
</script>

<template>
<v-container class="text-center">
  <v-container>
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
    <h1 class="text-4xl m-10">{{ pageLabel }}</h1>
    <following-list />
  </v-container>
</v-container>
</template>
