import { defineEventHandler } from 'h3'
import { proxy } from '../_utils/fetchProxy'

export default defineEventHandler(async (event) => {
     try {
          const config = useRuntimeConfig()
          // استفاده از پروکسی استاندارد تا JWT از کوکی به Authorization منتقل شود
          const data = await proxy(event, `${config.public.goApiBase}/api/products`, { method: 'GET' })
          return data
     } catch (error) {
          console.error('Error fetching products:', error)

          // در صورت خطا، آرایه خالی برگردان تا UI گیر نکند
          return []
     }
}) 