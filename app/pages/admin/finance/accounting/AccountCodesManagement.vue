<template>
  <div>
    <div class="flex items-center justify-between mb-4">
      <h4 class="text-base font-semibold text-gray-900">مدیریت کدهای حساب</h4>
      <button class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg" @click="showAddModal = true">افزودن کد حساب</button>
    </div>
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">کد حساب</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">نام حساب</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">نوع</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">شرح</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="code in accountCodes" :key="code.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4 font-mono">{{ code.code }}</td>
            <td class="py-3 px-4">{{ code.name }}</td>
            <td class="py-3 px-4">{{ code.type }}</td>
            <td class="py-3 px-4">{{ code.description }}</td>
            <td class="py-3 px-4">
              <button class="text-yellow-600 hover:text-yellow-800 mx-1" @click="editCode(code)">ویرایش</button>
              <button class="text-red-600 hover:text-red-800 mx-1" @click="deleteCode(code)">حذف</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <!-- مودال افزودن/ویرایش -->
    <div v-if="showAddModal || showEditModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md mx-4">
        <h5 class="font-bold mb-4">{{ showEditModal ? 'ویرایش کد حساب' : 'افزودن کد حساب' }}</h5>
        <form class="space-y-4" @submit.prevent="showEditModal ? updateCode() : addCode()">
          <div>
            <label class="block text-sm mb-1">کد حساب</label>
            <input v-model="form.code" required class="w-full px-3 py-2 border rounded-lg font-mono" />
          </div>
          <div>
            <label class="block text-sm mb-1">نام حساب</label>
            <input v-model="form.name" required class="w-full px-3 py-2 border rounded-lg" />
          </div>
          <div>
            <label class="block text-sm mb-1">نوع</label>
            <select v-model="form.type" required class="w-full px-3 py-2 border rounded-lg">
              <option value="فروش">فروش</option>
              <option value="خرید">خرید</option>
              <option value="مالیاتی">مالیاتی</option>
              <option value="بانکی">بانکی</option>
              <option value="سایر">سایر</option>
            </select>
          </div>
          <div>
            <label class="block text-sm mb-1">شرح</label>
            <input v-model="form.description" class="w-full px-3 py-2 border rounded-lg" />
          </div>
          <div class="flex gap-2 mt-4">
            <button type="submit" class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg">ذخیره</button>
            <button type="button" class="flex-1 px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-700 rounded-lg" @click="closeModal">انصراف</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const accountCodes = ref([
  { id: 1, code: '4001', name: 'فروش کالا', type: 'فروش', description: 'کد فروش کالا' },
  { id: 2, code: '5001', name: 'خرید کالا', type: 'خرید', description: 'کد خرید کالا' },
  { id: 3, code: '2201', name: 'مالیات بر ارزش افزوده', type: 'مالیاتی', description: 'کد مالیات بر ارزش افزوده' }
])
const showAddModal = ref(false)
const showEditModal = ref(false)
const form = ref({ id: null, code: '', name: '', type: '', description: '' })
const editCode = (code: any) => {
  form.value = { ...code }
  showEditModal.value = true
}
const addCode = () => {
  accountCodes.value.push({ ...form.value, id: Date.now() })
  closeModal()
}
const updateCode = () => {
  const idx = accountCodes.value.findIndex(a => a.id === form.value.id)
  if (idx !== -1) accountCodes.value[idx] = { ...form.value }
  closeModal()
}
const deleteCode = (code: any) => {
  if (confirm('آیا از حذف این کد اطمینان دارید؟'))
    accountCodes.value = accountCodes.value.filter(a => a.id !== code.id)
}
const closeModal = () => {
  showAddModal.value = false
  showEditModal.value = false
  form.value = { id: null, code: '', name: '', type: '', description: '' }
}
</script>
<!-- مدیریت کدهای حساب --> 
