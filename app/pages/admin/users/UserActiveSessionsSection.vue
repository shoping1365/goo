<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-red-700 mb-2 flex justify-between items-center">
      <span>سشن‌های فعال</span>
      <button v-if="activeSessions.length > 0" class="bg-red-100 hover:bg-red-200 text-red-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAllSessions = !showAllSessions">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAllSessions ? 'نمایش کمتر' : 'مشاهده همه' }}
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
        @click="fetchActiveSessions"
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="activeSessions.length === 0" class="text-center py-4">
      <div class="text-green-500 mb-2">
        <i class="fas fa-check-circle text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">هیچ سشن فعالی ثبت نشده است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر در حال حاضر سشن فعالی ندارد</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">دستگاه</th>
          <th class="p-2 text-right font-bold">مرورگر</th>
          <th class="p-2 text-right font-bold">IP</th>
          <th class="p-2 text-right font-bold">آخرین فعالیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="session in displayedSessions" :key="session.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ session.deviceType }}</td>
          <td class="p-2 text-right">{{ session.browser }}</td>
          <td class="p-2 text-right">{{ session.ipAddress }}</td>
          <td class="p-2 text-right">{{ session.lastActivity }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAllSessions" title="سشن‌های فعال">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">دستگاه</th>
            <th class="p-2 text-right font-bold">مرورگر</th>
            <th class="p-2 text-right font-bold">IP</th>
            <th class="p-2 text-right font-bold">آخرین فعالیت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="session in activeSessions" :key="session.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ session.deviceType }}</td>
            <td class="p-2 text-right">{{ session.browser }}</td>
            <td class="p-2 text-right">{{ session.ipAddress }}</td>
            <td class="p-2 text-right">{{ session.lastActivity }}</td>
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
const showAllSessions = ref(false);

// Real data for active sessions
const activeSessions = ref([]);
const loading = ref(false);
const error = ref('');

const displayedSessions = computed(() => {
  return showAllSessions.value ? activeSessions.value : activeSessions.value.slice(0, 5);
});

const fetchActiveSessions = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/active-sessions`, {
      query: {
        page: 1,
        limit: 100
      }
    }) as any;

    if (response.success) {
      activeSessions.value = response.data.activeSessions.map((session: any) => ({
        id: session.id,
        deviceType: session.deviceType,
        browser: session.browser,
        ipAddress: session.ipAddress,
        lastActivity: formatDateTime(session.lastUsedAt || session.createdAt),
        token: session.token,
        authMethod: session.authMethod,
        isVerified: session.isVerified,
        expiresAt: session.expiresAt
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: any) {
    console.error('خطا در دریافت سشن‌های فعال:', err);
    error.value = err.data?.message || 'خطا در دریافت سشن‌های فعال';
  } finally {
    loading.value = false;
  }
};

const formatDateTime = (dateString: string) => {
  if (!dateString) return 'نامشخص';
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
  fetchActiveSessions();
});
</script>
