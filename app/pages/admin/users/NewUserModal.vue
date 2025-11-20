<template>
  <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
    <div class="bg-white rounded-lg shadow-lg px-4 py-4 w-full max-w-md relative">
      <button class="absolute left-4 topx-4 py-4 text-gray-400 hover:text-gray-600 text-xl" @click="$emit('close')">&times;</button>
      <h2 class="text-lg font-bold mb-4 text-center">ثبت کاربر جدید</h2>
      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">نام و نام خانوادگی</label>
          <input
            v-model="form.full_name"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">شماره موبایل</label>
          <input
            v-model="form.mobile"
            type="tel"
            required
            pattern="[0-9]{11}"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
          />
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button
            type="button"
            class="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200"
            @click="$emit('close')"
          >
            انصراف
          </button>
          <button
            type="submit"
            :disabled="loading"
            class="px-4 py-2 text-white bg-primary-600 rounded-lg hover:bg-primary-700 disabled:opacity-50"
          >
            {{ loading ? 'در حال ثبت...' : 'ثبت کاربر' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown }) => Promise<T>
</script>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
  open: boolean;
}>();

const emit = defineEmits(['close', 'success']);

const loading = ref(false);
const form = ref({
  full_name: '',
  mobile: '',
});

async function handleSubmit() {
  try {
    loading.value = true;
    await $fetch('/api/users', {
      method: 'POST',
      body: form.value
    });
    emit('success');
    emit('close');
    form.value = {
      full_name: '',
      mobile: '',
    };
  } catch (error) {
    // console.error('Error registering user:', error);
  } finally {
    loading.value = false;
  }
}
</script> 
