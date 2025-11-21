<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="flex items-center justify-between">
      <div>
        <h4 class="text-lg font-medium text-gray-900">سوالات متداول</h4>
        <p class="text-sm text-gray-600 mt-1">پاسخ سوالات رایج کاربران در مورد اتصال حسابداری</p>
      </div>
      <div class="flex items-center space-x-3 space-x-reverse">
        <button
          class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="suggestQuestion"
        >
          <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
          </svg>
          پیشنهاد سوال
        </button>
      </div>
    </div>

    <!-- جستجو -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <div class="relative">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="جستجو در سوالات متداول..."
          class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
        />
        <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
          <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
      </div>
    </div>

    <!-- دسته‌بندی‌ها -->
    <div class="bg-white rounded-lg border border-gray-200 p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">دسته‌بندی‌ها</h5>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="category in categories"
          :key="category.id"
          class="px-4 py-2 text-sm font-medium rounded-full transition-colors"
          :class="selectedCategory === category.id 
            ? 'bg-blue-100 text-blue-700 border border-blue-200' 
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200 border border-gray-200'"
          @click="selectCategory(category.id)"
        >
          {{ category.name }}
          <span class="text-xs text-gray-500 mr-1">({{ category.count }})</span>
        </button>
      </div>
    </div>

    <!-- سوالات متداول -->
    <div class="space-y-4">
      <div
        v-for="faq in filteredFAQs"
        :key="faq.id"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden"
      >
        <button
          class="w-full px-6 py-4 text-right flex items-center justify-between hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
          @click="toggleFAQ(faq.id)"
        >
          <div class="flex items-center space-x-3 space-x-reverse">
            <span
              class="px-2 py-1 text-xs font-medium rounded-full"
              :class="getCategoryClass(faq.category)"
            >
              {{ getCategoryLabel(faq.category) }}
            </span>
            <h6 class="text-sm font-medium text-gray-900">{{ faq.question }}</h6>
          </div>
          <div class="flex items-center space-x-2 space-x-reverse">
            <div class="flex items-center space-x-1 space-x-reverse text-xs text-gray-500">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15l7-7 7 7" />
              </svg>
              <span>{{ faq.views }} بازدید</span>
            </div>
            <svg
              class="w-5 h-5 text-gray-400 transition-transform"
              :class="{ 'rotate-180': openFAQs.includes(faq.id) }"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </button>
        
        <div
          v-show="openFAQs.includes(faq.id)"
          class="px-6 pb-4 border-t border-gray-100"
        >
          <div class="pt-4">
            <p class="text-sm text-gray-700 leading-relaxed">{{ faq.answer }}</p>
            
            <!-- راه‌حل‌های اضافی -->
            <div v-if="faq.solutions && faq.solutions.length > 0" class="mt-4">
              <h6 class="text-sm font-medium text-gray-900 mb-2">راه‌حل‌های اضافی:</h6>
              <ul class="list-disc pr-6 text-sm text-gray-600 space-y-1">
                <li v-for="(solution, index) in faq.solutions" :key="index">
                  {{ solution }}
                </li>
              </ul>
            </div>

            <!-- لینک‌های مفید -->
            <div v-if="faq.links && faq.links.length > 0" class="mt-4">
              <h6 class="text-sm font-medium text-gray-900 mb-2">لینک‌های مفید:</h6>
              <div class="space-y-2">
                <a
                  v-for="link in faq.links"
                  :key="link.url"
                  :href="link.url"
                  class="block text-sm text-blue-600 hover:text-blue-800 hover:underline"
                >
                  {{ link.title }}
                </a>
              </div>
            </div>

            <!-- رای‌دهی -->
            <div class="mt-4 pt-4 border-t border-gray-100">
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-4 space-x-reverse">
                  <span class="text-sm text-gray-600">آیا این پاسخ مفید بود؟</span>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button
                      class="inline-flex items-center px-3 py-1 text-xs font-medium rounded-md text-green-700 bg-green-100 hover:bg-green-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                      @click="rateFAQ(faq.id, 'helpful')"
                    >
                      <svg class="w-3 h-3 ml-1" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M14 9V5a3 3 0 00-6 0v4H4a2 2 0 00-2 2v8a2 2 0 002 2h12a2 2 0 002-2v-8a2 2 0 00-2-2h-2z"/>
                      </svg>
                      مفید ({{ faq.helpfulVotes }})
                    </button>
                    <button
                      class="inline-flex items-center px-3 py-1 text-xs font-medium rounded-md text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                      @click="rateFAQ(faq.id, 'not_helpful')"
                    >
                      <svg class="w-3 h-3 ml-1" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M10 15v4a3 3 0 006 0v-4h2a2 2 0 002-2V9a2 2 0 00-2-2h-2V3a3 3 0 00-6 0v4H8a2 2 0 00-2 2v4a2 2 0 002 2h2z"/>
                      </svg>
                      غیرمفید ({{ faq.notHelpfulVotes }})
                    </button>
                  </div>
                </div>
                <button
                  class="text-xs text-gray-500 hover:text-gray-700"
                  @click="reportFAQ(faq.id)"
                >
                  گزارش مشکل
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- آمار -->
    <div class="bg-gray-50 rounded-lg p-6">
      <h5 class="text-md font-medium text-gray-900 mb-4">آمار سوالات متداول</h5>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <div class="text-center">
          <div class="text-2xl font-bold text-blue-600">{{ totalFAQs }}</div>
          <div class="text-sm text-gray-600">کل سوالات</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-green-600">{{ totalViews }}</div>
          <div class="text-sm text-gray-600">کل بازدیدها</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ totalHelpfulVotes }}</div>
          <div class="text-sm text-gray-600">رای‌های مفید</div>
        </div>
        <div class="text-center">
          <div class="text-2xl font-bold text-purple-600">{{ categories.length }}</div>
          <div class="text-sm text-gray-600">دسته‌بندی‌ها</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

