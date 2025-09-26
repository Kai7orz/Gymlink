<script setup lang="ts">
    import { useUserStore } from '~/stores/userStore';
    const user = useUserStore()
    const {data} = await useFetch(
        `/api/user_profiles/${user.userId}`
    );
    console.log("get profile:",data.value.name,data.value.profileImage)
</script>

<template>
    <ui-profile 
                :id="data.id" 
                :name="data.name" 
                :profileImage="data.profile_image"
                :followCount="data.follow_count"
                :followerCount="data.follower_count"
                />
</template>
