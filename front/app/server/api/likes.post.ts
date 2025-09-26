export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeaders(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch(`http://swagger-api:4010/likes`,
                                {
                                    method: 'POST',
                                    headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                    body: {
                                        exercise_record_id: Number(body.exercise_record_id)
                                    }
                                }
    )
    return data

})