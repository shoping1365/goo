<template>
  <div class="space-y-6 overflow-hidden" dir="rtl">
    <!-- توضیحات کوتاه -->
    <div v-if="product?.description" class="bg-yellow-50 rounded-lg border border-yellow-200 px-4 py-4 overflow-hidden">
      <h4 class="text-base font-bold text-yellow-800 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6"></path>
        </svg>
        توضیحات کوتاه
      </h4>
      <!-- 
        ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
        
        این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
        
        ✅ راه حل صحیح:
        1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
        2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
        3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
        
        مثال صحیح:
        import DOMPurify from 'dompurify'
        const sanitizedDescription = DOMPurify.sanitize(product.description)
        <div v-html="sanitizedDescription"></div>
      -->
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div class="prose prose-sm max-w-none text-gray-800 leading-relaxed overflow-hidden" v-html="product.description"></div>
    </div>

    <!-- توضیحات کامل -->
    <div v-if="product?.full_description" class="bg-white rounded-lg border border-gray-200 px-4 py-4 overflow-hidden">
      <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        توضیحات کامل محصول
      </h3>
      <!-- 
        ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
        
        این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
        
        ✅ راه حل صحیح:
        1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
        2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
        3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
        
        مثال صحیح:
        import DOMPurify from 'dompurify'
        const sanitizedFullDescription = DOMPurify.sanitize(product.full_description)
        <div v-html="sanitizedFullDescription"></div>
      -->
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div class="prose prose-sm max-w-none text-gray-700 leading-relaxed overflow-hidden" v-html="product.full_description"></div>
    </div>
    <!-- اگر توضیحات کامل نباشد اما کوتاه هست همان را نشان بده -->
    <div v-else-if="!product.full_description && product.description" class="bg-white rounded-lg border border-gray-200 px-4 py-4 overflow-hidden">
      <h3 class="text-lg font-bold text-gray-900 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        توضیحات محصول
      </h3>
      <!-- 
        ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
        
        این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
        
        ✅ راه حل صحیح:
        1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
        2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
        3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
        
        مثال صحیح:
        import DOMPurify from 'dompurify'
        const sanitizedDescription = DOMPurify.sanitize(product.description)
        <div v-html="sanitizedDescription"></div>
      -->
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div class="prose prose-sm max-w-none text-gray-700 leading-relaxed overflow-hidden" v-html="product.description"></div>
    </div>

    <!-- نقاط قوت و ضعف -->
    <div v-if="product?.strengths?.length || product?.weaknesses?.length" class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
      <!-- نقاط قوت -->
      <div v-if="product?.strengths?.length" class="bg-green-50 rounded-lg border border-green-200 px-4 py-4">
        <h4 class="text-lg font-bold text-green-800 mb-4 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          نقاط قوت
        </h4>
        <ul class="space-y-3">
          <li v-for="(strength, index) in product.strengths" :key="index" class="flex items-start gap-3 text-green-700">
            <svg class="w-4 h-4 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
            </svg>
            <span>{{ strength }}</span>
          </li>
        </ul>
      </div>

      <!-- نقاط ضعف -->
      <div v-if="product?.weaknesses?.length" class="bg-red-50 rounded-lg border border-red-200 px-4 py-4">
        <h4 class="text-lg font-bold text-red-800 mb-4 flex items-center gap-2">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          نقاط ضعف
        </h4>
        <ul class="space-y-3">
          <li v-for="(weakness, index) in product.weaknesses" :key="index" class="flex items-start gap-3 text-red-700">
            <svg class="w-4 h-4 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
            </svg>
            <span>{{ weakness }}</span>
          </li>
        </ul>
      </div>
    </div>

    <!-- راهنمای استفاده -->
    <div v-if="product?.usage_guide" class="bg-blue-50 rounded-lg border border-blue-200 px-4 py-4 overflow-hidden">
      <h4 class="text-lg font-bold text-blue-800 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"></path>
        </svg>
        راهنمای استفاده
      </h4>
      <!-- 
        ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
        
        این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
        
        ✅ راه حل صحیح:
        1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
        2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
        3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
        
        مثال صحیح:
        import DOMPurify from 'dompurify'
        const sanitizedUsageGuide = DOMPurify.sanitize(product.usage_guide)
        <div v-html="sanitizedUsageGuide"></div>
      -->
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div class="prose prose-sm max-w-none text-blue-700 leading-relaxed overflow-hidden" v-html="product.usage_guide"></div>
    </div>

    <!-- مشخصات کلی -->
    <div v-if="product?.features?.length" class="bg-gray-50 rounded-lg border border-gray-200 px-4 py-4 overflow-hidden">
      <h4 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
        </svg>
        ویژگی‌های کلیدی
      </h4>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gapx-4 py-4">
        <div v-for="(feature, index) in product.features" :key="index" class="flex items-center gap-3 bg-white rounded-lg px-4 py-4 border border-gray-200">
          <div class="w-2 h-2 bg-blue-500 rounded-full flex-shrink-0"></div>
          <span class="text-gray-700">{{ feature }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Product {
  description?: string
  strengths?: string[]
  weaknesses?: string[]
  usage_guide?: string
  features?: string[]
}

interface Props {
  product: Product
}

defineProps<Props>()
</script>

<style scoped>
/* اضافه کردن استایل‌های اضافی برای جلوگیری از overflow ویدیوها */
.prose iframe,
.prose video {
  max-width: 100% !important;
  height: auto !important;
  max-height: 400px !important;
  border-radius: 0.5rem;
  overflow: hidden;
}

.prose img {
  max-width: 100% !important;
  height: auto !important;
  border-radius: 0.5rem;
}

/* برای موبایل */
@media (max-width: 768px) {
  .prose iframe,
  .prose video {
    max-height: 250px !important;
  }
}
</style> 
