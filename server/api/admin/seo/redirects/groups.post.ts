export default defineEventHandler(async (event) => {
     try {
          const body = await readBody(event)

          if (!body.groupName) {
               throw createError({
                    statusCode: 400,
                    message: 'نام دسته‌بندی الزامی است'
               })
          }

          // دسته‌بندی‌های پیش‌فرض که می‌توان ایجاد کرد
          const defaultGroups = [
               {
                    name: 'انتقال محصولات',
                    description: 'ریدایرکت‌های محصولات انتقال یافته از وردپرس',
                    color: 'blue'
               },
               {
                    name: 'تغییرات URL',
                    description: 'تغییر مسیرهای ناشی از تغییر ساختار URL',
                    color: 'green'
               },
               {
                    name: 'حذف صفحات',
                    description: 'صفحاتی که حذف شده‌اند و به صفحه 404 یا صفحه جایگزین هدایت می‌شوند',
                    color: 'red'
               },
               {
                    name: 'SEO Optimization',
                    description: 'بهینه‌سازی URL برای سئو',
                    color: 'purple'
               },
               {
                    name: 'Mobile Redirects',
                    description: 'ریدایرکت‌های مخصوص موبایل',
                    color: 'orange'
               },
               {
                    name: 'Landing Pages',
                    description: 'صفحات فرود و کمپین‌ها',
                    color: 'yellow'
               }
          ]

          // بررسی وجود دسته‌بندی
          const existingGroup = defaultGroups.find(g => g.name === body.groupName)

          if (existingGroup) {
               return {
                    success: true,
                    message: `دسته‌بندی "${body.groupName}" آماده استفاده است`,
                    data: existingGroup
               }
          }

          // ایجاد دسته‌بندی جدید
          const newGroup = {
               name: body.groupName,
               description: body.description || `دسته‌بندی ${body.groupName}`,
               color: body.color || 'gray'
          }

          return {
               success: true,
               message: `دسته‌بندی "${body.groupName}" ایجاد شد`,
               data: newGroup
          }

     } catch (error: unknown) {
          throw createError({
               statusCode: 500,
               message: 'خطا در ایجاد دسته‌بندی: ' + ((error as { message?: string }).message || 'خطای نامشخص')
          })
     }
})
