<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-blue-700 mb-2 flex justify-between items-center">
      <span>فعالیت کاربر</span>
      <button v-if="visitedPages.length > 5" class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">صفحه</th>
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">مدت زمان</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="page in displayedPages" :key="page.url + page.date" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ page.title }}</td>
          <td class="p-2 text-right">{{ page.date }}</td>
          <td class="p-2 text-right">{{ page.duration }}</td>
        </tr>
      </tbody>
    </table>
    <div class="mt-4 grid grid-cols-1 md:grid-cols-3 gap-2 text-xs text-gray-600">
      <div><b>مدت حضور کل:</b> <span class="text-blue-700">{{ totalTime }}</span></div>
      <div><b>ساعات اوج فعالیت:</b> <span class="text-blue-700">{{ peakHours }}</span></div>
      <div><b>نوع دستگاه‌ها:</b> <span class="text-blue-700">{{ devices.join(', ') }}</span></div>
    </div>
    <div class="mt-4">
      <div class="font-semibold text-gray-700 mb-1">مسیر حرکت کاربر (User Flow)</div>
      <ol class="list-decimal list-inside text-xs text-gray-600">
        <li v-for="step in userFlow" :key="step">{{ step }}</li>
      </ol>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue';
import type { User } from '~/types/user';

defineProps<{ user: User }>();
const showAll = ref(false);

// Mock data for activity
const totalTime = '2 ساعت و 35 دقیقه';
const peakHours = '10:00 تا 13:00';
const devices = ['موبایل', 'دسکتاپ'];
const visitedPages = [
  { title: 'صفحه محصول: لپ‌تاپ', url: '/product/1', date: '1402/03/10', duration: '12 دقیقه' },
  { title: 'سبد خرید', url: '/cart', date: '1402/03/10', duration: '3 دقیقه' },
  { title: 'پروفایل', url: '/profile', date: '1402/03/09', duration: '7 دقیقه' },
  { title: 'صفحه محصول: گوشی', url: '/product/2', date: '1402/03/09', duration: '5 دقیقه' },
  { title: 'صفحه محصول: تبلت', url: '/product/3', date: '1402/03/08', duration: '8 دقیقه' },
  { title: 'صفحه محصول: هدفون', url: '/product/4', date: '1402/03/08', duration: '4 دقیقه' },
  { title: 'صفحه محصول: شارژر', url: '/product/5', date: '1402/03/07', duration: '2 دقیقه' },
];

const displayedPages = computed(() => {
  return showAll.value ? visitedPages : visitedPages.slice(0, 5);
});

const userFlow = [
  'صفحه اصلی',
  'جستجو: لپ‌تاپ',
  'صفحه محصول: لپ‌تاپ',
  'افزودن به سبد خرید',
  'سبد خرید',
  'تکمیل خرید',
];
</script> 
