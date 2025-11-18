<template>
  <div class="p-6" dir="rtl">
    <!-- هدر بنفش -->
    <div class="bg-gradient-to-r from-purple-600 to-purple-700 text-white p-6 rounded-lg mb-6">
      <h1 class="text-2xl font-bold mb-2">پایگاه دانش</h1>
      <p class="text-purple-100">مدیریت سوالات متداول و شخصی‌سازی پایگاه دانش</p>
    </div>

    <!-- تب‌ها -->
    <div class="bg-white rounded-lg shadow-md border border-gray-200 mb-6">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 space-x-reverse px-6">
          <button 
            @click="activeTab = 'faq'"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
              activeTab === 'faq' 
                ? 'border-purple-500 text-purple-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            سوالات متداول
          </button>
          <button 
            @click="activeTab = 'customization'"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
              activeTab === 'customization' 
                ? 'border-purple-500 text-purple-600' 
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            شخصی‌سازی
          </button>
        </nav>
      </div>

      <!-- محتوای تب سوالات متداول -->
      <div v-if="activeTab === 'faq'" class="p-6">
        <!-- دکمه افزودن -->
        <div class="flex justify-end mb-6">
          <button @click="addNewFaq" class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700 transition-colors flex items-center space-x-2 space-x-reverse">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
            </svg>
            <span>افزودن</span>
          </button>
        </div>

        <!-- لیست سوالات متداول -->
        <div class="space-y-4">
          <div v-for="item in faqItems" :key="item.id" class="bg-gray-50 rounded-lg p-6 flex items-center justify-between">
            <div class="flex items-center space-x-4 space-x-reverse">
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
              </svg>
              <label class="relative inline-flex items-center cursor-pointer">
                <input 
                  type="checkbox" 
                  :checked="item.isActive" 
                  @change="toggleFaqStatus(item.id)"
                  class="sr-only peer"
                >
                <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
              </label>
                          <span class="text-gray-900">{{ item.title }}</span>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <button @click="editFaqItem(item.id)" class="text-blue-600 hover:text-blue-900 transition-colors" title="ویرایش">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </button>
            <button @click="deleteFaqItem(item.id)" class="text-red-500 hover:text-red-700" title="حذف">
              <span class="text-xl">🗑️</span>
            </button>
          </div>
          </div>
        </div>
      </div>

      <!-- محتوای تب شخصی‌سازی -->
      <div v-if="activeTab === 'customization'" class="p-6">
        <div class="space-y-8">
          <!-- انتخاب رنگ -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">رنگ اصلی پایگاه دانش خود را انتخاب کنید.</h3>
            <div class="flex space-x-4 space-x-reverse">
              <button 
                v-for="color in colors" 
                :key="color.name"
                @click="selectedColor = color.value"
                :class="[
                  'w-12 h-12 rounded-full border-2 transition-all',
                  selectedColor === color.value ? 'border-gray-400 scale-110' : 'border-gray-200 hover:border-gray-300'
                ]"
                :style="{ backgroundColor: color.value }"
                :title="color.name"
              ></button>
            </div>
          </div>

          <!-- لوگو پایگاه دانش -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">لوگو پایگاه دانش</h3>
            <div class="flex items-center space-x-4 space-x-reverse">
              <div class="w-16 h-16 bg-purple-100 rounded-full flex items-center justify-center">
                <svg class="w-8 h-8 text-purple-600" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>
              </div>
              <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors flex items-center space-x-2 space-x-reverse">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                <span>آپلود تصویر</span>
              </button>
            </div>
            <p class="text-sm text-gray-500 mt-2">ابعاد پیشنهادی لوگو ۵۵*۵۵ پیکسل و در فرمت png باشد.</p>
          </div>

          <!-- عنوان پایگاه دانش -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">عنوان پایگاه دانش</h3>
            <input 
              v-model="settings.title" 
              type="text" 
              placeholder="پایگاه دانش شما" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            >
          </div>

          <!-- توضیحات -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">توضیحات</h3>
            <textarea 
              v-model="settings.description" 
              rows="4" 
              placeholder="توضیحات را وارد نمایید" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            ></textarea>
          </div>

          <!-- نام دامنه -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">نام دامنه</h3>
            <input 
              v-model="settings.domain" 
              type="text" 
              placeholder="نام دامنه را وارد نمایید" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            >
          </div>

          <!-- آدرس وبسایت -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">آدرس وبسایت</h3>
            <input 
              v-model="settings.website" 
              type="url" 
              placeholder="https://" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            >
          </div>

          <!-- متن جستجو -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">متن جستجو</h3>
            <input 
              v-model="settings.searchText" 
              type="text" 
              placeholder="متن جستجو را وارد نمایید" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            >
          </div>

          <!-- شخصی‌سازی متون -->
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-4">شخصی‌سازی متون</h3>
            <div class="space-y-4">
              <div>
                <h4 class="text-sm font-medium text-gray-700 mb-2">عنوان سر صفحه</h4>
                <input 
                  v-model="settings.headerTitle" 
                  type="text" 
                  placeholder="عنوان سر صفحه را وارد نمایید" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
                >
              </div>
            </div>
          </div>

          <!-- دکمه ذخیره -->
          <div class="flex justify-end pt-6">
            <button @click="saveSettings" class="bg-purple-600 text-white px-8 py-3 rounded-lg hover:bg-purple-700 transition-colors">
              ذخیره
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({ layout: 'admin-main' })

