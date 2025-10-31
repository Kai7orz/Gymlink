<script setup lang="ts">

const emits = defineEmits<{
  record: [];
  close: [];
  select: [imageId: string];
}>();

const cleanTimeLabel = "片付け時間（min）";
const cleanCommentLabel = "コメント";
const isShownMenu = defineModel<boolean>("isShownMenu");
const cleanUpTime = defineModel<string>("cleanUpTime");
const cleanUpDate = defineModel<Date>("cleanUpDate");
const comment = defineModel<string>("comment");

const recordAndClose = () => {
  emits("close");
};
</script>

<template>
    <v-dialog v-model="isShownMenu" max-width="500">
        <v-card class="bg-grey-darken-3">
            <v-btn color="primary" text @click="emits('close')">Close</v-btn>
            <v-card-title class="text-h5">片付け記録</v-card-title>
            <v-text-field v-model="cleanUpTime" :label="cleanTimeLabel"/>
            <v-text-field v-model="comment" :label="cleanCommentLabel"/>
            <v-container>
                <v-row justify="space-around">
                <v-date-picker v-model="cleanUpDate" class="bg-black m-3" show-adjacent-months/>
                </v-row>
            </v-container>
            <v-spacer/>
            <v-btn class="bg-black" @click="recordAndClose">情報セット</v-btn>
        </v-card>
    </v-dialog>
</template>
