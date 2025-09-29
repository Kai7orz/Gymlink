export default defineEventHandler(async (event)=>{
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    return $fetch("http://host.docker.internal:3001/login",{
        method: 'POST',
        headers: {
            'Authorization': idToken,
            'Content-Type': 'application/json'
        },
    })
})