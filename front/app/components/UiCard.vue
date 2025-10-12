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

const onAccount = async (uid: number) => {
    emits('account',uid)
} 

</script>

<template>
    <v-hover v-slot="{ isHovering, props: hoverProps }">
        <v-card class="d-flex flex-column w-1/2 h-45 mx-auto gap-2 p-10 px-7 pb-15  my-auto bg-grey-darken-4 rounded-xl card" :class=" isHovering? 'on-hover':'non-hover'" v-bind="hoverProps" @click="() => clicked(props.id)">
            <img :src="props.image" class="w-60 h-50 rounded-xl" />
            <v-card-title >{{ props.date }}</v-card-title>
            <v-card-subtitle class="pb-0">運動時間: {{ props.time }}分</v-card-subtitle>
            <div class="d-flex mt-10 gap-5">
                <v-icon class="mr-20" size="30" icon="mdi-account-file-outline" @click="()=>onAccount(props.userId)" />
                <div class="d-flex">
                    <v-icon class="mx-3" icon="mdi-thumb-up" @click.stop="()=>onLike(props.id)"></v-icon>
                    <div>{{ props.likesCount }}</div>
                </div>
            </div>        
        </v-card>
    </v-hover>
</template>

<style scoped>

.card {

  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);

  border: 1px solid rgba(255, 255, 255, 0.25); /* 縁取りでガラス感 */
  border-radius: 20px;

  color: #ffffff; /* 背景に溶けないよう白字 */
  padding: 1rem;
  box-shadow: 0 8px 32px rgba(0,0,0,0.3); /* 浮かせる影 */
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card:hover {
  transform: translateY(-6px);
  box-shadow: 0 12px 40px rgba(0,0,0,0.35);
}
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
