<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-green-700 mb-2 flex justify-between items-center">
      <span>ورودهای موفق اخیر</span>
      <button v-if="successfulLogins.length > 0" class="bg-green-100 hover:bg-green-200 text-green-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllLogins = !showAllLogins">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAllLogins ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-4">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-green-600"></div>
      <p class="mt-2 text-sm text-gray-500">در حال بارگذاری...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-4">
      <div class="text-red-500 mb-2">
        <i class="fas fa-exclamation-triangle"></i>
      </div>
      <p class="text-red-600 text-sm">{{ error }}</p>
      <button
        @click="fetchSuccessfulLogins"
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="successfulLogins.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <i class="fas fa-info-circle text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">هیچ ورود موفقی ثبت نشده است</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">IP</th>
          <th class="p-2 text-right font-bold">دستگاه</th>
          <th class="p-2 text-right font-bold">مرورگر</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="login in displayedLogins" :key="login.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ login.date }}</td>
          <td class="p-2 text-right">{{ login.ip }}</td>
          <td class="p-2 text-right">{{ login.device }}</td>
          <td class="p-2 text-right">{{ login.browser }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAllLogins" title="ورودهای موفق اخیر">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
            <th class="p-2 text-right font-bold">مرورگر</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="login in successfulLogins" :key="login.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ login.date }}</td>
            <td class="p-2 text-right">{{ login.ip }}</td>
            <td class="p-2 text-right">{{ login.device }}</td>
            <td class="p-2 text-right">{{ login.browser }}</td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';

const props = defineProps<{ user: any }>();
const showAllLogins = ref(false);

// Real data for successful logins
const successfulLogins = ref([]);
const loading = ref(false);
const error = ref('');

const displayedLogins = computed(() => {
  return showAllLogins.value ? successfulLogins.value : successfulLogins.value.slice(0, 5);
});

const fetchSuccessfulLogins = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/successful-logins`, {
      query: {
        page: 1,
        limit: 100
      }
    }) as any;

    if (response.success) {
      successfulLogins.value = response.data.successfulLogins.map((login: any) => ({
        id: login.id,
        date: formatDateTime(login.createdAt),
        ip: login.ipAddress,
        device: login.deviceType,
        browser: login.browser
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: any) {
    console.error('خطا در دریافت ورودهای موفق:', err);
    error.value = err.data?.message || 'خطا در دریافت ورودهای موفق';
  } finally {
    loading.value = false;
  }
};

const formatDateTime = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleString('fa-IR', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  });
};

onMounted(() => {
  fetchSuccessfulLogins();
});
</script>

