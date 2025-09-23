<script setup lang="ts">
import { mergeProps } from 'vue';

const props = defineProps<{
            id: number,
            image: string
            time: number,
            date: string,
            comment: string,
}>();

const emits = defineEmits<{
    detail: [id:number],
}>();

const clicked = (id:number) => {
    emits('detail',id);
}
</script>

<template>
    <v-hover v-slot="{ isHovering, props: hoverProps }">
        <v-card class="d-flex flex-column justify-center w-1/4 h-1/2 mx-auto gap-2 p-10 px-7 pb-15 mx-8 my-2 bg-black rounded-xl" :class=" isHovering? 'on-hover':'non-hover'" v-bind="hoverProps" @click="() => clicked(props.id)">
            <img :src="props.image" class="w-60 h-50" />
            <v-card-title >{{ props.date }}</v-card-title>
            <v-card-subtitle class="pb-0">運動時間: {{ props.time }}分</v-card-subtitle>
        </v-card>
    </v-hover>
</template>

<style scoped>
    .on-hover {
        min-width:300px;
        min-height: 250px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
     }

    .non-hover {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
</style>
