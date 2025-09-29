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
  <!-- フル画面グラデ背景 -->
  <div class="min-h-screen bg-gradient-to-br from-indigo-600 via-violet-600 to-fuchsia-500 pt-6 pb-16">
    <v-container class="max-w-4xl">
      <!-- 戻るボタン -->
      <div class="flex items-center gap-3 mb-4 text-white/90">
        <v-btn
          density="comfortable"
          variant="text"
          class="rounded-full hover:opacity-90"
          icon="mdi-arrow-left"
          @click="toBack"
        />
        <span class="text-sm opacity-80">戻る</span>
      </div>

      <!-- ガラスカード -->
      <v-card
        class="backdrop-blur-md bg-white/10 border border-white/20 text-white rounded-3xl shadow-xl overflow-hidden"
        elevation="10"
      >
        <!-- 画像ヘッダー：日付チップをオーバーレイ -->
        <div class="relative">
          <v-img
            :src="detailStore.detailImage"
            aspect-ratio="16/9"
            cover
            class="opacity-95"
          />
          <v-chip
            class="absolute left-4 bottom-4 bg-black/60 text-white"
            size="small"
            prepend-icon="mdi-calendar-month"
          >
            {{ detailStore.detailDate }}
          </v-chip>
        </div>

        <!-- タイトル -->
        <v-card-title class="text-2xl md:text-3xl font-semibold tracking-wide py-4">
          {{ detailStore.detailName }}
        </v-card-title>

        <v-divider class="opacity-30" />

        <!-- 情報グリッド -->
        <v-card-text class="py-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <v-sheet class="bg-white/5 rounded-2xl p-4 border border-white/10">
              <div class="text-sm opacity-80 mb-1 flex items-center gap-2">
                <v-icon size="18" icon="mdi-timer-sand" /> 運動時間
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

        <!-- アクション（必要ならイベント紐づけ） -->
        <v-card-actions class="justify-end gap-2 py-4">
          <v-btn variant="tonal" class="bg-white/10" prepend-icon="mdi-share-variant">シェア</v-btn>
          <v-btn variant="elevated" color="amber-accent-3" class="text-black font-medium" prepend-icon="mdi-heart-outline">
            いいね
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </div>
</template>

<style scoped>
</style>
