export default defineEventHandler(async (event) => {
    const  idToken  = getRequestHeader(event,'authorization')
    const userId = getRouterParam(event,'user_id')
    const data = $fetch(`http://host.docker.internal:3001/user_profiles/${userId}`)
    return data
})