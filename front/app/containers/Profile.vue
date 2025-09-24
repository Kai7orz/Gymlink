<script setup lang="ts">
    console.log()
    const {data} = await useFetch(
        '/api/users/profile'
    );
    console.log("get profile:",data.value.name,data.value.profileImage)
</script>

<template>
    <ui-profile 
                :id="data.id" 
                :name="data.name" 
                :profileImage="data.profileImage"
                :followCount="data.followCount"
                :followerCount="data.followerCount"
                />
</template>
