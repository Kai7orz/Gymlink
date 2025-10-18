<script setup lang="ts">
import { useDetailStore } from '~/stores/detailStore'
import { useUserStore } from '~/stores/userStore'

const detailStore = useDetailStore()
const user = useUserStore()

const toBack = () => {
  if (user.userName === detailStore.detailName) navigateTo({ name: 'home' })
  else navigateTo({ name: 'share' })
}

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

        <v-card-title class="d-flex items-center text-2xl md:text-3xl font-semibold tracking-wide py-4">
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
        <v-container class="d-flex justify-center">
          <v-card class="img-design">
            <img class="mx-auto bg-black m-10 rounded" :src="detailStore.detailPresignedImage" />
          </v-card>
        </v-container>
        <v-card-text class="py-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <v-sheet class="bg-white/5 rounded-2xl p-4 border border-white/10">
              <div class="text-sm opacity-80 mb-1 flex items-center gap-2">
                <v-icon size="18" icon="mdi-timer-sand" /> 片付け時間
              </div>
              <div class="text-xl font-semibold">{{ detailStore.detailTime }} 分</div>
            </v-sheet>

            <v-sheet class="bg-white/5 rounded-2xl p-4 border border-white/10 md:col-span-2">
              <div class="text-sm opacity-80 mb-1 flex items-center gap-2">
                <v-icon size="18" icon="mdi-comment-text-outline" /> コメント
              </div>
              <div class="leading-relaxed">
                {{ detailStore.detailComment || 'コメントはありません' }}
              </div>
            </v-sheet>
          </div>
        </v-card-text>

        <v-divider class="opacity-30" />

        <v-card-actions class="justify-end gap-2 py-4">
          <v-btn variant="elevated" color="amber-accent-3" class="text-black font-medium" prepend-icon="mdi-heart-outline">
            いいね
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </div>
</template>

<style>

.img-design{
  display: flex;
  justify-content: center;
  text-align: center;
  height: 600px;
  width: 600px;
}

</style>