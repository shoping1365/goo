import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const body = await readBody(event)
     return await fetchGo(event, '/api/admin/warehouses', {
          method: 'POST',
          body
     })
})


