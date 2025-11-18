import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface User {
  id: number | string
  [key: string]: unknown
}

interface CreditResult {
  score: number
  maxAmount: number
  status: string
  [key: string]: unknown
}

interface InstallmentBody {
  nationalId?: string
  mobile?: string
  amount?: number
  installmentCount?: number
  productId?: number
  creditResult?: CreditResult
  [key: string]: unknown
}

interface ProductRow {
  id: number
  name: string
  sku: string
  price: number
  installment_price: number
  [key: string]: unknown
}

interface UserRow {
  id: string
  [key: string]: unknown
}

interface InstallmentRow {
  id: number
  [key: string]: unknown
}

declare const requireAdminAuth: (event: H3Event) => Promise<User | null>
declare const $db: {
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
  connect: () => Promise<{ query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>; release: () => void }>
}

export default defineEventHandler(async (event) => {
  try {
    // بررسی دسترسی ادمین
    const user = await requireAdminAuth(event)
    if (!user) {
      throw createError({
        statusCode: 401,
        message: 'دسترسی غیرمجاز'
      })
    }

    // دریافت داده‌های درخواست
    const body = await readBody(event) as InstallmentBody
    const { nationalId, mobile, amount, installmentCount, productId, creditResult } = body

    // اعتبارسنجی داده‌های ورودی
    if (!nationalId || !mobile || !amount || !installmentCount || !productId) {
      throw createError({
        statusCode: 400,
        message: 'تمام فیلدهای ضروری باید پر شوند'
      })
    }

    // بررسی وضعیت اعتبارسنجی
    if (!creditResult || creditResult.status === 'رد شده') {
      throw createError({
        statusCode: 400,
        message: 'اعتبارسنجی تایید نشده است'
      })
    }

    // بررسی وجود محصول
    const productQuery = `
      SELECT id, name, sku, price, installment_price 
      FROM products 
      WHERE id = $1 AND installment_enabled = true
    `
    const productResult = await $db.query(productQuery, [productId])

    if (productResult.rows.length === 0) {
      throw createError({
        statusCode: 404,
        message: 'محصول یافت نشد یا امکان خرید اقساطی ندارد'
      })
    }

    const product = productResult.rows[0] as ProductRow

    // محاسبه مبلغ هر قسط
    const monthlyPayment = Math.ceil(amount / installmentCount)

    // بررسی محدودیت‌های اعتباری
    if (amount > creditResult.maxAmount) {
      throw createError({
        statusCode: 400,
        message: `مبلغ درخواستی بیشتر از حداکثر مبلغ مجاز (${creditResult.maxAmount.toLocaleString()} تومان) است`
      })
    }

    // شروع تراکنش دیتابیس
    const client = await $db.connect()

    try {
      await client.query('BEGIN')

      // ایجاد یا بروزرسانی کاربر
      let userId: string
      const userQuery = `
        SELECT id FROM users WHERE national_id = $1 OR mobile = $2
      `
      const userResult = await client.query(userQuery, [nationalId, mobile])

      if (userResult.rows.length > 0) {
        userId = (userResult.rows[0] as UserRow).id

        // بروزرسانی امتیاز اعتباری
        await client.query(`
          UPDATE users 
          SET credit_score = $1, updated_at = NOW()
          WHERE id = $2
        `, [creditResult.score, userId])
      } else {
        // ایجاد کاربر جدید
        const newUserQuery = `
          INSERT INTO users (
            name, national_id, mobile, credit_score, created_at, updated_at
          ) VALUES ($1, $2, $3, $4, NOW(), NOW())
          RETURNING id
        `
        const newUserResult = await client.query(newUserQuery, [
          `مشتری ${nationalId}`,
          nationalId,
          mobile,
          creditResult.score
        ])
        userId = (newUserResult.rows[0] as UserRow).id
      }

      // ایجاد خرید اقساطی
      const installmentQuery = `
        INSERT INTO installment_payments (
          user_id,
          customer_name,
          national_id,
          mobile,
          product_id,
          product_name,
          amount,
          installment_count,
          monthly_payment,
          credit_score,
          status,
          approved_by,
          approved_at,
          created_at,
          updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, NOW(), NOW(), NOW())
        RETURNING id
      `

      const installmentResult = await client.query(installmentQuery, [
        userId,
        `مشتری ${nationalId}`,
        nationalId,
        mobile,
        productId,
        product.name,
        amount,
        installmentCount,
        monthlyPayment,
        creditResult.score,
        'approved',
        user.id
      ])

      const installmentId = (installmentResult.rows[0] as InstallmentRow).id

      // ایجاد اقساط
      const installments = []
      const currentDate = new Date()

      for (let i = 1; i <= installmentCount; i++) {
        const dueDate = new Date(currentDate)
        dueDate.setMonth(dueDate.getMonth() + i)

        installments.push({
          installment_id: installmentId,
          installment_number: i,
          amount: monthlyPayment,
          due_date: dueDate,
          status: i === 1 ? 'current' : 'pending',
          created_at: new Date()
        })
      }

      // ذخیره اقساط
      for (const installment of installments) {
        await client.query(`
          INSERT INTO installment_schedules (
            installment_id,
            installment_number,
            amount,
            due_date,
            status,
            created_at
          ) VALUES ($1, $2, $3, $4, $5, $6)
        `, [
          installment.installment_id,
          installment.installment_number,
          installment.amount,
          installment.due_date,
          installment.status,
          installment.created_at
        ])
      }

      // ثبت لاگ
      await client.query(`
        INSERT INTO installment_logs (
          installment_id,
          action,
          description,
          performed_by,
          created_at
        ) VALUES ($1, $2, $3, $4, NOW())
      `, [
        installmentId,
        'approve',
        `خرید اقساطی تایید شد - مبلغ: ${amount.toLocaleString()} تومان - تعداد اقساط: ${installmentCount}`,
        user.id
      ])

      await client.query('COMMIT')

      return {
        success: true,
        data: {
          installmentId,
          message: 'خرید اقساطی با موفقیت تایید شد',
          installmentDetails: {
            amount,
            installmentCount,
            monthlyPayment,
            firstDueDate: installments[0].due_date
          }
        }
      }

    } catch (error: unknown) {
      await client.query('ROLLBACK')
      throw error
    } finally {
      client.release()
    }

  } catch (error: unknown) {
    console.error('خطا در تایید خرید اقساطی:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.statusMessage || 'خطا در تایید خرید اقساطی'
    })
  }
}) 
