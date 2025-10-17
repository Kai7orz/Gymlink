<script setup lang="ts">
definePageMeta({ layout: 'navigator' })

const url = "/api/records/record"

const selectedFile = ref<File | null>(null)

const responsedUrl = ref<string>("")
const previewUrl = ref<string>("")

const postData = reactive({
  userId: "",
  categoryId: "",
  recordName: "",
  imageUrl: "",
  imageDescription: "",
})

const getIllustration = async (event: Event) => {
  event.preventDefault()
  if (!selectedFile.value) return

  const auth = useAuthStore()
  const TOKEN = auth.idToken

  const formData = new FormData()
  formData.append('file', selectedFile.value, selectedFile.value.name)
  formData.append('s3_key',"raw_image")

  console.log("formdata.get:",formData.get('s3_key'))
  await useFetch("/api/upload", {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${TOKEN}`,
    },
    body: formData,
  })
}

const getPreview = async (event: Event) => {
  event.preventDefault()
  if (!selectedFile.value) return
  previewUrl.value = URL.createObjectURL(selectedFile.value)
}

const createNewRecord = async () => {
  if (postData.imageUrl == "") {
    console.log("Image is Empty")
    return
  }
  await useFetch(url, { method: 'POST', body: postData })
}

onMounted(()=>{
  responsedUrl.value = "https://katazuke.s3.ap-northeast-1.amazonaws.com/katazuke.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=AKIA2CUNLZ5LROQUFMZJ%2F20251017%2Fap-northeast-1%2Fs3%2Faws4_request&X-Amz-Date=20251017T042100Z&X-Amz-Expires=60&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=404a1ff2fa9ef4e8ce1415a2949125f7ec36c9920980cd35d294165bbd2db4b6"
})
</script>

<template>
  <v-container>
    <v-container class="flex flex-row justify-center m-10 p-2">
      <form class="w-1/3">
        <v-file-input
          v-model="selectedFile"
          class="w-100 m-5 p-3"
          label="Select Image to Illustrate"
          @change="getPreview"
          :show-size="true"
          accept="image/*"
        />
      </form>
    </v-container>

    <v-sheet class="flex m-10 bg-black text-center rounded-lg">
      <v-btn class="m-5" variant="outlined" @click="getIllustration">
        イラスト生成
      </v-btn>
      <v-btn class="m-5" variant="outlined" @click="createNewRecord">
        画像の保存
      </v-btn>
    </v-sheet>
    <v-container class="flex flex-row justify-center">
      <div v-if='previewUrl!=""' class="w-1/3 flex flex-col md:flex-row md:justify-center m-10">
        <ui-image-card :image_url="previewUrl"><template #title/></ui-image-card>
      </div>
      <div v-if='responsedUrl!=""' class="w-1/3 flex flex-col md:flex-row md:justify-center m-10">
          <img :src="responsedUrl" />
      </div>
    </v-container>
  </v-container>
</template>
