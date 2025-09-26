<script setup lang="ts">
import { mergeProps } from 'vue';

const props = defineProps<{
            id: number,
            userId: number,
            userName: string,
            image: string
            time: number,
            date: string,
            comment: string,
            likesCount: number,
}>();

const emits = defineEmits<{
    detail: [id:number],
    like: [id:number], //exercise record のいいねボタン押した際にいいねした人一覧見えるようにするために，クリック時にロジック側にexercise record の id を渡すようにする
    account: [id:number],
}>();

const isProcessing = ref(false) // 二重でいいねを押せないようにする状態管理フラグ

const sleep = (ms: number) => new Promise<void>( resolve => setTimeout(resolve,ms))

const clicked = (id:number) => {
    emits('detail',id);
}

const onLike = async (id:number) => {
    if(isProcessing.value){
        return 
    }

    isProcessing.value = true 

    emits('like',id);
    
    try{
        await sleep(3000)
    } finally {
        isProcessing.value = false
    }

}

const onAccount = async (uid:number) => {
    emits('account',uid);
}

</script>

<template>
    <v-hover v-slot="{ isHovering, props: hoverProps }">
        <v-card class="d-flex flex-column justify-center w-1/4 h-1/2 mx-auto gap-2 p-10 px-7 pb-15 mx-8 my-2 bg-grey-lighten-3 rounded-xl" :class=" isHovering? 'on-hover':'non-hover'" v-bind="hoverProps" @click="() => clicked(props.id)">
            <img :src="props.image" class="w-60 h-50" />
            <v-card-title >{{ props.date }}</v-card-title>
            <v-card-subtitle class="pb-0">運動時間: {{ props.time }}分</v-card-subtitle>
            <div class="d-flex w-100 mt-10">
                <v-icon class="mr-15" size="30" icon="mdi-account-file-outline" @click="()=>onAccount(props.userId)" />  
                <v-icon class="ml-auto" icon="mdi-thumb-up" @click.stop="()=>onLike(props.id)"></v-icon>
                <div class="mx-2">{{ props.likesCount }}</div>
            </div> 
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
        border: 1px solid
     }

    .non-hover {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        border: 1px dotted
    }
</style>
