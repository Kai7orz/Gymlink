<script lang="ts" setup>
import CardList from "~/containers/CardList.vue";
const isOwner = ref(false);
const route = useRoute();
const router = useRouter();
const toBack = () => {
  router.back();
};
const toBackLabel = "戻る";
</script>

<template>
    <v-container  class="p-100 min-h-screen" >
      <div class="flex items-center gap-3 m-4 text-white/90 mr-auto">
              <v-btn
              density="comfortable"
              variant="text"
              class="rounded-full hover:opacity-90"
              icon="mdi-arrow-left"
              @click="toBack"
              />
              <span class="text-sm opacity-80">{{ toBackLabel }}</span>
      </div>
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
