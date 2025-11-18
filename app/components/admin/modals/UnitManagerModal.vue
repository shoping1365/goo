<template>
  <transition name="fade-scale">
    <div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40 backdrop-blur-sm">
      <div class="relative w-full max-w-xl bg-white rounded-2xl shadow-2xl border border-blue-200 p-6 animate-fade-in" dir="rtl">
        <!-- Close -->
        <button @click="$emit('close')" class="absolute left-4 top-6 text-gray-400 hover:text-red-500 text-2xl focus:outline-none">×</button>
        <!-- Title -->
        <h3 class="text-xl font-extrabold text-gray-800 mb-4 text-center">مدیریت واحدهای اندازه‌گیری</h3>
        <hr class="mb-4 border-blue-100"/>

        <div class="max-h-[60vh] overflow-y-auto" dir="ltr">
          <table class="min-w-full divide-y divide-gray-200 text-sm" dir="rtl">
            <thead class="bg-gray-50 sticky top-0">
              <tr>
                <th class="px-4 py-2">کد</th>
                <th class="px-4 py-2">برچسب</th>
                <th class="px-4 py-2 text-center">عملیات</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(u, idx) in units" :key="idx" class="divide-x divide-gray-200">
                <td class="px-3 py-2 whitespace-nowrap"><input v-model="u.value" class="w-full border border-gray-300 rounded-md px-2 py-1 text-xs"/></td>
                <td class="px-3 py-2 whitespace-nowrap"><input v-model="u.label" class="w-full border border-gray-300 rounded-md px-2 py-1 text-xs"/></td>
                <td class="px-3 py-2 text-center">
                  <button 
                    v-if="canDeleteUnit"
                    @click="remove(idx)" 
                    class="text-red-600 hover:text-red-800 text-sm"
                  >
                    ��️
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
          <button @click="addRow" class="mt-3 inline-flex items-center px-3 py-1.5 bg-emerald-500 hover:bg-emerald-600 text-white text-sm rounded shadow">➕ افزودن ردیف</button>
        </div>

        <div class="mt-6 flex flex-row-reverse gap-3">
          <button @click="save" class="inline-flex items-center px-6 py-2 rounded-lg text-white bg-gradient-to-r from-blue-500 to-blue-700 hover:from-blue-600 hover:to-blue-800 shadow-md font-bold">ذخیره</button>
          <button @click="$emit('close')" class="inline-flex items-center px-6 py-2 rounded-lg text-gray-700 bg-gray-200 hover:bg-gray-300 shadow font-bold">انصراف</button>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { ref, computed } from 'vue'


const props = defineProps({
  modelValue: { type: Array, default: () => [] }
})
const emit = defineEmits(['close', 'saved'])

const units = ref(JSON.parse(JSON.stringify(props.modelValue)))

// const { user, hasPermission } = useAuth() // احراز هویت غیرفعال شده است

const canDeleteUnit = computed(() => hasPermission('unit.delete'))

const addRow = () => {
  units.value.push({ value: '', label: '' })
}

const remove = (idx) => {
  units.value.splice(idx, 1)
}

const save = () => {
  // فیلتر ورودی‌های خالی
  const cleaned = units.value.filter(u => u.value.trim() && u.label.trim())
  localStorage.setItem('measurementUnits', JSON.stringify(cleaned))
  emit('saved')
  emit('close')
}
</script>

<style scoped>
.fade-scale-enter-active, .fade-scale-leave-active { transition: all 0.2s ease; }
.fade-scale-enter-from, .fade-scale-leave-to { opacity: 0; transform: scale(0.95); }
</style> 
