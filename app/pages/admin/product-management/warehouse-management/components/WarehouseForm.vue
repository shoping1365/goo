<template>
  <div class="rounded-xl p-6 mb-4 bg-white shadow">
    <div class="grid grid-cols-1 md:grid-cols-5 gap-3">
      <input v-model="model.code" type="text" placeholder="کد" class="rounded px-3 py-2 text-sm bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200 w-full"/>
      <input v-model="model.name" type="text" placeholder="نام" class="rounded px-3 py-2 text-sm bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200 w-full"/>
      <input v-model="model.city" type="text" placeholder="شهر" class="rounded px-3 py-2 text-sm bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200 w-full"/>
      <input v-model="model.address" type="text" placeholder="آدرس" class="rounded px-3 py-2 text-sm bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200 w-full"/>
      <input v-model.number="model.priority" type="number" min="1" placeholder="اولویت" class="rounded px-3 py-2 text-sm bg-white shadow-sm border border-transparent focus:border-blue-300 focus:ring-2 focus:ring-blue-200 w-full"/>
    </div>
    <div class="flex flex-wrap items-center gap-6 mt-3 text-sm">
      <label class="flex items-center gap-2"><input v-model="model.is_default" type="checkbox"/> انبار پیش‌فرض</label>
      <label class="flex items-center gap-2"><input v-model="model.is_active" type="checkbox"/> فعال</label>
      <span class="text-xs text-gray-500">انتخاب انبار پیش‌فرض، سایر انبارها را از پیش‌فرض بودن خارج می‌کند.</span>
    </div>
  </div>
  
</template>

<script setup lang="ts">
const model = defineModel({
  type: Object,
  default: () => ({ code: '', name: '', city: '', address: '', is_default: false, is_active: true, priority: 100 })
})
</script>


