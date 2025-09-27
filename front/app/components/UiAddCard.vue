<script setup lang="ts">
    import ExerciseModal from '~/components/UiExerciseModal.vue';

    const emits = defineEmits<{
        add: [],
        close: [],
        open: [],
    }>();

    const isShownMenu = defineModel('isShownMenu')
    const imageUrl = defineModel<string>('imageUrl')
    const date = defineModel<string>('date')
    const exerciseTime = defineModel<number>('exerciseTime')
    const comment = defineModel<string>('comment')

</script>

<template>
        <v-icon class="mx-5 my-auto" size="50" icon="mdi-plus-box" @click="emits('open')"></v-icon>
        <exercise-modal
                        v-model:imageUrl="imageUrl"
                        v-model:exerciseTime="exerciseTime"
                        v-model:date="date"
                        v-model:comment="comment"
                        v-model:is-shown-menu="isShownMenu" 
                        @record="emits('add')" 
                        @close="emits('close')"></exercise-modal>

</template>
