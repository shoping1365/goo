<template>
  <div class="bg-white rounded-xl shadow p-6 flex flex-col gap-6">
    <div class="flex items-center gap-2 mb-2">
      <span class="inline-block w-8 h-8 rounded-lg bg-blue-100 flex items-center justify-center text-xl text-blue-600">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V4a2 2 0 10-4 0v1.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" /></svg>
      </span>
      <span class="font-bold text-blue-700 text-base">فعالیت کاربر</span>
    </div>
    <div class="flex flex-wrap gap-6 text-sm">
      <div><b>مدت حضور کل:</b> <span class="text-blue-700">{{ totalTime }}</span></div>
      <div><b>ساعات اوج فعالیت:</b> <span class="text-blue-700">{{ peakHours }}</span></div>
      <div><b>نوع دستگاه‌ها:</b> <span class="text-blue-700">{{ devices.join(', ') }}</span></div>
    </div>
    <div>
      <div class="font-semibold text-gray-700 mb-1">صفحات بازدید شده اخیر</div>
      <table class="w-full text-xs bg-gray-50 rounded-lg overflow-hidden">
        <thead><tr class="bg-gray-100"><th class="p-1">صفحه</th><th class="p-1">تاریخ</th><th class="p-1">مدت زمان</th></tr></thead>
        <tbody>
          <tr v-for="page in visitedPages" :key="page.url + page.date" class="border-b">
            <td class="p-1">{{ page.title }}</td>
            <td class="p-1">{{ page.date }}</td>
            <td class="p-1">{{ page.duration }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div>
      <div class="font-semibold text-gray-700 mb-1">مسیر حرکت کاربر (User Flow)</div>
      <ol class="list-decimal list-inside text-xs text-gray-600">
        <li v-for="step in userFlow" :key="step">{{ step }}</li>
      </ol>
    </div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ user: any }>();
// Mock data for activity
const totalTime = '2 ساعت و 35 دقیقه';
const peakHours = '10:00 تا 13:00';
const devices = ['موبایل', 'دسکتاپ'];
const visitedPages = [
  { title: 'صفحه محصول: لپ‌تاپ', url: '/product/1', date: '1402/03/10', duration: '12 دقیقه' },
  { title: 'سبد خرید', url: '/cart', date: '1402/03/10', duration: '3 دقیقه' },
  { title: 'پروفایل', url: '/profile', date: '1402/03/09', duration: '7 دقیقه' },
];
const userFlow = [
  'صفحه اصلی',
  'جستجو: لپ‌تاپ',
  'صفحه محصول: لپ‌تاپ',
  'افزودن به سبد خرید',
  'سبد خرید',
  'تکمیل خرید',
];
</script> 
