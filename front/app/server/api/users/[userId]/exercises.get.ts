export default defineEventHandler(async (event)=> {
    const  userId  = getRouterParam(event,'userId')
    const  idToken  = getRequestHeaders(event,'authorization')
    const data = await $fetch(`http://swagger-api:4010/users/${userId}/exercises-example`,
                                {
                                    method: 'GET',
                                    headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    }
                                }
    )

    return data

})