<script lang="ts" setup>
import UiCleanModal from "~/components/UiCleanModal.vue";

const isShownMenu = defineModel<boolean>("isShownMenu");
const cleanUpTime = defineModel<string>("cleanUpTime");
const cleanUpDate = defineModel<Date>("cleanUpDate");
const comment = defineModel<string>("comment");
const emits = defineEmits<{
  open: [];
  close: [];
}>();
</script>

<template>
    <ui-clean-modal
                       v-model:clean-up-time="cleanUpTime"
                       v-model:clean-up-date="cleanUpDate"
                       v-model:comment="comment"
                       v-model:is-shown-menu="isShownMenu"
                       @close="emits('close')"
                       />
</template>
