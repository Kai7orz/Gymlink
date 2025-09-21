<script setup lang="ts">
    import ExerciseModal from '~/components/UiExerciseModal.vue';

    const emits = defineEmits<{
        add: [],
        close: [],
        open: [],
    }>();

    const isShownMenu = defineModel('isShownMenu')

</script>

<template>
    <v-card class="d-flex flex-column justify-center mx-auto gap-2 p-10 m-10">
        <v-card-title @click="emits('open')">+</v-card-title>
        <exercise-modal v-model:is-shown-menu="isShownMenu" @close="emits('close')"></exercise-modal>
    </v-card>
</template>
