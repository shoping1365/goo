<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-green-700 mb-2 flex justify-between items-center">
      <span>سفارشات کاربر</span>
      <button v-if="orders.length > 0" class="bg-green-100 hover:bg-green-200 text-green-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="showAll = !showAll">
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
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
        @click="fetchOrders"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="orders.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <i class="fas fa-shopping-cart text-xl"></i>
      </div>
      <p class="text-gray-600 text-sm">هیچ سفارشی ثبت نشده است</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر تا کنون سفارشی نداشته است</p>
    </div>

    <!-- Data Table -->
    <table v-else class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">کد سفارش</th>
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
          <th class="p-2 text-right font-bold">مبلغ</th>
          <th class="p-2 text-right font-bold">روش پرداخت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="order in displayedOrders" :key="order.id" class="border-b border-blue-200 hover:bg-gray-50">
          <td class="p-2 text-right">
            <button 
              class="text-blue-600 hover:text-blue-800 underline cursor-pointer" 
              @click="openOrderModal(order)"
            >
              {{ order.orderNumber }}
            </button>
          </td>
          <td class="p-2 text-right">{{ order.date }}</td>
          <td class="p-2 text-right">
            <span :class="getStatusClass(order.status)">{{ getStatusText(order.status) }}</span>
          </td>
          <td class="p-2 text-right">{{ formatPrice(order.finalAmount) }}</td>
          <td class="p-2 text-right">{{ getPaymentMethodText(order.paymentMethod) }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Order Detail Modal -->
    <ViewAllModal v-model="orderModalOpen" :title="`جزئیات سفارش ${selectedOrder?.orderNumber || ''}`">
      <div v-if="selectedOrder" class="space-y-4">
        <!-- Order Info -->
        <div class="grid grid-cols-2 gap-6 p-6 bg-gray-50 rounded-lg">
          <div>
            <label class="text-sm font-bold text-gray-600">کد سفارش:</label>
            <p class="text-lg font-bold">{{ selectedOrder.orderNumber }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">تاریخ سفارش:</label>
            <p>{{ selectedOrder.date }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">وضعیت:</label>
            <p :class="getStatusClass(selectedOrder.status)">{{ getStatusText(selectedOrder.status) }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">وضعیت پرداخت:</label>
            <p :class="getPaymentStatusClass(selectedOrder.paymentStatus)">{{ getPaymentStatusText(selectedOrder.paymentStatus) }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">روش پرداخت:</label>
            <p>{{ getPaymentMethodText(selectedOrder.paymentMethod) }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">مبلغ کل:</label>
            <p class="font-bold">{{ formatPrice(selectedOrder.totalAmount) }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">مبلغ نهایی:</label>
            <p class="font-bold text-green-600">{{ formatPrice(selectedOrder.finalAmount) }}</p>
          </div>
        </div>

        <!-- User Information -->
        <div class="grid grid-cols-2 gap-6 p-6 bg-blue-50 rounded-lg">
          <h3 class="col-span-2 text-lg font-bold text-blue-800 mb-2">اطلاعات کاربر</h3>
          <div>
            <label class="text-sm font-bold text-gray-600">نام کاربر:</label>
            <p>{{ selectedOrder.userName || 'نامشخص' }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">تلفن همراه:</label>
            <p>{{ selectedOrder.userMobile || 'نامشخص' }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">ایمیل:</label>
            <p>{{ selectedOrder.userEmail || 'نامشخص' }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">کد ملی:</label>
            <p>{{ selectedOrder.userNationalCode || 'نامشخص' }}</p>
          </div>
          <div>
            <label class="text-sm font-bold text-gray-600">آخرین IP:</label>
            <p>{{ selectedOrder.userLastLoginIP || 'نامشخص' }}</p>
          </div>
        </div>

        <!-- Address Information -->
        <div class="grid grid-cols-1 gap-6 p-6 bg-green-50 rounded-lg">
          <h3 class="text-lg font-bold text-green-800 mb-2">آدرس‌های سفارش</h3>
          <div v-if="selectedOrder.shippingAddress">
            <label class="text-sm font-bold text-gray-600">آدرس ارسال:</label>
            <p class="text-sm">{{ selectedOrder.shippingAddress }}</p>
          </div>
          <div v-if="selectedOrder.billingAddress">
            <label class="text-sm font-bold text-gray-600">آدرس صورتحساب:</label>
            <p class="text-sm">{{ selectedOrder.billingAddress }}</p>
          </div>
          <div v-if="!selectedOrder.shippingAddress && !selectedOrder.billingAddress">
            <p class="text-gray-500">آدرسی ثبت نشده است</p>
          </div>
        </div>

        <!-- Order Items -->
        <div>
          <h3 class="text-lg font-bold mb-3">محصولات سفارش</h3>
          <div v-if="orderItems.length > 0" class="space-y-2">
            <div v-for="item in orderItems" :key="item.id" class="flex items-center justify-between p-3 bg-white border rounded-lg">
              <div class="flex items-center space-x-3 space-x-reverse">
                <img v-if="item.image" :src="item.image" :alt="item.name" class="w-12 h-12 object-cover rounded">
                <div>
                  <p class="font-bold">{{ item.name }}</p>
                  <p class="text-sm text-gray-600">SKU: {{ item.sku }}</p>
                  <p class="text-sm text-gray-600">تعداد: {{ item.quantity }}</p>
                </div>
              </div>
              <div class="text-left">
                <p class="font-bold">{{ formatPrice(item.totalPrice) }}</p>
                <p class="text-sm text-gray-600">قیمت واحد: {{ formatPrice(item.unitPrice) }}</p>
              </div>
            </div>
          </div>
          <div v-else class="text-center py-8 text-gray-500">
            <p>در حال بارگذاری محصولات...</p>
          </div>
        </div>
      </div>
    </ViewAllModal>
  </div>
</template>
<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';
import type { Order, OrderItem, OrderItemsResponse, OrderResponse } from '~/types/order';
import type { User } from '~/types/user';

const props = defineProps<{ user: User }>();
const showAll = ref(false);
const orderModalOpen = ref(false);
const selectedOrder = ref<Order | null>(null);
const orderItems = ref<OrderItem[]>([]);

// Real data for orders
const orders = ref<Order[]>([]);

const loading = ref(false);
const error = ref('');

const displayedOrders = computed(() => {
  return showAll.value ? orders.value : orders.value.slice(0, 5);
});

const fetchOrders = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch<OrderResponse>(`/api/admin/users/${props.user.id}/orders`, {
      query: {
        page: 1,
        limit: 100
      }
    });

    if (response.success) {
      orders.value = response.data.orders.map((order) => {
        const mappedOrder: Order = {
          id: order.id as number,
          orderNumber: order.orderNumber as string,
          date: formatDateTime(order.createdAt as string),
          status: order.status as string,
          paymentStatus: order.paymentStatus as string,
          paymentMethod: order.paymentMethod as string,
          finalAmount: order.finalAmount as number,
          totalAmount: order.totalAmount as number,
          // اطلاعات کاربر
          userName: order.userName as string,
          userMobile: order.userMobile as string,
          userEmail: order.userEmail as string,
          userNationalCode: order.userNationalCode as string,
          userLastLoginIP: order.userLastLoginIP as string,
          // آدرس‌ها
          shippingAddress: order.shippingAddress as string,
          billingAddress: order.billingAddress as string
        };
        
        return mappedOrder;
      });
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: unknown) {
    const e = err as { data?: { message?: string } };
    error.value = e.data?.message || 'خطا در دریافت سفارشات';
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

const formatPrice = (amount: number) => {
  return new Intl.NumberFormat('fa-IR').format(amount) + ' تومان';
};

const getStatusText = (status: string) => {
  const statusMap: { [key: string]: string } = {
    'pending': 'در انتظار',
    'processing': 'در حال پردازش',
    'shipped': 'ارسال شده',
    'delivered': 'تحویل داده شده',
    'cancelled': 'لغو شده',
    'refunded': 'بازگردانده شده'
  };
  return statusMap[status] || status;
};

const getStatusClass = (status: string) => {
  const classMap: { [key: string]: string } = {
    'pending': 'text-yellow-600',
    'processing': 'text-blue-600',
    'shipped': 'text-purple-600',
    'delivered': 'text-green-600',
    'cancelled': 'text-red-600',
    'refunded': 'text-gray-600'
  };
  return classMap[status] || 'text-gray-600';
};

const getPaymentMethodText = (method: string) => {
  const methodMap: { [key: string]: string } = {
    'card_to_card': 'کارت به کارت',
    'bank_gateway': 'درگاه بانکی',
    'cash_on_delivery': 'پرداخت در محل',
    'wallet': 'کیف پول',
    'installment': 'قسطی'
  };
  return methodMap[method] || method || 'نامشخص';
};

const getPaymentStatusText = (status: string) => {
  const statusMap: { [key: string]: string } = {
    'pending': 'در انتظار',
    'paid': 'پرداخت شده',
    'failed': 'ناموفق',
    'refunded': 'بازگردانده شده'
  };
  return statusMap[status] || status;
};

const getPaymentStatusClass = (status: string) => {
  const classMap: { [key: string]: string } = {
    'pending': 'text-yellow-600',
    'paid': 'text-green-600',
    'failed': 'text-red-600',
    'refunded': 'text-gray-600'
  };
  return classMap[status] || 'text-gray-600';
};

const openOrderModal = async (order: Order) => {
  selectedOrder.value = order;
  orderModalOpen.value = true;
  
  // Fetch order items
  try {
    const response = await $fetch<OrderItemsResponse>(`/api/admin/orders/${order.id}/items`);
    
    if (response.success) {
      orderItems.value = response.data.items;
    } else {
      orderItems.value = [];
    }
  } catch {
    orderItems.value = [];
  }
};

onMounted(() => {
  fetchOrders();
});
</script> 
