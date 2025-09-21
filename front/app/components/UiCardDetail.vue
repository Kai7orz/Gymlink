<script setup lang="ts">
import { useDetailStore } from '~/stores/detailStore';

const route = useRoute();
const detailStore = useDetailStore();

const toHome = () => {
    navigateTo({ name: 'Prototype' });
}

</script>

<template>
        <v-hover v-slot="{ isHovering, props: hoverProps}">
            <v-icon icon="mdi-arrow-left" class="m-3" :size="isHovering ? 'x-large':'small'" v-bind="hoverProps" @click="toHome"></v-icon>     
        </v-hover>
        <v-card class="d-flex flex-column justify-center align-center mx-auto gap-2 p-10 m-10">
            <div>exercise id:{{ route.params.id }} </div>
            <img :src="detailStore.detailImage" class="w-60 h-50" />
            <v-card-title >{{ detailStore.detailDate }}</v-card-title>
            <v-card-subtitle class="pb-0">運動時間：{{ detailStore.detailTime }}分</v-card-subtitle>
            <v-card-text class="pt-0" color="grey">コメント：{{ detailStore.detailComment }}</v-card-text>
        </v-card>
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
</style>
