export default defineEventHandler(async (event) => {
    const  idToken  = getRequestHeader(event,'authorization')
    const userId = getRouterParam(event,'user_id')
    const data = await $fetch(`http://host.docker.internal:3001/user_profiles/${userId}`,{
                         headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                }
    )
    return data
})