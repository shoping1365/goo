<template>
  <!-- هدر صفحه -->
  <div class="header-bg-full mb-6">
    <div class="page-header-flex">
      <div>
        <h1 class="page-title">ایجاد ناوبری جدید</h1>
        <p class="page-subtitle">ایجاد ناوبری جدید برای موبایل و اپلیکیشن</p>
      </div>
      <NuxtLink
        to="/admin/content/mobile-app-navigation-management"
        class="btn btn-secondary"
      >
        <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
        </svg>
        بازگشت
      </NuxtLink>
    </div>
  </div>

  <!-- محتوای اصلی -->
  <div class="p-6 w-full">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- فرم ایجاد ناوبری (سمت چپ) -->
      <div>
        <div class="bg-white rounded-lg shadow p-6 h-full">
          <form class="space-y-6" @submit.prevent="createNavigation">
            <!-- اطلاعات اصلی -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- نام ناوبری -->
              <div>
                <label for="name" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                  نام ناوبری <span class="text-red-500">*</span>
                </label>
                <input
                  id="name"
                  v-model="formData.name"
                  type="text"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                  placeholder="نام ناوبری را وارد کنید"
                />
              </div>

              <!-- پلتفرم -->
              <div>
                <label for="platform" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                  پلتفرم <span class="text-red-500">*</span>
                </label>
                <select
                  id="platform"
                  v-model="formData.platform"
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                >
                  <option value="mobile">موبایل</option>
                  <option value="app">اپلیکیشن</option>
                  <option value="both">هر دو</option>
                </select>
              </div>
            </div>

            <!-- توضیحات -->
            <div>
              <label for="description" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                توضیحات
              </label>
              <textarea
                id="description"
                v-model="formData.description"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                placeholder="توضیحات ناوبری را وارد کنید"
              ></textarea>
            </div>

            <!-- تنظیمات رنگ -->
            <div class="border-t pt-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">تنظیمات رنگ</h3>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- رنگ پس‌زمینه -->
                <div>
                  <label for="background_color" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                    رنگ پس‌زمینه
                  </label>
                  <div class="flex gap-2">
                    <input
                      id="background_color"
                      v-model="formData.background_color"
                      type="color"
                      class="w-12 h-10 border border-gray-300 rounded-md cursor-pointer"
                    />
                    <input
                      v-model="formData.background_color"
                      type="text"
                      placeholder="#f8fafc"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                    />
                  </div>
                </div>

                <!-- رنگ متن -->
                <div>
                  <label for="text_color" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                    رنگ متن
                  </label>
                  <div class="flex gap-2">
                    <input
                      id="text_color"
                      v-model="formData.text_color"
                      type="color"
                      class="w-12 h-10 border border-gray-300 rounded-md cursor-pointer"
                    />
                    <input
                      v-model="formData.text_color"
                      type="text"
                      placeholder="#000000"
                      class="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                    />
                  </div>
                </div>
              </div>
            </div>

            <!-- تنظیمات نمایش -->
            <div class="border-t pt-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4 text-right">تنظیمات نمایش</h3>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- انتخاب صفحات -->
                <div>
                  <label for="page_selection" class="block text-sm font-medium text-gray-700 mb-2 text-right">
                    انتخاب صفحات
                  </label>
                  <select
                    id="page_selection"
                    v-model="formData.page_selection"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-right"
                  >
                    <option value="all">همه صفحات</option>
                    <option value="specific">صفحات خاص</option>
                    <option value="excluded">به جز صفحات خاص</option>
                  </select>
                </div>

                <!-- وضعیت فعال -->
                <div class="flex items-center justify-end">
                  <input
                    id="is_active"
                    v-model="formData.is_active"
                    type="checkbox"
                    class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  />
                  <label for="is_active" class="mr-2 block text-sm text-gray-900">
                    فعال
                  </label>
                </div>
              </div>

              <!-- صفحات خاص -->
              <div v-if="formData.page_selection === 'specific'" class="mt-4">
                <label for="specific_pages" class="block text-sm font-medium text-gray-700 mb-2">
                  صفحات خاص (هر خط یک صفحه)
                </label>
                <textarea
                  id="specific_pages"
                  v-model="formData.specific_pages"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="/home&#10;/products&#10;/about"
                ></textarea>
              </div>

              <!-- صفحات مستثنی -->
              <div v-if="formData.page_selection === 'excluded'" class="mt-4">
                <label for="excluded_pages" class="block text-sm font-medium text-gray-700 mb-2">
                  صفحات مستثنی (هر خط یک صفحه)
                </label>
                <textarea
                  id="excluded_pages"
                  v-model="formData.excluded_pages"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="/admin&#10;/login&#10;/register"
                ></textarea>
              </div>
            </div>

            <!-- انتخاب آیتم‌های ناوبری -->
            <NavigationItemSelector v-model="formData.navigation_items" />

            <!-- دکمه‌های عملیات -->
            <div class="flex justify-end space-x-4 space-x-reverse pt-6 border-t">
              <NuxtLink
                to="/admin/content/mobile-app-navigation-management"
                class="px-6 py-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition-colors"
              >
                انصراف
              </NuxtLink>
              <button
                type="submit"
                :disabled="loading"
                class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                <span v-if="loading">در حال ایجاد...</span>
                <span v-else>ایجاد ناوبری</span>
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- پیش‌نمایش ناوبری (سمت راست) -->
      <div>
        <MobileNavigationPreview :navigation-data="formData" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import NavigationItemSelector from '~/components/admin/NavigationItemSelector.vue'
import MobileNavigationPreview from '~/components/widgets/MobileNavigationPreview.vue'

// Meta
definePageMeta({
  layout: 'admin-main'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const loading = ref(false)
const formData = reactive({
  name: '',
  description: '',
  platform: 'mobile',
  background_color: '#f8fafc',
  text_color: '#000000',
  page_selection: 'all',
  specific_pages: '',
  excluded_pages: '',
  is_active: true,
  navigation_items: ['home', 'categories', 'cart']
})

// Methods
const createNavigation = async () => {
  try {
    loading.value = true

    const payload = {
      ...formData,
      specific_pages: formData.page_selection === 'specific' ? formData.specific_pages : '',
      excluded_pages: formData.page_selection === 'excluded' ? formData.excluded_pages : ''
    }

    const response = await fetch('/api/admin/mobile-app-navigation-settings/', {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const responseData = await response.json()

    if (responseData.success) {
      // نمایش پیام موفقیت
      await navigateTo('/admin/content/mobile-app-navigation-management')
    }
  } catch (error) {
    console.error('خطا در ایجاد ناوبری:', error)
    // نمایش پیام خطا
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.header-bg-full {
  background: white;
  color: #374151;
  padding: 12px 20px;
  border-radius: 0;
  margin-bottom: 15px;
  border: 1px solid #e5e7eb;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  width: 100%;
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 1.25rem;
  font-weight: bold;
  margin: 0;
  color: #111827;
}

.page-subtitle {
  color: #6b7280;
  margin: 2px 0 0 0;
  font-size: 0.8rem;
}

.btn {
  padding: 10px 20px;
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.3s ease;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
  text-decoration: none;
  font-size: 0.875rem;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}
</style>

