<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <h1 class="text-3xl font-bold mb-4">با تشکر از خرید شما!</h1>
    <p class="mb-2">سفارش شما با موفقیت ثبت شد.</p>
    <p>شماره پیگیری سفارش: <span class="font-mono font-bold">{{ orderId }}</span></p>
    <NuxtLink to="/" class="mt-6 btn btn-primary">بازگشت به صفحه اصلی</NuxtLink>
  </div>
</template>

<script setup>
// دریافت پارامتر id از مسیر
const route = useRoute()
const orderId = route.params.id
</script>

<!--
  این صفحه پس از ثبت موفق سفارش نمایش داده می‌شود و شماره پیگیری سفارش را به کاربر نشان می‌دهد.
  اگر نیاز به نمایش اطلاعات بیشتر سفارش دارید، می‌توانید با استفاده از useFetch اطلاعات سفارش را نیز دریافت و نمایش دهید.
--> 