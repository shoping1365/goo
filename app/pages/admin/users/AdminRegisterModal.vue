<template>
  <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
    <div class="bg-white rounded-lg shadow-lg px-4 py-4 w-full max-w-md relative">
      <button class="absolute left-4 topx-4 py-4 text-gray-400 hover:text-gray-600 text-xl" @click="$emit('close')">&times;</button>
      <h2 class="text-lg font-bold mb-4 text-center bg-blue-100 text-blue-700 rounded-lg py-2">ثبت کاربر جدید</h2>
      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نام کامل</label>
          <input v-model="form.fullName" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نام کاربری</label>
          <input v-model="form.username" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">ایمیل</label>
          <input v-model="form.email" type="email" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">شماره موبایل</label>
          <input v-model="form.mobile" type="tel" required pattern="[0-9]{11}" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نقش</label>
          <select v-model="form.role" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500">
            <option value="developer">توسعه‌دهنده</option>
            <option value="customer">مشتری</option>
          </select>
        </div>
        <div v-if="form.role === 'developer'">
          <label class="block text-sm font-medium text-gray-700 mb-1">رمز عبور</label>
          <input v-model="form.password" type="password" required class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500" />
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button type="button" class="px-4 py-2 text-pink-700 bg-pink-100 rounded-lg hover:bg-pink-200 transition-colors" @click="$emit('close')">انصراف</button>
          <button type="submit" :disabled="loading" class="px-4 py-2 text-green-800 bg-green-100 rounded-lg hover:bg-green-200 transition-colors disabled:opacity-50">{{ loading ? 'در حال ثبت...' : 'ثبت کاربر' }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{ open: boolean }>();
const emit = defineEmits(['close', 'success']);

interface RegisterForm {
  fullName: string;
  username: string;
  email: string;
  mobile: string;
  role: string;
  password?: string;
}

const loading = ref(false);
const form = ref<RegisterForm>({
  fullName: '',
  username: '',
  email: '',
  mobile: '',
  role: 'customer',
  password: '',
});

async function handleSubmit() {
  try {
    loading.value = true;
    const payload: RegisterForm = { ...form.value };
    if (payload.role === 'customer') {
      delete payload.password;
    }
    // ارسال درخواست به سرور
    await $fetch('/api/admin/users/admin-register', {
      method: 'POST',
      body: payload,
    });
    emit('success');
    emit('close');
    form.value = { fullName: '', username: '', email: '', mobile: '', role: 'customer', password: '' };
  } catch {
    // نمایش خطا به کاربر (در صورت نیاز)
    alert('خطا در ثبت کاربر جدید!');
  } finally {
    loading.value = false;
  }
}
</script> 
