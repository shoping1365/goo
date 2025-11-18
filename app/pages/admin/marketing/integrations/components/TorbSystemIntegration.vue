<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold">یکپارچه‌سازی سیستم ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید تنظیمات و وضعیت یکپارچه‌سازی سرویس ترب با سایر سیستم‌ها را مدیریت کنید.</p>
    <!-- فرم یکپارچه‌سازی -->
    <form class="space-y-3">
      <div>
        <label class="block mb-1">انتخاب سیستم مقصد</label>
        <select class="w-full border rounded px-3 py-2">
          <option>سیستم سفارشات</option>
          <option>سیستم کاربران</option>
          <option>سیستم گزارش‌گیری</option>
        </select>
      </div>
      <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">ذخیره تنظیمات</button>
    </form>
  </div>
</template>

<script setup lang="ts">
// این کامپوننت یکپارچه‌سازی سیستم ترب را مدیریت می‌کند
</script>

<!-- توضیحات: این کامپوننت برای مدیریت یکپارچه‌سازی سیستم ترب استفاده می‌شود. --> 
