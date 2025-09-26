<script setup lang="ts">
    import UiAddCard from '~/components/UiAddCard.vue';
    import { useUserStore } from '~/stores/userStore';

    const isShownMenu = ref(false);
    const image_url = ref("/images/sportImage.png")
    const time = ref(0)
    const date = ref("2025/09/27")
    const comment = ref('')

    const user = useUserStore()

    const addCard = ()=>{
        console.log("add card")
        // ここでカード追加の処理を実装 API 絡み
        console.log("image->",image_url.value,time.value,date.value,comment.value)
        const TOKEN = user.idToken
        $fetch('/api/userExercise',{
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + TOKEN,
                'Content-Type': 'application/json'
            },
            body: {
                "image_url": image_url.value,
                "time": time.value,
                "date": date.value,
                "comment": comment.value,
            }
        }
        );
    }

    const closeMenu = ()=>{
        isShownMenu.value = false;
        console.log("close menu")
    }

    const openMenu = ()=>{
        isShownMenu.value = true;
        console.log("open menu")
    }


</script>

<template>
    <ui-add-card 
                v-model:imageUrl="image_url"
                v-model:exerciseTime="time"
                v-model:date="date"
                v-model:comment="comment"
                v-model:is-shown-menu=isShownMenu 
                @open="openMenu" 
                @add="addCard" 
                @close="closeMenu" />
</template>
