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
  
  console.log("liked value -> ",data.value.liked)

  const toGood = (id:number) =>{

  }

  const toDelete = (id:number) => {

  }

</script>

<template>
        <div>
          <ui-card-detail :id="props.id" :liked="data.liked" @good="toGood" @delete="toDelete"/>
        </div>
</template>
