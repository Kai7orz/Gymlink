<script setup lang="ts">
    interface MemberProfile {
        name: string
        profileImage: string
        exerciseCount: number
    }

    const props = defineProps<{
            teamName: string
            teamTarget: number
            teamMember: MemberProfile[]
        }>();

    const teamExerciseCounts = ref<number[]>(props.teamMember.map((item:MemberProfile)=> item.exerciseCount))        
</script>

<template>
    <v-container class="m-10 rounded bg-black">
        <v-sheet class="d-flex mx-auto my-10">
            <v-card v-for="(item,index) in props.teamMember" class="d-flex flex-row m-5 mx-auto">
                <v-card class="d-flex flex-column px-10">
                    <v-avatar :image="item.profileImage" size="200" class="mx-auto my-3 p-5" />
                    <div class="text-h3 mx-auto my-10">{{ item.name }}</div>
                    <div class="d-flex justify-center items-center mb-5">
                        <v-icon  v-if="index==0" class="mx-5" color="red" icon="mdi-circle"></v-icon>
                        <v-icon  v-if="index==1" class="mx-5" color="blue" icon="mdi-circle"></v-icon>
                        <v-icon  v-if="index==2" class="mx-5" color="green" icon="mdi-circle"></v-icon>
                        <div >{{ item.exerciseCount }} times</div>
                    </div>
                </v-card>
            </v-card>
        </v-sheet>
        <ui-team-bar :team-exercise-counts="teamExerciseCounts"></ui-team-bar>
    </v-container>
</template>
