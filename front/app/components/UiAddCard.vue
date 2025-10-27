<script setup lang="ts">
import CleanModal from "~/containers/CleanModal.vue";
import type { Illustration } from "~/type";

const props = defineProps<{
  illustObjs: Record<string, Illustration>;
  imageUrl: string;
}>();

const emits = defineEmits<{
  add: [];
  close: [];
  open: [];
  select: [imageId: string];
}>();

const isShownMenu = defineModel("isShownMenu");
const date = defineModel<string>("date");
const cleanTime = defineModel<string>("cleanTime");
const comment = defineModel<string>("comment");
const onSelect = (imageId: string) => {
  emits("select", imageId);
};

</script>

<template>
        <v-icon class="mx-5 my-auto" size="50" icon="mdi-plus-box" @click="emits('open')"/>
        <clean-modal
                        v-model:clean-time="cleanTime"
                        v-model:date="date"
                        v-model:comment="comment"
                        v-model:is-shown-menu="isShownMenu"
                        :image-url="props.imageUrl"
                        :illust-objs="props.illustObjs"
                        @record="emits('add')"
                        @close="emits('close')"
                        @select="onSelect"/>

</template>
