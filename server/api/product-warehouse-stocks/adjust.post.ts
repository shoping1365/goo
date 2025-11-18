import { defineEventHandler, readBody, createError } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const body = await readBody(event)
     // body: { product_id: number, warehouse_id: number, delta: number, reason?: string }
     const { product_id, warehouse_id, delta, reason } = body || {}
     if (!product_id || !warehouse_id || (!Number.isInteger(delta))) {
          throw createError({ statusCode: 400, message: 'invalid payload' })
     }
     return await fetchGo(event, `/api/product-warehouse-stocks/${product_id}/adjust`, {
          method: 'POST',
          body: { warehouse_id, delta, reason }
     })
})
