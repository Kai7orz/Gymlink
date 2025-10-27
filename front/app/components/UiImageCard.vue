<script setup lang="ts">
    const props = defineProps<{
        image_url: string
    }>();
</script>

<template>
    <v-sheet class="p-5 w-100 bg-black" rounded="rounded">
        <img  :src=props.image_url >
    </v-sheet>
</template>