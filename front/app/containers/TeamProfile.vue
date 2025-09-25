<script setup lang="ts">
    const { data } = await useFetch(
        '/api/teams'
    )


    console.log(data.value)

</script>

<template>
  <ui-team-profile
                  :data="data"
                  />
</template>
