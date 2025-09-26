export default defineEventHandler((event)=>{
    console.log("loginPost is called")
    return $fetch("http://swagger-api:4010/login",{
        method: 'POST',
    })
})