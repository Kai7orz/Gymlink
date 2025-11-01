<script lang="ts" setup>
import UploadBtn from "~/containers/UploadBtn.vue";
import CardList from "~/containers/CardList.vue";
const isOwner = ref(true);
</script>

<template>
    <v-container  class="p-100 min-h-screen" >
        <upload-btn />
        <card-list :is-owner="isOwner" :is-target-exist="false" :target-id="0"/>
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
