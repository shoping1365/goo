<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-semibold text-gray-800">مدیریت دسته‌ای سفارشات</h3>
        <p class="text-gray-600 text-sm">انتخاب چند سفارش و انجام عملیات گروهی</p>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm hover:bg-blue-700" @click="refreshOrders">بروزرسانی لیست</button>
    </div>
    <!-- Bulk Actions -->
    <div class="flex flex-wrap gap-2 items-center mb-2">
      <button 
        class="px-3 py-1 bg-gray-200 rounded text-sm hover:bg-gray-300"
        @click="selectAll"
      >انتخاب همه</button>
      <button 
        class="px-3 py-1 bg-gray-200 rounded text-sm hover:bg-gray-300"
        @click="clearSelection"
      >لغو انتخاب</button>
      <span class="text-sm text-gray-700">تعداد انتخاب شده: <b>{{ selected.length }}</b></span>
      <button 
        :disabled="!selected.length"
        class="px-3 py-1 bg-green-600 text-white rounded text-sm hover:bg-green-700 disabled:bg-gray-300 disabled:text-gray-500"
        @click="bulkSendSMS"
      >ارسال پیامک گروهی</button>
      <button 
        :disabled="!selected.length"
        class="px-3 py-1 bg-red-600 text-white rounded text-sm hover:bg-red-700 disabled:bg-gray-300 disabled:text-gray-500"
        @click="bulkDelete"
      >حذف گروهی</button>
      <button 
        :disabled="!selected.length"
        class="px-3 py-1 bg-purple-600 text-white rounded text-sm hover:bg-purple-700 disabled:bg-gray-300 disabled:text-gray-500"
        @click="showTagModal = true"
      >برچسب‌گذاری</button>
    </div>
    <!-- Orders Table -->
    <div class="bg-white rounded-lg border border-gray-200 overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3"><input type="checkbox" :checked="allSelected" @change="toggleAll"></th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">سفارش</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">مشتری</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">وضعیت</th>
            <th class="px-4 py-3 text-right text-xs font-medium text-gray-500">برچسب‌ها</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id" class="hover:bg-gray-50">
            <td class="px-4 py-3"><input v-model="selected" type="checkbox" :value="order.id"></td>
            <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">#{{ order.orderNumber }}</td>
            <td class="px-4 py-3 whitespace-nowrap text-sm text-gray-900">{{ order.customerName }}</td>
            <td class="px-4 py-3 whitespace-nowrap text-xs">
              <span :class="getStatusClass(order.status)">{{ getStatusText(order.status) }}</span>
            </td>
            <td class="px-4 py-3 whitespace-nowrap">
              <span v-for="tag in order.tags" :key="tag" class="inline-block bg-purple-100 text-purple-800 rounded-full px-2 py-0.5 text-xs font-medium mr-1">{{ tag }}</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <!-- Tag Modal -->
    <div v-if="showTagModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50" @click.self="showTagModal = false">
      <div class="bg-white rounded-lg px-4 py-4 w-full max-w-md mx-4">
        <h4 class="text-lg font-semibold text-gray-800 mb-4">برچسب‌گذاری گروهی</h4>
        <div class="flex flex-wrap gap-2 mb-4">
          <span v-for="tag in availableTags" :key="tag" class="inline-block bg-purple-100 text-purple-800 rounded-full px-2 py-0.5 text-xs font-medium">{{ tag }}</span>
        </div>
        <input v-model="newTag" type="text" class="w-full px-3 py-2 border border-gray-300 rounded-lg mb-4" placeholder="برچسب جدید">
        <div class="flex items-center gap-2">
          <button class="px-4 py-2 bg-purple-600 text-white rounded-lg text-sm hover:bg-purple-700" @click="addTagToSelection">افزودن برچسب</button>
          <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-lg text-sm hover:bg-gray-400" @click="showTagModal = false">انصراف</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue'
const orders = ref([
  { id: 1, orderNumber: '1001', customerName: 'علی احمدی', status: 'delivered', tags: ['VIP'] },
  { id: 2, orderNumber: '1002', customerName: 'زهرا رضایی', status: 'pending', tags: [] },
  { id: 3, orderNumber: '1003', customerName: 'محمد موسوی', status: 'delivered', tags: ['شکایت'] },
  { id: 4, orderNumber: '1004', customerName: 'سارا کریمی', status: 'canceled', tags: [] },
  { id: 5, orderNumber: '1005', customerName: 'رضا مرادی', status: 'pending', tags: ['جدید'] }
])
const selected = ref<number[]>([])
const showTagModal = ref(false)
const newTag = ref('')
const availableTags = ref(['VIP', 'شکایت', 'جدید'])
const allSelected = computed(() => selected.value.length === orders.value.length && orders.value.length > 0)
const selectAll = () => { selected.value = orders.value.map(o => o.id) }
const clearSelection = () => { selected.value = [] }
const toggleAll = () => { allSelected.value ? clearSelection() : selectAll() }
const refreshOrders = () => { /* فرضی: دریافت مجدد لیست */ alert('لیست سفارشات بروزرسانی شد!') }
const getStatusClass = (status: string) => {
  return status === 'delivered' ? 'bg-green-100 text-green-800 px-2 py-0.5 rounded-full' :
         status === 'pending' ? 'bg-yellow-100 text-yellow-800 px-2 py-0.5 rounded-full' :
         status === 'canceled' ? 'bg-red-100 text-red-800 px-2 py-0.5 rounded-full' :
         'bg-gray-100 text-gray-800 px-2 py-0.5 rounded-full'
}
const getStatusText = (status: string) => {
  return status === 'delivered' ? 'تحویل شده' :
         status === 'pending' ? 'در انتظار' :
         status === 'canceled' ? 'لغو شده' :
         status
}
const bulkSendSMS = () => { alert('ارسال پیامک گروهی برای سفارشات انتخاب شده!') }
const bulkDelete = () => {
  if (confirm('آیا مطمئن هستید که می‌خواهید سفارشات انتخاب شده حذف شوند؟')) {
    orders.value = orders.value.filter(o => !selected.value.includes(o.id))
    clearSelection()
    alert('سفارشات حذف شدند.')
  }
}
const addTagToSelection = () => {
  if (!newTag.value.trim()) return
  selected.value.forEach(id => {
    const order = orders.value.find(o => o.id === id)
    if (order && !order.tags.includes(newTag.value)) order.tags.push(newTag.value)
  })
  if (!availableTags.value.includes(newTag.value)) availableTags.value.push(newTag.value)
  newTag.value = ''
  showTagModal.value = false
}
</script> 
