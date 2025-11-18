<template>
  <div class="rounded-xl bg-white shadow p-6" dir="rtl">
    <div class="flex items-center justify-between mb-3">
      <h3 class="font-semibold">حرکات موجودی</h3>
      <div class="flex gap-2 text-xs">
        <select v-model="filters.type" class="rounded px-2 py-1 bg-white shadow">
          <option value="">همه نوع‌ها</option>
          <option value="in">افزایش</option>
          <option value="out">کاهش</option>
          <option value="transfer">انتقال</option>
          <option value="adjust">اصلاح</option>
        </select>
        <input v-model.number="filters.product" placeholder="ID محصول" class="rounded px-2 py-1 bg-white shadow w-28"/>
        <input v-model.number="filters.warehouse" placeholder="ID انبار" class="rounded px-2 py-1 bg-white shadow w-28"/>
        <button @click="load" class="rounded px-3 py-1 bg-gray-100 hover:bg-gray-200">اعمال فیلتر</button>
      </div>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full text-sm">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-3 py-2 text-right">زمان</th>
            <th class="px-3 py-2 text-right">محصول</th>
            <th class="px-3 py-2 text-right">انبار</th>
            <th class="px-3 py-2 text-right">تغییر</th>
            <th class="px-3 py-2 text-right">نوع</th>
            <th class="px-3 py-2 text-right">یادداشت</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="m in items" :key="m.id" class="border-b">
            <td class="px-3 py-2">{{ new Date(m.created_at || m.createdAt).toLocaleString('fa-IR') }}</td>
            <td class="px-3 py-2">#{{ m.product_id }}</td>
            <td class="px-3 py-2">#{{ m.warehouse_id }}</td>
            <td class="px-3 py-2" :class="m.delta>=0?'text-green-600':'text-red-600'">{{ m.delta }}</td>
            <td class="px-3 py-2">{{ m.type }}</td>
            <td class="px-3 py-2 truncate max-w-[400px]" :title="m.notes">{{ m.notes || '-' }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  
</template>

<script setup>
import { ref, onMounted } from 'vue'

const items = ref([])
const filters = ref({ type: '', product: undefined, warehouse: undefined })

async function load() {
  try {
    const q = new URLSearchParams()
    if (filters.value.type) q.append('type', filters.value.type)
    if (filters.value.product) q.append('product_id', String(filters.value.product))
    if (filters.value.warehouse) q.append('warehouse_id', String(filters.value.warehouse))
    items.value = await $fetch(`/api/admin/inventory-movements?${q.toString()}`)
  } catch (e) {
    items.value = []
  }
}

onMounted(load)
</script>
