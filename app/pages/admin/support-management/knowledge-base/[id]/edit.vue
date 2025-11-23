<template>
  <div class="p-6" dir="rtl">
    <!-- دکمه بازگشت -->
    <div class="flex justify-end mb-6">
      <button class="bg-purple-600 text-white px-4 py-2 rounded-lg hover:bg-purple-700 transition-colors flex items-center space-x-2 space-x-reverse" @click="$router.back()">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
        </svg>
        <span>بازگشت</span>
      </button>
    </div>

    <!-- فرم اصلی -->
    <div class="space-y-6">
      <!-- بخش جزئیات سوال -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- عنوان سوال -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">عنوان سوال</label>
            <input 
              v-model="questionData.title" 
              type="text" 
              placeholder="قیمت و موجودی" 
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
            >
          </div>
          
          <!-- وضعیت -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
            <label class="relative inline-flex items-center cursor-pointer">
              <input 
                v-model="questionData.isActive" 
                type="checkbox" 
                class="sr-only peer"
              >
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-purple-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-purple-600"></div>
              <span class="mr-3 text-sm font-medium text-gray-900">{{ questionData.isActive ? 'فعال' : 'غیرفعال' }}</span>
            </label>
          </div>
        </div>
        
        <!-- عنوان سایت -->
        <div class="mt-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">عنوان سوال در سایت</label>
          <input 
            v-model="questionData.siteTitle" 
            type="text" 
            placeholder="قیمت و موجودی سایت" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
          >
        </div>
        
        <!-- دکمه ویرایش جزئیات -->
        <div class="flex justify-end mt-6">
          <button class="bg-purple-600 text-white px-6 py-2 rounded-lg hover:bg-purple-700 transition-colors" @click="saveBasicInfo">
            ویرایش
          </button>
        </div>
      </div>

      <!-- بخش سوالات -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">سوالات</h3>
        
        <!-- فیلد سوال -->
        <div class="mb-6">
          <textarea 
            v-model="questionData.question" 
            rows="3" 
            placeholder="آیا قیمت و موجودی سایت به روز است ؟" 
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500"
          ></textarea>
        </div>
      </div>

      <!-- بخش متن پیام -->
      <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">متن پیام</h3>
        
        <!-- ویرایشگر متن غنی -->
        <div class="border border-gray-300 rounded-md">
          <!-- نوار ابزار -->
          <div class="bg-gray-50 px-4 py-2 border-b border-gray-300 flex items-center space-x-2 space-x-reverse flex-wrap">
            <!-- استایل متن -->
            <select class="px-2 py-1 border border-gray-300 rounded text-sm">
              <option>عادی</option>
              <option>عنوان 1</option>
              <option>عنوان 2</option>
              <option>عنوان 3</option>
            </select>
            
            <!-- اندازه فونت -->
            <div class="flex items-center space-x-1 space-x-reverse">
              <button class="p-1 hover:bg-gray-200 rounded">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
                </svg>
              </button>
              <button class="p-1 hover:bg-gray-200 rounded">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"></path>
                </svg>
              </button>
            </div>
            
            <!-- تراز متن -->
            <div class="flex items-center space-x-1 space-x-reverse">
              <button class="p-1 hover:bg-gray-200 rounded" title="تراز راست">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                </svg>
              </button>
              <button class="p-1 hover:bg-gray-200 rounded" title="تراز وسط">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8M4 18h16"></path>
                </svg>
              </button>
              <button class="p-1 hover:bg-gray-200 rounded" title="تراز چپ">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                </svg>
              </button>
              <button class="p-1 hover:bg-gray-200 rounded" title="تراز کامل">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                </svg>
              </button>
            </div>
            
            <!-- لیست -->
            <div class="flex items-center space-x-1 space-x-reverse">
              <button class="p-1 hover:bg-gray-200 rounded" title="لیست نقطه‌ای">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
                </svg>
              </button>
              <button class="p-1 hover:bg-gray-200 rounded" title="لیست شماره‌ای">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M6 15h14"></path>
                </svg>
              </button>
            </div>
            
            <!-- پاک کردن فرمت -->
            <button class="p-1 hover:bg-gray-200 rounded" title="پاک کردن فرمت">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
            
            <!-- استایل‌های متن -->
            <div class="flex items-center space-x-1 space-x-reverse">
              <button class="p-1 hover:bg-gray-200 rounded font-bold" title="ضخیم">G</button>
              <button class="p-1 hover:bg-gray-200 rounded italic" title="ایتالیک">I</button>
              <button class="p-1 hover:bg-gray-200 rounded underline" title="زیرخط">U</button>
              <button class="p-1 hover:bg-gray-200 rounded line-through" title="خط خورده">B</button>
            </div>
            
            <!-- نقل قول -->
            <button class="p-1 hover:bg-gray-200 rounded" title="نقل قول">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
            </button>
            
            <!-- لینک -->
            <button class="p-1 hover:bg-gray-200 rounded" title="درج لینک">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
              </svg>
            </button>
            
            <!-- تصویر -->
            <button class="p-1 hover:bg-gray-200 rounded" title="درج تصویر">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </button>
            
            <!-- ویدیو -->
            <button class="p-1 hover:bg-gray-200 rounded" title="درج ویدیو">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
            </button>
            
            <!-- جدول -->
            <button class="p-1 hover:bg-gray-200 rounded" title="درج جدول">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M3 14h18m-9-4v8m-7 0h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
              </svg>
            </button>
          </div>
          
          <!-- ناحیه ویرایش -->
          <div class="p-6">
            <textarea 
              v-model="questionData.answer" 
              rows="8" 
              placeholder="بله. قیمت و موجودی با ارسال فوری در سایت به روز است و می توانید با خیال راحت سفارش دهید." 
              class="w-full border-none outline-none resize-none"
            ></textarea>
          </div>
        </div>
        
        <!-- دکمه‌های عملیات -->
        <div class="flex justify-end space-x-4 space-x-reverse mt-6">
          <button class="bg-purple-600 text-white px-6 py-2 rounded-lg hover:bg-purple-700 transition-colors" @click="saveAnswer">
            ویرایش
          </button>
        </div>
      </div>
    </div>

    <!-- دکمه افزودن در پایین -->
    <div class="fixed bottom-6 left-6">
      <button class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700 transition-colors flex items-center space-x-2 space-x-reverse shadow-lg" @click="addNewQuestion">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
        </svg>
        <span>افزودن</span>
      </button>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// دریافت پارامترهای URL
