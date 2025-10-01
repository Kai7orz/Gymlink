export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch(`http://host.docker.internal:3001/exercises`,
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