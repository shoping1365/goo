<template>
  <div class="flex gap-2 justify-center pb-2">
    <div
      v-for="(image, index) in images"
      :key="index"
      :class="[
        'w-12 h-12 rounded-lg border-2 overflow-hidden transition-all cursor-pointer relative group',
        currentIndex === index ? 'border-blue-500' : 'border-gray-200 hover:border-gray-300'
      ]"
      @click="index === 0 ? $emit('modal-click') : $emit('thumbnail-click', index)"
    >
      <img
        :src="image"
        :alt="`تصویر ${index + 1}`"
        class="w-full h-full object-cover"
      />

      <!-- Blur overlay for first image -->
      <div v-if="index === 0" class="absolute inset-0 flex items-center justify-center bg-black bg-opacity-50">
        <span class="text-white text-xs font-bold">---</span>
      </div>

      <!-- Admin Edit Button on first thumbnail - only visible on hover -->
      <div v-if="index === 0 && isAdmin" class="absolute top-0 right-0 z-10 opacity-0 group-hover:opacity-100 transition-opacity">
        <NuxtLink
          :to="editUrl"
          target="_blank"
          class="bg-blue-600 text-white p-0.5 rounded-full flex items-center justify-center hover:bg-blue-700 transition-colors"
          style="width: 1rem; height: 1rem;"
          title="ویرایش محصول"
        >
          <svg class="w-2 h-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
          </svg>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAdmin } from '~/composables/useAdmin'

interface Props {
  images: string[]
  currentIndex: number
  productId: number
}

const props = defineProps<Props>()
const emit = defineEmits(['thumbnail-click', 'modal-click'])

const { isAdmin } = useAdmin()

const editUrl = computed(() => `/admin/product-management/products/edit?id=${props.productId}`)
</script>
