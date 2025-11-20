import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface User {
  id: number | string
  [key: string]: unknown
}

interface CreditCheckBody {
  nationalId?: string
  mobile?: string
  amount?: number
  installmentCount?: number
  productId?: number
  [key: string]: unknown
}

interface CustomerRow {
  id: number
  name: string
  national_id: string
  mobile: string
  birth_date: string
  created_at: string
  credit_score: number
  total_loans: number
  overdue_loans: number
  membership_duration: number
  [key: string]: unknown
}

declare const requireAdminAuth: (event: H3Event) => Promise<User | null>
declare const $db: {
  query: (sql: string, params?: unknown[]) => Promise<{ rows: unknown[]; rowCount: number }>
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
    const body = await readBody(event) as CreditCheckBody
        const { nationalId, mobile, amount, installmentCount } = body

    // اعتبارسنجی داده‌های ورودی
    if (!nationalId || !mobile || !amount || !installmentCount) {
      throw createError({
        statusCode: 400,
        message: 'تمام فیلدهای ضروری باید پر شوند'
      })
    }

    // بررسی فرمت کد ملی
    if (!/^\d{10}$/.test(nationalId)) {
      throw createError({
        statusCode: 400,
        message: 'کد ملی باید 10 رقم باشد'
      })
    }

    // بررسی فرمت شماره موبایل
    if (!/^09\d{9}$/.test(mobile)) {
      throw createError({
        statusCode: 400,
        message: 'شماره موبایل باید با 09 شروع شود و 11 رقم باشد'
      })
    }

    // بررسی مبلغ
    if (amount <= 0) {
      throw createError({
        statusCode: 400,
        message: 'مبلغ باید بیشتر از صفر باشد'
      })
    }

    // بررسی تعداد اقساط
    const validInstallmentCounts = [3, 6, 12, 24]
    if (!validInstallmentCounts.includes(installmentCount)) {
      throw createError({
        statusCode: 400,
        message: 'تعداد اقساط باید یکی از مقادیر 3، 6، 12 یا 24 باشد'
      })
    }

    // دریافت اطلاعات مشتری از دیتابیس
    const customerQuery = `
      SELECT 
        id,
        name,
        national_id,
        mobile,
        birth_date,
        created_at,
        credit_score,
        total_loans,
        overdue_loans,
        membership_duration
      FROM users 
      WHERE national_id = $1 OR mobile = $2
    `
    const customerResult = await $db.query(customerQuery, [nationalId, mobile])
    const customer = customerResult.rows[0] as CustomerRow | undefined

    // محاسبه امتیاز اعتباری
    let creditScore = 500 // امتیاز پایه
    // const recommendations: string[] = []

    if (customer) {
      // محاسبه سن
      const birthDate = new Date(customer.birth_date)
      const age = new Date().getFullYear() - birthDate.getFullYear()

      // امتیاز بر اساس سن
      if (age >= 25 && age <= 55) {
        creditScore += 100
      } else if (age >= 18 && age < 25) {
        creditScore += 50
      } else {
        creditScore -= 50
        recommendations.push('سن شما برای دریافت وام مناسب نیست')
      }

      // امتیاز بر اساس مدت عضویت
      const membershipYears = Math.floor(customer.membership_duration / 12)
      creditScore += membershipYears * 20

      // امتیاز بر اساس تاریخچه وام‌ها
      if (customer.total_loans > 0) {
        const repaymentRate = ((customer.total_loans - customer.overdue_loans) / customer.total_loans) * 100
        if (repaymentRate >= 90) {
          creditScore += 150
        } else if (repaymentRate >= 70) {
          creditScore += 50
        } else {
          creditScore -= 100
          recommendations.push('تاریخچه بازپرداخت شما مناسب نیست')
        }
      }

      // امتیاز بر اساس امتیاز قبلی
      if (customer.credit_score) {
        creditScore = Math.floor((creditScore + customer.credit_score) / 2)
      }
    } else {
      // مشتری جدید
      creditScore -= 100
      recommendations.push('شما مشتری جدید هستید، امتیاز اعتباری پایین‌تری دارید')
    }

    // بررسی نسبت درآمد به بدهی
    const monthlyPayment = amount / installmentCount
    const debtToIncomeRatio = (monthlyPayment / 5000000) * 100 // فرض درآمد ماهانه 5 میلیون

    if (debtToIncomeRatio > 50) {
      creditScore -= 100
      recommendations.push('نسبت بدهی به درآمد شما بالا است')
    } else if (debtToIncomeRatio > 30) {
      creditScore -= 50
      recommendations.push('نسبت بدهی به درآمد شما در حد متوسط است')
    }

    // محدود کردن امتیاز بین 0 تا 1000
    creditScore = Math.max(0, Math.min(1000, creditScore))

    // تعیین وضعیت
    let status = 'رد شده'
    let maxAmount = 0

    if (creditScore >= 700) {
      status = 'تایید شده'
      maxAmount = amount * 2
    } else if (creditScore >= 500) {
      status = 'نیاز به بررسی بیشتر'
      maxAmount = amount
    } else {
      status = 'رد شده'
      maxAmount = 0
      recommendations.push('امتیاز اعتباری شما برای دریافت وام کافی نیست')
    }

    // محاسبه تاریخچه اعتباری
    const history = {
      totalLoans: customer?.total_loans || 0,
      overdueLoans: customer?.overdue_loans || 0,
      membershipDuration: customer?.membership_duration || 0
    }

    // ذخیره نتیجه اعتبارسنجی در دیتابیس
    const creditCheckQuery = `
      INSERT INTO credit_checks (
        national_id,
        mobile,
        requested_amount,
        installment_count,
        credit_score,
        status,
        max_amount,
        recommendations,
        checked_by,
        created_at
      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
      RETURNING id
    `

    await $db.query(creditCheckQuery, [
      nationalId,
      mobile,
      amount,
      installmentCount,
      creditScore,
      status,
      maxAmount,
      JSON.stringify(recommendations),
      user.id
    ])

    return {
      success: true,
      data: {
        score: creditScore,
        status,
        maxAmount,
        history,
        recommendations
      }
    }

  } catch (error: unknown) {
    console.error('خطا در اعتبارسنجی:', error)
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string }
    throw createError({
      statusCode: errorWithStatus?.statusCode || 500,
      message: errorWithStatus?.statusMessage || 'خطا در اعتبارسنجی'
    })
  }
}) 
