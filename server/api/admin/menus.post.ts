import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const config = useRuntimeConfig()

  return await proxy(event, `${config.public.goApiBase}/api/menus`, {
    method: 'POST',
    body: JSON.stringify(body)
  })
}) 