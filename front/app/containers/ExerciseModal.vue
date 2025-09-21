<script lang="ts" setup>
    import UiExerciseModal from '~/components/UiExerciseModal.vue';

    const emits = defineEmits<{
        close: [],
    }>();

    const isShownMenu = defineModel('isShownMenu')
</script>

<template>
    <ui-exercise-modal v-model:is-shown-menu="isShownMenu" @close="emits('close')"/>
</template>
