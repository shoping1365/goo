<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-blue-700 mb-2 flex justify-between items-center">
      <span>لیست علاقه‌مندی‌ها</span>
      <button v-if="wishlist.length > 0" class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-4">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
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
        @click="fetchWishlist"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="wishlist.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <i class="fas fa-heart text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">لیست علاقه‌مندی‌ها خالی است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر هنوز محصولی به علاقه‌مندی‌ها اضافه نکرده است</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام محصول</th>
          <th class="p-2 text-right font-bold">قیمت</th>
          <th class="p-2 text-right font-bold">تاریخ اضافه شدن</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in displayedWishlist" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">{{ item.productName }}</td>
          <td class="p-2 text-right">{{ formatPrice(item.productPrice) }}</td>
          <td class="p-2 text-right">{{ item.dateAdded }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Modal -->
    <ViewAllModal v-model="showAll" title="لیست علاقه‌مندی‌ها">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="p-2 text-right font-bold">نام محصول</th>
            <th class="p-2 text-right font-bold">قیمت</th>
            <th class="p-2 text-right font-bold">تاریخ اضافه شدن</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in wishlist" :key="item.id" class="border-b border-blue-200 hover:bg-gray-50">
            <td class="p-2 text-right">{{ item.productName }}</td>
            <td class="p-2 text-right">{{ formatPrice(item.productPrice) }}</td>
            <td class="p-2 text-right">{{ item.dateAdded }}</td>
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

const props = defineProps<{ user: User }>();
const showAll = ref(false);

interface WishlistItem {
  id: number;
  productName: string;
  productPrice: number;
  dateAdded: string;
}

interface WishlistResponse {
  success: boolean;
  data: {
    wishlistItems: { id: number; productName: string; productPrice: number; createdAt: string }[];
  };
}

// Real data for wishlist
const wishlist = ref<WishlistItem[]>([]);
const loading = ref(false);
const error = ref('');

const displayedWishlist = computed(() => {
  return showAll.value ? wishlist.value : wishlist.value.slice(0, 5);
});

const fetchWishlist = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch<WishlistResponse>(`/api/admin/users/${props.user.id}/wishlist`, {
      query: {
        page: 1,
        limit: 100
      }
    });

    if (response.success) {
      wishlist.value = response.data.wishlistItems.map((item) => ({
        id: item.id,
        productName: item.productName,
        productPrice: item.productPrice,
        dateAdded: formatDateTime(item.createdAt)
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: unknown) {
    const e = err as { data?: { message?: string } };
    console.error('خطا در دریافت لیست علاقه‌مندی‌ها:', err);
    error.value = e.data?.message || 'خطا در دریافت لیست علاقه‌مندی‌ها';
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
  fetchWishlist();
});
</script> 
