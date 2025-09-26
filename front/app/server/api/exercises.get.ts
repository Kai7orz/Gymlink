export default defineEventHandler(async (event)=> {
    const data = await $fetch(`http://swagger-api:4010/exercises`)
    console.log("testtest",data)
    return data
})