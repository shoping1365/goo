<template>
  <div class="p-6" dir="rtl">
    <div class="flex items-center justify-between mb-4">
      <h1 class="text-xl font-bold text-gray-900">مدیریت نظرات محصولات</h1>
      <div class="flex gap-3">
        <input v-model="filters.search" type="text" placeholder="جستجو در نظرها" class="px-3 py-2 border rounded-md" />
        <select v-model="filters.status" class="px-3 py-2 border rounded-md">
          <option value="">همه وضعیت‌ها</option>
          <option value="pending">در انتظار</option>
          <option value="approved">تایید شده</option>
          <option value="rejected">رد شده</option>
        </select>
        <select v-model="filters.rating" class="px-3 py-2 border rounded-md">
          <option value="">همه امتیازها</option>
          <option v-for="r in [5,4,3,2,1]" :key="r" :value="r">{{ r }} ستاره</option>
        </select>
        <button @click="fetchReviews" class="px-4 py-2 bg-blue-600 text-white rounded-md">اعمال فیلتر</button>
      </div>
    </div>

    <div class="bg-white border rounded-lg overflow-hidden">
      <table class="min-w-full">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">محصول</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">کاربر</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">امتیاز</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">وضعیت</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">تاریخ</th>
            <th class="px-4 py-3 text-right text-sm font-medium text-gray-600">اقدامات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.id" class="border-t">
            <td class="px-4 py-3">
              <div class="flex items-center gap-3">
                <img :src="item.product?.thumbnail || item.product?.image" alt="thumb" class="w-10 h-10 rounded object-cover" />
                <div>
                  <div class="font-medium text-gray-900">{{ item.product?.name }}</div>
                  <div class="text-xs text-gray-500">#{{ item.product?.id }}</div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3">
              <div class="font-medium">{{ item.customer?.name }}</div>
              <div class="text-xs text-gray-500">ID: {{ item.customer?.id }}</div>
            </td>
            <td class="px-4 py-3">
              <div class="flex items-center gap-1 text-yellow-500">
                <span v-for="i in 5" :key="i">
                  <svg :class="['w-4 h-4', i <= item.rating ? 'fill-current' : 'fill-gray-200']" viewBox="0 0 24 24"><path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/></svg>
                </span>
              </div>
            </td>
            <td class="px-4 py-3">
              <span :class="statusClass(item.status)" class="px-2 py-1 text-xs rounded">{{ humanStatus(item.status) }}</span>
            </td>
            <td class="px-4 py-3 text-sm text-gray-600">{{ formatDate(item.created_at) }}</td>
            <td class="px-4 py-3">
              <div class="flex flex-wrap gap-2">
                <button v-if="item.status==='pending'" @click="approve(item)" class="px-3 py-1 text-sm bg-green-600 text-white rounded">تایید</button>
                <button v-if="item.status!=='rejected'" @click="reject(item)" class="px-3 py-1 text-sm bg-red-600 text-white rounded">رد</button>
                <button @click="reply(item)" class="px-3 py-1 text-sm bg-blue-600 text-white rounded">پاسخ</button>
                <button @click="remove(item)" class="px-3 py-1 text-sm bg-gray-600 text-white rounded">حذف</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="!items.length" class="p-6 text-center text-gray-500">موردی یافت نشد</div>
    </div>
  </div>
  
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main' })

const items = ref<any[]>([])
const filters = ref<{ status: string | number | ''; rating: string | number | ''; search: string }>({ status: '', rating: '', search: '' })

const fetchReviews = async () => {
  const params = new URLSearchParams()
  if (filters.value.status) params.append('status', String(filters.value.status))
  if (filters.value.rating) params.append('rating', String(filters.value.rating))
  if (filters.value.search) params.append('search', String(filters.value.search))
  const res = await $fetch('/api/admin/reviews?' + params.toString()) as any
  items.value = res?.reviews || []
}

const approve = async (row: any) => {
  await $fetch(`/api/admin/reviews/${row.id}/approve`, { method: 'POST' })
  fetchReviews()
}

const reject = async (row: any) => {
  await $fetch(`/api/admin/reviews/${row.id}/reject`, { method: 'POST', body: { reason: '' } })
  fetchReviews()
}

const reply = async (row: any) => {
  const content = prompt('پاسخ مدیر:', row.admin_reply || '')
  if (content == null) return
  await $fetch(`/api/admin/reviews/${row.id}/reply`, { method: 'POST', body: { reply: content } })
  fetchReviews()
}

const remove = async (row: any) => {
  if (!confirm('حذف این نظر؟')) return
  await $fetch(`/api/admin/reviews/${row.id}`, { method: 'DELETE' })
  fetchReviews()
}

const formatDate = (d: string) => new Date(d).toLocaleDateString('fa-IR')
const statusClass = (s: string) => s === 'approved' ? 'bg-green-100 text-green-700' : s === 'rejected' ? 'bg-red-100 text-red-700' : 'bg-yellow-100 text-yellow-700'
const humanStatus = (s: string) => s === 'approved' ? 'تایید شده' : s === 'rejected' ? 'رد شده' : 'در انتظار'

onMounted(fetchReviews)
</script>


