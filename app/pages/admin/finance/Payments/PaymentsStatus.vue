<template>
  <div v-if="hasAccess" class="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
    <div class="flex items-center justify-between mb-6">
      <div>
        <h3 class="text-lg font-semibold text-gray-900">وضعیت پرداخت‌ها</h3>
        <p class="text-sm text-gray-500 mt-1">Payment Status Distribution</p>
      </div>
    </div>

    <div class="space-y-4">
      <div 
        v-for="item in status" 
        :key="item.status"
        class="flex items-center justify-between p-3 rounded-lg"
        :class="getStatusBgClass(item.status)"
      >
        <div class="flex items-center">
          <div 
            class="w-3 h-3 rounded-full ml-3"
            :class="getStatusColorClass(item.status)"
          ></div>
          <div>
            <div class="font-medium" :class="getStatusTextClass(item.status)">
              {{ getStatusName(item.status) }}
            </div>
            <div class="text-sm" :class="getStatusSubTextClass(item.status)">
              {{ item.percentage }}% از کل
            </div>
          </div>
        </div>
        <div class="text-left">
          <div class="font-semibold" :class="getStatusTextClass(item.status)">
            {{ item.count }}
          </div>
          <div class="text-xs" :class="getStatusSubTextClass(item.status)">
            تراکنش
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6 pt-4 border-t border-gray-200">
      <div class="text-center">
        <div class="text-lg font-semibold text-gray-900">{{ totalTransactions }}</div>
        <div class="text-sm text-gray-500">کل تراکنش‌ها</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface Props {
  status?: Array<{
    status: string
    count: number
    percentage: number
  }>
}

const props = withDefaults(defineProps<Props>(), {
  status: () => []
})

const totalTransactions = computed(() => 
  props.status.reduce((sum, item) => sum + item.count, 0)
)

const getStatusName = (status: string) => {
  const names = {
    'success': 'موفق',
    'failed': 'ناموفق',
    'pending': 'در انتظار',
    'cancelled': 'لغو شده'
  }
  return names[status as keyof typeof names] || status
}

const getStatusColorClass = (status: string) => {
  const colors = {
    'success': 'bg-green-500',
    'failed': 'bg-red-500',
    'pending': 'bg-yellow-500',
    'cancelled': 'bg-gray-500'
  }
  return colors[status as keyof typeof colors] || 'bg-gray-500'
}

const getStatusBgClass = (status: string) => {
  const colors = {
    'success': 'bg-green-50',
    'failed': 'bg-red-50',
    'pending': 'bg-yellow-50',
    'cancelled': 'bg-gray-50'
  }
  return colors[status as keyof typeof colors] || 'bg-gray-50'
}

const getStatusTextClass = (status: string) => {
  const colors = {
    'success': 'text-green-800',
    'failed': 'text-red-800',
    'pending': 'text-yellow-800',
    'cancelled': 'text-gray-800'
  }
  return colors[status as keyof typeof colors] || 'text-gray-800'
}

const getStatusSubTextClass = (status: string) => {
  const colors = {
    'success': 'text-green-600',
    'failed': 'text-red-600',
    'pending': 'text-yellow-600',
    'cancelled': 'text-gray-600'
  }
  return colors[status as keyof typeof colors] || 'text-gray-600'
}
</script> 
