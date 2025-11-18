<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold">تنظیمات سرویس ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید تنظیمات مربوط به اتصال و مدیریت سرویس ترب را انجام دهید.</p>
    <!-- فرم تنظیمات ترب -->
    <form class="space-y-3">
      <div>
        <label class="block mb-1">کلید API ترب</label>
        <input type="text" class="w-full border rounded px-3 py-2" placeholder="API Key" />
      </div>
      <div>
        <label class="block mb-1">وضعیت فعال بودن سرویس</label>
        <select class="w-full border rounded px-3 py-2">
          <option>فعال</option>
          <option>غیرفعال</option>
        </select>
      </div>
      <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">ذخیره تنظیمات</button>
    </form>
  </div>
</template>

<script setup lang="ts">
// این کامپوننت تنظیمات سرویس ترب را مدیریت می‌کند
</script>

<!-- توضیحات: این کامپوننت برای مدیریت تنظیمات سرویس ترب استفاده می‌شود. --> 
