import { proxy } from '../_utils/fetchProxy'

// خواندن تنظیمات واقعی از Go backend (دسته‌بندی compression)
export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  return await proxy(event, `${config.public.goApiBase}/api/admin/settings/category/compression`)
})