<script setup lang="ts">
import { useUserStore } from '~/stores/userStore';
const isError = ref(false)
const user = useUserStore()
const props = defineProps<{
            id: number,
            isFollow: boolean,
            name: string,
            profileImage: string,
            followCount: number,
            followerCount: number
    }>();

const emits = defineEmits<{
    follow: [id:number],
    unfollow: [id:number],
    following: [id:number],
    followed: [id:number],
}>();

const state = ref(props.isFollow)
const toggleFollow = () => {
    if(props.id == user.userId){
        isError.value = true
        return
    }
    state.value = !state.value
}

const toBack = () =>{
    navigateTo("/home")
}

onUnmounted(()=>{
    if(props.id == user.userId || props.state == props.isFollow){
        return 
    }
    else if(state.value){
        emits('follow',props.id)
    }
    else{
        emits('unfollow',props.id)
    }
})

console.log("isFollow::",props.isFollow)

</script>

<template>
    <v-container class="d-flex flex-column justify-center items-center">
        <div class="flex items-center gap-3 m-4 text-white/90 mr-auto">
            <v-btn
            density="comfortable"
            variant="text"
            class="rounded-full hover:opacity-90"
            icon="mdi-arrow-left"
            @click="toBack"
            />
            <span class="text-sm opacity-80">ホームへ</span>
        </div>
        <v-avatar :image="props.profileImage" size="200" class="mx-auto my-3" />
        <div class="text-h3">{{ props.name }}</div>
        <v-container class="d-flex mx-auto my-10 gap-10">
            <div class="mx-auto bg-grey-darken-3 p-5 px-10 rounded-lg" @click="emits('following',props.id)">
                フォロー中 : {{ props.followCount }}
            </div>
            <div class="mx-auto bg-grey-darken-3 p-5 px-10 rounded-lg" @click="emits('followed',props.id)">
                フォロワー : {{ props.followerCount }}
            </div>
        </v-container>
        <v-container class="d-flex mx-auto my-10 gap-10">
            <v-btn class="bg-red mx-auto p-5 px-10 rounded-lg" color="primary" v-bind:class="{'bg-red': state===true, 'bg-blue': state=== false}" @click="toggleFollow">
                {{ state ? "フォロー解除" : "フォロー" }}
            </v-btn>
            <v-snackbar class="mb-20" v-model="isError"
                    multi-line>
                    cannot follow yourself 
                <template v-slot:actions>
                    <v-btn
                        color="red"
                        variant="text"
                        @click="isError = false">
                        Close    
                    </v-btn>
                </template>        
            </v-snackbar>
        </v-container>
    </v-container>
</template>
