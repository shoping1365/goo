import { defineEventHandler } from 'h3'
import { proxy } from './_utils/fetchProxy'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const response = await proxy(event, `${config.public.goApiBase}/api/roles`)
  
  // فیلتر کردن نقش customer از لیست نمایشی پنل ادمین
  if (response && response.data && Array.isArray(response.data)) {
    response.data = response.data.filter((role: { name?: string }) => role.name !== 'customer')
  }
  
  return response
}) 