import { createError, defineEventHandler } from 'h3'
import { getDatabase } from '../_utils/database'

export default defineEventHandler(async (event) => {
  try {
    // Fetching wishlist items

    const userId = event.context?.user?.id

    // User ID check

    if (!userId) {
      throw createError({
        statusCode: 401,
        message: 'User not authenticated'
      })
    }

    // Get database connection
    const db = await getDatabase()
    // Connected to database

    // Get wishlist items
    const wishlistItems = await db.query(`
      SELECT 
        uci.id,
        uci.product_id,
        uci.created_at,
        p.name,
        p.price AS base_price,
        p.sale_price,
        p.sale_start_at,
        p.sale_end_at,
        p.image,
        p.stock_quantity,
        p.show_stock_to_customer,
        p.track_inventory,
        p.call_for_price,
        p.disable_buy_button,
        CASE
          WHEN p.call_for_price = TRUE THEN NULL
          WHEN p.sale_price IS NOT NULL AND p.sale_price > 0
               AND (p.sale_start_at IS NULL OR p.sale_start_at <= NOW())
               AND (p.sale_end_at IS NULL OR p.sale_end_at >= NOW())
            THEN p.sale_price
          ELSE p.price
        END AS effective_price
      FROM user_collection_items uci
      LEFT JOIN user_collections uc ON uc.id = uci.collection_id
      LEFT JOIN products p ON p.id = uci.product_id
      WHERE uc.user_id = $1 AND uc.is_default = true
      ORDER BY uci.created_at DESC
    `, [userId])

    // Found wishlist items

    // Convert to product objects
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const products = wishlistItems.map((item: any) => {
      const available = !item.disable_buy_button && !item.call_for_price && (item.stock_quantity > 0 || !item.track_inventory)
      const rawPrice = item.effective_price ?? item.sale_price ?? item.base_price
      const price = item.call_for_price ? null : (rawPrice == null ? null : Number(rawPrice))

      return {
        id: item.product_id,
        name: item.name || `Product ${item.product_id}`,
        price,
        image: item.image || '/default-product.svg',
        stock_quantity: Number(item.stock_quantity || 0),
        available,
        stock_status: item.call_for_price ? 'Call for price' :
          item.disable_buy_button ? 'Out of stock' :
            (item.track_inventory && Number(item.stock_quantity) <= 0) ? 'Out of stock' :
              (item.track_inventory && Number(item.stock_quantity) > 0 && Number(item.stock_quantity) <= 5) ? `In stock ${item.stock_quantity} units` : null
      }
    })

    return {
      success: true,
      data: products,
      count: products.length
    }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (error: any) {
    console.error('Error fetching wishlist items:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'Error fetching wishlist items'
    })
  }
})
