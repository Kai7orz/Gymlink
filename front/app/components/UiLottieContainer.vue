<script setup lang="ts">
import { ref, onMounted } from 'vue'
import lottie from 'lottie-web'

type AnimationJSON = Record<string, any>
const props = withDefaults(
  defineProps<{
    animationData: AnimationJSON
    loop?: boolean
    autoplay?: boolean
    width?: string
    height?: string
  }>(),
  { loop: true, autoplay: true, width: '100%', height: '100%' }
)

const lavContainer = ref<HTMLDivElement | null>(null)
onMounted(async () => {
  if (!lavContainer.value || !props.animationData) return

  // loadAnimation の仕様は lottie-web に乗っているのでそれを参照する
  // container はアニメーションを引っ掛けるための DOM 要素で， ref を使ってあらかじめ用意した変数 lavContainer にその要素情報を入れておく． onMounted の中では DOM 要素が確実に存在するので template に記述したdiv の情報がlavContainerに代入されて，
  // それをcontainer に指定することで， その要素に対して animation を描画してくれる流れ 
  const anim = lottie.loadAnimation({
    container: lavContainer.value,
    renderer: 'svg',
    loop: props.loop,
    autoplay: props.autoplay,
    animationData: props.animationData,
  })
})
</script>

<template>
  <div
    ref="lavContainer"
    :style="{ width: props.width, height: props.height }"
  />
</template>
