export default defineEventHandler(async (event)=> {
    const  idToken  = getRequestHeader(event,'authorization')
    const data = await $fetch(`http://host.docker.internal:3001/exercises`,
            {
                method: 'GET',
                headers: {
                    'Authorization': idToken,
                    'Content-Type': 'application/json'
                }
            }
    )
    console.log("testtest",data)
    return data
})