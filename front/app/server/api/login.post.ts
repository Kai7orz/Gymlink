export default defineEventHandler(async (event)=>{
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    return $fetch("http://go:8080/users",{
        method: 'POST',
        headers: {
            'Authorization': idToken,
            'Content-Type': 'application/json'
        },
    })
})