<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-red-700 mb-2 flex justify-between items-center">
      <span>تلاش‌های ناموفق اخیر</span>
      <button v-if="failedAttempts.length > 0" class="bg-red-100 hover:bg-red-200 text-red-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllAttempts = !showAllAttempts">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAllAttempts ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-4">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-red-600"></div>
      <p class="mt-2 text-sm text-gray-500">در حال بارگذاری...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-4">
      <div class="text-red-500 mb-2">
        <i class="fas fa-exclamation-triangle"></i>
      </div>
      <p class="text-red-600 text-sm">{{ error }}</p>
      <button
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
        @click="fetchFailedAttempts"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="failedAttempts.length === 0" class="text-center py-4">
      <div class="text-green-500 mb-2">
        <i class="fas fa-shield-check text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">هیچ تلاش ناموفقی ثبت نشده است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر تا کنون با موفقیت وارد شده است</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">IP</th>
          <th class="p-2 text-right font-bold">دستگاه</th>
          <th class="p-2 text-right font-bold">مرورگر</th>
          <th class="p-2 text-right font-bold">دلیل شکست</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="attempt in displayedAttempts" :key="attempt.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ attempt.date }}</td>
          <td class="p-2 text-right">{{ attempt.ip }}</td>
          <td class="p-2 text-right">{{ attempt.device }}</td>
          <td class="p-2 text-right">{{ attempt.browser }}</td>
          <td class="p-2 text-right">{{ attempt.failureReason }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAllAttempts" title="تلاش‌های ناموفق اخیر">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">تاریخ</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">دستگاه</th>
            <th class="p-2 text-right font-bold">مرورگر</th>
            <th class="p-2 text-right font-bold">دلیل شکست</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="attempt in failedAttempts" :key="attempt.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ attempt.date }}</td>
            <td class="p-2 text-right">{{ attempt.ip }}</td>
            <td class="p-2 text-right">{{ attempt.device }}</td>
            <td class="p-2 text-right">{{ attempt.browser }}</td>
            <td class="p-2 text-right">{{ attempt.failureReason }}</td>
          </tr>
        </tbody>
      </table>
    </ViewAllModal>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';
import type { User } from '~/types/user';

interface FailedAttempt {
  id: number;
  date: string;
  ip: string;
  device: string;
  browser: string;
  failureReason: string;
}

const props = defineProps<{ user: User }>();
const showAllAttempts = ref(false);

// Real data for failed login attempts
const failedAttempts = ref<FailedAttempt[]>([]);
const loading = ref(false);
const error = ref('');

const displayedAttempts = computed(() => {
  return showAllAttempts.value ? failedAttempts.value : failedAttempts.value.slice(0, 5);
});

const fetchFailedAttempts = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/failed-attempts`, {
      query: {
        page: 1,
        limit: 100
      }
    }) as { success: boolean; data: { failedAttempts: Record<string, unknown>[] } };

    if (response.success) {
      failedAttempts.value = response.data.failedAttempts.map((attempt) => ({
        id: attempt.id as number,
        date: formatDateTime(attempt.createdAt as string),
        ip: attempt.ipAddress as string,
        device: attempt.deviceType as string,
        browser: attempt.browser as string,
        failureReason: attempt.failureReason as string
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: unknown) {
    const e = err as { data?: { message?: string } };
    error.value = e.data?.message || 'خطا در دریافت تلاش‌های ناموفق';
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
  fetchFailedAttempts();
});
</script>

