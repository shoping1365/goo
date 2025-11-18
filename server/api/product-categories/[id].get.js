export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { id } = event.context.params
  const { proxy } = await import('../_utils/fetchProxy')
  return proxy(event, `${base}/api/product-categories/${id}`)
})