const route = useRoute()
const isEdit = computed(() => route.params.id !== 'new')

// داده‌های سوال
const questionData = ref({
  title: '',
  siteTitle: '',
  isActive: true,
  question: '',
  answer: ''
})

// بارگذاری داده‌های موجود در صورت ویرایش
onMounted(() => {
  if (isEdit.value) {
    // در اینجا می‌توانید داده‌های سوال را از API بارگذاری کنید
    questionData.value = {
      title: 'قیمت و موجودی',
      siteTitle: 'قیمت و موجودی سایت',
      isActive: true,
      question: 'آیا قیمت و موجودی سایت به روز است ؟',
      answer: 'بله. قیمت و موجودی با ارسال فوری در سایت به روز است و می توانید با خیال راحت سفارش دهید.'
    }
  }
})

// تابع ذخیره اطلاعات پایه
function saveBasicInfo() {
  // نمایش پیام موفقیت
  alert('اطلاعات پایه با موفقیت ذخیره شد!')
}

// تابع ذخیره پاسخ
function saveAnswer() {
  
  // نمایش پیام موفقیت
  alert('پاسخ با موفقیت ذخیره شد!')
}

// تابع افزودن سوال جدید
function addNewQuestion() {
  // پاک کردن فرم
  questionData.value = {
    title: '',
    siteTitle: '',
    isActive: true,
    question: '',
    answer: ''
  }
  
  // تغییر مسیر به صفحه جدید
  navigateTo('/admin/support-management/knowledge-base/edit/new')
}
</script> 