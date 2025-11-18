// Validation utilities for frontend
export const validationUtils = {
     /**
      * پاک کردن شماره موبایل از کاراکترهای اضافی و نرمال‌سازی
      */
     cleanPhoneNumber: (phone: string): string => {
          // حذف فاصله، خط تیره و کاراکترهای غیر عددی
          let cleanPhone = phone.replace(/[\s\-\(\)]/g, '')

          // حذف +98 از ابتدای شماره
          if (cleanPhone.startsWith('+98')) {
               cleanPhone = cleanPhone.substring(3)
          }

          // اگر شماره با 0 شروع نمی‌شود، 0 اضافه کن
          if (!cleanPhone.startsWith('0')) {
               cleanPhone = '0' + cleanPhone
          }

          return cleanPhone
     },

     /**
      * اعتبارسنجی شماره موبایل ایرانی
      */
     isValidPhoneNumber: (phone: string): boolean => {
          const cleanPhone = validationUtils.cleanPhoneNumber(phone)
          const phoneRegex = /^09\d{9}$/
          return phoneRegex.test(cleanPhone) && cleanPhone.length === 11
     }
}
