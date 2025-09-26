export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeaders(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch(`http://swagger-api:4010/follows`,
                                {
                                    method: 'POST',
                                    headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                    body: {
                                        follower_id: Number(body.follower_id),
                                        followed_id: Number(body.followed_id)
                                    }
                                }
    )
    return data

})