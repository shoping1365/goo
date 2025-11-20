<template>
  <div class="bg-gray-50 rounded-lg p-6 mb-6">
    <h4 class="font-medium text-gray-900 mb-3">پیوند سفارشی</h4>
    <div class="space-y-3">
      <input
        v-model="localCustomLink.title"
        type="text"
        placeholder="عنوان پیوند"
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <input
        v-model="localCustomLink.url"
        type="text"
        placeholder="آدرس پیوند (مثال: /about)"
        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
        :disabled="!localCustomLink.title || !localCustomLink.url"
        class="w-full bg-green-200 hover:bg-green-300 text-green-900 px-4 py-2 rounded-lg font-medium transition-colors duration-150 border border-green-300 disabled:opacity-50 disabled:cursor-not-allowed"
        style="box-shadow: 0 2px 12px 0 rgba(34,197,94,0.10);"
        @click="addLink"
      >
        افزودن پیوند
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  customLink: Object
})

const emit = defineEmits(['add-custom-link'])

const localCustomLink = ref({
  title: '',
  url: ''
})

const addLink = () => {
  if (localCustomLink.value.title && localCustomLink.value.url) {
    emit('add-custom-link', { ...localCustomLink.value })
    localCustomLink.value = { title: '', url: '' }
  }
}

watch(() => props.customLink, (newVal) => {
  if (newVal) {
    localCustomLink.value = { ...newVal }
  }
}, { deep: true })
</script>
