<script setup lang="ts">
const emits = defineEmits<{
  home: [];
  account: [];
  monster: [];
  share: [];
  signout: [];
}>();
const appTitle = "KMonitor"
const shareTip = "みんなの投稿";
const homeTip = "自分の投稿";
const motiveTip = "エールをもらう";
const profileTip = "マイプロフィール";
const signOutTip = "再ログイン";

</script>

<template>
      <v-app-bar scroll-behavior="hide" color="black" class="flex mb-4 ">
        <v-app-bar-title class="hidden md:inline-block">
          {{ appTitle }}
        </v-app-bar-title>
        <v-tooltip location="bottom" :text="shareTip">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-book-open-page-variant" @click="emits('share')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" :text="homeTip">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-home" @click="emits('home')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" :text="motiveTip">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-sheep" @click="emits('monster')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" :text="profileTip">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5" icon="mdi-account" @click="emits('account')"/>
          </template>
        </v-tooltip>
        <v-tooltip location="bottom" :text="signOutTip">
          <template #activator="{ props }">
            <v-icon v-bind:="props" class="mx-5 bg-red" icon="mdi-exit-run" @click="emits('signout')"/>
          </template>
        </v-tooltip>
      </v-app-bar>
</template>
