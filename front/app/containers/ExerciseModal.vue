<script lang="ts" setup>
    import UiExerciseModal from '~/components/UiExerciseModal.vue';
    import type { Illustration } from '~/type';

    const props = defineProps<{
        illustObjs: Record<string,Illustration>
        imageUrl: string
    }>();
    
    const isShownMenu = defineModel('isShownMenu')

    const date = defineModel<string>('date')
    const exerciseTime = defineModel<string>('exerciseTime')
    const comment = defineModel<string>('comment')
    const emits = defineEmits<{
        close: [],
        record: [],
        select: [imageId:string],
    }>();

    const onSelect = (imageId: string) => {
        emits('select',imageId)
    }

</script>

<template>
    <ui-exercise-modal 
                       v-model:exerciseTime="exerciseTime"
                       v-model:date="date"
                       v-model:comment="comment"
                       v-model:is-shown-menu="isShownMenu"
                       :imageUrl="props.imageUrl"
                       :illustObjs="props.illustObjs"
                       @record="emits('record')"
                       @close="emits('close')"
                       @select="onSelect"
                       />
</template>
