<script setup lang="ts">
import type { ExerciseRecordType } from '~/type';

    const props = defineProps<{
        exerciseMocksList: ExerciseRecordType[]
    }>();

    const emits = defineEmits<{
        detail: [id: number],
        like: [id: number],
    }>();

    const toDetail = (id:number) => {
        emits('detail',id)
    }

    const like = (id:number) => {
        emits('like',id)
    }


</script>

<template>
    <v-container class="d-flex flex-wrap justify-center gap-5">
        <ui-card v-for="(exerciseMock, index) in props.exerciseMocksList" 
                                                                         :key="index" 
                                                                         :id="exerciseMock.id" 
                                                                         :image="exerciseMock.image_url" 
                                                                         :time="exerciseMock.time"
                                                                         :date="exerciseMock.date" 
                                                                         :comment="exerciseMock.comment" 
                                                                         :likesCount="exerciseMock.likes_count" 
                                                                         @detail="toDetail"
                                                                         @like="like"
                                                                         />
    </v-container>
</template>
