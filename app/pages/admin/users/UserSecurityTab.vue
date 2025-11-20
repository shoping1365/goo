<template>
  <div>
    <div class="bg-white rounded-xl shadow-lg p-4 mb-4">
      <div class="font-bold text-red-700 mb-2 flex justify-between items-center">
        <span>سشن‌های فعال</span>
        <button v-if="activeSessions.length > 5" class="bg-red-100 hover:bg-red-200 text-red-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllSessions = !showAllSessions">
          <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          مشاهده همه
        </button>
      </div>
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">دستگاه</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">آخرین فعالیت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="session in displayedSessions" :key="session.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ session.device }}</td>
            <td class="p-2 text-right">{{ session.ip }}</td>
            <td class="p-2 text-right">{{ session.lastActivity }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="bg-white rounded-xl shadow-lg p-4 mb-4">
      <div class="font-bold text-green-700 mb-2 flex justify-between items-center">
        <span>ورودهای موفق اخیر</span>
        <button v-if="successfulLogins.length > 5" class="bg-green-100 hover:bg-green-200 text-green-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllLogins = !showAllLogins">
          <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          مشاهده همه
        </button>
      </div>
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="login in displayedLogins" :key="login.date + login.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ login.date }}</td>
            <td class="p-2 text-right">{{ login.ip }}</td>
            <td class="p-2 text-right">{{ login.device }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="bg-white rounded-xl shadow-lg p-4">
      <div class="font-bold text-red-700 mb-2 flex justify-between items-center">
        <span>تلاش‌های ناموفق اخیر</span>
        <button v-if="failedAttempts.length > 5" class="bg-red-100 hover:bg-red-200 text-red-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllAttempts = !showAllAttempts">
          <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          مشاهده همه
        </button>
      </div>
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="attempt in displayedAttempts" :key="attempt.date + attempt.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ attempt.date }}</td>
            <td class="p-2 text-right">{{ attempt.ip }}</td>
            <td class="p-2 text-right">{{ attempt.device }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modals -->
    <ViewAllModal v-model="showAllSessions" title="سشن‌های فعال">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">دستگاه</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">آخرین فعالیت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="session in activeSessions" :key="session.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ session.device }}</td>
            <td class="p-2 text-right">{{ session.ip }}</td>
            <td class="p-2 text-right">{{ session.lastActivity }}</td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>

    <ViewAllModal v-model="showAllLogins" title="ورودهای موفق اخیر">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="login in successfulLogins" :key="login.date + login.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ login.date }}</td>
            <td class="p-2 text-right">{{ login.ip }}</td>
            <td class="p-2 text-right">{{ login.device }}</td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>

    <ViewAllModal v-model="showAllAttempts" title="تلاش‌های ناموفق اخیر">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="attempt in failedAttempts" :key="attempt.date + attempt.ip" class="border-b hover:bg-gray-50">
            <td class="p-2 text-right">{{ attempt.date }}</td>
            <td class="p-2 text-right">{{ attempt.ip }}</td>
            <td class="p-2 text-right">{{ attempt.device }}</td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>
  </div>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';

const showAllSessions = ref(false);
const showAllLogins = ref(false);
const showAllAttempts = ref(false);

// Mock data for active sessions
const activeSessions = [
  { device: 'موبایل', ip: '192.168.1.10', lastActivity: '1402/03/10 09:12' },
  { device: 'دسکتاپ', ip: '192.168.1.11', lastActivity: '1402/03/09 18:22' },
  { device: 'تبلت', ip: '192.168.1.13', lastActivity: '1402/03/09 15:30' },
  { device: 'موبایل', ip: '192.168.1.14', lastActivity: '1402/03/09 12:45' },
  { device: 'دسکتاپ', ip: '192.168.1.15', lastActivity: '1402/03/09 10:20' },
  { device: 'موبایل', ip: '192.168.1.16', lastActivity: '1402/03/08 22:15' },
];

// Mock data for successful logins
const successfulLogins = [
  { date: '1402/03/10 09:12', ip: '192.168.1.10', device: 'موبایل' },
  { date: '1402/03/09 18:22', ip: '192.168.1.11', device: 'دسکتاپ' },
  { date: '1402/03/09 15:30', ip: '192.168.1.13', device: 'تبلت' },
  { date: '1402/03/09 12:45', ip: '192.168.1.14', device: 'موبایل' },
  { date: '1402/03/09 10:20', ip: '192.168.1.15', device: 'دسکتاپ' },
  { date: '1402/03/08 22:15', ip: '192.168.1.16', device: 'موبایل' },
];

// Mock data for failed login attempts
const failedAttempts = [
  { date: '1402/03/08 15:30', ip: '192.168.1.12', device: 'موبایل' },
  { date: '1402/03/08 15:35', ip: '192.168.1.12', device: 'موبایل' },
  { date: '1402/03/08 15:40', ip: '192.168.1.12', device: 'موبایل' },
  { date: '1402/03/08 15:45', ip: '192.168.1.12', device: 'موبایل' },
  { date: '1402/03/08 15:50', ip: '192.168.1.12', device: 'موبایل' },
  { date: '1402/03/08 15:55', ip: '192.168.1.12', device: 'موبایل' },
];

const displayedSessions = computed(() => {
  return showAllSessions.value ? activeSessions : activeSessions.slice(0, 5);
});

const displayedLogins = computed(() => {
  return showAllLogins.value ? successfulLogins : successfulLogins.slice(0, 5);
});

const displayedAttempts = computed(() => {
  return showAllAttempts.value ? failedAttempts : failedAttempts.slice(0, 5);
});
</script> 
