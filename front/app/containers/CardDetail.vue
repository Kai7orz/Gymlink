<script setup lang="ts">
  import { useAuthStore } from '~/stores/auth';
  const  route = useRoute();
  const props = defineProps<{
    id: number,
  }>();

    
  const auth = useAuthStore()
  const TOKEN = auth.idToken
  const data = ref(false)
  const record_id = props.id
  
  data.value = await $fetch(`/api/likes/${record_id}`,{
      method: 'GET',
      headers: {
                    'Authorization': 'Bearer ' + TOKEN,
                    'Content-Type': 'application/json'  
      }
  })
  
  const toGood = async (id:number) =>{
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

  const toDelete = async (id:number) => {
    const record_id = id
    try{
      await $fetch(`/api/likes/${record_id}`,{
        method: "DELETE",
        headers: {
                    'Authorization': 'Bearer ' + TOKEN,
                    'Content-Type': 'application/json'
                  },
      })
    } catch(e){
      console.log("delete error: ",e)
    }
  }

</script>

<template>
        <div>
          <ui-card-detail :id="props.id" :liked="data.liked" @good="toGood" @delete="toDelete"/>
        </div>
</template>
