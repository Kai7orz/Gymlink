export default defineEventHandler((event)=>{
    const data = $fetch("http://swagger-api:4010/teams")
    return data
})