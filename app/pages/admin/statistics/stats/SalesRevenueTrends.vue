<template>
  <div class="bg-white rounded-xl shadow p-6">
    <div class="font-bold text-gray-700 mb-4">روند درآمد</div>
    <div class="space-y-4">
      <div class="flex justify-between items-center p-3 bg-green-50 rounded-lg">
        <div>
          <div class="font-semibold text-green-700">درآمد این ماه</div>
          <div class="text-sm text-gray-600">نسبت به ماه قبل</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-green-600">+12.5%</div>
          <div class="text-xs text-green-500">رشد</div>
        </div>
      </div>
      <div class="flex justify-between items-center p-3 bg-blue-50 rounded-lg">
        <div>
          <div class="font-semibold text-blue-700">درآمد این هفته</div>
          <div class="text-sm text-gray-600">نسبت به هفته قبل</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-blue-600">+8.3%</div>
          <div class="text-xs text-blue-500">رشد</div>
        </div>
      </div>
      <div class="flex justify-between items-center p-3 bg-orange-50 rounded-lg">
        <div>
          <div class="font-semibold text-orange-700">میانگین روزانه</div>
          <div class="text-sm text-gray-600">30 روز گذشته</div>
        </div>
        <div class="text-right">
          <div class="font-bold text-orange-600">29.7K</div>
          <div class="text-xs text-orange-500">تومان</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ trends: any[] }>()
</script> 
