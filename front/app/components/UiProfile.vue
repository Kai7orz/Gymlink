<script setup lang="ts">
const props = defineProps<{
            id: number,
            name: string,
            profileImage: string,
            followCount: number,
            followerCount: number
    }>();

</script>

<template>
    <v-container class="d-flex flex-column justify-center items-center">
        <v-avatar :image="props.profileImage" size="200" class="mx-auto my-3" />
        <div class="text-h3">{{ props.name }}</div>
        <v-container class="d-flex mx-auto my-10 gap-10">
            <div class="mx-auto bg-grey-darken-3 p-5 px-10 rounded-lg">
                フォロー中 : {{ props.followCount }}
            </div>
            <div class="mx-auto bg-grey-darken-3 p-5 px-10 rounded-lg">
                フォロワー : {{ props.followerCount }}
            </div>
        </v-container>
    </v-container>
</template>
