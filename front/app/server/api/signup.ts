export default defineEventHandler(async (event)=>{
    const body = await readBody(event)
    console.log("body:::",body.name,body.email,body.avatar_url)
    const data = await $fetch(`http://host.docker.internal:3001/users`,
        {
            method: 'POST',
            body: {
                name: body.name,
                email: body.email,
                avatar_url: body.avatar_url,
            }
        }
    )
    }
)