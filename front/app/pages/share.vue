<script lang="ts" setup>
import CardList from "~/containers/CardList.vue";
const isOwner = ref(false);

const toUpload = () => {
  navigateTo("/upload");
};
</script>

<template>
    <v-container  class="p-100 min-h-screen" >
        <div class="upload-pos">
          <v-btn color="light-blue-accent-4" @click="toUpload">upload</v-btn>
        </div>
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
