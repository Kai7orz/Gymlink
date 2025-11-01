<script setup lang="ts">

import AddCard from "~/containers/AddCard";

const color = ref("grey-darken-3");
const isError = ref(false);
const selectedFile = ref<File | null>(null);
const responsedUrl = ref<string>("");
const previewUrl = ref<string>("");
const isLoading = ref(false);
const isShownMenu = ref(false);
const cleanUpTime = ref("");
const cleanUpDate = ref(new Date());
const comment = ref("");
const auth = useAuthStore();
const router = useRouter();
const TOKEN = auth.idToken;
const takeTimeLabel = "※レコードの保存から表示まで1～2分ほど反映にかかります";
const warningLabel = "Invalid Form : form is empty or file isn't png";
const saveRecordLabel = "レコード保存";
const loadingLabel = "Loading...";
const invalidDateLabel = "日付：未入力";
const invalidCleanTimeLabel = "片付け時間：未入力";
const validImageLabel = "画像選択済み";
const invalidImageLabel = "画像：未選択";
const invalidCommentLabel = "コメント：未入力";
const toBackLabel = "戻る";
// イラスト一覧を読み込んで propsとして渡して表示する

const formatDate = (d: Date | string) => {
  const date = new Date(d);
  const yyyy = date.getFullYear();
  const mm = String(date.getMonth() + 1).padStart(2, "0");
  const dd = String(date.getDate()).padStart(2, "0");
  return `${yyyy}${mm}${dd}`;
};

const getIllustration = async (event: Event) => {
  event.preventDefault();
  if (isError.value || cleanUpTime.value == "" || comment.value == "" || selectedFile.value == undefined) {
    isError.value = true;
    return;
  }

  const formData = new FormData();
  formData.append("file", selectedFile.value, selectedFile.value.name);
  formData.append("s3_key", "raw_image");
  formData.append("clean_up_time", cleanUpTime.value);
  formData.append("clean_up_date", formatDate(cleanUpDate.value));
  formData.append("comment", comment.value);
  isLoading.value = true;

  try {
    await useFetch("/api/upload", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${TOKEN}`,
      },
      body: formData,
    });

    await navigateTo("/home");
  }
  catch (error) {
    console.log("error: response is invalid", error);
  }
  finally {
    isLoading.value = false;
  }
};

const toBack = () => {
  router.back();
};

watch(selectedFile, () => {
  if (selectedFile.value == undefined || selectedFile.value.type != "image/png") {
    isError.value = true;
  }
});

onMounted(() => {
  responsedUrl.value = "https://katazuke.s3.ap-northeast-1.amazonaws.com/katazuke.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=AKIA2CUNLZ5LROQUFMZJ%2F20251017%2Fap-northeast-1%2Fs3%2Faws4_request&X-Amz-Date=20251017T042100Z&X-Amz-Expires=60&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=404a1ff2fa9ef4e8ce1415a2949125f7ec36c9920980cd35d294165bbd2db4b6";
});
</script>

<template>
  <v-container class="d-flex flex-column items-center">
            <div class="flex items-center gap-3 m-4 text-white/90 mr-auto">
            <v-btn
            density="comfortable"
            variant="text"
            class="rounded-full hover:opacity-90"
            icon="mdi-arrow-left"
            @click="toBack"
            />
            <span class="text-sm opacity-80">{{ toBackLabel }}</span>
        </div>
    <v-snackbar
                v-model="isError" class="mb-20"
                multi-line>
                {{ warningLabel }}
        <template #actions>
            <v-btn
                  color="red"
                  variant="text"
                  @click="isError = false">
                Close
            </v-btn>
        </template>
    </v-snackbar>
    <v-container class="add-item">
      <add-card
                v-model:clean-up-time="cleanUpTime"
                v-model:clean-up-date="cleanUpDate"
                v-model:comment="comment"
                v-model:is-shown-menu="isShownMenu"
                v-model:preview-url="previewUrl"
                v-model:selected-file="selectedFile"
               />
    </v-container>
    <v-container class="flex flex-row justify-center">
      <div v-if='previewUrl!=""' class="w-2/3 flex flex-col md:flex-row md:justify-center m-5">
        <ui-image-card :image_url="previewUrl"><template #title/></ui-image-card>
      </div>
    </v-container>
    <v-sheet class="flex w-2/3 p-1 bg-black text-center rounded-lg">
      <v-container class="d-flex flex-column justify-center items-center bg-black gap-4">
        <div class="w-full d-flex justify-center ">
          <v-card class="d-flex justify-center items-center w-full h-20" :color="color">{{ cleanUpDate ? "✅ " + String((cleanUpDate.getMonth() + 1) + '月'+cleanUpDate.getDate() + '日') :invalidDateLabel}}</v-card>
        </div>
        <div class="mini-boxes">
          <v-card class="d-flex justify-center items-center w-full sm:w-2/3 h-20" :color="color">{{ cleanUpTime ? "✅ " + cleanUpTime : invalidCleanTimeLabel }}</v-card>
          <v-card class="d-flex justify-center items-center w-full sm:w-2/3 h-20" :color="color">{{ selectedFile ? "✅ " + validImageLabel : invalidImageLabel}}</v-card>
        </div>
        <div class="w-full d-flex justify-center ">
          <v-card class="d-flex justify-center items-center w-full h-20" :color="color">{{comment ? "✅ "+comment : invalidCommentLabel }}</v-card>
        </div>
      </v-container>
      <v-btn class="m-5" variant="outlined" @click="getIllustration">
        {{ saveRecordLabel }}
      </v-btn>
      <v-overlay
              v-model="isLoading"
              location-strategy="connected"
              class="d-flex justify-center items-center mx-auto my-auto"
            >
            <v-card class="d-flex items-center justify-center bg-black text-white mx-auto" min-width="150" min-height="100">
              {{ loadingLabel }}
            </v-card>
      </v-overlay>
      <v-container>
        <div class="text-design">{{ takeTimeLabel }}</div>
      </v-container>
    </v-sheet>
  </v-container>
</template>

<style>
  .add-item {
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 12px;
    padding: 8px;
    @media (max-width: 700px){
      flex-direction: column;
    }
  }

  .upload-form {
    width: 50%;
    @media (max-width: 700px){
      width: 100%;
    }
  }

  .mini-boxes {
    display: flex;
    width: 100%;
    gap: 16px;
    @media (max-width: 700px){
      flex-direction: column;
    }
  }

  .text-design{
    font-size:12px;
    @media (max-width: 400px){
      font-size: 7px;
    }
  }
</style>
