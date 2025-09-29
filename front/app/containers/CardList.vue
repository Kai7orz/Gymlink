<script setup lang="ts">
    import { useDetailStore } from '~/stores/detailStore';
    import type { ExerciseRecordType } from '~/type';
    import { useAuthStore } from '~/stores/auth';
    import { useUserStore } from '~/stores/userStore';
    //CardList の親は，ユーザー自身のカードリストを呼ぶ場合は isOwner=true, 他の人のカードリスト呼ぶ場合は isOwner=false で呼び出す必要がある
    
    const auth = useAuthStore()
    const user = useUserStore()

    const props = defineProps<{
        isOwner: boolean,
    }>();

    const TOKEN = auth.idToken
    const url =   props.isOwner? '/api/users/' + String(user.userId) + '/exercises' : '/api/exercises'
    const { data, pending, error, refresh } = await useFetch( url,
        {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + TOKEN,
                'Content-Type': 'application/json'
            }
        }
    );


    const toAccount = (uid:number) => {
         navigateTo({name: 'Account-id', params: {id:uid}})
    }

    const exerciseMocksList:ExerciseRecordType[] = data.value
    console.log("valuedata -.",exerciseMocksList)
    const detailStore = useDetailStore();
    const toDetail = (id:number) => {
        // Store に運動記録の情報をセットしてから遷移して，詳細画面で Store　から取り出す
        const detailRecord = exerciseMocksList[id-1]; //mockのlist は id=0 からスタートしているが，mockオブジェクト自体のidは1からスタートしているため，-1している　バックエンドから受け取るexerciseList のid=1からスタートすればid-1 は不要
        if(!detailRecord) return;
        detailStore.setDetail(detailRecord.id,detailRecord.user_id,detailRecord.user_name,detailRecord.image_url,detailRecord.time,detailRecord.date,detailRecord.comment,detailRecord.likes_count)
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
                await $fetch("/api/likes",
                    {
                        method: 'POST',
                        headers: {
                            'Authorization': 'Bearer ' + TOKEN,
                            'Content-Type': 'application/json'
                        },
                        body: {
                            exercise_record_id: id
                        },
                    }
                )
            } catch(e){
                console.log("likes post error: ",e)
            }
        }
    }



</script>

<template>
    <ui-card-list :exerciseMocksList="exerciseMocksList" @detail="toDetail" @like="like" @account="toAccount"/>
</template>
