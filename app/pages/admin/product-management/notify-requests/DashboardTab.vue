<template>
  <!-- جداول اصلی -->
  <div v-if="props.activeNavItem === 'dashboard'" class="grid grid-cols-1 lg:grid-cols-2 gap-6">
<!-- آخرین درخواست‌ها -->
<div class="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
  <div class="bg-gray-50 px-6 py-4 border-b border-gray-100">
    <h3 class="text-lg font-semibold text-gray-900">آخرین درخواست‌ها</h3>
  </div>
  <div class="p-0">
    <table class="w-full">
      <thead class="bg-blue-100">
      <tr>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">محصول</th>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">نوع</th>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">وضعیت</th>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">تاریخ</th>
      </tr>
      </thead>
      <tbody class="divide-y divide-gray-100">
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">ذره بین چراغ قوه در شیائومی مدل XIAODA XD-FDJ01</td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-blue-100 text-blue-800 rounded-full">موجودی</span>
        </td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-orange-100 text-orange-800 rounded-full">در انتظار</span>
        </td>
        <td class="py-3 px-4 text-sm text-gray-600">
          <div>1404/03/14</div>
          <div class="text-xs text-gray-400">01:43</div>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">محلول کن شیائومی شیائومی Deerma NU05 - اصالت و سلامت فیزیکی کالا</td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-blue-100 text-blue-800 rounded-full">موجودی</span>
        </td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-orange-100 text-orange-800 rounded-full">در انتظار</span>
        </td>
        <td class="py-3 px-4 text-sm text-gray-600">
          <div>1404/03/14</div>
          <div class="text-xs text-gray-400">00:41</div>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">تبلت شیائومی مدل Redmi Pad SE حافظه 256 رم 8 - خاکستری پلاگین</td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-purple-100 text-purple-800 rounded-full">تخفیف</span>
        </td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-orange-100 text-orange-800 rounded-full">در انتظار</span>
        </td>
        <td class="py-3 px-4 text-sm text-gray-600">
          <div>1404/03/13</div>
          <div class="text-xs text-gray-400">15:25</div>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">چارو رباتیک شیائومی مدل S20 مقید، اصالت و سلامت فیزیکی کالا</td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-purple-100 text-purple-800 rounded-full">تخفیف</span>
        </td>
        <td class="py-3 px-4">
          <span class="inline-flex px-3 py-1 text-xs font-semibold bg-orange-100 text-orange-800 rounded-full">در انتظار</span>
        </td>
        <td class="py-3 px-4 text-sm text-gray-600">
          <div>1404/03/13</div>
          <div class="text-xs text-gray-400">13:00</div>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</div>

<!-- پرتقاضاترین محصولات -->
<div class="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
  <div class="bg-gray-50 px-6 py-4 border-b border-gray-100">
    <h3 class="text-lg font-semibold text-gray-900">پرتقاضاترین محصولات</h3>
  </div>
  <div class="p-0">
    <table class="w-full">
      <thead class="bg-blue-100">
      <tr>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">محصول</th>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">تعداد درخواست</th>
        <th class="text-right py-3 px-4 text-sm font-medium text-blue-900">در انتظار</th>
      </tr>
      </thead>
      <tbody class="divide-y divide-gray-100">
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">چارو رباتیک روبراک Saros Z70</td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-800 rounded-full text-sm font-bold">9</span>
        </td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-orange-100 text-orange-800 rounded-full text-sm font-bold">9</span>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">سبد خرید تاشو شیائومی مدل MaiWei سایز 35L (دارای ایراد ظاهری جزئی)</td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-800 rounded-full text-sm font-bold">5</span>
        </td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-orange-100 text-orange-800 rounded-full text-sm font-bold">5</span>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">محلول کن شیائومی شیائومی Deerma NU05 - اصالت و سلامت فیزیکی کالا</td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-800 rounded-full text-sm font-bold">4</span>
        </td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-orange-100 text-orange-800 rounded-full text-sm font-bold">4</span>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">کوار آبی میکروهو شیائومی مدل MH01R</td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-800 rounded-full text-sm font-bold">4</span>
        </td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-orange-100 text-orange-800 rounded-full text-sm font-bold">2</span>
        </td>
      </tr>
      <tr class="hover:bg-gray-50">
        <td class="py-3 px-4 text-sm text-gray-900">ساعت هوشمند شیائومی Redmi Watch 5 Lite مشکی، اصالت و سلامت فیزیکی کالا</td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-800 rounded-full text-sm font-bold">4</span>
        </td>
        <td class="py-3 px-4 text-center">
          <span class="inline-flex items-center justify-center w-8 h-8 bg-orange-100 text-orange-800 rounded-full text-sm font-bold">4</span>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

interface Props {
  activeNavItem: string
}

const props = defineProps<Props>()
</script>
