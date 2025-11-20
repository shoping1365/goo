import { createError, defineEventHandler } from 'h3';

export default defineEventHandler(async (_) => {
     // بررسی احراز هویت و مجوزها
     try {
          // در اینجا باید دستور توقف worker به Go backend ارسال شود
          // این یک پیاده‌سازی ساده است

          // فعلاً یک پاسخ نمونه برمی‌گردانیم
          // در آینده باید با سرویس Go backend ارتباط برقرار کرد

          return {
               success: true,
               running: false,
               message: 'Worker stopped successfully'
          }
     } catch (error: unknown) {
          const errorWithStatus = error as { statusCode?: number; message?: string }
          throw createError({
               statusCode: errorWithStatus?.statusCode || 500,
               message: errorWithStatus?.message || 'Error stopping worker'
          })
     }
})


