<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold">کرالر سرویس ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید وضعیت و تنظیمات کرالر ترب را مدیریت کنید.</p>
    <!-- فرم مدیریت کرالر ترب -->
    <form class="space-y-3">
      <div>
        <label class="block mb-1">وضعیت کرالر</label>
        <select class="w-full border rounded px-3 py-2">
          <option>فعال</option>
          <option>غیرفعال</option>
        </select>
      </div>
      <div>
        <label class="block mb-1">زمان‌بندی اجرا</label>
        <input type="text" class="w-full border rounded px-3 py-2" placeholder="مثال: هر روز ساعت ۲" />
      </div>
      <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">ذخیره تنظیمات</button>
    </form>
  </div>
</template>

<script setup lang="ts">
// این کامپوننت کرالر سرویس ترب را مدیریت می‌کند
</script>

<!-- توضیحات: این کامپوننت برای مدیریت کرالر ترب استفاده می‌شود. --> 
