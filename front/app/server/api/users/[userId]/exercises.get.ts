export default defineEventHandler(async (event)=> {
    const  userId  = getRouterParam(event,'userId')
    const  idToken  = getRequestHeader(event,'authorization')
    const data = await $fetch(`http://host.docker.internal:3001/users/${userId}/exercises`,
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