<script setup lang="ts">
    // excerciseRecord[] を取得 fetch で API から取得する想定
    import { exerciseMocks } from '~/mocks/api/exerciseMock';
    import { useDetailStore } from '~/stores/detailStore';
    import type { exerciseRecordType } from '~/type';

    const exerciseMocksList:exerciseRecordType[] = exerciseMocks;
    const detailStore = useDetailStore();

    const toDetail = (id:number) => {
        // Store に運動記録の情報をセットしてから遷移して，詳細画面で Store　から取り出す
        const detailRecord = exerciseMocksList[id-1]; //mockのlist は id=0 からスタートしているが，mockオブジェクト自体のidは1からスタートしているため，-1している　バックエンドから受け取るexerciseList のid=1からスタートすればid-1 は不要
        if(!detailRecord) return;
        detailStore.setDetail(detailRecord.id,detailRecord.image,detailRecord.time,detailRecord.date,detailRecord.comment)
        navigateTo({name: 'Detail-id', params: {id: id }})
    }
</script>

<template>
    <ui-card-list :exerciseMocksList="exerciseMocksList" @detail="toDetail"/>
</template>
