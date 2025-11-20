<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 max-h-[80vh] overflow-hidden">
      <!-- هدر مودال -->
      <div class="flex items-center justify-between p-6 border-b border-gray-200">
        <h3 class="text-lg font-medium text-gray-900">انتخاب آیتم فوتر</h3>
        <button
          class="text-gray-400 hover:text-gray-600 transition-colors"
          @click="$emit('close')"
        >
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- محتوای مودال -->
      <div class="p-6 overflow-y-auto max-h-[60vh]">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div
            v-for="item in availableFooterItems"
            :key="item.id"
            class="item-card cursor-pointer border border-gray-200 rounded-lg p-6 hover:border-blue-300 hover:shadow-md transition-all duration-200"
            @click="selectItem(item.id)"
          >
            <!-- آیکون آیتم -->
            <div class="flex items-center space-x-3 space-x-reverse mb-3">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                <Icon :name="item.icon" class="w-6 h-6 text-blue-600" />
              </div>
              <div>
                <h4 class="font-medium text-gray-900">{{ item.name }}</h4>
                <p class="text-sm text-gray-500">{{ getItemDescription(item.id) }}</p>
              </div>
            </div>

            <!-- پیش‌نمایش آیتم -->
            <div class="item-preview">
              <div v-if="item.id === 'logo'" class="logo-preview">
                <div class="w-16 h-8 bg-gray-200 rounded flex items-center justify-center">
                  <span class="text-xs text-gray-600">LOGO</span>
                </div>
              </div>
              
              <div v-else-if="item.id === 'links'" class="links-preview">
                <div class="space-y-1">
                  <div class="h-3 bg-gray-200 rounded w-16"></div>
                  <div class="h-3 bg-gray-200 rounded w-20"></div>
                  <div class="h-3 bg-gray-200 rounded w-14"></div>
                </div>
              </div>
              
              <div v-else-if="item.id === 'contact'" class="contact-preview">
                <div class="space-y-1">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <div class="w-3 h-3 bg-gray-300 rounded-full"></div>
                    <div class="h-3 bg-gray-200 rounded w-24"></div>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <div class="w-3 h-3 bg-gray-300 rounded-full"></div>
                    <div class="h-3 bg-gray-200 rounded w-20"></div>
                  </div>
                </div>
              </div>

              <div v-else-if="item.id === 'about'" class="about-preview">
                <div class="space-y-2">
                  <div class="h-3 bg-gray-300 rounded w-20"></div>
                  <div class="h-3 bg-gray-200 rounded w-28"></div>
                  <div class="h-3 bg-gray-200 rounded w-24"></div>
                </div>
              </div>
              
              <div v-else-if="item.id === 'newsletter'" class="newsletter-preview">
                <div class="flex space-x-2 space-x-reverse">
                  <div class="h-6 bg-gray-200 rounded flex-1"></div>
                  <div class="h-6 w-16 bg-blue-500 rounded"></div>
                </div>
              </div>
              
              <div v-else-if="item.id === 'social'" class="social-preview">
                <div class="flex space-x-2 space-x-reverse">
                  <div class="w-6 h-6 bg-blue-400 rounded"></div>
                  <div class="w-6 h-6 bg-green-400 rounded"></div>
                  <div class="w-6 h-6 bg-pink-400 rounded"></div>
                </div>
              </div>
              
              <div v-else-if="item.id === 'trust'" class="trust-preview">
                <div class="flex space-x-2 space-x-reverse">
                  <div class="w-8 h-8 bg-yellow-200 rounded border-2 border-yellow-300"></div>
                  <div class="w-8 h-8 bg-blue-200 rounded border-2 border-blue-300"></div>
                </div>
              </div>
              
              <div v-else-if="item.id === 'custom'" class="custom-preview">
                <div class="w-full h-8 bg-gray-200 rounded border-2 border-dashed border-gray-300 flex items-center justify-center">
                  <span class="text-xs text-gray-500">Custom</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- فوتر مودال -->
      <div class="flex justify-end p-6 border-t border-gray-200">
        <button
          class="px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          @click="$emit('close')"
        >
          انصراف
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue'

interface Props {
  availableFooterItems?: any[]
}

const props = withDefaults(defineProps<Props>(), {
  availableFooterItems: () => []
})

// دریافت آیتم‌های قابل استفاده
const availableFooterItems = inject('availableItems', props.availableFooterItems)

// Emit events
const emit = defineEmits<{
  close: []
  'item-selected': [itemId: string]
}>()

// انتخاب آیتم
const selectItem = (itemId: string) => {
  emit('item-selected', itemId)
}

// توضیحات آیتم‌ها
function getItemDescription(itemId: string): string {
  const descriptions: { [key: string]: string } = {
    'logo': 'لوگوی شرکت یا برند',
    'links': 'لینک‌های مفید و مهم',
    'contact': 'اطلاعات تماس و آدرس',
    'about': 'متن درباره ما از تنظیمات سایت',
    'newsletter': 'عضویت در خبرنامه',
    'social': 'شبکه‌های اجتماعی',
    'trust': 'نشان‌های اعتماد و گواهی‌ها',
    'custom': 'محتوای سفارشی'
  }
  return descriptions[itemId] || 'آیتم فوتر'
}
</script>

<style scoped>
.item-card {
  transition: all 0.2s ease;
}

.item-card:hover {
  transform: translateY(-2px);
}

.item-preview {
  min-height: 40px;
  display: flex;
  align-items: center;
}

.logo-preview,
.links-preview,
.contact-preview,
.about-preview,
.newsletter-preview,
.social-preview,
.trust-preview,
.custom-preview {
  width: 100%;
}
</style>
