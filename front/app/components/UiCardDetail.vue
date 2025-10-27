<script setup lang="ts">
import { useDetailStore } from "~/stores/detailStore";
import { useUserStore } from "~/stores/userStore";

const props = defineProps<{
  id: number;
  liked: boolean;
}>();
const state = ref(props.liked);

const emits = defineEmits<{
  good: [id: number];
  delete: [id: number];
}>();
const detailStore = useDetailStore();
const user = useUserStore();

const toBack = () => {
  if (user.userName === detailStore.detailName) navigateTo({ name: "home" });
  else navigateTo({ name: "share" });
};

const GoodOrDelete = () => {
  state.value = !state.value;
};

onUnmounted(() => {
  if (props.liked == state.value) {
    return;
  }
  else if (state.value) {
    emits("good", props.id);
  }
  else {
    emits("delete", props.id);
  }
});
</script>

<template>
  <div>
    <v-container class="max-w-4xl">
      <div class="flex items-center gap-3 m-4 text-white/90">
        <v-btn
          density="comfortable"
          variant="text"
          class="rounded-full hover:opacity-90"
          icon="mdi-arrow-left"
          @click="toBack"
        />
        <span class="text-sm opacity-80">戻る</span>
      </div>

      <v-card
        class="backdrop-blur-md shadow-xl bg-black overflow-hidden"
        elevation="10"
      >

        <v-card-title class="d-flex h-20 items-center text-2xl md:text-3xl font-semibold tracking-wide">
          <v-chip
            class="absolute left-4 m-10 bg-black/60 text-white"
            size="small"
            prepend-icon="mdi-calendar-month"
          >
            {{ detailStore.detailDate }}
          </v-chip>
          {{ detailStore.detailName }}
        </v-card-title>
        
        <v-divider class="opacity-30" />
        <div class="responsive-switch">
          <div class="image-content-box">
            <v-card class="img-box-design">
              <img class="img-design mx-auto bg-black rounded" :src="detailStore.detailPresignedImage" >
            </v-card>
          </div>
          <v-card-text class="flex justify-center items-center py-6 w-1/3">
            <div class="items gap-4">
                <v-sheet class="item-content bg-white/5 rounded-lg border border-white/10">
                  <div class="item text-sm mb-1 flex items-center gap-2">
                    <v-icon size="18" icon="mdi-timer-sand" /> 片付け時間
                  </div>
                  <div class="mx-2">{{ detailStore.detailTime }} 分</div>
                </v-sheet>
                <v-sheet class="item-content bg-white/5 rounded-lg border border-white/10 md:col-span-2">
                  <div class="item text-sm mb-1 flex items-center gap-2">
                    <v-icon size="18" icon="mdi-comment-text-outline" /> コメント
                  </div>
                  <div class="mx-2 leading-relaxed">
                    {{ detailStore.detailComment || 'コメントはありません' }}
                  </div>
                </v-sheet>
            </div>
          </v-card-text>
        </div>
        <v-divider class="opacity-30" />

        <v-card-actions class="justify-end gap-2 py-4">
          <!--  click 複数回押しても対応できるように，clicke のたびに liked の false ,true 切り替わる処理を入れる何回もAPI 叩かないようにする -->
          <v-btn v-if="state" variant="elevated" color="amber-accent-3" class="text-black font-medium" @click ="() => GoodOrDelete(props.id)">
            <v-icon size="20" class="m-2" icon="mdi-heart" color="red"/>
            いいね
          </v-btn>
          <v-btn v-if="state==false" variant="elevated" color="amber-accent-3" class="text-black font-medium" @click ="() => GoodOrDelete(props.id)">
            <v-icon size="20" class="m-2" icon="mdi-heart-outline" color="red"/>
            いいね
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </div>
</template>

<style>

.responsive-switch {
  display: flex;
  @media(max-width:700px){
    flex-direction: column;
    align-items: center;
  }
}

.image-content-box{
  display:flex;
  justify-content: center;
  width:600px;
  @media(max-width:700px){
    width:100%;
    display:flex;
    flex-direction: column;
    justify-content:center;
    align-items: center;
  }
}
.img-box-design{
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  height: 400px;
  width: 400px;
  margin: 20px;
  @media(max-width:700px){
    width: 300px;
    height: 300px;
    margin: 10px;
  }
}

.img-design{
  margin:4px;
  width: 90%;
  height: 90%;
  @media(max-width:700px){
    width: 290px;
    height: 290px;
  }
}

.items-box{
  display:flex;
}
.item-content{
  padding:16px;
  @media(max-width:700px){
    padding:0px;
    margin:4px;
    width:300px;
  }
}

.item{
  margin:0px 4px;
}

.items{
  display: flex;
  flex-direction: column;
  width: 400px;
  justify-content:center;
  @media(max-width:700px){
    width:350px;
  }
}

</style>