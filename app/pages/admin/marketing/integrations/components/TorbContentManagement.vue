<template>
  <div class="space-y-4">
    <h2 class="text-xl font-bold">مدیریت محتوای ترب</h2>
    <p class="text-gray-600">در این بخش می‌توانید محتوای محصولات و اطلاعات ارسالی به ترب را مدیریت کنید.</p>
    <!-- فرم مدیریت محتوا -->
    <form class="space-y-3">
      <div>
        <label class="block mb-1">عنوان محصول</label>
        <input type="text" class="w-full border rounded px-3 py-2" placeholder="عنوان محصول" />
      </div>
      <div>
        <label class="block mb-1">توضیحات</label>
        <textarea class="w-full border rounded px-3 py-2" placeholder="توضیحات محصول"></textarea>
      </div>
      <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">ذخیره محتوا</button>
    </form>
  </div>
</template>

<script setup lang="ts">
// این کامپوننت مدیریت محتوای سرویس ترب را انجام می‌دهد
</script>

<!-- توضیحات: این کامپوننت برای مدیریت محتوای محصولات ترب استفاده می‌شود. --> 
