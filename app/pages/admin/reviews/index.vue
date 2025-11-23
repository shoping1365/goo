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
        <button class="px-4 py-2 bg-blue-600 text-white rounded-md" @click="fetchReviews">اعمال فیلتر</button>
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
                <button v-if="item.status==='pending'" class="px-3 py-1 text-sm bg-green-600 text-white rounded" @click="approveReview(item)">تایید</button>
                <button v-if="item.status!=='rejected'" class="px-3 py-1 text-sm bg-red-600 text-white rounded" @click="rejectReview(item)">رد</button>
                <button class="px-3 py-1 text-sm bg-blue-600 text-white rounded" @click="openReplyForm(item)">پاسخ</button>
                <button class="px-3 py-1 text-sm bg-gray-600 text-white rounded" @click="requestDeleteConfirmation(item)">حذف</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="!items.length" class="p-6 text-center text-gray-500">موردی یافت نشد</div>
    </div>

    <div v-if="replyContext.review" class="mt-6 border rounded-lg bg-white p-6">
      <div class="flex items-center justify-between mb-4">
        <div>
          <h2 class="text-lg font-semibold text-gray-900">ارسال پاسخ مدیر</h2>
          <p class="text-sm text-gray-500">نظر انتخاب شده: {{ replyContext.review.product?.name || replyContext.review.id }}</p>
        </div>
        <button class="text-sm text-gray-500" @click="resetReplyForm">بستن</button>
      </div>
      <textarea v-model="replyContext.content" rows="4" class="w-full border rounded-md p-3" placeholder="متن پاسخ را وارد کنید"></textarea>
      <div class="flex gap-3 mt-4">
        <button class="px-4 py-2 rounded-md bg-blue-600 text-white disabled:opacity-50" :disabled="isReplySubmitDisabled" @click="submitReply">
          ارسال پاسخ
        </button>
        <button class="px-4 py-2 rounded-md bg-gray-100 text-gray-700" @click="resetReplyForm">انصراف</button>
      </div>
    </div>

    <div v-if="deleteContext.review" class="mt-6 border rounded-lg bg-white p-6">
      <h2 class="text-lg font-semibold text-gray-900">تایید حذف نظر</h2>
      <p class="text-sm text-gray-600 mt-2">آیا از حذف نظر مربوط به "{{ deleteContext.review.product?.name || deleteContext.review.id }}" مطمئن هستید؟</p>
      <div class="flex gap-3 mt-4">
        <button class="px-4 py-2 rounded-md bg-red-600 text-white disabled:opacity-50" :disabled="deleteContext.isSubmitting" @click="confirmDelete">حذف</button>
        <button class="px-4 py-2 rounded-md bg-gray-100 text-gray-700" @click="resetDeleteContext">لغو</button>
      </div>
    </div>
  </div>
  
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

interface ReviewProduct {
  id?: number | string
  name?: string
  image?: string
  thumbnail?: string
}

interface ReviewCustomer {
  id?: number | string
  name?: string
}

interface Review {
  id?: number | string
  status?: string
  rating?: number
  created_at?: string
  admin_reply?: string
  product?: ReviewProduct
  customer?: ReviewCustomer
  [key: string]: unknown
}

interface ReviewFilters {
  status: string | number | ''
  rating: string | number | ''
  search: string
}

interface ReviewsResponse {
  reviews?: Review[]
  [key: string]: unknown
}

interface ReplyContextState {
  review: Review | null
  content: string
  isSubmitting: boolean
}

interface DeleteContextState {
  review: Review | null
  isSubmitting: boolean
}

const items = ref<Review[]>([])
const filters = ref<ReviewFilters>({ status: '', rating: '', search: '' })
const replyContext = ref<ReplyContextState>({ review: null, content: '', isSubmitting: false })
const deleteContext = ref<DeleteContextState>({ review: null, isSubmitting: false })

