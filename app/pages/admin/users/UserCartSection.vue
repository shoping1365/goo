<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-green-700 mb-2 flex justify-between items-center">
      <span>سبد خرید کاربر</span>
      <button v-if="cart.length > 0" class="bg-green-100 hover:bg-green-200 text-green-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
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
        @click="fetchCart"
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="cart.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <i class="fas fa-shopping-cart text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">سبد خرید خالی است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر هنوز محصولی به سبد خرید اضافه نکرده است</p>
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
        <tr v-for="item in displayedCart" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.productName }}</td>
          <td class="p-2 text-right">{{ item.quantity }}</td>
          <td class="p-2 text-right">{{ formatPrice(item.finalPrice) }}</td>
          <td class="p-2 text-right">{{ item.dateAdded }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAll" title="سبد خرید کاربر">
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
          <tr v-for="item in cart" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
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

// Real data for cart
const cart = ref([]);
const loading = ref(false);
const error = ref('');

const displayedCart = computed(() => {
  return showAll.value ? cart.value : cart.value.slice(0, 5);
});

const fetchCart = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/cart`, {
      query: {
        page: 1,
        limit: 100
      }
    }) as any;

    if (response.success) {
      cart.value = response.data.cartItems.map((item: any) => ({
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
    console.error('خطا در دریافت سبد خرید:', err);
    error.value = err.data?.message || 'خطا در دریافت سبد خرید';
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
  fetchCart();
});
</script> 
