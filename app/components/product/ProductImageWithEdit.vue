<template>
  <div class="relative bg-gray-100 rounded-lg overflow-hidden aspect-square group">
    <img
      :src="imageUrl"
      :alt="productName"
      class="w-full h-full object-cover transition-transform group-hover:scale-105"
    />

    <!-- Admin Edit Button - Only visible for admins -->
    <div v-if="isAdmin" class="absolute top-2 right-2 z-20 opacity-0 group-hover:opacity-100 transition-opacity">
      <NuxtLink
        :to="editUrl"
        target="_blank"
        class="bg-blue-600 hover:bg-blue-700 text-white p-2 rounded-full shadow-lg hover:shadow-xl transition-all duration-200 flex items-center justify-center"
        title="ویرایش محصول"
        style="width: 2.25rem; height: 2.25rem;"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
        </svg>
      </NuxtLink>
    </div>

    <!-- Optional: Click overlay for modal or other actions -->
    <div
      v-if="clickable"
      class="absolute inset-0 cursor-pointer z-10"
      @click="emit('click')"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAdmin } from '~/composables/useAdmin'

interface Props {
  productId: number
  productName: string
  imageUrl: string
  clickable?: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{
  'click': []
}>()

const { isAdmin } = useAdmin()

const editUrl = computed(() => `/admin/product-management/products/edit?id=${props.productId}`)
</script>

<style scoped>
/* Smooth hover transition for edit button */
.group:hover .group-hover\:opacity-100 {
  opacity: 1;
}
</style>