// تب فعال
const activeTab = ref('faq')

// داده‌های نمونه برای سوالات متداول
const faqItems = ref([
  {
    id: 1,
    title: 'روش های ارسال',
    isActive: true
  },
  {
    id: 2,
    title: 'قیمت و موجودی',
    isActive: true
  },
  {
    id: 3,
    title: 'راه های ار تباط با کارشناس فروش',
    isActive: true
  },
  {
    id: 4,
    title: 'آدرس فروشگاه حضوری و ساعت کاری',
    isActive: true
  }
])

// رنگ‌های موجود
const colors = ref([
  { name: 'پیش‌فرض', value: '#8B5CF6' },
  { name: 'قرمز', value: '#EF4444' },
  { name: 'صورتی', value: '#EC4899' },
  { name: 'بنفش', value: '#8B5CF6' },
  { name: 'آبی', value: '#3B82F6' },
  { name: 'بنفش تیره', value: '#7C3AED' }
])

// رنگ انتخاب شده
const selectedColor = ref('#8B5CF6')

// تنظیمات شخصی‌سازی
const settings = ref({
  title: '',
  description: '',
  domain: '',
  website: 'https://',
  searchText: '',
  headerTitle: ''
})

// تابع حذف سوال متداول
function deleteFaqItem(id) {
  const index = faqItems.value.findIndex(item => item.id === id)
  if (index > -1) {
    faqItems.value.splice(index, 1)
  }
}

// تابع تغییر وضعیت فعال/غیرفعال
function toggleFaqStatus(id) {
  const item = faqItems.value.find(item => item.id === id)
  if (item) {
    item.isActive = !item.isActive
  }
}

// تابع ذخیره تنظیمات
function saveSettings() {
  // اینجا می‌توانید کد ذخیره تنظیمات را اضافه کنید
  console.log('تنظیمات ذخیره شد:', {
    color: selectedColor.value,
    settings: settings.value
  })
  
  // نمایش پیام موفقیت
  alert('تنظیمات با موفقیت ذخیره شد!')
}

// تابع ویرایش سوال متداول
function editFaqItem(id) {
  navigateTo(`/admin/support-management/knowledge-base/edit/${id}`)
}

// تابع افزودن سوال جدید
function addNewFaq() {
  navigateTo('/admin/support-management/knowledge-base/edit/new')
}
</script> 
