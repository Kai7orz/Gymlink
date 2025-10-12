export default defineEventHandler(async (event)=> {
    const  userId  = getRouterParam(event,'userId')
    const  idToken  = getRequestHeader(event,'authorization')
    const data = await $fetch(`http://go:8080/users/${userId}/exercises`,
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