<script setup lang="ts">
import authGlobal from '~/middleware/auth.global'


    definePageMeta({
        layout: 'navigator'
    })

    const url = "/api/records/record"
    const fileInput = ref<HTMLInputElement | null >(null)
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
        event.preventDefault();
        const file = fileInput.value?.files?.[0];
        if(!file) return 
        const auth = useAuthStore()
        const formData = new FormData(fileInput.value)
        // バックエンドへファイルを送信する処理を実装
        const TOKEN = auth.idToken
        await useFetch("/api/upload",
            {
                method: 'POST',
                headers: {
                    'Authorization': 'Bearer ' + TOKEN,
                    'Content-Type': 'multipart'
                },
                body: formData,
            }
        );
    };

    const getPreview = async (event: Event) => {
        event.preventDefault();
        const file = fileInput.value?.files?.[0];
        if(!file) return 
        previewUrl.value = URL.createObjectURL(file);
    }

    const createNewRecord = async () => {
        if(postData.imageUrl == ""){
            console.log("Image is Empty")
            return 
        }
        
        const res = await useFetch(url,{
            method: 'POST',
            body: postData,
        })



    }
</script>

<template>
    <v-container>
        <v-container class="flex flex-row justify-center m-10 p-2">
                <v-file-input 
                ref="fileInput" 
                class="max-w-xs m-5 p-3" 
                label="Select Image to Illustrate"
                @change="getPreview" />
                <!-- <v-text-field
                v-model="postData.imageUrl"
                class="w-1/3 m-5"
                label="イラストURL手動入力"/> -->
        </v-container>
        <v-sheet class="flex m-10 bg-black text-center rounded-lg">
                <v-btn  
                class="m-5"
                prepend-icon="$vuetify" 
                append-icon="$vuetify" 
                variant="outlined"
                @click="getIllustration" >
                    イラスト生成
                </v-btn>
                <v-btn 
                class="m-5" 
                prepend-icon="$vuetify" 
                append-icon="$vuetify" 
                variant="outlined"
                @click="createNewRecord" >
                    画像の保存
                </v-btn>
        </v-sheet>

        <v-container class="flex flex-row justify-center">
            <!-- プレビュー画像-->
            <div v-if='previewUrl!=""' class=" w-1/3 flex flex-col flex-wrap md:flex-row md:justify-center m-10">
                <ui-image-card :image_url=previewUrl>
                    <template #title/>
                </ui-image-card>
            </div>
        <!-- レスポンス -->
            <div v-if='responsedUrl!="" ' class="w-1/3 flex flex-col flex-wrap md:flex-row md:justify-center m-10">
                <ui-image-card :image_url=responsedUrl>
                    <template #title>
                        <h1> カード名 </h1>
                    </template>
                </ui-image-card>
            </div>
        </v-container>
    </v-container>
</template>
