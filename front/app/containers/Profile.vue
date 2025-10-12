<script setup lang="ts">
    
    import { useUserStore } from '~/stores/userStore';
    const props = defineProps<{
        userId: string,
    }>();

    const user = useUserStore()
    const auth = useAuthStore()
    const TOKEN = auth.idToken
    const uid = Number(props.userId)
    const data = await $fetch(
        `/api/user_profiles/${uid}`,{
          headers: {
                        'Authorization': 'Bearer ' + TOKEN,
                        'Content-Type': 'application/json'
                    },
        }
    );

    console.log("profile id and data ",data.id," data->",data)

    const follow = async (id:number) => {
        await navigateTo('/following')
        console.log("follower_id:",user.userId)
        console.log("followed_id:",data.id)
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
