import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const id = event.context.params?.id
     return await fetchGo(event, `/api/admin/warehouses/${id}`, { method: 'DELETE' })
})


