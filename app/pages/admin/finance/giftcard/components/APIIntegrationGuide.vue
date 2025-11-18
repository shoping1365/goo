<template>
  <div class="space-y-8">
    <!-- بخش API ها -->
    <section>
      <h2 class="text-xl font-bold mb-4">APIهای مورد نیاز گیفت کارت</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="api in apis" :key="api.title" class="bg-white rounded-lg shadow p-6 border border-gray-200">
          <h3 class="text-lg font-semibold mb-2">{{ api.title }}</h3>
          <p class="text-gray-600 mb-2">{{ api.desc }}</p>
          <div class="text-xs bg-yellow-100 text-yellow-800 rounded px-2 py-1 inline-block mb-2">در حال توسعه</div>
          <div class="bg-gray-50 rounded p-3 mt-2">
            <span class="font-mono text-xs text-gray-700">{{ api.example }}</span>
          </div>
        </div>
      </div>
    </section>

    <!-- بخش یکپارچه‌سازی -->
    <section>
      <h2 class="text-xl font-bold mb-4">یکپارچه‌سازی با سیستم‌های دیگر</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div v-for="sys in integrations" :key="sys.title" class="bg-white rounded-lg shadow p-6 border border-gray-200 flex items-center">
          <div class="flex-1">
            <h3 class="text-lg font-semibold mb-2">{{ sys.title }}</h3>
            <p class="text-gray-600 mb-2">{{ sys.desc }}</p>
            <div class="text-xs bg-yellow-100 text-yellow-800 rounded px-2 py-1 inline-block">در حال توسعه</div>
          </div>
          <div class="ml-4">
            <span class="inline-block w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center">
              <span class="text-2xl">{{ sys.icon }}</span>
            </span>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// توضیحات و نمونه‌های API
const apis = [
  {
    title: 'API ایجاد کارت',
    desc: 'ایجاد گیفت کارت جدید با اطلاعات کامل کارت و گیرنده.',
    example: 'POST /api/giftcards'
  },
  {
    title: 'API بررسی اعتبار',
    desc: 'بررسی وضعیت و اعتبار گیفت کارت بر اساس کد کارت.',
    example: 'GET /api/giftcards/validate?code=XXXX'
  },
  {
    title: 'API استفاده از کارت',
    desc: 'ثبت استفاده از گیفت کارت و کاهش اعتبار آن.',
    example: 'POST /api/giftcards/use'
  },
  {
    title: 'API گزارش‌گیری',
    desc: 'دریافت گزارش‌های مختلف مربوط به گیفت کارت‌ها.',
    example: 'GET /api/giftcards/reports'
  }
]

// توضیحات یکپارچه‌سازی
const integrations = [
  {
    title: 'اتصال به سیستم پرداخت',
    desc: 'امکان استفاده از گیفت کارت به عنوان روش پرداخت یا شارژ مجدد کارت.',
    icon: '💳'
  },
  {
    title: 'اتصال به سیستم انبار',
    desc: 'مدیریت موجودی کارت‌های فیزیکی در انبار و ثبت تراکنش‌های ورود/خروج.',
    icon: '📦'
  },
  {
    title: 'اتصال به CRM',
    desc: 'ثبت و مشاهده اطلاعات مشتریان و تراکنش‌های مرتبط با گیفت کارت.',
    icon: '👥'
  },
  {
    title: 'اتصال به سیستم حسابداری',
    desc: 'ثبت تراکنش‌های مالی مرتبط با گیفت کارت در سیستم حسابداری.',
    icon: '📊'
  }
]
</script>

<!--
  مستندسازی:
  این کامپوننت صرفاً برای نمایش راهنمای API و وضعیت یکپارچه‌سازی است و هیچ اتصال واقعی به بک‌اند ندارد.
  تمام توضیحات به زبان فارسی و با رعایت ریسپانسیو ارائه شده است.
--> 
