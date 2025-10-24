export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const data = await $fetch(`http://go:8080/records`,
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
