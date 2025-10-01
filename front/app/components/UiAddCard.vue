<script setup lang="ts">
    import ExerciseModal from '~/components/UiExerciseModal.vue';
    import type { Illustration } from '~/type';

    const props = defineProps<{
        illustObjs: Record<string,Illustration>
        imageUrl: string
    }>();

    const emits = defineEmits<{
        add: [],
        close: [],
        open: [],
        select: [imageId:string]
    }>();

    const isShownMenu = defineModel('isShownMenu')
    const date = defineModel<string>('date')
    const exerciseTime = defineModel<string>('exerciseTime')
    const comment = defineModel<string>('comment')
    const onSelect = (imageId: string) => {
        emits('select',imageId)
    }
   

</script>

<template>
        <v-icon class="mx-5 my-auto" size="50" icon="mdi-plus-box" @click="emits('open')"></v-icon>
        <exercise-modal
                        v-model:exerciseTime="exerciseTime"
                        v-model:date="date"
                        v-model:comment="comment"
                        v-model:is-shown-menu="isShownMenu"
                        :imageUrl="props.imageUrl"
                        :illustObjs="props.illustObjs"
                        @record="emits('add')" 
                        @close="emits('close')"
                        @select="onSelect"></exercise-modal>

</template>
