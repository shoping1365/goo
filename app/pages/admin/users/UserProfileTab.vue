<template>
  <div class="bg-white rounded-xl shadow-lg p-2 md:p-6 flex items-center justify-center gap-6 relative">
    <!-- Edit button -->
    <button
      class="absolute right-4 top-6 bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-full p-2 transition"
      title="ویرایش پروفایل"
      @click="$emit('edit')"
    >
      <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536M9 13l6.586-6.586a2 2 0 112.828 2.828L11.828 15.828a2 2 0 01-2.828 0L9 13zm-6 6h12" />
      </svg>
    </button>
    <!-- Profile info row -->
    <div class="flex flex-row items-center gap-6 flex-1 min-w-0 flex-wrap justify-center">
      <!-- Name, status, username -->
      <div class="flex flex-col items-end gap-1 min-w-[120px]">
        <div class="font-bold text-lg">{{ user.name }}</div>
        <div class="flex items-center gap-2">
          <span :class="user.online ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-500'" class="px-2 py-0.5 rounded-full text-xs">
            {{ user.online ? 'آنلاین' : 'آفلاین' }}
          </span>
          <span class="text-xs text-gray-400">{{ user.username }}</span>
        </div>
      </div>
      <!-- Info fields -->
      <div class="flex flex-row flex-wrap gap-x-8 gap-y-1 text-sm flex-1 justify-center">
        <div class="flex items-center gap-1">
          <span class="font-semibold">ایمیل:</span>
          <span>{{ user.email }}</span>
        </div>
        <div class="flex items-center gap-1">
          <span class="font-semibold">شماره همراه:</span>
          <span>{{ user.mobile }}</span>
        </div>
        <div class="flex items-center gap-1">
          <span class="font-semibold">شماره ثابت:</span>
          <span>{{ user.landline }}</span>
        </div>
        <div class="flex items-center gap-1">
          <span class="font-semibold">کد ملی:</span>
          <span>{{ user.nationalCode }}</span>
        </div>
        <div class="flex items-center gap-1">
          <span class="font-semibold">آدرس منتخب:</span>
          <span>{{ user.selectedAddress }}</span>
        </div>
      </div>
      <!-- Avatar -->
      <div class="w-16 h-16 rounded-full bg-blue-100 flex items-center justify-center text-2xl text-blue-600 ml-2">
        <span>{{ user.name?.charAt(0) || '?' }}</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { User } from '~/types/user';

defineProps<{ user: User }>();
defineEmits(['edit']);
</script> 
