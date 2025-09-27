import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware(async (to,from) => {
    if(process.server) return 

    const auth = useAuthStore()
    const allowlist = new Set<string>(['/','/login','/signin','/signup'])

    console.log("middleware: auth uid email",auth.uid,auth.email,auth.isAuthenticated,allowlist.has(to.path),auth.loading)

    if(allowlist.has(to.path)) return 

    if(auth.loading){
        await new Promise<void>((resolve)=>{
            const stop = watch(()=> auth.loading, (v)=>{
                if(!v){
                    stop();
                    resolve()
                }
            },{
                immediate:true
            })
        })
    }

    if(!auth.isAuthenticated){
        console.log("middleware is blocking")
        return navigateTo('/login',{replace:true})
    }  
})