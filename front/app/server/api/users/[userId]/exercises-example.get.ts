export default defineEventHandler(async (event)=> {
    const  userId  = getRouterParam(event,'userId')
    const data = await $fetch(`http://swagger-api:4010/users/${userId}/exercises-example`)
    console.log("$saide:",data)
    return data
})