import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const id = event.context.params?.id
     const body = await readBody(event)
     return await fetchGo(event, `/api/admin/warehouses/${id}`, {
          method: 'PUT',
          body
     })
})


