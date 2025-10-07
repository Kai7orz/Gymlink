export default defineEventHandler(async (event)=>{
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch(`http://go:8080/users`,
        {
            method: 'POST',
            headers: {
                'Authorization': idToken,
                'Content-Type': 'application/json'
            },
            body: {
                name: body.name,
                avatar_url: body.avatar_url,
            }
        }
    )
    }
)