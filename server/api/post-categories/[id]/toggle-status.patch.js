export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { proxy } = await import('../../_utils/fetchProxy')
  const { id } = event.context.params
  return proxy(event, `${base}/api/post-categories/${id}/toggle-status`, {
    method: 'PATCH'
  })
})