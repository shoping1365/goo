<template>
  <div class="seo-content-generation" dir="rtl">
    <div class="bg-white rounded-lg shadow-md p-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-bold text-gray-800">تولید محتوا</h1>
        <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors">
          تولید محتوای جدید
        </button>
      </div>

      <!-- آمار کلی -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
        <div class="bg-gradient-to-r from-blue-500 to-blue-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۵۰</div>
          <div class="text-sm">محتوای تولید شده</div>
        </div>
        <div class="bg-gradient-to-r from-green-500 to-green-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۸۵</div>
          <div class="text-sm">امتیاز کیفیت</div>
        </div>
        <div class="bg-gradient-to-r from-yellow-500 to-yellow-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۲۵</div>
          <div class="text-sm">در انتظار بررسی</div>
        </div>
        <div class="bg-gradient-to-r from-red-500 to-red-600 text-white p-6 rounded-lg">
          <div class="text-2xl font-bold">۱۲</div>
          <div class="text-sm">نیازمند بازنویسی</div>
        </div>
      </div>

      <!-- ابزار تولید محتوا -->
      <div class="bg-gray-50 p-6 rounded-lg mb-6">
        <h2 class="text-lg font-semibold text-gray-800 mb-4">ابزار تولید محتوا</h2>
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- تنظیمات محتوا -->
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">نوع محتوا</label>
              <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option>مقاله وبلاگ</option>
                <option>توضیحات محصول</option>
                <option>صفحه لندینگ</option>
                <option>محتوا شبکه‌های اجتماعی</option>
                <option>متا تگ</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">کلمات کلیدی اصلی</label>
              <input type="text" placeholder="کلمات کلیدی را وارد کنید..." class="w-full border border-gray-300 rounded-lg px-3 py-2">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">طول محتوا</label>
              <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option>کوتاه (۳۰۰-۵۰۰ کلمه)</option>
                <option>متوسط (۵۰۰-۱۰۰۰ کلمه)</option>
                <option>بلند (۱۰۰۰-۲۰۰۰ کلمه)</option>
                <option>خیلی بلند (۲۰۰۰+ کلمه)</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">تن (Tone)</label>
              <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option>حرفه‌ای</option>
                <option>دوستانه</option>
                <option>آموزشی</option>
                <option>ترغیب‌کننده</option>
                <option>اطلاع‌رسانی</option>
              </select>
            </div>
          </div>

          <!-- تنظیمات پیشرفته -->
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">سطح تخصص</label>
              <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option>مبتدی</option>
                <option>متوسط</option>
                <option>پیشرفته</option>
                <option>متخصص</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">زبان</label>
              <select class="w-full border border-gray-300 rounded-lg px-3 py-2">
                <option>فارسی</option>
                <option>انگلیسی</option>
                <option>عربی</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">ساختار محتوا</label>
              <div class="space-y-2">
                <label class="flex items-center">
                  <input type="checkbox" class="ml-2" checked>
                  <span class="text-sm text-gray-700">عنوان اصلی (H1)</span>
                </label>
                <label class="flex items-center">
                  <input type="checkbox" class="ml-2" checked>
                  <span class="text-sm text-gray-700">عنوان‌های فرعی (H2, H3)</span>
                </label>
                <label class="flex items-center">
                  <input type="checkbox" class="ml-2" checked>
                  <span class="text-sm text-gray-700">لیست‌ها</span>
                </label>
                <label class="flex items-center">
                  <input type="checkbox" class="ml-2">
                  <span class="text-sm text-gray-700">نقل قول</span>
                </label>
              </div>
            </div>
            <div class="pt-4">
              <button class="w-full bg-green-600 text-white px-6 py-3 rounded-lg hover:bg-green-700 transition-colors">
                تولید محتوا
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- نمونه محتوای تولید شده -->
      <div class="bg-white border border-gray-200 rounded-lg p-6 mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">نمونه محتوای تولید شده</h3>
        <div class="bg-gray-50 p-6 rounded-lg">
          <h4 class="text-lg font-semibold text-gray-800 mb-2">بهترین لپ تاپ‌های گیمینگ ۲۰۲۴</h4>
          <p class="text-gray-700 leading-relaxed">
            در دنیای امروز، لپ تاپ‌های گیمینگ به یکی از محبوب‌ترین ابزارهای سرگرمی تبدیل شده‌اند. 
            با پیشرفت تکنولوژی، این دستگاه‌ها حالا قابلیت اجرای بازی‌های سنگین را با کیفیت بالا دارند.
          </p>
          <div class="mt-4 flex items-center justify-between">
            <div class="flex items-center space-x-4 space-x-reverse">
              <span class="text-sm text-gray-500">طول: ۸۵۰ کلمه</span>
              <span class="text-sm text-gray-500">امتیاز سئو: ۹۲%</span>
            </div>
            <div class="flex space-x-2 space-x-reverse">
              <button class="text-blue-600 hover:text-blue-900 text-sm">ویرایش</button>
              <button class="text-green-600 hover:text-green-900 text-sm">کپی</button>
              <button class="text-red-600 hover:text-red-900 text-sm">حذف</button>
            </div>
          </div>
        </div>
      </div>

      <!-- جدول محتوای تولید شده -->
      <div class="bg-white border border-gray-200 rounded-lg overflow-hidden">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800">محتوای تولید شده</h3>
        </div>
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عنوان</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نوع</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">طول</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز سئو</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="content in contents" :key="content.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{{ content.title }}</div>
                <div class="text-sm text-gray-500">{{ content.keywords }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getTypeClass(content.type)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ content.type }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ content.length }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="w-16 bg-gray-200 rounded-full h-2 mr-2">
                    <div :class="getScoreClass(content.seoScore)" :style="{ width: content.seoScore + '%' }" class="h-2 rounded-full"></div>
                  </div>
                  <span class="text-sm text-gray-900">{{ content.seoScore }}%</span>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="getStatusClass(content.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                  {{ content.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 ml-3">مشاهده</button>
                <button class="text-green-600 hover:text-green-900">ویرایش</button>
                <button class="text-red-600 hover:text-red-900">حذف</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useAuth: () => { user: any; hasPermission: (perm: string) => boolean }
