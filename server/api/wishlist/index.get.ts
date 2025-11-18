import { defineEventHandler, createError } from 'h3'
import { getDatabase } from '../_utils/database'

export default defineEventHandler(async (event) => {
  try {
    console.log('œ—ŒÊ«”  œ—Ì«›  ·Ì”  ⁄·«ﬁÂù„‰œÌùÂ«Ì ò«—»—')

    const userId = event.context?.user?.id

    console.log('ò«—»— «Õ—«“ ÂÊÌ  ‘œÂ:', userId ? { id: userId } : 'null')

    if (!userId) {
      console.log('ò«—»— «Õ—«“ ÂÊÌ  ‰‘œÂ «” ')
      throw createError({
        statusCode: 401,
        message: 'ò«—»— «Õ—«“ ÂÊÌ  ‰‘œÂ «” '
      })
    }

    // œ—Ì«›  « ’«· œÌ «»Ì”
    const db = await getDatabase()
    console.log('« ’«· œÌ «»Ì” »—ﬁ—«— ‘œ')

    // œ—Ì«›  „Õ’Ê·«  ⁄·«ﬁÂù„‰œÌ ò«—»— «“ œÌ «»Ì”
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

    console.log(' ⁄œ«œ „Õ’Ê·«  ⁄·«ﬁÂù„‰œÌ Ì«›  ‘œÂ:', wishlistItems.length)

    //  »œÌ· œ«œÂùÂ« »Â ›—„  „‰«”»
    const products = wishlistItems.map((item: any) => {
      const available = !item.disable_buy_button && !item.call_for_price && (item.stock_quantity > 0 || !item.track_inventory)
      const rawPrice = item.effective_price ?? item.sale_price ?? item.base_price
      // «ê—  „«” »—«Ì ﬁÌ„  »«‘œ° ﬁÌ„  —« null »—„Ìùê—œ«‰Ì„  « UI "ﬁÌ„  ‰«„‘Œ’" ‰„«Ì‘ œÂœ
      const price = item.call_for_price ? null : (rawPrice == null ? null : Number(rawPrice))

      return {
        id: item.product_id,
        name: item.name || `„Õ’Ê· ${item.product_id}`,
        price,
        image: item.image || '/default-product.svg',
        stock_quantity: Number(item.stock_quantity || 0),
        available,
        stock_status: item.call_for_price ? ' „«” »—«Ì ﬁÌ„ ' :
          item.disable_buy_button ? '€Ì—ﬁ«»· ›—Ê‘' :
            (item.track_inventory && Number(item.stock_quantity) <= 0) ? '‰«„ÊÃÊœ' :
              (item.track_inventory && Number(item.stock_quantity) > 0 && Number(item.stock_quantity) <= 5) ? ` ‰Â« ${item.stock_quantity} ⁄œœ œ— «‰»«— »«ﬁÌ „«‰œÂ` : null
      }
    })

    return {
      success: true,
      data: products,
      count: products.length
    }

  } catch (error: any) {
    console.error('Œÿ« œ— œ—Ì«›  ·Ì”  ⁄·«ﬁÂù„‰œÌùÂ«:', error)

    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'Œÿ« œ— œ—Ì«›  ·Ì”  ⁄·«ﬁÂù„‰œÌùÂ«'
    })
  }
})
