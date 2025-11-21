<template>
  <div>
    <div class="flex items-center justify-between mb-4">
      <h4 class="text-base font-semibold text-gray-900">تعریف حساب‌های بانکی</h4>
      <button class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg" @click="showAddModal = true">افزودن حساب بانکی</button>
    </div>
    <div class="overflow-x-auto">
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 bg-gray-50">
            <th class="text-right py-3 px-4 font-medium text-gray-600">نام بانک</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">شماره حساب</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">کد شبا</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">شرح</th>
            <th class="text-right py-3 px-4 font-medium text-gray-600">عملیات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="account in bankAccounts" :key="account.id" class="border-b border-gray-100 hover:bg-gray-50">
            <td class="py-3 px-4">{{ account.bankName }}</td>
            <td class="py-3 px-4 font-mono">{{ account.accountNumber }}</td>
            <td class="py-3 px-4 font-mono">{{ account.sheba }}</td>
            <td class="py-3 px-4">{{ account.description }}</td>
            <td class="py-3 px-4">
              <button class="text-yellow-600 hover:text-yellow-800 mx-1" @click="editAccount(account)">ویرایش</button>
              <button class="text-red-600 hover:text-red-800 mx-1" @click="deleteAccount(account)">حذف</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <!-- مودال افزودن/ویرایش -->
    <div v-if="showAddModal || showEditModal" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl p-6 w-full max-w-md mx-4">
        <h5 class="font-bold mb-4">{{ showEditModal ? 'ویرایش حساب بانکی' : 'افزودن حساب بانکی' }}</h5>
        <form class="space-y-4" @submit.prevent="showEditModal ? updateAccount() : addAccount()">
          <div>
            <label class="block text-sm mb-1">نام بانک</label>
            <input v-model="form.bankName" required class="w-full px-3 py-2 border rounded-lg" />
          </div>
          <div>
            <label class="block text-sm mb-1">شماره حساب</label>
            <input v-model="form.accountNumber" required class="w-full px-3 py-2 border rounded-lg font-mono" />
          </div>
          <div>
            <label class="block text-sm mb-1">کد شبا</label>
            <input v-model="form.sheba" required class="w-full px-3 py-2 border rounded-lg font-mono" />
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

interface BankAccount {
  id?: number | string | null
  bankName?: string
  accountNumber?: string
  sheba?: string
  description?: string
  [key: string]: unknown
}

const bankAccounts = ref<BankAccount[]>([
  { id: 1, bankName: 'ملت', accountNumber: '1234567890', sheba: 'IR820540102680020817909002', description: 'حساب جاری شرکت' },
  { id: 2, bankName: 'ملی', accountNumber: '9876543210', sheba: 'IR120170000000123456789001', description: 'حساب پس‌انداز' }
])
const showAddModal = ref(false)
const showEditModal = ref(false)
const form = ref<BankAccount>({ id: null, bankName: '', accountNumber: '', sheba: '', description: '' })
const editAccount = (account: BankAccount) => {
  form.value = { ...account }
  showEditModal.value = true
}
const addAccount = () => {
  bankAccounts.value.push({ ...form.value, id: Date.now() })
  closeModal()
}
const updateAccount = () => {
  const idx = bankAccounts.value.findIndex(a => a.id === form.value.id)
  if (idx !== -1) bankAccounts.value[idx] = { ...form.value }
  closeModal()
}
const deleteAccount = (account: BankAccount) => {
  if (confirm('آیا از حذف این حساب اطمینان دارید؟'))
    bankAccounts.value = bankAccounts.value.filter(a => a.id !== account.id)
}
const closeModal = () => {
  showAddModal.value = false
  showEditModal.value = false
  form.value = { id: null, bankName: '', accountNumber: '', sheba: '', description: '' }
}
</script>
<!-- تعریف حساب‌های بانکی --> 
