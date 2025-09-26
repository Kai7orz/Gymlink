import { defineStore, skipHydrate } from 'pinia'

export const useAuthStore = defineStore('auth',() =>{
    const idTokenRef = ref<string | null>(null)
    const uid = ref<string | null>(null)
    const email = ref<string | null>(null)
    const loading = ref(true)

    const isAuthenticated = computed(()=> !!idTokenRef.value && !!uid.value)

    const setAuth = (p: {idToken?:string|null,uid?: string|null,email?: string|null})=>{
        if(p.idToken !== undefined) idTokenRef.value = p.idToken
        if(p.uid !== undefined) uid.value = p.uid
        if(p.email !== undefined) email.value = p.email
    }

    const clearAuth = ()=>{
        idTokenRef.value = null;
        uid.value = null;
        email.value = null;
    }

    const setLoading = (v: boolean)=>{
        loading.value = v
    }

    return {
        idToken: skipHydrate(idTokenRef),
        uid,email,loading,isAuthenticated,setAuth,clearAuth,setLoading
    }
})