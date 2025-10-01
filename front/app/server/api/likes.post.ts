export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch(`http://host.docker.internal:3001/likes`,
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