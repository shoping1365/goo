<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="max-w-7xl mx-auto">
      <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
        <div class="px-6 py-4">
          <div class="flex items-center justify-between">
            <div>
              <h1 class="text-2xl font-bold text-gray-900">مدیریت اسکیما</h1>
              <p class="text-sm text-gray-500 mt-1">تنظیم و مدیریت اسکیماهای ساختاریافته</p>
            </div>
            
            <div class="flex items-center gap-3">
              <NuxtLink 
                to="/admin/seo/SchemaManagement/SchemaAdd"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              >
                <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                </svg>
                افزودن اسکیما جدید
              </NuxtLink>
              
              <NuxtLink 
                to="/admin"
                class="inline-flex items-center px-4 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md"
              >
                <svg class="w-5 h-5 ml-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
                </svg>
                <span class="text-gray-700">بازگشت</span>
              </NuxtLink>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content -->
      <div class="p-6">
        <!-- آمار کارت‌ها در بالای صفحه -->
        <div class="relative">
          <div class="relative grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6 max-w-5xl mx-auto">
            <div class="bg-gradient-to-br from-blue-50 to-blue-100 overflow-hidden shadow-md rounded-lg border border-blue-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
              <div class="p-3 flex items-center justify-between">
                <div class="flex-1 text-right">
                  <dl>
                    <dt class="text-xs font-medium text-blue-600 truncate">کل اسکیماها</dt>
                    <dd class="text-base font-bold text-blue-800">{{ schemas.length }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-blue-400 to-blue-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gradient-to-br from-green-50 to-emerald-100 overflow-hidden shadow-md rounded-lg border border-green-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
              <div class="p-3 flex items-center justify-between">
                <div class="flex-1 text-right">
                  <dl>
                    <dt class="text-xs font-medium text-green-600 truncate">اسکیماهای فعال</dt>
                    <dd class="text-base font-bold text-green-800">{{ activeSchemasCount }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-green-400 to-emerald-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gradient-to-br from-amber-50 to-yellow-100 overflow-hidden shadow-md rounded-lg border border-amber-200 transition-all duration-200 hover:shadow-lg hover:scale-105">
              <div class="p-3 flex items-center justify-between">
                <div class="flex-1 text-right">
                  <dl>
                    <dt class="text-xs font-medium text-amber-600 truncate">انواع مختلف</dt>
                    <dd class="text-base font-bold text-amber-800">{{ uniqueSchemaTypes.length }}</dd>
                  </dl>
                </div>
                <div class="flex-shrink-0 mr-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-amber-400 to-yellow-500 rounded-lg flex items-center justify-center shadow-md">
                    <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- لیست اسکیماهای موجود -->
          <div class="lg:col-span-3">
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <h2 class="text-lg font-semibold text-gray-900 mb-4">اسکیماهای موجود</h2>
              <div v-if="loading" class="flex justify-center items-center py-8">
                <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
                <span class="mr-2 text-gray-600">در حال بارگذاری...</span>
              </div>
              
              <div v-else-if="error" class="flex justify-center items-center py-8">
                <div class="text-red-600 text-center">
                  <svg class="w-12 h-12 mx-auto mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                  </svg>
                  <p class="text-lg font-medium">{{ error }}</p>
                </div>
              </div>
              
              <div v-else-if="schemas.length > 0" class="overflow-x-auto">
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">نوع</th>
                      <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">نام</th>
                      <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">توضیحات</th>
                      <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">وضعیت</th>
                      <th class="px-4 py-2 text-right text-xs font-medium text-gray-500 uppercase">عملیات</th>
                    </tr>
                  </thead>
                  <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-for="schema in schemas" :key="schema.id">
                      <td class="px-4 py-2 whitespace-nowrap">
                        <span class="text-sm font-medium text-gray-700">{{ getSchemaTypeLabel(schema.type) }}</span>
                        <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800 ml-2">
                          {{ schema.type }}
                        </span>
                      </td>
                      <td class="px-4 py-2 whitespace-nowrap">
                        <span class="text-sm text-gray-900 font-medium">{{ schema.name || 'بدون نام' }}</span>
                        <div class="text-xs text-gray-500 mt-1">ID: {{ schema.id }}</div>
                      </td>
                      <td class="px-4 py-2 whitespace-nowrap">
                        <span class="text-sm text-gray-900">{{ schema.description || 'بدون توضیحات' }}</span>
                      </td>
                      <td class="px-4 py-2 whitespace-nowrap">
                        <span :class="schema.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium">
                          {{ schema.isActive ? 'فعال' : 'غیرفعال' }}
                        </span>
                      </td>
                      <td class="px-4 py-2 whitespace-nowrap">
                        <div class="flex items-center gap-2">
                          <button class="text-blue-500 hover:text-blue-700 transition-colors text-sm" @click="viewSchema(schema)">
                            مشاهده
                          </button>
                          <button class="text-amber-500 hover:text-amber-700 transition-colors text-sm" @click="editSchema(schema)">
                            ویرایش
                          </button>
                          <button class="text-green-500 hover:text-green-700 transition-colors text-sm" @click="testSchema(schema)">
                            تست
                          </button>
                          <button class="text-red-500 hover:text-red-700 transition-colors text-sm" @click="deleteSchema(schema)">
                            حذف
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div v-else class="flex flex-col items-center justify-center py-16 text-gray-500">
                <svg class="w-16 h-16 mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                <p class="text-lg font-medium">هیچ اسکیمایی وجود ندارد</p>
                <p class="text-sm text-gray-400 mt-2">برای شروع، اسکیما جدید اضافه کنید</p>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const useHead: (head: { title?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useSchema: () => any
</script>

<script setup lang="ts">
import { computed, onMounted } from 'vue'

definePageMeta({
  layout: 'admin-main'
})

useHead({
  title: 'مدیریت اسکیما - پنل مدیریت'
})

// استفاده از composable
const { 
  schemas, 
  loading, 
  error, 
  fetchAllTemplates,
  activeTemplates
} = useSchema()

// Computed properties
const activeSchemasCount = computed(() => activeTemplates.value.length)

const uniqueSchemaTypes = computed(() => {
  const types = schemas.value.map(s => s.type)
  return [...new Set(types)]
})

// Methods
const getSchemaTypeLabel = (type: string): string => {
  const labels: Record<string, string> = {
    'Article': 'مقاله',
    'Product': 'محصول',
    'Organization': 'سازمان',
    'WebPage': 'صفحه وب',
    'BreadcrumbList': 'مسیر ناوبری',
    'FAQPage': 'سوالات متداول',
    'LocalBusiness': 'کسب و کار محلی',
    'Review': 'نظر'
  }
  return labels[type] || type
}

interface Schema {
  id?: number | string
  name?: string
  template?: Record<string, unknown>
  [key: string]: unknown
}

const viewSchema = (schema: Schema) => {
  // نمایش اسکیمای JSON در یک modal یا صفحه جدید
  const schemaJson = JSON.stringify(schema.template, null, 2)
  alert(`اسکیمای ${schema.name || 'نامشخص'}:\n\n${schemaJson}`)
}

const editSchema = (schema: Schema) => {
  // هدایت به صفحه ویرایش اسکیما
  navigateTo(`/admin/seo/SchemaManagement/SchemaEdit?id=${schema.id}`)
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const testSchema = async (schema: any) => {
  try {
    // تست اسکیما با استفاده از Google's Rich Results Test
    const schemaJson = JSON.stringify(schema.template, null, 2)
    const testUrl = `https://search.google.com/test/rich-results?url=${encodeURIComponent(window.location.origin)}&user_agent=DESKTOP`
    
    // باز کردن صفحه تست در تب جدید
    window.open(testUrl, '_blank')
    
    // نمایش پیام راهنما
    alert(`صفحه تست Google Rich Results در تب جدید باز شد.\n\nبرای تست اسکیمای ${schema.name}، کد JSON زیر را در قسمت "Code" کپی کنید:\n\n${schemaJson}`)
  } catch (err) {
    console.error('خطا در تست اسکیما:', err)
    alert('خطا در باز کردن صفحه تست')
  }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const deleteSchema = async (schema: any) => {
  // تأیید حذف اسکیما
  const confirmed = confirm(`آیا از حذف اسکیمای "${schema.name}" اطمینان دارید؟\n\nاین عملیات قابل بازگشت نیست.`)
  
  if (confirmed) {
    try {
      // حذف فایل اسکیما از سیستم فایل
      const response = await $fetch(`/api/schemas/${schema.id}`, {
        method: 'DELETE'
      }) as { success?: boolean }
      
      if (response.success) {
        alert('اسکیما با موفقیت حذف شد!')
        // بارگذاری مجدد لیست اسکیماها
        await fetchAllTemplates()
      } else {
        alert('خطا در حذف اسکیما')
      }
    } catch (err) {
      console.error('خطا در حذف اسکیما:', err)
      alert('خطا در حذف اسکیما')
    }
  }
}

// Lifecycle
onMounted(async () => {
  await fetchAllTemplates()
})
</script> 
