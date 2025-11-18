<template>
  <div class="cart-summary bg-white/90 backdrop-blur-lg rounded-2xl shadow-lg p-6" style="font-family: 'IranYekan', 'Vazir', 'IRANSansX', 'Tahoma', sans-serif;">
    <h3 class="text-lg font-bold mb-4 text-gray-800">خلاصه سفارش</h3>
    <div class="flex flex-col gap-3 mb-4">
      <div class="flex justify-between text-gray-600">
        <span>جمع کل کالاها</span>
        <span>{{ formatPrice(total) }}</span>
      </div>
      <div class="flex justify-between text-green-600">
        <span>تخفیف</span>
        <span>-{{ formatPrice(discount) }}</span>
      </div>
      <div class="flex justify-between text-blue-600">
        <span>هزینه ارسال</span>
        <span>{{ formatPrice(shipping) }}</span>
      </div>
      <div class="flex justify-between font-bold text-lg text-gray-800 border-t pt-3 mt-2">
        <span>مبلغ نهایی</span>
        <span>{{ formatPrice(finalAmount) }}</span>
      </div>
    </div>
    <button class="w-full py-3 rounded-xl bg-gradient-to-r from-blue-500 to-purple-500 text-white font-bold text-lg shadow-lg hover:from-blue-600 hover:to-purple-600 transition">ثبت سفارش</button>
  </div>
</template>

<script setup>
const total = 37500000
const discount = 2500000
const shipping = 120000
const finalAmount = total - discount + shipping
function formatPrice(val) { return val.toLocaleString('fa-IR') + ' تومان' }
</script> 
