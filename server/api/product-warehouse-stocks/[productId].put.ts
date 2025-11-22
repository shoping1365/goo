import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const { productId } = event.context.params as { productId?: string }
     const body = await readBody(event)
     // body: { entries: [{ warehouse_id, quantity, min_qty, max_qty }], default_warehouse_id?: number }
     return fetchGo(event, `/api/product-warehouse-stocks/${productId}`, {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(body)
     })
})
