<script setup lang="ts">
import CleanModal from "~/containers/CleanModal.vue";

const emits = defineEmits<{
  add: [];
  close: [];
  open: [];
}>();

const isShownMenu = defineModel<boolean>("isShownMenu");
const cleanUpTime = defineModel<string>("cleanUpTime");
const cleanUpDate = defineModel<Date>("cleanUpDate");
const comment = defineModel<string>("comment");
</script>

<template>
        <v-icon class="mx-5 my-auto" size="50" icon="mdi-plus-box" @click="emits('open')"/>
        <clean-modal
                        v-model:clean-up-time="cleanUpTime"
                        v-model:clean-up-date="cleanUpDate"
                        v-model:comment="comment"
                        v-model:is-shown-menu="isShownMenu"
                        @close="emits('close')"
                        />

</template>