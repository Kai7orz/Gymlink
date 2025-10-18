export default defineEventHandler(async (event) => {
    const  idToken  = getRequestHeader(event,'authorization')
    const recordId = getRouterParam(event,'record_id')
    const data = await $fetch(`http://go:8080/likes/${recordId}`,{
                         method: 'DELETE',
                         headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                }
    )
    return data
})