<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold">گزارشات سرویس ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید گزارشات مربوط به سرویس ترب را مشاهده کنید.</p>
    <!-- جدول گزارشات ترب -->
    <table class="w-full border rounded">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2">تاریخ</th>
          <th class="p-2">نوع گزارش</th>
          <th class="p-2">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="p-2">1403/02/01</td>
          <td class="p-2">گزارش قیمت</td>
          <td class="p-2 text-green-600">موفق</td>
        </tr>
        <tr>
          <td class="p-2">1403/02/02</td>
          <td class="p-2">گزارش موجودی</td>
          <td class="p-2 text-red-600">خطا</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
// این کامپوننت گزارشات سرویس ترب را نمایش می‌دهد
</script>

<!-- توضیحات: این کامپوننت برای نمایش گزارشات سرویس ترب استفاده می‌شود. --> 
