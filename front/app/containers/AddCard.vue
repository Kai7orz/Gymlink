<script setup lang="ts">
import UiAddCard from "~/components/UiAddCard.vue";

const isShownMenu = defineModel<boolean>("isShownMenu");
const cleanUpTime = defineModel<string>("cleanUpTime");
const cleanUpDate = defineModel<Date>("cleanUpDate");
const comment = defineModel<string>("comment");
const previewUrl = defineModel<string>("previewUrl");
const selectedFile = defineModel<File | null>("selectedFile");
const openMenu = () => {
  isShownMenu.value = true;
};

const closeMenu = () => {
  isShownMenu.value = false;
};

const getPreview = async (event: Event) => {
  event.preventDefault();
  if (!selectedFile.value) return;
  previewUrl.value = URL.createObjectURL(selectedFile.value);
};
</script>

<template>
    <ui-add-card
                v-model:clean-up-time="cleanUpTime"
                v-model:clean-up-date="cleanUpDate"
                v-model:comment="comment"
                v-model:is-shown-menu=isShownMenu
                @open="openMenu"
                @close="closeMenu"
            />
    <form class="upload-form">
        <v-file-input
          v-model="selectedFile"
          class="w-100 mx-auto p-3"
          accept="image/png"
          label="Select PNG Image"
          :show-size="true"
          @change="getPreview"
        />
    </form>
</template>