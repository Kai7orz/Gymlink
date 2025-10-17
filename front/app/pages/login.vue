<script setup lang="ts">
    // middleware からのログインページへの遷移時にハイドレーションミスマッチがおきるので ClientOnly で対応している
    import {signIn}  from '~/composables/SignInUser';
    import { useUserStore } from '~/stores/userStore';

    const email = ref('')
    const password = ref('')
    const isLoading = ref(false)
    const firebaseData = ref({})
    const userData = ref({})
    const user = useUserStore()
    const auth = useAuthStore()

    await signOutUser()

    watch(isLoading, val => {
            val && setTimeout(() => {
            isLoading.value = false
            console.log("firebase data:",firebaseData.value)
            }, 3000)
        })

    const toSignUp = () => {
            navigateTo('/signup')
    }
    const signInUser = async () => {
        isLoading.value = true
        const minLoadingPromise = new Promise(resolve => setTimeout(resolve, 1000));

        try{

            firebaseData.value = await signIn(email.value,password.value)
            await minLoadingPromise;
            const TOKEN = auth.idToken
            await new Promise(resolve => setTimeout(resolve, 100));
            userData.value = await $fetch("/api/login",
                {
                    method: 'POST',
                    headers: {
                        'Authorization': 'Bearer ' + TOKEN,
                        'Content-Type': 'application/json'
                    },
                }
            )
            user.setUser(userData.value.id,userData.value.name)
            
            await navigateTo('/home')
        } catch (error) {
            console.error('SignIn Error:', error);
        } finally {
            isLoading.value = false
        }
    }
</script>

<template>
    <div>
    <ClientOnly>
        <v-card class="d-flex flex-column justify-center mx-auto w-50 m-20 border-lg bg-grey-darken-1 rounded-lg">
            <v-card-title class="d-flex justify-center">サインイン</v-card-title>
            <v-text-field v-model="email" class="w-1/2 mx-auto m-5 " label="メールアドレス" />
            <v-text-field v-model="password" class="w-1/2 mx-auto m-5" label="パスワード" type="password" />
            <v-btn class="d-flex justify-center m-5" @click="signInUser" color="primary">
                サインイン
                <v-overlay v-model="isLoading"
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
            <v-btn class="bg-black text-blue m-5"  @click="toSignUp">未登録の方 サインアップ </v-btn>
        </v-card>
    </ClientOnly>
    </div>
</template>