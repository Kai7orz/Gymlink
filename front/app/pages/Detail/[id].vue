<script setup lang="ts">
import CardDetail from "~/containers/CardDetail.vue";
const route = useRoute();
console.log("isowner:::", route.params.isOwner);

</script>

<template>
        <div>
          <card-detail :id="Number(route.params.id)" :is-owner="Boolean(route.params.isOwner)"/>
        </div>
</template>
