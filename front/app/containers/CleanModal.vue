<script lang="ts" setup>
import UiCleanModal from "~/components/UiCleanModal.vue";
import type { Illustration } from "~/type";

const props = defineProps<{
  illustObjs: Record<string, Illustration>;
  imageUrl: string;
}>();

const isShownMenu = defineModel<boolean>("isShownMenu");

const date = defineModel<string>("date");
const cleanTime = defineModel<string>("cleanTime");
const comment = defineModel<string>("comment");
const emits = defineEmits<{
  close: [];
  record: [];
  select: [imageId: string];
}>();

const onSelect = (imageId: string) => {
  emits("select", imageId);
};

</script>

<template>
    <ui-clean-modal
                       v-model:clean-time="cleanTime"
                       v-model:date="date"
                       v-model:comment="comment"
                       v-model:is-shown-menu="isShownMenu"
                       :image-url="props.imageUrl"
                       :illust-objs="props.illustObjs"
                       @record="emits('record')"
                       @close="emits('close')"
                       @select="onSelect"
                       />
</template>
