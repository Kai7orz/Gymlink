export default defineEventHandler(async (event) => {
    const  idToken  = getRequestHeader(event,'authorization')
    const userId = getRouterParam(event,'user_id')
    console.log("user?oid--->>",userId)
    const data = $fetch(`http://host.docker.internal:3001/user_profiles/${userId}`,{
                         headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                }
    )
    console.log("dataadata:::",data)
    return data
})