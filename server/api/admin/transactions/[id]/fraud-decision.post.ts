import { defineEventHandler, readBody } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const id = event.context.params!.id
  const body = await readBody(event)
  return await $fetch(`${config.public.goApiBase}/api/admin/fraud/cases/${id}/decision`, {
    method: 'POST',
    body
  })
})


