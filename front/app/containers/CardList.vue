<script setup lang="ts">
    import { useDetailStore } from '~/stores/detailStore';
    import type { RecordType } from '~/type';
    import { useAuthStore } from '~/stores/auth';
    import { useUserStore } from '~/stores/userStore';
    //CardList の親は，ユーザー自身のカードリストを呼ぶ場合は isOwner=true, 他の人のカードリスト呼ぶ場合は isOwner=false で呼び出す必要がある
    
    const auth = useAuthStore()
    const user = useUserStore()

    const props = defineProps<{
        isOwner: boolean,
    }>();

    const tempUserId = ref("")
    const tempUserName = ref("")
    const TOKEN = auth.idToken
    let recordList = ref([])
    // login 時にセットした id,name を localstorage から取得してくる処理
    onMounted(async ()=>{
        const tempUserIdRaw = localStorage.getItem("userId")
        const tempUserNameRaw = localStorage.getItem("userName")
        if(tempUserIdRaw != null && tempUserNameRaw != null){
        tempUserId.value = tempUserIdRaw
        tempUserName.value = tempUserNameRaw
        user.setUser(Number(tempUserId.value),tempUserName.value)
        }

        const url =   props.isOwner? '/api/users/' + String(user.userId) + '/exercises' : '/api/exercises'


        const data = await $fetch(url,{
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + TOKEN,
                    'Content-Type': 'application/json'
                }
        })
 
        recordList.value = data
    })
    const toAccount = (uid:number) => {
         navigateTo({name: 'Account-id', params: {id:uid}})
    }

    const detailStore = useDetailStore();
    const toDetail = (id:number) => {
        // Store に運動記録の情報をセットしてから遷移して，詳細画面で Store　から取り出す
    if(props.isOwner){
        const tempRecord = recordList.value.find((record : RecordType)=>{
        return record.id == id
        })
        if(tempRecord){
        detailStore.setDetail(tempRecord.id,tempRecord.user_id,tempRecord.user_name,tempRecord.presigned_image,tempRecord.clean_up_time,tempRecord.clean_up_date,tempRecord.comment,tempRecord.likes_count)
        }
        else {
            return
        }
    }
    else{
        const  detailRecord = recordList.value[ recordList.value.length - id];
        if(!detailRecord) return;
        detailStore.setDetail(detailRecord.id,detailRecord.user_id,detailRecord.user_name,detailRecord.presigned_image,detailRecord.clean_up_time,detailRecord.clean_up_date,detailRecord.comment,detailRecord.likes_count)
    }
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
                            record_id: id
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
    <ui-card-list :recordList="recordList" @detail="toDetail" @like="like" @account="toAccount"/>
</template>