/**
 * این تابع تمام فیلترهای اعمال‌شده را به کوئری تبدیل کرده و فهرست نظرات را از سرور دریافت می‌کند.
 */
const fetchReviews = async () => {
  const params = new URLSearchParams()
  if (filters.value.status) params.append('status', String(filters.value.status))
  if (filters.value.rating) params.append('rating', String(filters.value.rating))
  if (filters.value.search) params.append('search', String(filters.value.search))
  const res = await $fetch<ReviewsResponse>('/api/admin/reviews?' + params.toString())
  items.value = res?.reviews || []
}

/**
 * این تابع یک نظر را تایید کرده و پس از موفقیت، فهرست به‌روزرسانی می‌شود.
 */
const approveReview = async (row: Review) => {
  await $fetch(`/api/admin/reviews/${row.id}/approve`, { method: 'POST' })
  await fetchReviews()
}

/**
 * این تابع یک نظر را با وضعیت رد به‌روز می‌کند.
 */
const rejectReview = async (row: Review) => {
  await $fetch(`/api/admin/reviews/${row.id}/reject`, { method: 'POST', body: { reason: '' } })
  await fetchReviews()
}

/**
 * این تابع زمینه‌ی لازم برای نوشتن پاسخ را فعال می‌کند.
 */
const openReplyForm = (row: Review) => {
  replyContext.value = {
    review: row,
    content: typeof row.admin_reply === 'string' ? row.admin_reply : '',
    isSubmitting: false
  }
}

/**
 * این تابع فرم پاسخ را ریست می‌کند.
 */
const resetReplyForm = () => {
  replyContext.value = { review: null, content: '', isSubmitting: false }
}

/**
 * این تابع پاسخ مدیر را برای نظر انتخاب شده ارسال می‌کند.
 */
const submitReply = async () => {
  if (!replyContext.value.review) return
  const content = replyContext.value.content.trim()
  if (!content) return
  replyContext.value.isSubmitting = true
  try {
    await $fetch(`/api/admin/reviews/${replyContext.value.review.id}/reply`, {
      method: 'POST',
      body: { reply: content }
    })
    await fetchReviews()
    resetReplyForm()
  } finally {
    replyContext.value.isSubmitting = false
  }
}

/**
 * این تابع درخواست حذف را برای تایید ثانویه ثبت می‌کند.
 */
const requestDeleteConfirmation = (row: Review) => {
  deleteContext.value = { review: row, isSubmitting: false }
}

/**
 * این تابع درخواست حذف را لغو می‌کند.
 */
const resetDeleteContext = () => {
  deleteContext.value = { review: null, isSubmitting: false }
}

/**
 * این تابع حذف قطعی نظر را انجام می‌دهد.
 */
const confirmDelete = async () => {
  if (!deleteContext.value.review) return
  deleteContext.value.isSubmitting = true
  try {
    await $fetch(`/api/admin/reviews/${deleteContext.value.review.id}`, { method: 'DELETE' })
    await fetchReviews()
    resetDeleteContext()
  } finally {
    deleteContext.value.isSubmitting = false
  }
}

const isReplySubmitDisabled = computed(() => {
  return replyContext.value.isSubmitting || !replyContext.value.content.trim()
})

/**
 * این تابع تاریخ ذخیره شده را به فرمت شمسی نمایش می‌دهد.
 */
const formatDate = (d: string) => new Date(d).toLocaleDateString('fa-IR')

/**
 * این تابع کلاس‌های نمایشی وضعیت را بازمی‌گرداند.
 */
const statusClass = (s: string) => s === 'approved' ? 'bg-green-100 text-green-700' : s === 'rejected' ? 'bg-red-100 text-red-700' : 'bg-yellow-100 text-yellow-700'

/**
 * این تابع وضعیت قابل خواندن برای مدیر را تامین می‌کند.
 */
const humanStatus = (s: string) => s === 'approved' ? 'تایید شده' : s === 'rejected' ? 'رد شده' : 'در انتظار'

onMounted(fetchReviews)
</script>


