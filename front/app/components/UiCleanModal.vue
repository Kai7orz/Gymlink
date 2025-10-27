<script setup lang="ts">

const emits = defineEmits<{
  record: [];
  close: [];
  select: [imageId: string];
}>();

const selected = ref("");
const cleanTimeLabel = "片付け時間（min）";
const cleanCommentLabel = "コメント";
const isShownMenu = defineModel<boolean>("isShownMenu");
const cleanTime = defineModel<string>("cleanTime");
const comment = defineModel<string>("comment");
const date = defineModel<string>("date");

const recordAndClose = () => {
  emits("record");
  emits("close");
};

watch(selected, (val) => {
  emits("select", val);
});
</script>

<template>
    <v-dialog v-model="isShownMenu" max-width="500">
        <v-card class="bg-grey-darken-3">
            <v-btn color="primary" text @click="emits('close')">Close</v-btn>
            <v-card-title class="text-h5">片付け記録</v-card-title>
            <v-text-field v-model="cleanTime" :label="cleanTimeLabel"/>
            <v-text-field v-model="comment" :label="cleanCommentLabel"/>
            <v-container>
                <v-row justify="space-around">
                <v-date-picker v-model="date" class="bg-black m-3" show-adjacent-months/>
                </v-row>
            </v-container>
            <v-spacer/>
            <v-btn class="bg-black" @click="recordAndClose">情報セット</v-btn>
        </v-card>
    </v-dialog>
</template>
