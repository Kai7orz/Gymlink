<script setup lang="ts">
    import {signUp}  from '~/composables/SignUpUser';

    const name = ref('')
    const email = ref('')
    const password = ref('')
    const avatarUrl = ref('/images/test.png')
    const isLoading = ref(false)

    const signUpUser = async () => {
        isLoading.value = true

        try{
            await signUp(email.value,password.value)
            await $fetch("/api/signup",
                {
                    method: 'POST',
                    body: {
                        name:  name.value,
                        email: email.value,
                        avatar_url: avatarUrl.value
                    }
                }
            )
            await navigateTo('/login')
        } catch (error) {
            console.error('Error signing up:', error);
        } finally {
            isLoading.value = false
        }

    }
</script>

<template>
    <v-card class="d-flex flex-column justify-center mx-auto w-50 m-20 border-lg rounded-lg">
        <v-card-title class="d-flex justify-center">サインアップ</v-card-title>
        <v-text-field v-model="name" class="w-1/2 mx-auto m-5 " label="ユーザーネーム" />
        <v-text-field v-model="email" class="w-1/2 mx-auto m-5 " label="メールアドレス" />
        <v-text-field v-model="password" class="w-1/2 mx-auto m-5" label="パスワード" type="password" />
        <v-btn class="d-flex justify-center" @click="signUpUser" color="primary">
            サインアップ
            <v-overlay v-model="isLoading"
                activator="parent"
                location-strategy="connected"
                class="d-flex justify-canter mx-auto my-auto"
            >
                <v-card class="d-flex items-center justify-center p-2 bg-black text-white mx-auto" max-width="200" min-height="100">
                    loading...
                </v-card>
            </v-overlay>
        </v-btn>
    </v-card>
</template>