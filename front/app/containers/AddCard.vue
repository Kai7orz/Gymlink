<script setup lang="ts">
    import UiAddCard from '~/components/UiAddCard.vue';
import { illustrations } from '~/data/illustrations';
    import { useUserStore } from '~/stores/userStore';

    const isShownMenu = ref(false);
    const time = ref("")
    const date = ref("2025/09/27")
    const comment = ref('')
    const imageUrl = ref('')
    const user = useUserStore()
    const auth = useAuthStore()
    // イラスト一覧を読み込んで propsとして渡して表示する
    const illustrationsObjs = illustrations

    const addCard = async ()=>{
        console.log(imageUrl.value)
        const TOKEN = auth.idToken
        await $fetch('/api/userExercise',{
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + TOKEN,
                'Content-Type': 'application/json'
            },
            body: {
                "exercise_image": imageUrl.value,
                "exercise_time": Number(time.value),
                "exercise_date": date.value,
                "comment": comment.value,
            }
        }
        );

        navigateTo('/home')

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
