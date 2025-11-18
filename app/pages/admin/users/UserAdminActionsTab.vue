<template>
  <div>
    <div class="font-bold mb-2 flex items-center gap-2">
      <svg class="w-5 h-5 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17v-2a4 4 0 014-4h4a4 4 0 014 4v2" /></svg>
      عملیات ادمین
    </div>
    <div class="flex flex-row gap-6">
      <button
        class="flex-1 flex items-center gap-2 p-6 rounded-xl shadow bg-red-50 hover:bg-red-100 transition"
        @click="handleAction('block')"
      >
        <span class="bg-red-200 text-red-700 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-12.728 12.728M5.636 5.636l12.728 12.728" /></svg>
        </span>
        <span class="font-bold text-red-700">مسدودسازی کاربر</span>
      </button>
      <button
        class="flex-1 flex items-center gap-2 p-6 rounded-xl shadow bg-blue-50 hover:bg-blue-100 transition"
        @click="handleAction('resetPassword')"
      >
        <span class="bg-blue-200 text-blue-700 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11c0-2.21 1.79-4 4-4s4 1.79 4 4-1.79 4-4 4-4-1.79-4-4zm-8 8v-1a4 4 0 014-4h4a4 4 0 014 4v1" /></svg>
        </span>
        <span class="font-bold text-blue-700">ریست رمز عبور</span>
      </button>
      <button
        class="flex-1 flex items-center gap-2 p-6 rounded-xl shadow bg-purple-50 hover:bg-purple-100 transition"
        @click="handleAction('changeRole')"
      >
        <span class="bg-purple-200 text-purple-700 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.5 3.5a2.121 2.121 0 113 3L7 19.5 3 21l1.5-4L16.5 3.5z" /></svg>
        </span>
        <span class="font-bold text-purple-700">تغییر نقش</span>
      </button>
      <button
        class="flex-1 flex items-center gap-2 p-6 rounded-xl shadow bg-orange-50 hover:bg-orange-100 transition"
        @click="handleAction('forceLogout')"
      >
        <span class="bg-orange-200 text-orange-700 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7" /></svg>
        </span>
        <span class="font-bold text-orange-700">خروج اجباری</span>
      </button>
      <button
        class="flex-1 flex items-center gap-2 p-6 rounded-xl shadow bg-yellow-50 hover:bg-yellow-100 transition"
        @click="handleAction('sendWarning')"
      >
        <span class="bg-yellow-200 text-yellow-700 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01" /></svg>
        </span>
        <span class="font-bold text-yellow-700">ارسال هشدار</span>
      </button>
    </div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ user: any }>()

async function post(url: string, body?: any) {
  try {
    await $fetch(url, { method: 'POST', body, credentials: 'include' })
    // optionally toast success
  } catch (e: any) {
    // optionally toast error
  }
}

async function handleAction(action: string) {
  const id = props.user?.id
  if (!id) return
  if (action === 'block') return post(`/api/users/${id}/block`, {})
  if (action === 'forceLogout') return post(`/api/users/${id}/unblock`, {}) // موقت: به‌جای خروج اجباری، رفع مسدودی
  if (action === 'changeRole') {
    const newRole = prompt('Role ID جدید را وارد کنید', '1')
    if (!newRole) return
    return post(`/api/users/${id}/role`, { role_id: Number(newRole) })
  }
  if (action === 'resetPassword') {
    // جای‌خالی: endpoint اختصاصی بعداً اضافه می‌شود
    return
  }
  if (action === 'sendWarning') {
    // جای‌خالی: استفاده از user_notifications بعداً اضافه می‌شود
    return
  }
}
</script> 
