<script setup lang="ts">
    interface memberProfile {
        teamName: string
        memberName: string
        memberProfileImage: string
        memberExerciseCount: number
    }
    const props = defineProps<{
            data: memberProfile[]
        }>();

    console.log("team info: ",props.data)
</script>

<template>
    <v-card v-for="(item,index) in props.data" class="d-flex flex-column m-5">
        <v-avatar :image="item.memberProfileImage" size="200" class="mx-auto my-3" />
        <div class="text-h3 mx-auto my-10">{{ item.memberName }}</div>
    </v-card>

</template>
