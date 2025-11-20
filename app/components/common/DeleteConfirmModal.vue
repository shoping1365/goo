<template>
  <!-- مودال تایید حذف تک آیتم -->
  <div v-if="deleteModalOpen" class="fixed inset-0 bg-transparent z-50 flex items-center justify-center">
    <div role="dialog" aria-modal="true" class="bg-blue-50 rounded-xl shadow-2xl ring-4 ring-purple-300 p-6 sm:p-7 w-full max-w-md mx-4">
      <h3 class="text-base font-semibold text-gray-900 mb-2 flex items-center gap-2">
        <span class="i-heroicons-exclamation-triangle text-red-500 text-xl"></span>
        {{ singleDeleteTitle }}
      </h3>
      <p class="text-sm leading-6 text-gray-600 mb-6">{{ singleDeleteMessage }}</p>
      <div class="flex justify-end gap-2">
        <button class="px-4 py-2 text-sm rounded-md bg-gray-100 text-gray-700 hover:bg-gray-200" @click="closeDeleteConfirm">انصراف</button>
        <button :disabled="deleting" class="px-4 py-2 text-sm rounded-md bg-red-600 text-white hover:bg-red-700 disabled:opacity-60 disabled:cursor-not-allowed" @click="confirmDelete">
          <span v-if="deleting" class="i-heroicons-arrow-path animate-spin mr-1"></span>
          حذف
        </button>
      </div>
    </div>
  </div>

  <!-- مودال تایید حذف گروهی -->
  <div v-if="bulkDeleteModalOpen" class="fixed inset-0 bg-transparent z-50 flex items-center justify-center">
    <div role="dialog" aria-modal="true" class="bg-blue-50 rounded-xl shadow-2xl ring-4 ring-purple-300 p-6 sm:p-7 w-full max-w-md mx-4">
      <h3 class="text-base font-semibold text-gray-900 mb-2 flex items-center gap-2">
        <span class="i-heroicons-exclamation-triangle text-red-500 text-xl"></span>
        {{ bulkDeleteTitle }}
      </h3>
      <p class="text-sm leading-6 text-gray-600 mb-6">{{ bulkDeleteMessage }}</p>
      <div class="flex justify-end gap-2">
        <button class="px-4 py-2 text-sm rounded-md bg-gray-100 text-gray-700 hover:bg-gray-200" @click="closeBulkDeleteConfirm">انصراف</button>
        <button :disabled="deletingBulk" class="px-4 py-2 text-sm rounded-md bg-red-600 text-white hover:bg-red-700 disabled:opacity-60 disabled:cursor-not-allowed" @click="confirmBulkDelete">
          <span v-if="deletingBulk" class="i-heroicons-arrow-path animate-spin mr-1"></span>
          حذف
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// Props برای سفارشی‌سازی متن‌ها
const props = defineProps<{
  singleDeleteTitle?: string
  singleDeleteMessage?: string
  bulkDeleteTitle?: string
  bulkDeleteMessage?: string
  selectedCount?: number
}>()

// Emits برای اطلاع‌رسانی به کامپوننت والد
const emit = defineEmits<{
  'single-delete': [id: number | string]
  'bulk-delete': []
  'close-single': []
  'close-bulk': []
}>()

// وضعیت مودال‌ها
const deleteModalOpen = ref(false)
const deleting = ref(false)
const deleteItemId = ref<number | string | null>(null)
const bulkDeleteModalOpen = ref(false)
const deletingBulk = ref(false)

// متن‌های پیش‌فرض
const singleDeleteTitle = computed(() => props.singleDeleteTitle || 'تایید حذف')
const singleDeleteMessage = computed(() => props.singleDeleteMessage || 'آیا از حذف این مورد اطمینان دارید؟ این عملیات غیرقابل بازگشت است.')
const bulkDeleteTitle = computed(() => props.bulkDeleteTitle || 'حذف گروهی')
const bulkDeleteMessage = computed(() => {
  const count = props.selectedCount || 0
  return props.bulkDeleteMessage || `آیا از حذف ${count} مورد انتخاب شده اطمینان دارید؟ این عملیات غیرقابل بازگشت است.`
})

// باز کردن مودال تأیید حذف تک آیتم
function openDeleteConfirm(itemId: number | string) {
  deleteItemId.value = itemId
  deleteModalOpen.value = true
}

// بستن مودال تأیید حذف تک آیتم
function closeDeleteConfirm() {
  deleteModalOpen.value = false
  deleteItemId.value = null
  emit('close-single')
}

// تأیید حذف تک آیتم
async function confirmDelete() {
  console.log('confirmDelete فراخوانی شد، ID:', deleteItemId.value)
  if (!deleteItemId.value) return
  try {
    deleting.value = true
    console.log('emit single-delete با ID:', deleteItemId.value)
    emit('single-delete', deleteItemId.value)
    closeDeleteConfirm()
  } finally {
    deleting.value = false
  }
}

// باز کردن مودال تأیید حذف گروهی
function openBulkDeleteConfirm() {
  bulkDeleteModalOpen.value = true
}

// بستن مودال تأیید حذف گروهی
function closeBulkDeleteConfirm() {
  bulkDeleteModalOpen.value = false
  emit('close-bulk')
}

// تأیید حذف گروهی
async function confirmBulkDelete() {
  try {
    deletingBulk.value = true
    emit('bulk-delete')
    closeBulkDeleteConfirm()
  } finally {
    deletingBulk.value = false
  }
}

// Expose methods برای استفاده از خارج
defineExpose({
  openDeleteConfirm,
  openBulkDeleteConfirm,
  closeDeleteConfirm,
  closeBulkDeleteConfirm
})
</script>

<style scoped>
/* استایل‌های دقیق از ProductTable.vue */
</style>
