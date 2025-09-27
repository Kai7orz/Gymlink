<script setup lang="ts">
    
    import { useUserStore } from '~/stores/userStore';
    const user = useUserStore()
    const TOKEN = user.idToken
    const {data} = await useFetch(
        `/api/user_profiles/${user.userId}`
    );
    console.log("get profile:",data.value.name,data.value.profileImage)

    const follow = async (id:number) => {
        await $fetch("/api/follows", {
                        method: 'POST',
                        headers: {
                            'Authorization': 'Bearer ' + TOKEN,
                            'Content-Type': 'application/json'
                        },
                        body: {
                            follower_id: user.userId,
                            followed_id: id
                        },
                    })
    }

    const unfollow = async (id:number) => {
        await $fetch("/api/unfollows", {
                        method: 'POST',
                        headers: {
                            'Authorization': 'Bearer ' + TOKEN,
                            'Content-Type': 'application/json'
                        },
                        body: {
                            follower_id: user.userId,
                            followed_id: id
                        },
                    })
    }


</script>

<template>
    <ui-profile 
                :id="data.id" 
                :name="data.name" 
                :profileImage="data.profile_image"
                :followCount="data.follow_count"
                :followerCount="data.follower_count"
                @follow="follow"
                @unfollow="unfollow"
                @following="following"
                @followd="followed"
                />
</template>
