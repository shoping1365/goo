<template>
  <div class="summary-box bg-white rounded-2xl shadow p-6 flex flex-col gap-6">
    <div v-if="count > 0" class="flex flex-col gap-2">
      <div class="flex justify-between text-gray-600 text-sm">
        <span>قیمت کالاها ({{ count }})</span>
        <span>{{ formatPrice(total) }}</span>
      </div>
      <div class="flex justify-between font-bold text-lg text-[#1a2341]">
        <span>جمع سبد خرید</span>
        <span>{{ formatPrice(total) }}</span>
      </div>
    </div>
    <button v-if="count > 0" @click="proceed" class="w-full py-3 rounded-xl bg-[#e60023] text-white font-bold text-lg shadow-lg hover:bg-[#c9001b] transition">تایید و تکمیل سفارش</button>
    <div v-if="count > 0" class="text-xs text-gray-500 mt-2">هزینه این سفارش هنوز پرداخت نشده و در صورت اتمام موجودی، کالاها از سبد حذف می‌شوند.</div>
  </div>
</template>

<script setup>
// const { isAuthenticated } = useAuth() // احراز هویت غیرفعال شده است

const props = defineProps({ total: { type: Number, default: 0 }, count: { type: Number, default: 0 } })
function formatPrice(val) {
  if (val == null || isNaN(val)) return '۰ تومان'
  return val.toLocaleString('fa-IR') + ' تومان'
}

const proceed = async () => {
  // احراز هویت غیرفعال شده است - همیشه به checkout هدایت می‌شود
  return navigateTo('/checkout')
}

</script>

<style scoped>
.summary-box {
  background: linear-gradient(120deg, #fff 60%, #f3f0ff 100%);
  border-radius: 1.5rem;
  box-shadow: 0 4px 16px 0 #a5b4fc22;
  border: 1.5px solid #ede9fe;
  padding: 2rem 1.5rem;
  font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;
}
.summary-box .summary-title {
  color: #5b21b6;
  font-weight: bold;
  font-size: 1.2rem;
}
.summary-box .summary-total {
  color: #1a2341;
  font-weight: bold;
  font-size: 1.4rem;
}
.summary-box .summary-btn {
  background: linear-gradient(90deg, #5b21b6 0%, #0284c7 100%);
  color: #fff;
  font-weight: bold;
  font-size: 1.1rem;
  border-radius: 1rem;
  box-shadow: 0 2px 8px 0 #a5b4fc33;
  transition: background .18s;
}
.summary-box .summary-btn:hover {
  background: linear-gradient(90deg, #0284c7 0%, #5b21b6 100%);
}
</style> 
