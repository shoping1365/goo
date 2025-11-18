import { proxy } from './_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  return await proxy(event, `${base}/api/media/upload`)
})