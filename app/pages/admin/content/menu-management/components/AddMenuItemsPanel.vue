<template>
  <div class="lg:col-span-1 lg:w-[35%]">
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-gray-900">افزودن آیتم‌های منو</h3>
        <button
          class="text-blue-600 hover:text-blue-800 transition-colors"
          title="بروزرسانی محتوا"
          @click="emit('refresh')"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>
      </div>
      
      <!-- Custom Link Section -->
  <CustomLinkPanel :custom-link="customLink" @add-custom-link="(linkData) => emit('add-custom-link', linkData)" />

      <!-- Content Panels -->
      <ContentPanels
        :open-panel="openPanel"
        :pages="pages"
        :posts="posts"
        :categories="categories"
        :product-categories="productCategories"
        @toggle-panel="togglePanel"
        @add-selected-pages="emit('add-selected-pages', $event)"
        @add-selected-posts="emit('add-selected-posts', $event)"
        @add-selected-categories="emit('add-selected-categories', $event)"
        @add-selected-product-categories="emit('add-selected-product-categories', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import CustomLinkPanel from './CustomLinkPanel.vue'
import ContentPanels from './ContentPanels.vue'

const props = defineProps({
  customLink: {
    type: Object,
    required: true,
  },
  pages: {
    type: Array,
    default: () => [],
  },
  posts: {
    type: Array,
    default: () => [],
  },
  categories: {
    type: Array,
    default: () => [],
  },
  productCategories: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits([
  'refresh',
  'add-custom-link',
  'add-selected-pages',
  'add-selected-posts',
  'add-selected-categories',
  'add-selected-product-categories',
])

const openPanel = ref('pages')

const togglePanel = (panel) => {
  openPanel.value = openPanel.value === panel ? null : panel
}
</script>
