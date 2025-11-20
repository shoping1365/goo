import { defineNuxtRouteMiddleware, useState } from 'nuxt/app'

// میدل‌ور مدیریت نمایش هدر
// این میدل‌ور بررسی می‌کند که آیا هدر باید در صفحه فعلی نمایش داده شود یا نه

export default defineNuxtRouteMiddleware(async (to, from) => {
     
     // اگر مسیر تغییر نکرده، اجرا نکن
     if (to.path === from.path) {
          return
     }

     // API routes را از middleware مستثنا کن
     if (to.path.startsWith('/api/')) {
          return
     }

     // فقط در صفحات خارج از پنل ادمین و صفحات احراز هویت اجرا شود
     if (to.path.startsWith('/admin') ||
          to.path.includes('/admin/') ||
          to.path === '/admin' ||
          to.path.startsWith('/auth') ||
          to.path.includes('/auth/') ||
          to.path === '/auth' ||
          to.path.includes('/auth/login') ||
          to.path.includes('/register') ||
          to.path.includes('/verify') ||
          to.path.includes('/forgot-password') ||
          to.path.includes('/reset-password') ||
          to.path === '/auth/login' ||
          to.path === '/register' ||
          to.path === '/verify') {
          return
     }

     // استفاده از state برای کش کردن نتیجه و جلوگیری از درخواست‌های مکرر
     const headerDisplayState = useState('headerDisplayState', () => ({
          shouldShow: true,
          headerData: null,
          lastFetch: 0,
          isFetching: false
     }))

     // اگر در حال fetch هستیم، از درخواست جدید جلوگیری کنیم
     if (headerDisplayState.value.isFetching) {
          return
     }

     // اگر کمتر از 5 دقیقه از آخرین fetch گذشته، از کش استفاده کنیم
     const now = Date.now()
     const CACHE_DURATION = 5 * 60 * 1000 // 5 دقیقه
     if (now - headerDisplayState.value.lastFetch < CACHE_DURATION && headerDisplayState.value.headerData) {
          return
     }

     try {
          headerDisplayState.value.isFetching = true

          // بارگذاری هدرها از سرور
          const response = await $fetch('/api/header-settings') as { success: boolean; data: Record<string, any>[] }

          if (response && response.success && response.data) {
               const headers = response.data
               const activeHeader = headers.find((header: Record<string, any>) => header.is_active)

               if (activeHeader) {
                    // استفاده از نام‌های صحیح فیلدهای دیتابیس
                    const pageSelection = activeHeader.page_selection
                    const specificPages = activeHeader.specific_pages
                    const excludedPages = activeHeader.excluded_pages
                    const currentPath = to.path

                    let shouldShowHeader = true

                    // بررسی منطق نمایش هدر
                    if (pageSelection === 'specific') {
                         if (specificPages && specificPages.trim() !== '') {
                              // تبدیل رشته صفحات خاص به آرایه و حذف موارد تکراری
                              const pages = specificPages
                                   .split(/[\n,]/) // تقسیم بر اساس خط جدید یا کاما
                                   .map((page: string) => page.trim())
                                   .filter((page: string) => page.length > 0) // حذف خطوط خالی
                                   .filter((page: string, index: number, arr: string[]) => arr.indexOf(page) === index) // حذف موارد تکراری

                              shouldShowHeader = pages.some((page: string) => {
                                   // حذف http://localhost:3000 از ابتدای مسیر اگر وجود داشته باشد
                                   let cleanPage = page
                                   if (cleanPage.startsWith('http://localhost:3000')) {
                                        cleanPage = cleanPage.replace('http://localhost:3000', '')
                                   }
                                   if (cleanPage.startsWith('https://localhost:3000')) {
                                        cleanPage = cleanPage.replace('https://localhost:3000', '')
                                   }

                                   // بررسی دقیق مسیر
                                   if (cleanPage === currentPath) {
                                        return true
                                   }

                                   // بررسی مسیرهای نسبی (مثل /admin/*)
                                   if (cleanPage.endsWith('*')) {
                                        const basePath = cleanPage.slice(0, -1)
                                        if (currentPath.startsWith(basePath)) {
                                             return true
                                        }
                                   }

                                   return false
                              })
                         } else {
                              // اگر صفحات خاص تعریف نشده باشند، نمایش نده
                              shouldShowHeader = false
                         }
                    } else if (pageSelection === 'exclude') {
                         if (excludedPages && excludedPages.trim() !== '') {
                              // تبدیل رشته صفحات مستثنی به آرایه و حذف موارد تکراری
                              const pages = excludedPages
                                   .split(/[\n,]/) // تقسیم بر اساس خط جدید یا کاما
                                   .map((page: string) => page.trim())
                                   .filter((page: string) => page.length > 0) // حذف خطوط خالی
                                   .filter((page: string, index: number, arr: string[]) => arr.indexOf(page) === index) // حذف موارد تکراری

                              const shouldExclude = pages.some((page: string) => {
                                   // حذف http://localhost:3000 از ابتدای مسیر اگر وجود داشته باشد
                                   let cleanPage = page
                                   if (cleanPage.startsWith('http://localhost:3000')) {
                                        cleanPage = cleanPage.replace('http://localhost:3000', '')
                                   }
                                   if (cleanPage.startsWith('https://localhost:3000')) {
                                        cleanPage = cleanPage.replace('https://localhost:3000', '')
                                   }

                                   // بررسی دقیق مسیر
                                   if (cleanPage === currentPath) {
                                        return true
                                   }

                                   // بررسی مسیرهای نسبی (مثل /admin/*)
                                   if (cleanPage.endsWith('*')) {
                                        const basePath = cleanPage.slice(0, -1)
                                        if (currentPath.startsWith(basePath)) {
                                             return true
                                        }
                                   }

                                   return false
                              })

                              shouldShowHeader = !shouldExclude
                         } else {
                              // اگر صفحات مستثنی تعریف نشده باشند، همیشه نمایش بده
                              shouldShowHeader = true
                         }
                    }

                    // ذخیره نتیجه در state برای استفاده در کامپوننت Header
                    headerDisplayState.value = {
                         shouldShow: shouldShowHeader,
                         headerData: activeHeader,
                         lastFetch: now,
                         isFetching: false
                    }
               } else {
                    // اگر هدر فعالی نباشد، هدر را مخفی کن
                    headerDisplayState.value = {
                         shouldShow: false,
                         headerData: null,
                         lastFetch: now,
                         isFetching: false
                    }
               }
          }
     } catch (_error) {
          // خطا در میدل‌ور نمایش هدر

          // در صورت خطا، هدر را نمایش بده
          headerDisplayState.value = {
               shouldShow: true,
               headerData: null,
               lastFetch: now,
               isFetching: false
          }
     }
})
