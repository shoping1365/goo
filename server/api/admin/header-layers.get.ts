import { defineEventHandler, getQuery, createError } from 'h3'
import { $fetch } from 'ofetch'

export default defineEventHandler(async (event) => {
     try {
          const query = getQuery(event)
          const isTemporary = query.isTemporary === 'true'

          // دریافت لایه‌های موقت یا اصلی
          const response = await $fetch('/api/admin/header-layers', {
               method: 'GET',
               query: {
                    isTemporary: isTemporary
               }
          })

          return {
               success: true,
               data: response,
               message: 'لایه‌ها با موفقیت دریافت شدند'
          }

     } catch (error) {
          console.error('خطا در دریافت لایه‌ها:', error)

          throw createError({
               statusCode: 500,
               message: 'خطا در دریافت لایه‌ها'
          })
     }
})










