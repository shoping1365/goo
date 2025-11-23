<template>
  <div v-if="hasAccess" class="flex flex-wrap gap-6 items-center bg-white p-6 rounded-xl shadow mb-4">
    <input v-model="from" type="date" class="border rounded px-2 py-1" />
    <span>-</span>
    <input v-model="to" type="date" class="border rounded px-2 py-1" />
    <select v-model="paymentMethod" class="border rounded px-2 py-1">
      <option value="">روش پرداخت</option>
      <option value="online">آنلاین</option>
      <option value="offline">حضوری</option>
      <option value="installment">قسطی</option>
    </select>
    <select v-model="orderStatus" class="border rounded px-2 py-1">
      <option value="">وضعیت سفارش</option>
      <option value="pending">در انتظار</option>
      <option value="completed">تکمیل شده</option>
      <option value="cancelled">لغو شده</option>
    </select>
    <select v-model="customerType" class="border rounded px-2 py-1">
      <option value="">نوع مشتری</option>
      <option value="new">جدید</option>
      <option value="returning">بازگشتی</option>
      <option value="vip">VIP</option>
    </select>
    <input v-model="productName" placeholder="نام محصول..." class="border rounded px-2 py-1" />
    <button class="px-4 py-2 bg-green-500 text-white rounded-lg shadow hover:bg-green-600 transition" @click="emitFilter">اعمال فیلتر</button>
  </div>
</template>
<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useAuth } from '~/composables/useAuth'

// احراز هویت
const { user, isAuthenticated } = useAuth()

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false
  }

  const userRole = user.value?.role?.toLowerCase() || ''
  const adminRoles = ['admin', 'developer']
  return adminRoles.includes(userRole)
})

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async () => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false })
  }
}

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth()
})

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth()
  }
})

const from = ref('')
const to = ref('')
const paymentMethod = ref('')
const orderStatus = ref('')
const customerType = ref('')
const productName = ref('')
const emit = defineEmits(['filter'])
function emitFilter() {
  emit('filter', { 
    from: from.value, 
    to: to.value, 
    paymentMethod: paymentMethod.value, 
    orderStatus: orderStatus.value, 
    customerType: customerType.value, 
    productName: productName.value 
  })
}
</script> 
