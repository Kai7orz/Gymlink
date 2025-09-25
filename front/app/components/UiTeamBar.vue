<script setup lang="ts">
    const props = defineProps<{
        teamExerciseCounts: number[]
    }>();

    console.log("teamcount:",props.teamExerciseCounts)

</script>

<template>
        <div class="d-flex justify-center">
            <div  v-for="(item,index) in props.teamExerciseCounts" class="d-flex justify-center items-center">
                <v-icon v-if="index==0" v-for="i in item" class="my-5 mx-5" color="red" icon="mdi-circle"></v-icon>
                <v-icon v-if="index==1" v-for="i in item" class="my-5 mx-5" color="blue" icon="mdi-circle"></v-icon>
                <v-icon v-if="index==2" v-for="i in item" class="my-5 mx-5" color="green" icon="mdi-circle"></v-icon>            
            </div>
        </div>

</template>
