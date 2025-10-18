<script setup lang="ts">
  import CardDetail from '~/containers/CardDetail.vue';
  const  route = useRoute();
</script>

<template>
        <div>
          <card-detail :id="route.params.id"/>
        </div>
</template>
