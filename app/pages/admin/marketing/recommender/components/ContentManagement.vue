<template>
  <div class="space-y-8">
    <!-- هدر بخش -->
    <div class="flex justify-between items-center">
      <div>
        <h2 class="text-2xl font-bold text-gray-900">مدیریت محتوا</h2>
        <p class="mt-1 text-sm text-gray-600">مدیریت محتوای پیشنهادات و تنظیمات نمایش</p>
      </div>
      <button class="bg-teal-600 text-white px-4 py-2 rounded-lg hover:bg-teal-700 transition-colors">
        <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        افزودن محتوای جدید
      </button>
    </div>

    <!-- آمار محتوا -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-gradient-to-r from-teal-500 to-teal-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">کل محتوا</p>
            <p class="text-2xl font-bold">1,247</p>
            <p class="text-xs opacity-75">+23 از هفته گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">محتوای فعال</p>
            <p class="text-2xl font-bold">1,089</p>
            <p class="text-xs opacity-75">87.3% از کل محتوا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-green-500 to-green-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">نرخ تعامل</p>
            <p class="text-2xl font-bold">34.2%</p>
            <p class="text-xs opacity-75">+5.8% از ماه گذشته</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
            </svg>
          </div>
        </div>
      </div>

      <div class="bg-gradient-to-r from-purple-500 to-purple-600 rounded-xl p-6 text-white">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm opacity-90">محتوای در انتظار</p>
            <p class="text-2xl font-bold">158</p>
            <p class="text-xs opacity-75">12.7% از کل محتوا</p>
          </div>
          <div class="p-3 bg-white bg-opacity-20 rounded-lg">
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- قالب‌های پیشنهادات -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-teal-100 rounded-lg">
            <svg class="w-6 h-6 text-teal-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">قالب‌های پیشنهادات</h3>
        </div>
        <button class="text-teal-600 hover:text-teal-800 text-sm font-medium">افزودن قالب جدید</button>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="template in contentTemplates" :key="template.id" class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow">
          <div class="flex items-center justify-between mb-3">
            <h4 class="text-sm font-medium text-gray-900">{{ template.name }}</h4>
            <label class="relative inline-flex items-center cursor-pointer">
              <input v-model="template.enabled" type="checkbox" class="sr-only peer">
              <div class="w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-teal-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-teal-600"></div>
            </label>
          </div>
          
          <div class="space-y-3">
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">نوع:</span>
              <span class="font-medium">{{ template.type }}</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">استفاده:</span>
              <span class="font-medium">{{ template.usage }} بار</span>
            </div>
            <div class="flex items-center justify-between text-xs">
              <span class="text-gray-500">نرخ کلیک:</span>
              <span class="font-medium">{{ template.clickRate }}%</span>
            </div>
          </div>
          
          <div class="mt-4 flex space-x-2 space-x-reverse">
            <button class="text-xs px-2 py-1 bg-teal-100 text-teal-700 rounded hover:bg-teal-200">ویرایش</button>
            <button class="text-xs px-2 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200">پیش‌نمایش</button>
            <button class="text-xs px-2 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200">حذف</button>
          </div>
        </div>
      </div>
    </div>

    <!-- مدیریت متن‌ها -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- متن‌های پیشنهادات -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">متن‌های پیشنهادات</h3>
          </div>
          <button class="text-blue-600 hover:text-blue-800 text-sm font-medium">افزودن متن جدید</button>
        </div>
        
        <div class="space-y-4">
          <div v-for="text in recommendationTexts" :key="text.id" class="border border-gray-200 rounded-lg p-6">
            <div class="flex items-center justify-between mb-3">
              <h4 class="text-sm font-medium text-gray-900">{{ text.title }}</h4>
              <label class="relative inline-flex items-center cursor-pointer">
                <input v-model="text.enabled" type="checkbox" class="sr-only peer">
                <div class="w-8 h-4 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-3 after:w-3 after:transition-all peer-checked:bg-blue-600"></div>
              </label>
            </div>
            
            <p class="text-sm text-gray-600 mb-3">{{ text.content }}</p>
            
            <div class="flex items-center justify-between text-xs text-gray-500">
              <span>نوع: {{ text.type }}</span>
              <span>استفاده: {{ text.usage }} بار</span>
            </div>
            
            <div class="mt-3 flex space-x-2 space-x-reverse">
              <button class="text-xs px-2 py-1 bg-blue-100 text-blue-700 rounded hover:bg-blue-200">ویرایش</button>
              <button class="text-xs px-2 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200">کپی</button>
            </div>
          </div>
        </div>
      </div>

      <!-- تصاویر و آیکون‌ها -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-3">تصاویر و آیکون‌ها</h3>
          </div>
          <button class="text-green-600 hover:text-green-800 text-sm font-medium">آپلود فایل</button>
        </div>
        
        <div class="grid grid-cols-2 gap-6">
          <div v-for="image in contentImages" :key="image.id" class="border border-gray-200 rounded-lg p-3">
            <div class="aspect-square bg-gray-100 rounded-lg mb-3 flex items-center justify-center">
              <img v-if="image.url" :src="image.url" :alt="image.name" class="w-full h-full object-cover rounded-lg">
              <svg v-else class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
            
            <div class="space-y-2">
              <h4 class="text-sm font-medium text-gray-900 truncate">{{ image.name }}</h4>
              <div class="flex items-center justify-between text-xs text-gray-500">
                <span>{{ image.type }}</span>
                <span>{{ image.size }}</span>
              </div>
              
              <div class="flex space-x-1 space-x-reverse">
                <button class="text-xs px-2 py-1 bg-green-100 text-green-700 rounded hover:bg-green-200">استفاده</button>
                <button class="text-xs px-2 py-1 bg-gray-100 text-gray-700 rounded hover:bg-gray-200">حذف</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- تنظیمات نمایش -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 mr-3">تنظیمات نمایش</h3>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input v-model="displaySettings.enabled" type="checkbox" class="sr-only peer">
          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:right-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
        </label>
      </div>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">اندازه کارت</label>
            <select v-model="displaySettings.cardSize" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="small">کوچک</option>
              <option value="medium">متوسط</option>
              <option value="large">بزرگ</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">تعداد ستون</label>
            <input v-model="displaySettings.columns" type="range" min="1" max="6" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>1</span>
              <span>{{ displaySettings.columns }}</span>
              <span>6</span>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">فاصله بین کارت‌ها</label>
            <input v-model="displaySettings.spacing" type="range" min="8" max="32" class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer">
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>8px</span>
              <span>{{ displaySettings.spacing }}px</span>
              <span>32px</span>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">رنگ پس‌زمینه</label>
            <div class="flex space-x-2 space-x-reverse">
              <div
