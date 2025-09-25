<script setup lang="ts">
    const { data: team } = await useFetch(
        '/api/teams'
    )

    const teamName   = computed(() => team.value?.teamName  ?? '')
    const teamTarget = computed(() => team.value?.teamTarget ?? 0)
    const teamMember = computed(() => team.value?.member    ?? [])
    console.log("team:->:",team.value)
</script>

<template>
  <ui-team-profile
                   :team-name="teamName"
                   :team-target="teamTarget"
                   :team-member="teamMember"
                  />
</template>
