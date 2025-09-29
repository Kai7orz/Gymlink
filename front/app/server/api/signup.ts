export default defineEventHandler(async (event)=>{
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    console.log("idtoken-->",idToken)
    const data = await $fetch(`http://host.docker.internal:3001/users`,
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