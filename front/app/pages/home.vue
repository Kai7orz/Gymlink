<script lang="ts" setup>
    import CardList from '~/containers/CardList.vue';
    import AddCard from '~/containers/AddCard.vue';
    const isOwner=ref(true)
</script>

<template>
    <v-container class="d-flex">
        <add-card />
        <card-list :isOwner="isOwner" />
    </v-container>
</template>
