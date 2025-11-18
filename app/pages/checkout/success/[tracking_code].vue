<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <h1 class="text-3xl font-bold mb-4">با تشکر از خرید شما!</h1>
    <p class="mb-2">سفارش شما با موفقیت ثبت شد.</p>
    <p>کد رهگیری سفارش: <span class="font-mono font-bold">{{ trackingCode }}</span></p>
    <NuxtLink to="/" class="mt-6 btn btn-primary">بازگشت به صفحه اصلی</NuxtLink>
  </div>
</template>

<script setup>
const route = useRoute()
const trackingCode = route.params.tracking_code
</script>

<!--
  این صفحه پس از ثبت موفق سفارش نمایش داده می‌شود و کد رهگیری سفارش را به کاربر نشان می‌دهد.
  در صورت نیاز می‌توانید اطلاعات بیشتر سفارش را با useFetch و tracking_code دریافت و نمایش دهید.
--> 