v-for="color in displaySettings.backgroundColors" :key="color" 
                   class="w-8 h-8 rounded-lg cursor-pointer border-2" 
                   :class="displaySettings.backgroundColor === color ? 'border-gray-900' : 'border-gray-300'"
                   :style="{ backgroundColor: color }"
                   @click="displaySettings.backgroundColor = color">
              </div>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">سایه کارت</label>
            <select v-model="displaySettings.shadow" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-purple-500">
              <option value="none">بدون سایه</option>
              <option value="small">سایه کوچک</option>
              <option value="medium">سایه متوسط</option>
              <option value="large">سایه بزرگ</option>
            </select>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">انیمیشن</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="displaySettings.animations.hover" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">انیمیشن hover</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.animations.load" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">انیمیشن بارگذاری</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.animations.scroll" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">انیمیشن اسکرول</span>
              </label>
            </div>
          </div>
        </div>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">نمایش اطلاعات</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="displaySettings.showInfo.price" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">قیمت</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.showInfo.rating" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">امتیاز</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.showInfo.discount" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">تخفیف</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.showInfo.stock" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">موجودی</span>
              </label>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">دکمه‌های عملیات</label>
            <div class="space-y-2">
              <label class="flex items-center">
                <input v-model="displaySettings.buttons.addToCart" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">افزودن به سبد</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.buttons.wishlist" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">لیست علاقه</span>
              </label>
              <label class="flex items-center">
                <input v-model="displaySettings.buttons.quickView" type="checkbox" class="rounded border-gray-300 text-purple-600 focus:ring-purple-500">
                <span class="mr-2 text-sm text-gray-700">مشاهده سریع</span>
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- پیش‌نمایش -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
      <div class="flex items-center justify-between mb-6">
        <h3 class="text-lg font-semibold text-gray-900">پیش‌نمایش</h3>
        <div class="flex space-x-2 space-x-reverse">
          <button class="px-3 py-1 text-xs bg-blue-100 text-blue-800 rounded-lg">دسکتاپ</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">تبلت</button>
          <button class="px-3 py-1 text-xs bg-gray-100 text-gray-600 rounded-lg">موبایل</button>
        </div>
      </div>
      
      <div class="border border-gray-200 rounded-lg p-6 bg-gray-50">
        <div class="grid gap-6" :class="`grid-cols-${displaySettings.columns}`">
          <div v-for="i in 6" :key="i" class="bg-white rounded-lg p-6 shadow-sm">
            <div class="aspect-square bg-gray-200 rounded-lg mb-3"></div>
            <h4 class="text-sm font-medium text-gray-900 mb-2">محصول نمونه {{ i }}</h4>
            <div class="flex items-center justify-between text-sm">
              <span class="text-gray-600">قیمت: ۲۵۰,۰۰۰ تومان</span>
              <div class="flex items-center">
                <span class="text-yellow-500">★</span>
                <span class="text-gray-600 mr-1">4.5</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// قالب‌های محتوا
