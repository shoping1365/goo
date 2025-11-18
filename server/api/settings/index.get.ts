import { proxy } from '../_utils/fetchProxy'

// مسیردهی تنظیمات به Go backend
export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const url = new URL(event.node.req.url || '', 'http://localhost')
  const category = url.searchParams.get('category')

  if (category) {
    return await proxy(event, `${config.public.goApiBase}/api/admin/settings/category/${encodeURIComponent(category)}`)
  }

  // در صورت نبود category، تنظیمات عمومی
  return await proxy(event, `${config.public.goApiBase}/api/settings`)
})