import { defineEventHandler, getQuery } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     // عبور مستقیم به سرویس Go با حفظ کوکی‌ها و هدرهای Dev
     // فیلترهای احتمالی از کوئری هم فوروارد می‌شود
     const query = getQuery(event)
     const qs = new URLSearchParams(query as Record<string, string>).toString()
     const path = `/api/admin/warehouses${qs ? `?${qs}` : ''}`
     return await fetchGo(event, path, { method: 'GET' })
})