const contentTemplates = ref([
  {
    id: 1,
    name: 'قالب محصول ساده',
    type: 'محصول',
    enabled: true,
    usage: 1247,
    clickRate: 24.5
  },
  {
    id: 2,
    name: 'قالب محصول با تخفیف',
    type: 'محصول',
    enabled: true,
    usage: 892,
    clickRate: 31.2
  },
  {
    id: 3,
    name: 'قالب دسته‌بندی',
    type: 'دسته‌بندی',
    enabled: true,
    usage: 567,
    clickRate: 18.7
  },
  {
    id: 4,
    name: 'قالب برند',
    type: 'برند',
    enabled: false,
    usage: 234,
    clickRate: 15.3
  },
  {
    id: 5,
    name: 'قالب ویژه',
    type: 'ویژه',
    enabled: true,
    usage: 445,
    clickRate: 28.9
  },
  {
    id: 6,
    name: 'قالب جدید',
    type: 'جدید',
    enabled: true,
    usage: 123,
    clickRate: 22.1
  }
])

// متن‌های پیشنهادات
const recommendationTexts = ref([
  {
    id: 1,
    title: 'پیشنهاد ویژه',
    content: 'محصولات ویژه که ممکن است دوست داشته باشید',
    type: 'ویژه',
    enabled: true,
    usage: 1247
  },
  {
    id: 2,
    title: 'محصولات محبوب',
    content: 'محصولات پرطرفدار در دسته‌بندی مشابه',
    type: 'محبوب',
    enabled: true,
    usage: 892
  },
  {
    id: 3,
    title: 'پیشنهاد شخصی',
    content: 'بر اساس سلیقه شما پیشنهاد می‌شود',
    type: 'شخصی',
    enabled: true,
    usage: 567
  },
  {
    id: 4,
    title: 'محصولات جدید',
    content: 'جدیدترین محصولات اضافه شده',
    type: 'جدید',
    enabled: true,
    usage: 234
  }
])

// تصاویر محتوا
const contentImages = ref([
  {
    id: 1,
    name: 'آیکون پیشنهاد',
    type: 'PNG',
    size: '24KB',
    url: null
  },
  {
    id: 2,
    name: 'تصویر ویژه',
    type: 'JPG',
    size: '156KB',
    url: null
  },
  {
    id: 3,
    name: 'آیکون تخفیف',
    type: 'SVG',
    size: '8KB',
    url: null
  },
  {
    id: 4,
    name: 'تصویر جدید',
    type: 'PNG',
    size: '89KB',
    url: null
  }
])

// تنظیمات نمایش
const displaySettings = ref({
  enabled: true,
  cardSize: 'medium',
  columns: 3,
  spacing: 16,
  backgroundColor: '#ffffff',
  backgroundColors: ['#ffffff', '#f8fafc', '#f1f5f9', '#e2e8f0'],
  shadow: 'medium',
  animations: {
    hover: true,
    load: true,
    scroll: false
  },
  showInfo: {
    price: true,
    rating: true,
    discount: true,
    stock: false
  },
  buttons: {
    addToCart: true,
    wishlist: true,
    quickView: false
  }
})
</script> 
