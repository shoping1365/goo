<template>
  <div class="space-y-4" dir="rtl">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h3 class="text-lg font-bold text-gray-900 flex items-center gap-2">
        <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
        </svg>
        سوالات متداول
      </h3>
      
      <div class="flex gap-2">
        <button
          @click="expandAll"
          class="px-3 py-1 bg-blue-100 text-blue-700 rounded-lg hover:bg-blue-200 transition-colors text-sm"
        >
          باز کردن همه
        </button>
        <button
          @click="collapseAll"
          class="px-3 py-1 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors text-sm"
        >
          بستن همه
        </button>
      </div>
    </div>

    <!-- FAQ List -->
    <div v-if="faqs?.length" class="space-y-3">
      <div
        v-for="(faq, index) in faqs"
        :key="faq.id || index"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden"
      >
        <button
          @click="toggleFaq(index)"
          class="w-full px-6 py-4 text-right hover:bg-gray-50 transition-colors flex items-center justify-between"
        >
          <h4 class="font-medium text-gray-900 flex-1">
            {{ faq.question }}
          </h4>
          <svg
            :class="[
              'w-5 h-5 text-gray-500 transition-transform ml-4 flex-shrink-0',
              expandedFaqs[index] ? 'rotate-180' : ''
            ]"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
          </svg>
        </button>
        
        <div v-show="expandedFaqs[index]" class="px-6 pb-4">
          <div class="text-gray-700 leading-relaxed" v-html="faq.answer"></div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="bg-gray-50 rounded-lg p-12 text-center">
      <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
      </svg>
      <h3 class="text-lg font-medium text-gray-900 mb-2">سوالی ثبت نشده است</h3>
      <p class="text-gray-500">هنوز سوال متداولی برای این محصول ثبت نشده است.</p>
    </div>

    <!-- Ask Question -->
    <div class="bg-blue-50 rounded-lg px-4 py-4 border border-blue-200">
      <h4 class="font-medium text-blue-900 mb-2 flex items-center gap-2">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        سوال خود را بپرسید
      </h4>
      <p class="text-blue-700 text-sm mb-4">
        اگر سوال خاصی درباره این محصول دارید، می‌توانید از بخش نظرات یا تماس با ما استفاده کنید.
      </p>
      <div class="flex gap-3">
        <NuxtLink 
          to="/contact"
          class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors"
        >
          تماس با ما
        </NuxtLink>
        <button
          @click="scrollToReviews"
          class="bg-white border border-blue-300 text-blue-700 hover:bg-blue-50 px-4 py-2 rounded-lg text-sm font-medium transition-colors"
        >
          ثبت نظر
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'

interface Props {
  product: any
}

const props = defineProps<Props>()

// State
const expandedFaqs = reactive<Record<number, boolean>>({})

// Computed
const faqs = computed(() => {
  return props.product?.faqs || []
})

// Methods
const toggleFaq = (index: number) => {
  expandedFaqs[index] = !expandedFaqs[index]
}

const expandAll = () => {
  faqs.value.forEach((_: any, index: number) => {
    expandedFaqs[index] = true
  })
}

const collapseAll = () => {
  faqs.value.forEach((_: any, index: number) => {
    expandedFaqs[index] = false
  })
}

const scrollToReviews = () => {
  // Emit event to parent to change tab
  // For now, we'll just scroll to reviews section
  const reviewsTab = document.querySelector('[data-tab="reviews"]') as HTMLElement
  if (reviewsTab) {
    reviewsTab.click()
  }
}
</script> 
