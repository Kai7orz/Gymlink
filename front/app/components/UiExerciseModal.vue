<script setup lang="ts">
    const emits = defineEmits<{
        close: [],
    }>();

    const record = () => {
        const postData = {
            imageUrl: imageUrl.value,
            time: Number(exerciseTime.value),
            date: selectedDate.value,
            comment: comment.value
        }

        console.log("postData->:",postData)

        emits('close')
    }

    const exerciseTimeLabel = "運動時間（min）"
    const exerciseCommentLabel = "コメント"
    const imageUrl = ref(1)
    const isShownMenu = defineModel<boolean>('isShownMenu')
    const exerciseTime = ref("")
    const selectedDate = ref(null)
    const comment = ref("")

</script>

<template>
    <v-dialog v-model="isShownMenu" max-width="500">
        <v-card>
            <v-btn color="primary" text @click="emits('close')">Close</v-btn>
            <v-card-title class="text-h5">運動記録</v-card-title>
            <img src="/images/sportImage.png" class="w-60 h-50 mx-auto my-5" />
            <v-text-field v-model="exerciseTime" :label="exerciseTimeLabel"></v-text-field>
            <v-text-field v-model="comment" :label="exerciseCommentLabel"></v-text-field>
            <v-container>
                <v-row justify="space-around">
                <v-date-picker v-model="selectedDate" show-adjacent-months></v-date-picker>                    
                </v-row>
            </v-container>           
            <v-spacer></v-spacer>
            <v-btn @click="record">記録</v-btn>
        </v-card>
    </v-dialog>
</template>
