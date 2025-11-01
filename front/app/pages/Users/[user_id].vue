<script lang="ts" setup>
import CardList from "~/containers/CardList.vue";
const isOwner = ref(false);
const route = useRoute();
</script>

<template>
    <v-container  class="p-100 min-h-screen" >
        <card-list :is-owner="isOwner" :is-target-exist="true" :target-id="Number(route.params.user_id)"/>
    </v-container>
</template>

<style>
  .upload-pos {
    display: flex;
    justify-content: center;
    margin: 20px;
    @media (max-width:400px) {
      margin:8px 0px;
    }
  }
</style>
