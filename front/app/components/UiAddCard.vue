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
    <v-card class="d-flex flex-column justify-center mx-auto gap-2 p-10 m-10">
        <v-card-title @click="emits('open')">+</v-card-title>
        <exercise-modal
                        v-model:imageUrl="imageUrl"
                        v-model:exerciseTime="exerciseTime"
                        v-model:date="date"
                        v-model:comment="comment"
                        v-model:is-shown-menu="isShownMenu" 
                        @record="emits('add')" 
                        @close="emits('close')"></exercise-modal>
    </v-card>
</template>
