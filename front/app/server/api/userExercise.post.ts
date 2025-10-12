export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const body = await readBody(event)
    const data = await $fetch("http://go:8080/exercises",
                                {
                                    method: 'POST',
                                    headers: {
                                        'Authorization': idToken,
                                        'Content-Type': 'application/json'
                                    },
                                    body: {
                                        exercise_image: body.exercise_image,
                                        exercise_time: Number(body.exercise_time),
                                        exercise_date: new Date(body.exercise_date),
                                        comment: body.comment,
                                    }
                                }
    )
    return data

})