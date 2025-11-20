<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200">
    <!-- هدر ساده و زیبا -->
    <div class="px-6 py-4 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center space-x-3 rtl:space-x-reverse">
          <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-semibold text-gray-900">ادغام با سیستم حسابداری</h3>
            <p class="text-sm text-gray-500">مدیریت اتصال و همگام‌سازی با سیستم حسابداری</p>
          </div>
        </div>
        <div class="flex items-center space-x-2 rtl:space-x-reverse">
          <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
            <span class="w-2 h-2 bg-green-400 rounded-full mr-1 rtl:ml-1 rtl:mr-0"></span>
            متصل
          </span>
        </div>
      </div>
    </div>

    <!-- تب‌های ساده -->
    <div class="border-b border-gray-200">
      <nav class="flex space-x-8 rtl:space-x-reverse px-6" aria-label="Tabs">
        <button
          v-for="tab in tabs"
          :key="tab.name"
          :class="[
            activeTab === tab.name
              ? 'border-blue-500 text-blue-600'
              : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm transition-colors duration-200'
          ]"
          @click="activeTab = tab.name"
        >
          {{ tab.label }}
        </button>
      </nav>
    </div>

    <!-- محتوای تب‌ها -->
    <div class="px-4 py-4">
      <!-- تب اتصالات -->
      <div v-if="activeTab === 'connections'" class="space-y-6">
        <!-- کارت وضعیت اتصال -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg px-4 py-4 border border-blue-100">
          <div class="flex items-center justify-between">
            <div>
              <h4 class="text-lg font-medium text-gray-900 mb-2">وضعیت اتصال</h4>
              <p class="text-gray-600">سیستم حسابداری شما با موفقیت متصل شده است</p>
            </div>
            <div class="text-right">
              <div class="text-2xl font-bold text-green-600">فعال</div>
              <div class="text-sm text-gray-500">آخرین همگام‌سازی: ۲ دقیقه پیش</div>
            </div>
          </div>
        </div>

        <!-- کارت‌های اتصال -->
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <div class="bg-white border border-gray-200 rounded-lg px-4 py-4 hover:shadow-md transition-shadow">
            <div class="flex items-center space-x-3 rtl:space-x-reverse mb-4">
              <div class="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
              <h5 class="font-medium text-gray-900">همگام‌سازی خودکار</h5>
            </div>
            <p class="text-gray-600 text-sm mb-4">سفارشات به صورت خودکار در سیستم حسابداری ثبت می‌شوند</p>
            <div class="flex items-center justify-between">
              <span class="text-sm text-green-600 font-medium">فعال</span>
              <button class="text-blue-600 hover:text-blue-700 text-sm font-medium">تنظیمات</button>
            </div>
          </div>

          <div class="bg-white border border-gray-200 rounded-lg px-4 py-4 hover:shadow-md transition-shadow">
            <div class="flex items-center space-x-3 rtl:space-x-reverse mb-4">
              <div class="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center">
                <svg class="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <h5 class="font-medium text-gray-900">تأیید خودکار</h5>
            </div>
            <p class="text-gray-600 text-sm mb-4">پرداخت‌ها به صورت خودکار تأیید و ثبت می‌شوند</p>
            <div class="flex items-center justify-between">
              <span class="text-sm text-green-600 font-medium">فعال</span>
              <button class="text-blue-600 hover:text-blue-700 text-sm font-medium">تنظیمات</button>
            </div>
          </div>
        </div>
      </div>

      <!-- تب تنظیمات -->
      <div v-if="activeTab === 'settings'" class="space-y-6">
        <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
          <h4 class="text-lg font-medium text-gray-900 mb-4">تنظیمات همگام‌سازی</h4>
          
          <div class="space-y-4">
            <div class="flex items-center justify-between py-3 border-b border-gray-100">
              <div>
                <h5 class="font-medium text-gray-900">همگام‌سازی خودکار</h5>
                <p class="text-sm text-gray-500">سفارشات جدید به صورت خودکار همگام شوند</p>
              </div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer" checked>
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>

            <div class="flex items-center justify-between py-3 border-b border-gray-100">
              <div>
                <h5 class="font-medium text-gray-900">اعلان‌های خطا</h5>
                <p class="text-sm text-gray-500">در صورت بروز خطا در همگام‌سازی اطلاع دهید</p>
              </div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer" checked>
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>

            <div class="flex items-center justify-between py-3">
              <div>
                <h5 class="font-medium text-gray-900">گزارش‌گیری خودکار</h5>
                <p class="text-sm text-gray-500">گزارش‌های مالی به صورت خودکار تولید شوند</p>
              </div>
              <label class="relative inline-flex items-center cursor-pointer">
                <input type="checkbox" class="sr-only peer">
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- تب گزارش‌ها -->
      <div v-if="activeTab === 'reports'" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gapx-4 py-4">
          <div class="bg-white border border-gray-200 rounded-lg px-4 py-4 text-center">
            <div class="text-2xl font-bold text-blue-600 mb-2">۱۲۵</div>
            <div class="text-sm text-gray-600">سفارش همگام شده</div>
          </div>
          <div class="bg-white border border-gray-200 rounded-lg px-4 py-4 text-center">
            <div class="text-2xl font-bold text-green-600 mb-2">۹۸%</div>
            <div class="text-sm text-gray-600">نرخ موفقیت</div>
          </div>
          <div class="bg-white border border-gray-200 rounded-lg px-4 py-4 text-center">
            <div class="text-2xl font-bold text-orange-600 mb-2">۲</div>
            <div class="text-sm text-gray-600">خطاهای اخیر</div>
          </div>
        </div>

        <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
          <h4 class="text-lg font-medium text-gray-900 mb-4">آخرین فاکتورهای ثبت شده</h4>
          <div class="space-y-3">
            <div class="flex items-center justify-between py-2 border-b border-gray-100">
              <div>
                <div class="font-medium text-gray-900">فاکتور #۱۲۳۴۵</div>
                <div class="text-sm text-gray-500">سفارش #۶۷۸۹۰</div>
              </div>
              <div class="text-right">
                <div class="font-medium text-gray-900">۲,۵۰۰,۰۰۰ تومان</div>
                <div class="text-sm text-green-600">ثبت شده</div>
              </div>
            </div>
            <div class="flex items-center justify-between py-2 border-b border-gray-100">
              <div>
                <div class="font-medium text-gray-900">فاکتور #۱۲۳۴۴</div>
                <div class="text-sm text-gray-500">سفارش #۶۷۸۸۹</div>
              </div>
              <div class="text-right">
                <div class="font-medium text-gray-900">۱,۸۰۰,۰۰۰ تومان</div>
                <div class="text-sm text-green-600">ثبت شده</div>
              </div>
            </div>
            <div class="flex items-center justify-between py-2">
              <div>
                <div class="font-medium text-gray-900">فاکتور #۱۲۳۴۳</div>
                <div class="text-sm text-gray-500">سفارش #۶۷۸۸۸</div>
              </div>
              <div class="text-right">
                <div class="font-medium text-gray-900">۳,۲۰۰,۰۰۰ تومان</div>
                <div class="text-sm text-green-600">ثبت شده</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- تب لاگ‌ها -->
      <div v-if="activeTab === 'logs'" class="space-y-6">
        <div class="bg-white border border-gray-200 rounded-lg px-4 py-4">
          <h4 class="text-lg font-medium text-gray-900 mb-4">لاگ‌های سیستم</h4>
          <div class="space-y-3">
            <div class="flex items-center space-x-3 rtl:space-x-reverse p-3 bg-green-50 rounded-lg">
              <div class="w-2 h-2 bg-green-500 rounded-full"></div>
              <div class="flex-1">
                <div class="text-sm font-medium text-gray-900">همگام‌سازی موفق</div>
                <div class="text-xs text-gray-500">سفارش #۶۷۸۹۰ با موفقیت در سیستم حسابداری ثبت شد</div>
              </div>
              <div class="text-xs text-gray-400">۲ دقیقه پیش</div>
            </div>
            <div class="flex items-center space-x-3 rtl:space-x-reverse p-3 bg-blue-50 rounded-lg">
              <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
              <div class="flex-1">
                <div class="text-sm font-medium text-gray-900">اتصال برقرار شد</div>
                <div class="text-xs text-gray-500">اتصال با سیستم حسابداری با موفقیت برقرار شد</div>
              </div>
              <div class="text-xs text-gray-400">۵ دقیقه پیش</div>
            </div>
            <div class="flex items-center space-x-3 rtl:space-x-reverse p-3 bg-yellow-50 rounded-lg">
              <div class="w-2 h-2 bg-yellow-500 rounded-full"></div>
              <div class="flex-1">
                <div class="text-sm font-medium text-gray-900">هشدار همگام‌سازی</div>
                <div class="text-xs text-gray-500">تأخیر در همگام‌سازی سفارش #۶۷۸۸۸</div>
              </div>
              <div class="text-xs text-gray-400">۱۰ دقیقه پیش</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// تعریف تب‌ها
const tabs = [
  { name: 'connections', label: 'اتصالات' },
  { name: 'settings', label: 'تنظیمات' },
  { name: 'reports', label: 'گزارش‌ها' },
  { name: 'logs', label: 'لاگ‌ها' }
]

const activeTab = ref('connections')
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 
