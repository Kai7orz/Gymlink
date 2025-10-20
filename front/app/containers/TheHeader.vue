<script setup lang="ts">

    import { useUserStore } from "~/stores/userStore"
    // ユーザーid を params へ渡して遷移する
    const user = useUserStore()

    const toHome = () => {
        navigateTo('/home')
    }

    const toAccount = async () => {
        navigateTo({name: 'Account-id', params: {id: user.userId }})
    }

    const toMonster = () => {
        navigateTo('/monster')
    }

    const toShare = () => {
        navigateTo('/share')
    }

    const toSignout = ()=>{
        navigateTo('/login')
        localStorage.removeItem("userName")
        localStorage.removeItem("userId")
    }

</script>

<template>
    <ui-the-header 
                  @home="toHome" 
                  @account="toAccount" 
                  @monster="toMonster"
                  @share="toShare"
                  @signout="toSignout"
                  />
</template>
