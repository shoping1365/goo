import { defineEventHandler, createError, readBody, getRouterParam } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const id = getRouterParam(event, 'id')
    
    if (!id) {
      throw createError({
        statusCode: 400,
        message: 'شناسه حساب الزامی است'
      })
    }
    
    // اعتبارسنجی داده‌ها
    if (!body.bankName || !body.cardNumber || !body.accountNumber || !body.accountHolderName) {
      throw createError({
        statusCode: 400,
        message: 'نام بانک، شماره کارت، شماره حساب و نام صاحب حساب الزامی است'
      })
    }

    // ارسال درخواست به بک‌اند Go
    const response = await goApiFetch<{ success: boolean; data?: unknown }>(event, `/api/user/banking-info/${id}`, {
      method: 'PUT',
      body: {
        bank_name: body.bankName,
        card_number: body.cardNumber.replace(/\D/g, ''),
        account_number: body.accountNumber,
        sheba_number: body.shebaNumber || '',
        account_holder_name: body.accountHolderName || '',
        account_type: body.accountType || '',
        is_default: body.isDefault || false,
        notes: body.notes || ''
      }
    })

    return {
      success: true,
      message: 'حساب با موفقیت به‌روزرسانی شد',
      data: response.data
    }
  } catch (error: unknown) {
    console.error('خطا در به‌روزرسانی حساب بانکی:', error)
    
    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'خطا در به‌روزرسانی حساب بانکی'
    })
  }
})