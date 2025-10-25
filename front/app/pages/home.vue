<script lang="ts" setup>
    import CardList from '~/containers/CardList.vue';
    import AddCard from '~/containers/AddCard.vue';
    const isOwner=ref(true)
</script>

<template>
    <v-container  class="p-100 min-h-screen" >
        <add-card />
        <card-list :is-owner="isOwner" />
    </v-container>
</template>
