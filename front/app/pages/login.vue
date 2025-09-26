<script setup lang="ts">
    // middleware からのログインページへの遷移時にハイドレーションミスマッチがおきるので ClientOnly で対応している
    import {signIn}  from '~/composables/SignInUser';

    const email = ref('')
    const password = ref('')
    const isLoading = ref(false)

    watch(isLoading, val => {
            val && setTimeout(() => {
            isLoading.value = false
            }, 3000)
        })

    const signInUser = async () => {
        isLoading.value = true
        const minLoadingPromise = new Promise(resolve => setTimeout(resolve, 1000));
        try{
            await Promise.all([
                signIn(email.value, password.value),
                minLoadingPromise
            ]);
            await new Promise(resolve => setTimeout(resolve, 100)); 
            await navigateTo('/home')
        } catch (error) {
            console.error('SignIn Error:', error);
        } finally {
            isLoading.value = false
        }
    }
</script>

<template>
  <ClientOnly>
    <v-card class="d-flex flex-column justify-center mx-auto w-50 m-20 border-lg rounded-lg">
        <v-card-title class="d-flex justify-center">サインイン</v-card-title>
        <v-text-field v-model="email" class="w-1/2 mx-auto m-5 " label="メールアドレス" />
        <v-text-field v-model="password" class="w-1/2 mx-auto m-5" label="パスワード" type="password" />
        <v-btn class="d-flex justify-center" @click="signInUser" color="primary">
            サインイン
            <v-overlay v-model="isLoading"
                activator="parent"
                location-strategy="connected"
                class="d-flex justify-canter mx-auto my-auto"
            >
                <v-progress-circular
                    color="primary"
                    size="64"
                    indeterminate
                ></v-progress-circular>
            </v-overlay>
        </v-btn>
    </v-card>
  </ClientOnly>
</template>