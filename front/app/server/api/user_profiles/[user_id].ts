export default defineEventHandler(async (event) => {
    const userId = getRouterParam(event,'user_id')
    const data = $fetch(`http://swagger-api:4010/user_profiles/${userId}`)
    return data
})