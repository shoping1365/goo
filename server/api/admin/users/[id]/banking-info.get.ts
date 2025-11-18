import type { H3Event } from 'h3'
import { defineEventHandler, createError, getRouterParam } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

declare const requireAdminAuth: (event: H3Event) => Promise<{ id: number | string; role: string } | null>

interface BankingInfoResponse {
  success: boolean
  data?: {
    bankingInfos?: unknown[]
  }
}

export default defineEventHandler(async (event) => {
  try {
    // احراز هویت ادمین
    const adminUser = await requireAdminAuth(event)
    if (!adminUser) {
      throw createError({
        statusCode: 401,
        message: 'احراز هویت ادمین الزامی است'
      })
    }

    // دریافت پارامترهای URL
    const userId = getRouterParam(event, 'id')
    if (!userId) {
      throw createError({
        statusCode: 400,
        message: 'شناسه کاربر الزامی است'
      })
    }

    // ارسال درخواست به بک‌اند Go
    const data = await goApiFetch<BankingInfoResponse>(event, `/api/admin/user/${userId}/banking-info`, {
      method: 'GET'
    })
    
    return {
      success: true,
      data: {
        bankingInfos: data.data?.bankingInfos || []
      }
    }

  } catch (error: unknown) {
    console.error('خطا در دریافت اطلاعات بانکی کاربر:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در دریافت اطلاعات بانکی کاربر'
    })
  }
})



