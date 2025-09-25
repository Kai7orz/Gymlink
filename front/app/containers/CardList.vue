<script setup lang="ts">
    import { useDetailStore } from '~/stores/detailStore';
    import type { exerciseRecordType } from '~/type';
    //CardList の親は，ユーザー自身のカードリストを呼ぶ場合は isOwner=true, 他の人のカードリスト呼ぶ場合は isOwner=false で呼び出す必要がある
    const props = defineProps<{
        isOwner: boolean,
    }>();

    const userId = 1
    // cardList が自身のものか全体共有用のものかで URL分岐
    const url = computed(()=>
    props.isOwner
    ? '/api/users/' + String(userId) + '/exercises-example'
    : '/api/exercises' );
    const { data }  = await useFetch(url)


    console.log("swaggerObject:",data.value[1].id)
    const exerciseMocksList:exerciseRecordType[] = data.value
    console.log("data:",data.value)
    const detailStore = useDetailStore();
    const toDetail = (id:number) => {
        // Store に運動記録の情報をセットしてから遷移して，詳細画面で Store　から取り出す
        const detailRecord = exerciseMocksList[id-1]; //mockのlist は id=0 からスタートしているが，mockオブジェクト自体のidは1からスタートしているため，-1している　バックエンドから受け取るexerciseList のid=1からスタートすればid-1 は不要
        if(!detailRecord) return;
        detailStore.setDetail(detailRecord.id,detailRecord.imageUrl,detailRecord.time,detailRecord.date,detailRecord.comment,detailRecord.likesCount)
        navigateTo({name: 'Detail-id', params: {id: id }})
    }

    const like = async (id:number) => {
        // cardList が自身のものか否か判定して処理分岐
        console.log("isOwner:",props.isOwner)
        if(props.isOwner){
            await navigateTo({name: 'liked-id', params: {id: id}}) //exercise record id番　にいいねした人一覧ページへの遷移
        }
        else{
            try{
                await useFetch("/api/exerciseRecords/"+String(id)+"/likes")
            } catch(e){
                console.log("likes post error: ",e)
            }
        }
    }



</script>

<template>
    <ui-card-list :exerciseMocksList="exerciseMocksList" @detail="toDetail" @like="like"/>
</template>
