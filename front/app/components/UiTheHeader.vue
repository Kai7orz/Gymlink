<script setup lang="ts">
const emits = defineEmits<{
  home: [];
  account: [];
  monster: [];
  share: [];
  signout: [];
}>();
const appTitle = "KMonitor";
const shareTip = "みんなの投稿";
const homeTip = "自分の投稿";
const motiveTip = "エールをもらう";
const profileTip = "マイプロフィール";
const signOutTip = "再ログイン";

</script>

<template>
  <client-only>
      <v-app-bar scroll-behavior="hide" color="black" class="flex mb-4 ">
        <v-app-bar-title class="hidden md:inline-block">
          {{ appTitle }}
        </v-app-bar-title>
        <v-tooltip :key="'share'" location="bottom" :text="shareTip">
          <template #activator="{ props:shareProps }">
            <v-icon v-bind="shareProps" class="mx-5" icon="mdi-book-open-page-variant" @click="emits('share')"/>
          </template>
        </v-tooltip>
        <v-tooltip :key="'home'" location="bottom" :text="homeTip">
          <template #activator="{ props:propsHome }">
            <v-icon v-bind="propsHome" class="mx-5" icon="mdi-home" @click="emits('home')"/>
          </template>
        </v-tooltip>
        <v-tooltip :key="'motivate'" location="bottom" :text="motiveTip">
          <template #activator="{ props:propsMotive }">
            <v-icon v-bind="propsMotive" class="mx-5" icon="mdi-sheep" @click="emits('monster')"/>
          </template>
        </v-tooltip>
        <v-tooltip :key="'profile'" location="bottom" :text="profileTip">
          <template #activator="{ props: propsProfile }">
            <v-icon v-bind="propsProfile" class="mx-5" icon="mdi-account" @click="emits('account')"/>
          </template>
        </v-tooltip>
        <v-tooltip :key="'signout'" location="bottom" :text="signOutTip">
          <template #activator="{ props:propsSign }">
            <v-icon v-bind="propsSign" class="mx-5 bg-red" icon="mdi-exit-run" @click="emits('signout')"/>
          </template>
        </v-tooltip>
      </v-app-bar>
  </client-only>
</template>
