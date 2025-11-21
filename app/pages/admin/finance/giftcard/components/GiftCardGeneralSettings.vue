<template>
  <div class="bg-white rounded-lg border border-gray-200 p-6 max-w-xl mx-auto">
    <h3 class="text-lg font-semibold text-gray-900 mb-6">تنظیمات عمومی گیفت کارت</h3>
    <form class="space-y-6" @submit.prevent="handleSubmit">
      <!-- فعال/غیرفعال کردن سیستم -->
      <div class="flex items-center justify-between">
        <label class="text-sm font-medium text-gray-700">فعال بودن سیستم گیفت کارت</label>
        <input v-model="form.enabled" type="checkbox" class="h-5 w-5 text-green-600 focus:ring-green-500 border-gray-300 rounded" />
      </div>
      <!-- حداقل و حداکثر مبلغ کارت -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداقل مبلغ کارت <span class="text-red-500">*</span></label>
          <input
            v-model.number="form.minAmount"
            type="number"
            min="10000"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: ۱۰۰۰۰"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر مبلغ کارت <span class="text-red-500">*</span></label>
          <input
            v-model.number="form.maxAmount"
            type="number"
            :min="form.minAmount"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: ۵۰۰۰۰۰۰"
          />
        </div>
      </div>
      <!-- مدت اعتبار پیش‌فرض -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">مدت اعتبار پیش‌فرض <span class="text-red-500">*</span></label>
          <input
            v-model.number="form.defaultValidity"
            type="number"
            min="1"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="مثال: ۱۲"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">واحد مدت اعتبار</label>
          <select
            v-model="form.validityUnit"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="month">ماه</option>
            <option value="day">روز</option>
            <option value="year">سال</option>
          </select>
        </div>
      </div>
      <!-- محدودیت تعداد کارت در هر سفارش -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">حداکثر تعداد کارت در هر سفارش <span class="text-red-500">*</span></label>
        <input
          v-model.number="form.maxCardsPerOrder"
          type="number"
          min="1"
          required
          class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          placeholder="مثال: ۵"
        />
      </div>
      <!-- دکمه ذخیره -->
      <div class="flex justify-end">
        <button
          type="submit"
          :disabled="isSubmitting"
          class="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="isSubmitting">در حال ذخیره...</span>
          <span v-else>ذخیره تنظیمات</span>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
// متغیرهای reactive
const isSubmitting = ref(false)
const form = ref({
  enabled: true,
  minAmount: 10000,
  maxAmount: 5000000,
  defaultValidity: 12,
  validityUnit: 'month',
  maxCardsPerOrder: 5
})

const handleSubmit = async () => {
  // اعتبارسنجی
  if (form.value.minAmount < 10000) {
    alert('حداقل مبلغ کارت باید حداقل ۱۰٬۰۰۰ تومان باشد.')
    return
  }
  if (form.value.maxAmount <= form.value.minAmount) {
    alert('حداکثر مبلغ کارت باید بزرگتر از حداقل مبلغ باشد.')
    return
  }
  if (form.value.defaultValidity < 1) {
    alert('مدت اعتبار پیش‌فرض باید حداقل ۱ باشد.')
    return
  }
  if (form.value.maxCardsPerOrder < 1) {
    alert('حداکثر تعداد کارت در هر سفارش باید حداقل ۱ باشد.')
    return
  }
  isSubmitting.value = true
  try {
    // شبیه‌سازی ذخیره تنظیمات
    await new Promise(resolve => setTimeout(resolve, 1000))
    alert('تنظیمات با موفقیت ذخیره شد.')
  } catch (_e) {
    alert('خطا در ذخیره تنظیمات. لطفاً دوباره تلاش کنید.')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
/* استایل‌های خاص تنظیمات */
</style> 
