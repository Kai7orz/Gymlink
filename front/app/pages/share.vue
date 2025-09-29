<script lang="ts" setup>
    import CardList from '~/containers/CardList.vue';
    import AddCard from '~/containers/AddCard.vue';
    const isOwner=ref(false)
</script>

<template>
    <v-container  class="p-100 min-h-screen" style="background: linear-gradient(135deg, #0044ff, #ff5e5e);">
        <add-card />
        <card-list :is-owner="isOwner" />
    </v-container>
</template>
