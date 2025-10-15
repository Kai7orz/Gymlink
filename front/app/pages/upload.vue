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
</script>

<template>
  <v-container>
    <v-container class="flex flex-row justify-center m-10 p-2">
      <form>
        <v-file-input
          v-model="selectedFile"
          class="max-w-xs m-5 p-3"
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
        <ui-image-card :image_url="responsedUrl">
          <template #title><h1>カード名</h1></template>
        </ui-image-card>
      </div>
    </v-container>
  </v-container>
</template>
