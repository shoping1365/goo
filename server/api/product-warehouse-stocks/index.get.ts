import { defineEventHandler, getQuery } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const q = getQuery(event) as { product_id?: string }
     const productId = q.product_id
     if (!productId) {
          return []
     }
     return await fetchGo(event, `/api/product-warehouse-stocks/${productId}`, { method: 'GET' })
})









