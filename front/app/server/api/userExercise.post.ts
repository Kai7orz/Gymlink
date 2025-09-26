export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeaders(event,'authorization')
    const body = await readBody(event)
    console.log("body--.",body.image_url,body.time,body.date,body.comment)
    const data = await $fetch(`http://swagger-api:4010/exercises`,
                                {
                                    method: 'POST',
                                    headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                    body: {
                                        image_url: body.image_url,
                                        time: Number(body.time),
                                        date: body.date,
                                        comment: body.comment,
                                    }
                                }
    )
    return data

})