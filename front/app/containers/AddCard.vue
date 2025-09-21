<script setup lang="ts">
    import UiAddCard from '~/components/UiAddCard.vue';

    const addCard = ()=>{{
        console.log("add card")
        // ここでカード追加の処理を実装 API 絡み
    }}

    const closeMenu = ()=>{
        isShownMenu.value = false;
        console.log("close menu")
    }

    const openMenu = ()=>{
        isShownMenu.value = true;
        console.log("open menu")
    }

    const isShownMenu = ref(false);
</script>

<template>
    <ui-add-card v-model:is-shown-menu=isShownMenu @open="openMenu" @add="addCard" @close="closeMenu" />
</template>
