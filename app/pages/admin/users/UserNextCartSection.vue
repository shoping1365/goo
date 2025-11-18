<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-purple-700 mb-2 flex justify-between items-center">
      <span>سبد خرید بعدی</span>
      <button v-if="nextCart.length > 0" class="bg-purple-100 hover:bg-purple-200 text-purple-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-4">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-purple-600"></div>
      <p class="mt-2 text-sm text-gray-500">در حال بارگذاری...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-4">
      <div class="text-red-500 mb-2">
        <i class="fas fa-exclamation-triangle"></i>
      </div>
      <p class="text-red-600 text-sm">{{ error }}</p>
      <button
        @click="fetchNextCart"
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="nextCart.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <i class="fas fa-shopping-bag text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">سبد خرید بعدی خالی است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر هنوز محصولی برای خرید بعدی ذخیره نکرده است</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">تعداد</th>
          <th class="p-2 text-right font-bold">قیمت</th>
          <th class="p-2 text-right font-bold">تاریخ افزودن</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in displayedNextCart" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.productName }}</td>
          <td class="p-2 text-right">{{ item.quantity }}</td>
          <td class="p-2 text-right">{{ formatPrice(item.finalPrice) }}</td>
          <td class="p-2 text-right">{{ item.dateAdded }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAll" title="سبد خرید بعدی">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">نام محصول</th>
            <th class="p-2 text-right font-bold">تعداد</th>
            <th class="p-2 text-right font-bold">قیمت</th>
            <th class="p-2 text-right font-bold">تاریخ افزودن</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in nextCart" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ item.productName }}</td>
            <td class="p-2 text-right">{{ item.quantity }}</td>
            <td class="p-2 text-right">{{ formatPrice(item.finalPrice) }}</td>
            <td class="p-2 text-right">{{ item.dateAdded }}</td>
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
const showAll = ref(false);

// Real data for next cart
const nextCart = ref([]);
const loading = ref(false);
const error = ref('');

const displayedNextCart = computed(() => {
  return showAll.value ? nextCart.value : nextCart.value.slice(0, 5);
});

const fetchNextCart = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/next-cart`, {
      query: {
        page: 1,
        limit: 100
      }
    }) as any;

    if (response.success) {
      nextCart.value = response.data.nextCartItems.map((item: any) => ({
        id: item.id,
        productName: item.productName,
        quantity: item.quantity,
        finalPrice: item.finalPrice,
        dateAdded: formatDateTime(item.addedAt)
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: any) {
    console.error('خطا در دریافت سبد خرید بعدی:', err);
    error.value = err.data?.message || 'خطا در دریافت سبد خرید بعدی';
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

const formatPrice = (price: number) => {
  if (!price) return 'قیمت نامشخص';
  return new Intl.NumberFormat('fa-IR').format(price) + ' تومان';
};

onMounted(() => {
  fetchNextCart();
});
</script> 
