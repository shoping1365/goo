<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-cyan-700 mb-2 flex justify-between items-center">
      <span>پیام‌های کاربر</span>
      <button class="bg-cyan-100 hover:bg-cyan-200 text-cyan-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-messages')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">عنوان</th>
          <th class="p-2 text-right font-bold">فرستنده</th>
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
          <th class="p-2 text-right font-bold">عملیات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="message in messages" :key="message.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ message.title }}</td>
          <td class="p-2 text-right">{{ message.sender }}</td>
          <td class="p-2 text-right">{{ message.date }}</td>
          <td class="p-2 text-right">
            <span :class="message.isRead ? 'text-gray-600' : 'text-blue-600'">{{ message.isRead ? 'خوانده شده' : 'خوانده نشده' }}</span>
          </td>
          <td class="p-2 text-right">
            <button class="text-cyan-600 hover:text-cyan-700" @click="$emit('view-message', message)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="mt-4 flex justify-end">
      <button class="bg-cyan-600 hover:bg-cyan-700 text-white rounded-lg px-4 py-2 text-sm flex items-center gap-2 transition font-bold" @click="$emit('send-new-message')">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"/>
        </svg>
        ارسال پیام جدید
      </button>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { User } from '~/types/user';

defineProps<{ user: User }>();
defineEmits(['view-all-messages', 'view-message', 'send-new-message']);
// Mock data for messages
const messages = [
  { id: 1, title: 'خوش آمدگویی', sender: 'سیستم', date: '1402/03/10', isRead: true },
  { id: 2, title: 'تایید سفارش', sender: 'پشتیبانی', date: '1402/03/09', isRead: false },
  { id: 3, title: 'پیام خصوصی', sender: 'مدیر سیستم', date: '1402/03/08', isRead: true },
];
</script> 
