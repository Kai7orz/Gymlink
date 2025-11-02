<script setup lang="ts">
const emits = defineEmits<{
  home: [];
  account: [];
  monster: [];
  share: [];
  signout: [];
}>();
</script>

<template>
      <v-app-bar scroll-behavior="hide" color="black" class="flex mb-4 ">
        <v-app-bar-title class="hidden md:inline-block">
          KMonitor
        </v-app-bar-title>
        <v-tooltip location="bottom" text="みんなの投稿">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-book-open-page-variant" @click="emits('share')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" text="自分の投稿">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-home" @click="emits('home')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" text="エールをもらう">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-sheep" @click="emits('monster')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" text="マイプロフィール">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-account" @click="emits('account')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" text="再ログイン">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5 bg-red" icon="mdi-exit-run" @click="emits('signout')"/>
          </template>
        </v-tooltip>
      </v-app-bar>
</template>
