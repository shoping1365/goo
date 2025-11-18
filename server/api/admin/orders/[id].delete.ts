import type { H3Event } from 'h3';
import { createError, defineEventHandler, getRouterParam } from 'h3';

interface Database {
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
}

declare const requireAuth: (event: H3Event) => Promise<{ [key: string]: unknown } | null>
declare const getDatabase: () => Promise<Database>

export default defineEventHandler(async (event: H3Event) => {
  await requireAuth(event)

  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه سفارش الزامی است'
      })
    }

    // دریافت اتصال دیتابیس
    const db = await getDatabase()

    // حذف آیتم‌های سفارش
    await db.query('DELETE FROM order_items WHERE order_id = $1', [id])

    // حذف سفارش و برگرداندن شناسه حذف‌شده برای تشخیص موفقیت
    const deleted = await db.query('DELETE FROM orders WHERE id = $1 RETURNING id', [id])

    if (!Array.isArray(deleted.rows) || deleted.rows.length === 0) {
      throw createError({
        statusCode: 404,
        message: 'سفارش یافت نشد'
      })
    }

    return {
      success: true,
      message: 'سفارش با موفقیت حذف شد'
    }
  } catch (error: unknown) {
    console.error('خطا در حذف سفارش:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string }
    if (errorWithStatus?.statusCode) {
      throw error
    }

    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.message || 'خطا در حذف سفارش'
    })
  }
})
