<template>
  <div class="p-6" dir="rtl">
    <h1 class="text-2xl font-bold mb-6">تنظیمات امنیت سامانه</h1>

    <div v-if="loading" class="text-center py-10">در حال بارگذاری...</div>
    <div v-else>
      <div class="space-y-4 max-w-md">
        <div v-for="rule in rules" :key="rule.key" class="flex items-center justify-between p-6 bg-white rounded-lg border">
          <div class="text-gray-800 font-medium">{{ rule.label }}</div>
          <label class="inline-flex items-center cursor-pointer">
            <input type="checkbox" class="sr-only peer" :checked="!!model[rule.key]" @change="toggleRule(rule.key)" />
            <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer peer-checked:bg-green-600 transition"></div>
          </label>
        </div>
      </div>

      <!-- غیر فعال سازی کلی -->
      <div class="mt-8 bg-yellow-50 p-6 rounded-lg max-w-md">
        <h2 class="font-semibold mb-2">غیرفعال‌سازی موقت همه قوانین</h2>
        <div class="flex items-center space-x-2 space-x-reverse">
          <input v-model.number="tempHours" type="number" min="1" max="24" class="w-20 px-2 py-1 border rounded" />
          <span>ساعت</span>
          <button @click="disableTemporarily" class="px-4 py-1 bg-red-600 text-white rounded hover:bg-red-700">اعمال</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
</script>

<script setup lang="ts">
import { reactive, ref, watchEffect } from 'vue';
import { useSecuritySettings, type SecuritySettings } from '~/composables/useSecuritySettings';

// meta
definePageMeta({ layout: 'admin-main' })

const { settings, loading, fetchSettings, saveSettings } = useSecuritySettings()

const model = reactive<SecuritySettings>({
  csrf: true,
  rateLimiter: true,
  xss: true,
  frameOptions: true,
  hsts: true,
})

const rules: Array<{ key: keyof SecuritySettings; label: string }> = [
  { key: 'csrf', label: 'محافظت CSRF' },
  { key: 'rateLimiter', label: 'محدودکننده نرخ' },
  { key: 'xss', label: 'محافظت XSS' },
  { key: 'frameOptions', label: 'X-Frame-Options' },
  { key: 'hsts', label: 'HSTS' },
]

watchEffect(() => {
  if (settings.value) {
    Object.assign(model, settings.value)
  }
})

const toggleRule = (key: keyof SecuritySettings) => {
  const value = model[key] as boolean
  const newValue = !value
  // استفاده از type assertion برای تغییر مقدار reactive
  Object.assign(model, { [key]: newValue } as Partial<SecuritySettings>)
  saveSettings({ [key]: newValue })
}

// غیرفعال‌سازی موقت همه قوانین برای چند ساعت
const tempHours = ref(2)
const disableTemporarily = () => {
  const until = Date.now() + tempHours.value * 60 * 60 * 1000
  saveSettings({ disabledUntil: until })
  // همه قوانین را خاموش نشان دهیم
  rules.forEach(r => {
    Object.assign(model, { [r.key]: false } as Partial<SecuritySettings>)
  })
}

fetchSettings()
</script>

<style scoped>
/* سوییچ ساده */
input[type='checkbox'] + div {
  position: relative;
}
input[type='checkbox'] + div::after {
  content: '';
  position: absolute;
  top: 2px;
  left: 2px;
  width: 20px;
  height: 20px;
  background: #fff;
  border-radius: 50%;
  transition: transform 0.2s;
}
input[type='checkbox']:checked + div::after {
  transform: translateX(20px);
}
</style>
