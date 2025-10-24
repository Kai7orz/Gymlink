export default defineEventHandler(async (event) => {
    const  idToken  = getRequestHeader(event,'authorization')
    const userId = getRouterParam(event,'user_id')
    const recordId = getRouterParam(event,'record_id')

    const data = await $fetch(`http://go:8080/users/${userId}/records/${recordId}`,{
                         method: 'DELETE',
                         headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                }
    )
    return data
})
