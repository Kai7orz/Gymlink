<script setup lang="ts">
    // ユーザーid を params へ渡して遷移する
    const userId = 1

    const toHome = () => {
        navigateTo('/home')
    }

    const toAccount = async () => {
        navigateTo({name: 'Account-id', params: {id: userId }})
    }

    const toMonster = () => {
        navigateTo('/monster')
    }

    const toTeam = () => {
        navigateTo('/team')
    }

    const toShare = () => {
        navigateTo('/share')
    }

</script>

<template>
    <ui-the-header 
                  @home="toHome" 
                  @account="toAccount" 
                  @monster="toMonster" 
                  @team="toTeam"
                  @share="toShare" />
</template>
