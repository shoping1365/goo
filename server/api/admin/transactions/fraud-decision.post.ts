import { defineEventHandler, readBody, createError, getQuery } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const id = event.context.params?.id || (getQuery(event).id as string)
  if (!id) {
    throw createError({ statusCode: 400, message: 'شناسه پرونده نامعتبر است' })
  }
  return await goApiFetch(event, `/api/admin/fraud/cases/${id}/decision`, {
    method: 'POST',
    body
  })
})


