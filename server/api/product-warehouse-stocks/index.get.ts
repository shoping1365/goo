import { defineEventHandler, getQuery } from 'h3'
import { fetchGo } from '../_utils/fetchGo'

export default defineEventHandler(async (event) => {
     const q: any = getQuery(event)
     const productId = q.product_id
     if (!productId) {
          return []
     }
     return await fetchGo(event, `/api/product-warehouse-stocks/${productId}`, { method: 'GET' })
})









