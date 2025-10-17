<script setup lang="ts">
    import type { Illustration } from '~/type';

    const props = defineProps<{
        illustObjs: Record<string,Illustration>
        imageUrl: string
    }>();

    const emits = defineEmits<{
        record: [],
        close: [],
        select:[imageId:string],
    }>();

    const selected = ref('')
    const exerciseTimeLabel = "片付け時間（min）"
    const exerciseCommentLabel = "コメント"
    const isShownMenu = defineModel<boolean>('isShownMenu')
    const exerciseTime = defineModel<string>('exerciseTime')
    const comment = defineModel<string>('comment')
    const date = defineModel<string>('date')

    const recordAndClose = () => {
        emits('record')
        emits('close')
    }

    watch(selected,(val)=>{
        emits('select',val)
    })

    const imageItems = Object.keys(props.illustObjs)
</script>

<template>
    <v-dialog v-model="isShownMenu" max-width="500">
        <v-card class="bg-grey-darken-3">
            <v-btn color="primary" text @click="emits('close')">Close</v-btn>
            <v-card-title class="text-h5">片付け記録</v-card-title>
            <!-- <v-select
                     label="Select"
                     :items="imageItems"
                     v-model="selected"
            ></v-select>
            <img v-if="selected!==undefined" :src="props.imageUrl" class="w-60 h-50 mx-auto my-5" /> -->
            <v-text-field v-model="exerciseTime" :label="exerciseTimeLabel"></v-text-field>
            <v-text-field v-model="comment" :label="exerciseCommentLabel"></v-text-field>
            <v-container>
                <v-row justify="space-around">
                <v-date-picker class="bg-black m-3" v-model="date" show-adjacent-months></v-date-picker>                    
                </v-row>
            </v-container>           
            <v-spacer></v-spacer>
            <v-btn class="bg-black" @click="recordAndClose">情報セット</v-btn>
        </v-card>
    </v-dialog>
</template>