// جستجو و فیلتر
const searchQuery = ref('')
const selectedCategory = ref('')
const openFAQs = ref<number[]>([])

// دسته‌بندی‌ها
const categories = ref([
  { id: 'connection', name: 'اتصال', count: 15 },
  { id: 'configuration', name: 'پیکربندی', count: 12 },
  { id: 'sync', name: 'همگام‌سازی', count: 8 },
  { id: 'troubleshooting', name: 'عیب‌یابی', count: 10 },
  { id: 'security', name: 'امنیت', count: 6 },
  { id: 'performance', name: 'عملکرد', count: 7 }
])

interface FAQLink {
  title: string
  url: string
}

interface FAQ {
  id: number
  question: string
  answer: string
  category: string
  views: number
  helpfulVotes: number
  notHelpfulVotes: number
  solutions?: string[]
  links?: FAQLink[]
}

// سوالات متداول
const faqs = ref<FAQ[]>([
  {
    id: 1,
    question: 'چگونه می‌توانم به نرم‌افزار هلو متصل شوم؟',
    answer: 'برای اتصال به نرم‌افزار هلو، ابتدا باید کلید API را از پنل مدیریت هلو دریافت کنید و سپس آن را در تنظیمات اتصال وارد نمایید.',
    category: 'connection',
    views: 1250,
    helpfulVotes: 89,
    notHelpfulVotes: 3,
    solutions: [
      'بررسی فعال بودن سرویس هلو',
      'اطمینان از صحت کلید API',
      'بررسی تنظیمات فایروال'
    ],
    links: [
      { title: 'راهنمای کامل اتصال به هلو', url: '#' },
      { title: 'ویدیوی آموزشی', url: '#' }
    ]
  },
  {
    id: 2,
    question: 'چرا همگام‌سازی تراکنش‌ها کند است؟',
    answer: 'کندی در همگام‌سازی می‌تواند به دلایل مختلفی مانند حجم بالای داده‌ها، تنظیمات نادرست یا مشکلات شبکه باشد.',
    category: 'sync',
    views: 890,
    helpfulVotes: 67,
    notHelpfulVotes: 8,
    solutions: [
      'بررسی سرعت اینترنت',
      'کاهش حجم داده‌های همگام‌سازی',
      'بهینه‌سازی تنظیمات'
    ]
  },
  {
    id: 3,
    question: 'چگونه امنیت اتصال را افزایش دهم؟',
    answer: 'برای افزایش امنیت اتصال، از کلیدهای API قوی استفاده کنید، SSL را فعال کنید و دسترسی‌ها را محدود نمایید.',
    category: 'security',
    views: 650,
    helpfulVotes: 45,
    notHelpfulVotes: 2,
    solutions: [
      'استفاده از کلیدهای API قوی',
      'فعال‌سازی SSL',
      'محدود کردن IP های مجاز'
    ]
  }
])

// فیلتر کردن سوالات
const filteredFAQs = computed(() => {
  return faqs.value.filter(faq => {
    if (searchQuery.value && !faq.question.includes(searchQuery.value) && !faq.answer.includes(searchQuery.value)) return false
    if (selectedCategory.value && faq.category !== selectedCategory.value) return false
    return true
  })
})

// آمار
const totalFAQs = computed(() => faqs.value.length)
const totalViews = computed(() => faqs.value.reduce((sum, faq) => sum + faq.views, 0))
const totalHelpfulVotes = computed(() => faqs.value.reduce((sum, faq) => sum + faq.helpfulVotes, 0))

// کلاس دسته‌بندی
const getCategoryClass = (category: string) => {
  const classes = {
    connection: 'bg-blue-100 text-blue-700',
    configuration: 'bg-green-100 text-green-700',
    sync: 'bg-yellow-100 text-yellow-700',
    troubleshooting: 'bg-red-100 text-red-700',
    security: 'bg-purple-100 text-purple-700',
    performance: 'bg-orange-100 text-orange-700'
  }
  return classes[category] || 'bg-gray-100 text-gray-700'
}

// برچسب دسته‌بندی
const getCategoryLabel = (category: string) => {
  const labels = {
    connection: 'اتصال',
    configuration: 'پیکربندی',
    sync: 'همگام‌سازی',
    troubleshooting: 'عیب‌یابی',
    security: 'امنیت',
    performance: 'عملکرد'
  }
  return labels[category] || category
}

// انتخاب دسته‌بندی
const selectCategory = (categoryId: string) => {
  selectedCategory.value = selectedCategory.value === categoryId ? '' : categoryId
}

// باز/بسته کردن سوال
const toggleFAQ = (faqId: number) => {
  const index = openFAQs.value.indexOf(faqId)
  if (index > -1) {
    openFAQs.value.splice(index, 1)
  } else {
    openFAQs.value.push(faqId)
  }
}

// رای‌دهی
const rateFAQ = (faqId: number, type: 'helpful' | 'not_helpful') => {
  const faq = faqs.value.find(f => f.id === faqId)
  if (faq) {
    if (type === 'helpful') {
      faq.helpfulVotes++
    } else {
      faq.notHelpfulVotes++
    }
  }

}

// گزارش مشکل
const reportFAQ = (_faqId: number) => {

}

// پیشنهاد سوال
const suggestQuestion = () => {

}
</script>

<!--
  کامپوننت سوالات متداول
  شامل:
  1. جستجو و فیلتر
  2. دسته‌بندی‌ها
  3. آکاردئون سوالات
  4. رای‌دهی
  5. آمار
  6. طراحی مدرن و کاملاً ریسپانسیو
--> 
