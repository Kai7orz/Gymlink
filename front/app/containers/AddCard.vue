<script setup lang="ts">
import UiAddCard from "~/components/UiAddCard.vue";
import { illustrations } from "~/data/illustrations";
import { useUserStore } from "~/stores/userStore";

const isShownMenu = ref(false);
const time = ref("");
const date = ref("2025/09/27");
const comment = ref("");
const imageUrl = ref("");
const auth = useAuthStore();
// イラスト一覧を読み込んで propsとして渡して表示する
const illustrationsObjs = illustrations;

const addCard = async () => {
  console.log("add card of add idToken:", auth.idToken);
  const TOKEN = auth.idToken;
  await $fetch("/api/userRecord", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + TOKEN,
      "Content-Type": "application/json",
    },
    body: {
      object_key: imageUrl.value,
      clean_up_time: Number(time.value),
      clean_up_date: date.value,
      comment: comment.value,
    },
  },
  );

  navigateTo("/home");
};

const closeMenu = () => {
  isShownMenu.value = false;
};

const openMenu = () => {
  navigateTo("/upload");
};

const select = (imageId: string) => {
  console.log("ttest", imageId);
  imageUrl.value = illustrationsObjs[imageId]?.src;
};

</script>

<template>
    <ui-add-card
                v-model:clean-time="time"
                v-model:date="date"
                v-model:comment="comment"
                v-model:is-shown-menu=isShownMenu
                :image-url="imageUrl"
                :illust-objs="illustrationsObjs"
                @open="openMenu"
                @add="addCard"
                @close="closeMenu"
                @select="select" />
</template>
