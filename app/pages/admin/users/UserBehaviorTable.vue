<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-purple-700 mb-2 flex justify-between items-center">
      <span>رفتار کاربر</span>
      <button v-if="actions.length > 5" class="bg-purple-100 hover:bg-purple-200 text-purple-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">عملیات</th>
          <th class="p-2 text-right font-bold">زمان</th>
          <th class="p-2 text-right font-bold">دستگاه</th>
          <th class="p-2 text-right font-bold">مکان</th>
          <th class="p-2 text-right font-bold">IP</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="action in displayedActions" :key="action.time + action.ip" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ action.action }}</td>
          <td class="p-2 text-right">{{ action.time }}</td>
          <td class="p-2 text-right">{{ action.device }}</td>
          <td class="p-2 text-right">{{ action.location }}</td>
          <td class="p-2 text-right">{{ action.ip }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue';

const props = defineProps<{ user: any }>();
const showAll = ref(false);

// Mock data for behavior
const actions = [
  { action: 'ورود به سیستم', time: '1402/03/10 09:12', device: 'موبایل', location: 'تهران', ip: '192.168.1.10' },
  { action: 'تغییر رمز عبور', time: '1402/03/09 18:22', device: 'دسکتاپ', location: 'اصفهان', ip: '192.168.1.11' },
  { action: 'خروج', time: '1402/03/09 19:00', device: 'موبایل', location: 'تهران', ip: '192.168.1.10' },
  { action: 'ورود به سیستم', time: '1402/03/08 14:30', device: 'دسکتاپ', location: 'تهران', ip: '192.168.1.12' },
  { action: 'تغییر اطلاعات پروفایل', time: '1402/03/08 15:45', device: 'دسکتاپ', location: 'تهران', ip: '192.168.1.12' },
  { action: 'خروج', time: '1402/03/08 17:20', device: 'دسکتاپ', location: 'تهران', ip: '192.168.1.12' },
  { action: 'ورود به سیستم', time: '1402/03/07 10:15', device: 'موبایل', location: 'شیراز', ip: '192.168.1.13' },
];

const displayedActions = computed(() => {
  return showAll.value ? actions : actions.slice(0, 5);
});
</script> 
