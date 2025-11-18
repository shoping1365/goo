<template>
  <!-- نمایش خلاصه آدرس انتخاب‌شده -->
  <div class="border-2 border-teal-500 rounded-xl p-6 bg-white flex flex-col gap-2">
    <!-- جزئیات آدرس -->
    <p class="text-gray-900 font-medium leading-relaxed">{{ address.full_address }}</p>
    <p v-if="address.province || address.city" class="text-gray-600 text-sm">
      {{ address.province }} {{ address.city ? '، ' + address.city : '' }}
    </p>
    <p class="text-gray-600 text-sm">کد پستی: {{ address.postal_code }}</p>
    <p class="text-gray-600 text-sm">{{ address.recipient_name }} | {{ address.recipient_mobile }}</p>
    <p v-if="address.phone" class="text-gray-600 text-sm">تلفن ثابت: {{ address.phone }}</p>

    <button @click="$emit('edit')" class="self-start text-sm text-[#1a2341] font-bold flex items-center gap-1 mt-2 hover:underline">
      <span>تغییر آدرس</span>
    </button>
  </div>
</template>

<script setup lang="ts">
interface Props {
  address: any
}

defineProps<Props>()
</script>

<style scoped>
/* بدون استایل اضافی */
</style> 