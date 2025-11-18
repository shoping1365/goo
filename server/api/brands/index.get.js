export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { proxy } = await import('../_utils/fetchProxy')
  return proxy(event, `${base}/api/brands`)
})