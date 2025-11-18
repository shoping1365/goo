<template>
     <table class="min-w-full divide-y divide-gray-200">
       <thead class="bg-gray-50">
         <tr>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">شماره سفارش</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">مشتری</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">شماره تماس</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">نوع خطا</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">تاریخ خطا</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">تعداد تلاش مجدد</th>
           <th class="px-6 py-3 text-right text-xs font-medium text-gray-500">وضعیت</th>
         </tr>
       </thead>
       <tbody class="bg-white divide-y divide-gray-200">
         <tr v-for="row in data" :key="row.orderNumber">
           <td class="px-6 py-4 whitespace-nowrap">{{ row.orderNumber }}</td>
           <td class="px-6 py-4 whitespace-nowrap">{{ row.customerName }}</td>
           <td class="px-6 py-4 whitespace-nowrap">{{ row.phoneNumber }}</td>
           <td class="px-6 py-4 whitespace-nowrap">{{ row.errorType }}</td>
           <td class="px-6 py-4 whitespace-nowrap">{{ row.errorDate }}</td>
           <td class="px-6 py-4 whitespace-nowrap">{{ row.retryCount }}</td>
           <td class="px-6 py-4 whitespace-nowrap">
             <span :class="row.status === 'ناموفق' ? 'text-red-600' : 'text-yellow-600'">{{ row.status }}</span>
           </td>
         </tr>
       </tbody>
     </table>
   </template>
   
   <script setup lang="ts">
   defineProps<{ data: any[] }>()
   </script>