declare const $fetch: <T = any>(url: string, options?: { method?: string; body?: any; params?: any }) => Promise<T>
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()

// Lifecycle
onMounted(() => {
  fetchContents()
})

// داده‌های محتوا از API
const contents = ref([])

// دریافت محتوا از API
const fetchContents = async () => {
  try {
    const response = await $fetch('/api/posts?all=1') as { data?: any[] } | any[]
    const data = Array.isArray(response) ? response : (response?.data || [])
    if (Array.isArray(data) && data.length > 0) {
      contents.value = data.map((post: any) => ({
        id: post.id,
        title: post.title,
        keywords: post.meta_keywords || '',
        type: 'مقاله وبلاگ',
        length: `${post.content?.length || 0} کلمه`,
        seoScore: calculateSeoScore(post),
        status: getContentStatus(post.status)
      }))
    }
  } catch (error) {
    console.error('خطا در دریافت محتوا:', error)
    contents.value = []
  }
}

// محاسبه امتیاز SEO
const calculateSeoScore = (post: any) => {
  let score = 0
  
  // امتیاز برای عنوان
  if (post.title && post.title.length > 10 && post.title.length < 60) score += 20
  
  // امتیاز برای خلاصه
  if (post.excerpt && post.excerpt.length > 50) score += 15
  
  // امتیاز برای محتوا
  if (post.content && post.content.length > 300) score += 25
  
  // امتیاز برای کلمات کلیدی
  if (post.meta_keywords) score += 10
  
  // امتیاز برای توضیحات متا
  if (post.meta_description) score += 10
  
  // امتیاز برای تصویر شاخص
  if (post.featured_image) score += 10
  
  // امتیاز برای وضعیت منتشر شده
  if (post.status === 'published') score += 10
  
  return Math.min(score, 100)
}

// تعیین وضعیت محتوا
const getContentStatus = (status: string) => {
  switch (status) {
    case 'published': return 'تکمیل شده'
    case 'draft': return 'در انتظار بررسی'
    case 'scheduled': return 'زمان‌بندی شده'
    default: return 'در انتظار بررسی'
  }
}

// تابع تعیین کلاس نوع
const getTypeClass = (type: string) => {
  switch (type) {
    case 'مقاله وبلاگ':
      return 'bg-blue-100 text-blue-800'
    case 'توضیحات محصول':
      return 'bg-green-100 text-green-800'
    case 'صفحه لندینگ':
      return 'bg-purple-100 text-purple-800'
    case 'محتوا شبکه‌های اجتماعی':
      return 'bg-pink-100 text-pink-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// تابع تعیین کلاس امتیاز
const getScoreClass = (score: number) => {
  if (score >= 80) return 'bg-green-500'
  if (score >= 60) return 'bg-yellow-500'
  return 'bg-red-500'
}

// تابع تعیین کلاس وضعیت
const getStatusClass = (status: string) => {
  switch (status) {
    case 'تکمیل شده':
      return 'bg-green-100 text-green-800'
    case 'در انتظار بررسی':
      return 'bg-yellow-100 text-yellow-800'
    case 'نیازمند بازنویسی':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script> 

