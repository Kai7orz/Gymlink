export default defineEventHandler(async(event)=>{
const exerciseRecordId = getRouterParam(event,'id')
  const data = await $fetch(`http://swagger-api:4010/exerciseRecords/${exerciseRecordId}/likes`,
    {
        method: 'POST',
    }
  )
})