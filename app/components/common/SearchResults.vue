<template>
  <div class="bg-white rounded-lg shadow-lg border border-gray-200 max-h-96 overflow-y-auto">
    <div class="p-6">
      <div class="flex items-center justify-between mb-3">
        <h3 class="text-sm font-medium text-gray-900">نتایج جستجو</h3>
        <NuxtLink
          :to="{ path: '/search-results', query: { q: query } }"
          class="text-xs text-blue-600 hover:text-blue-800"
        >
          مشاهده همه
        </NuxtLink>
      </div>
      
      <div class="space-y-3">
        <div
          v-for="result in results"
          :key="`${result.type}-${result.id}`"
          class="flex items-center space-x-3 space-x-reverse p-2 rounded-lg hover:bg-gray-50 cursor-pointer transition-colors"
          @click="selectResult(result)"
        >
          <!-- تصویر -->
          <div class="flex-shrink-0">
            <img
              v-if="result.image"
              :src="result.image"
              :alt="result.title"
              class="w-10 h-10 object-cover rounded-lg"
            />
            <div
              v-else
              class="w-10 h-10 bg-gray-200 rounded-lg flex items-center justify-center"
            >
              <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
            </div>
          </div>

          <!-- محتوا -->
          <div class="flex-1 min-w-0">
            <div class="flex items-center space-x-2 space-x-reverse">
              <h4 class="text-sm font-medium text-gray-900 truncate">{{ result.title }}</h4>
              <span
                :class="[
                  'text-xs px-1.5 py-0.5 rounded-full',
                  getTypeBadgeClass(result.type)
                ]"
              >
                {{ getTypeLabel(result.type) }}
              </span>
            </div>
            
            <p v-if="result.description" class="text-xs text-gray-500 truncate">
              {{ result.description }}
            </p>
            
            <div v-if="result.price" class="text-xs text-green-600 font-medium">
              {{ formatPrice(result.price) }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Props
interface Props {
  results: any[]
  query: string
}

defineProps<Props>()

// Emits
const emit = defineEmits<{
  'result-selected': [result: any]
}>()

// Methods
const selectResult = (result: any) => {
  emit('result-selected', result)
}

const getTypeLabel = (type: string) => {
  const labels: Record<string, string> = {
    product: 'محصول',
    post: 'مقاله',
    category: 'دسته‌بندی',
    brand: 'برند'
  }
  return labels[type] || type
}

const getTypeBadgeClass = (type: string) => {
  const classes: Record<string, string> = {
    product: 'bg-blue-100 text-blue-800',
    post: 'bg-green-100 text-green-800',
    category: 'bg-purple-100 text-purple-800',
    brand: 'bg-orange-100 text-orange-800'
  }
  return classes[type] || 'bg-gray-100 text-gray-800'
}

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان'
}
</script> 