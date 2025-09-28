<script setup lang="ts">
import { useDetailStore } from '~/stores/detailStore';

const route = useRoute();
const detailStore = useDetailStore();
const user = useUserStore();
const toBack = () => {
    if(user.userName === detailStore.detailName) navigateTo({ name: 'home' });
    else navigateTo({ name: 'share'});
}

</script>

<template>
    <div>
        <v-hover v-slot="{ isHovering, props: hoverProps}">
            <v-icon icon="mdi-arrow-left" class="m-3" :size="isHovering ? 'x-large':'small'" v-bind="hoverProps" @click="toBack"></v-icon>     
        </v-hover>
        <div class="d-flex flex-column justify-center align-center mx-auto gap-2 m-10">
            <div class="text-h2 m-3">{{ detailStore.detailName }} </div>
            <img :src="detailStore.detailImage" class="w-1/2 h-50" />
            <div class="text-h2 m-5">{{ detailStore.detailDate }}</div>
            <div class="text-h2 m-5 pb-0">運動時間：{{ detailStore.detailTime }}分</div>
            <div class="text-h2 m-5" color="grey">コメント：{{ detailStore.detailComment }}</div>
        </div>
    </div>
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
