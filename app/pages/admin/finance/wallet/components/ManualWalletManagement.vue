<template>
  <div class="space-y-6">
    <!-- هدر بخش -->
    <div class="bg-gradient-to-r from-yellow-50 to-orange-50 rounded-lg p-6">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">🔧 مدیریت دستی کیف پول</h2>
      <p class="text-gray-600">انجام عملیات دستی بر روی کیف پول کاربران (شارژ، کسر، مسدودسازی، انتقال)</p>
    </div>

    <!-- فرم شارژ دستی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">شارژ دستی کیف پول</h3>
      <form class="grid grid-cols-1 md:grid-cols-3 gap-6" @submit.prevent="onCredit">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه کاربر</label>
          <input v-model.number="credit.userId" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 12345">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان)</label>
          <input v-model.number="credit.amount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 500000">
        </div>
        <div class="flex items-end">
          <button type="submit" class="w-full px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500">شارژ</button>
        </div>
      </form>
    </div>

    <!-- فرم کسر دستی -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">کسر دستی از کیف پول</h3>
      <form class="grid grid-cols-1 md:grid-cols-3 gap-6" @submit.prevent="onDebit">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه کاربر</label>
          <input v-model.number="debit.userId" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 12345">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان)</label>
          <input v-model.number="debit.amount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 200000">
        </div>
        <div class="flex items-end">
          <button type="submit" class="w-full px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500">کسر</button>
        </div>
      </form>
    </div>

    <!-- فرم مسدود/آزادسازی کیف پول -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">مسدود/آزادسازی کیف پول</h3>
      <form class="grid grid-cols-1 md:grid-cols-3 gap-6" @submit.prevent="onBlockToggle">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه کاربر</label>
          <input v-model.number="block.userId" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 12345">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
          <select v-model="block.action" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="block">مسدودسازی</option>
            <option value="unblock">آزادسازی</option>
          </select>
        </div>
        <div class="flex items-end">
          <button type="submit" class="w-full px-4 py-2 bg-yellow-600 text-white rounded-lg hover:bg-yellow-700 focus:outline-none focus:ring-2 focus:ring-yellow-500">ثبت</button>
        </div>
      </form>
    </div>

    <!-- فرم انتقال بین کیف پول‌ها -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">انتقال بین کیف پول‌ها</h3>
      <form class="grid grid-cols-1 md:grid-cols-4 gap-6" @submit.prevent="onTransfer">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه مبدا</label>
          <input v-model.number="transfer.fromUserId" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 12345">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">شناسه مقصد</label>
          <input v-model.number="transfer.toUserId" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 67890">
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مبلغ (تومان)</label>
          <input v-model.number="transfer.amount" type="number" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="مثال: 100000">
        </div>
        <div class="flex items-end">
          <button type="submit" class="w-full px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">انتقال</button>
        </div>
      </form>
    </div>

    <!-- جدول عملیات دستی اخیر -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">جدول عملیات دستی اخیر</h3>
      <div class="overflow-x-auto">
        <table class="w-full text-sm text-right">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-4 py-3 text-gray-700 font-medium">نوع عملیات</th>
              <th class="px-4 py-3 text-gray-700 font-medium">شناسه کاربر/مبدا</th>
              <th class="px-4 py-3 text-gray-700 font-medium">شناسه مقصد</th>
              <th class="px-4 py-3 text-gray-700 font-medium">مبلغ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">تاریخ</th>
              <th class="px-4 py-3 text-gray-700 font-medium">وضعیت</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-200">
            <tr v-for="op in manualOps" :key="op.id" class="hover:bg-gray-50">
              <td class="px-4 py-3">{{ op.type }}</td>
              <td class="px-4 py-3">{{ op.from }}</td>
              <td class="px-4 py-3">{{ op.to || '-' }}</td>
              <td class="px-4 py-3">{{ formatCurrency(op.amount) }}</td>
              <td class="px-4 py-3">{{ op.date }}</td>
              <td class="px-4 py-3">
                <span :class="getStatusClass(op.status)" class="px-2 py-1 text-xs rounded-full">{{ op.status }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
const credit = reactive({ userId: undefined as unknown as number, amount: undefined as unknown as number })
const debit = reactive({ userId: undefined as unknown as number, amount: undefined as unknown as number })
const block = reactive({ userId: undefined as unknown as number, action: 'block' })
const transfer = reactive({ fromUserId: undefined as unknown as number, toUserId: undefined as unknown as number, amount: undefined as unknown as number })

async function onCredit() {
  if (!credit.userId || !credit.amount) return
  await $fetch('/api/admin/wallet/credit', { method: 'POST', body: { user_id: credit.userId, amount: credit.amount, method: 'manual' }, credentials: 'include' })
}
async function onDebit() {
  if (!debit.userId || !debit.amount) return
  await $fetch('/api/admin/wallet/debit', { method: 'POST', body: { user_id: debit.userId, amount: debit.amount, method: 'manual', reason: 'manual' }, credentials: 'include' })
}
async function onBlockToggle() {
  if (!block.userId) return
  const path = block.action === 'block' ? 'block' : 'unblock'
  await $fetch(`/api/admin/wallet/${path}/${block.userId}`, { method: 'POST', credentials: 'include' })
}
async function onTransfer() {
  if (!transfer.fromUserId || !transfer.toUserId || !transfer.amount) return
  await $fetch('/api/admin/wallet/transfer', { method: 'POST', body: { from_user_id: transfer.fromUserId, to_user_id: transfer.toUserId, amount: transfer.amount }, credentials: 'include' })
}
// داده‌های نمونه عملیات دستی
const manualOps = [
  { id: 1, type: 'شارژ دستی', from: '12345', to: '', amount: 500000, date: '1402/10/20', status: 'موفق' },
  { id: 2, type: 'کسر دستی', from: '67890', to: '', amount: 200000, date: '1402/10/19', status: 'موفق' },
  { id: 3, type: 'مسدودسازی', from: '11223', to: '', amount: 0, date: '1402/10/18', status: 'موفق' },
  { id: 4, type: 'آزادسازی', from: '11223', to: '', amount: 0, date: '1402/10/18', status: 'موفق' },
  { id: 5, type: 'انتقال', from: '12345', to: '67890', amount: 100000, date: '1402/10/17', status: 'موفق' }
]

// تابع فرمت کردن ارز
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان'
}

// تابع کلاس وضعیت
const getStatusClass = (status: string) => {
  const classes = {
    'موفق': 'bg-green-100 text-green-800',
    'ناموفق': 'bg-red-100 text-red-800',
    'در انتظار': 'bg-yellow-100 text-yellow-800'
  }
  return classes[status] || 'bg-gray-100 text-gray-800'
}
</script>

<!--
  مستندسازی:
  این کامپوننت شامل فرم‌های شارژ دستی، کسر دستی، مسدود/آزادسازی و انتقال بین کیف پول‌ها و جدول عملیات اخیر است.
  تمام توضیحات به فارسی و با طراحی ریسپانسیو ارائه شده است.
--> 
