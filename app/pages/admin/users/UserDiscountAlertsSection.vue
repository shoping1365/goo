<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-orange-700 mb-2 flex justify-between items-center">
      <span>اعلان‌های تخفیف</span>
      <button v-if="alerts.length > 5" class="bg-orange-100 hover:bg-orange-200 text-orange-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">تاریخ درخواست</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="alert in displayedAlerts" :key="alert.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ alert.productName }}</td>
          <td class="p-2 text-right">{{ alert.requestDate }}</td>
          <td class="p-2 text-right">
            <span class="px-2 py-1 rounded text-xs font-bold" :class="{
              'bg-green-100 text-green-700': alert.status === 'active',
              'bg-red-100 text-red-700': alert.status === 'expired',
              'bg-yellow-100 text-yellow-700': alert.status === 'pending'
            }">{{ alert.status === 'active' ? 'فعال' : alert.status === 'expired' ? 'منقضی شده' : 'در انتظار' }}</span>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAll" title="اعلان‌های تخفیف">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">نام محصول</th>
            <th class="p-2 text-right font-bold">تاریخ درخواست</th>
            <th class="p-2 text-right font-bold">وضعیت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="alert in alerts" :key="alert.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ alert.productName }}</td>
            <td class="p-2 text-right">{{ alert.requestDate }}</td>
            <td class="p-2 text-right">
              <span class="px-2 py-1 rounded text-xs font-bold" :class="{
                'bg-green-100 text-green-700': alert.status === 'active',
                'bg-red-100 text-red-700': alert.status === 'expired',
                'bg-yellow-100 text-yellow-700': alert.status === 'pending'
              }">{{ alert.status === 'active' ? 'فعال' : alert.status === 'expired' ? 'منقضی شده' : 'در انتظار' }}</span>
            </td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';

const props = defineProps<{ user: any }>();
const showAll = ref(false);

// Mock data for discount alerts
const alerts = [
  { id: 1, productName: 'لپ‌تاپ اپل', requestDate: '1402/03/10', status: 'active' },
  { id: 2, productName: 'هدفون سونی', requestDate: '1402/03/09', status: 'expired' },
  { id: 3, productName: 'موبایل سامسونگ', requestDate: '1402/03/08', status: 'pending' },
  { id: 4, productName: 'تبلت اپل', requestDate: '1402/03/07', status: 'active' },
  { id: 5, productName: 'ساعت هوشمند', requestDate: '1402/03/06', status: 'expired' },
  { id: 6, productName: 'دوربین دیجیتال', requestDate: '1402/03/05', status: 'pending' },
];

const displayedAlerts = computed(() => {
  return showAll.value ? alerts : alerts.slice(0, 5);
});
</script> 
