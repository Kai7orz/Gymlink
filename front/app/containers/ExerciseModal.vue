<script lang="ts" setup>
    import UiExerciseModal from '~/components/UiExerciseModal.vue';

    
    const isShownMenu = defineModel('isShownMenu')
    const imageUrl = defineModel<string>('imageUrl')
    const date = defineModel<string>('date')
    const exerciseTime = defineModel<number>('exerciseTime')
    const comment = defineModel<string>('comment')

    const emits = defineEmits<{
        close: [],
        record: [],
    }>();
</script>

<template>
    <ui-exercise-modal v-model:imageUrl="imageUrl"
                       v-model:exerciseTime="exerciseTime"
                       v-model:date="date"
                       v-model:comment="comment"
                       v-model:is-shown-menu="isShownMenu"
                       @record="emits('record')"
                       @close="emits('close')"/>
</template>
