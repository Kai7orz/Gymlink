<script setup lang="ts">
    import UiAddCard from '~/components/UiAddCard.vue';
import { illustrations } from '~/data/illustrations';
    import { useUserStore } from '~/stores/userStore';

    const isShownMenu = ref(false);
    const time = ref(0)
    const date = ref("2025/09/27")
    const comment = ref('')
    const imageUrl = ref('')
    const user = useUserStore()

    // イラスト一覧を読み込んで propsとして渡して表示する
    const illustrationsObjs = illustrations


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

    const select = (imageId:string) => {
        console.log("ttest",imageId)
        imageUrl.value = illustrationsObjs[imageId]?.src
    }

</script>

<template>
    <ui-add-card 
                v-model:exerciseTime="time"
                v-model:date="date"
                v-model:comment="comment"
                v-model:is-shown-menu=isShownMenu
                :imageUrl="imageUrl"
                :illustObjs="illustrationsObjs"
                @open="openMenu" 
                @add="addCard" 
                @close="closeMenu"
                @select="select" />
</template>
