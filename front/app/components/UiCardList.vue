<script setup lang="ts">
import type { exerciseRecordType } from '~/type';

    const props = defineProps<{
        exerciseMocksList: exerciseRecordType[]
    }>();

    const emits = defineEmits<{
        detail: [id: number],
    }>();

    const toDetail = (id:number) => {
        emits('detail',id)
    }
    console.log("props.list:",props.exerciseMocksList)

</script>

<template>
    <v-container class="d-flex flex-wrap justify-center gap-5">
        <ui-card v-for="(exerciseMock, index) in props.exerciseMocksList" :key="index" :id="exerciseMock.id" :image="exerciseMock.imageUrl" :time="exerciseMock.time" :date="exerciseMock.date" :comment="exerciseMock.comment" @detail="toDetail"/>
    </v-container>
</template>
