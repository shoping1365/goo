<template>
  <div class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">جدول فروش‌ها</div>
    <table class="w-full text-sm">
      <thead>
        <tr class="bg-gray-50">
          <th class="p-2">شماره سفارش</th>
          <th class="p-2">تاریخ</th>
          <th class="p-2">مشتری</th>
          <th class="p-2">محصولات</th>
          <th class="p-2">مبلغ</th>
          <th class="p-2">وضعیت</th>
          <th class="p-2">روش پرداخت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in 10" :key="i" class="border-b">
          <td class="p-2 font-mono">#{{1000 + i}}</td>
          <td class="p-2">1403/0{{Math.ceil(i/3)}}/{{10 + i}}</td>
          <td class="p-2">مشتری {{i}}</td>
          <td class="p-2">{{Math.ceil(i/2)}} قلم</td>
          <td class="p-2 font-bold text-green-600">{{(50000 + i*10000).toLocaleString('fa-IR')}}</td>
          <td class="p-2">
            <span :class="i % 3 === 0 ? 'bg-green-100 text-green-700' : i % 3 === 1 ? 'bg-yellow-100 text-yellow-700' : 'bg-red-100 text-red-700'" class="px-2 py-1 rounded text-xs">
              {{ i % 3 === 0 ? 'تکمیل' : i % 3 === 1 ? 'در انتظار' : 'لغو شده' }}
            </span>
          </td>
          <td class="p-2">{{ i % 2 === 0 ? 'آنلاین' : 'حضوری' }}</td>
        </tr>
      </tbody>
    </table>
    <div class="flex justify-end mt-4">
      <button class="px-3 py-1 rounded bg-gray-100 mx-1">قبلی</button>
      <button class="px-3 py-1 rounded bg-green-500 text-white mx-1">1</button>
      <button class="px-3 py-1 rounded bg-gray-100 mx-1">2</button>
      <button class="px-3 py-1 rounded bg-gray-100 mx-1">بعدی</button>
    </div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ sales: unknown[] }>()
</script> 
