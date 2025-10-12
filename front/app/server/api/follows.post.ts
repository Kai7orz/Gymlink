export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    console.log("user id nitro:",body.follower_id)
    console.log("user id nitro:",body.followed_id)
    const data = await $fetch(`http://go:8080/follows`,
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