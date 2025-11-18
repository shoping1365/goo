<template>
  <div class="rounded-xl bg-white shadow overflow-x-auto">
    <table class="min-w-full text-sm">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-3 py-2"><input type="checkbox" :checked="allCurrentSelected" @change="$emit('toggleSelectAll', items.map(i=>i.id))"/></th>
          <th class="px-3 py-2 text-right select-none">
            <button class="inline-flex items-center gap-1" @click="$emit('sort','id')">
              <span>ID</span>
              <span v-if="sortKey==='id'" class="text-xs">{{ sortOrder==='asc' ? '▲' : '▼' }}</span>
            </button>
          </th>
          <th class="px-3 py-2 text-right select-none">
            <button class="inline-flex items-center gap-1" @click="$emit('sort','code')">
              <span>کد</span>
              <span v-if="sortKey==='code'" class="text-xs">{{ sortOrder==='asc' ? '▲' : '▼' }}</span>
            </button>
          </th>
          <th class="px-3 py-2 text-right select-none">
            <button class="inline-flex items-center gap-1" @click="$emit('sort','name')">
              <span>نام</span>
              <span v-if="sortKey==='name'" class="text-xs">{{ sortOrder==='asc' ? '▲' : '▼' }}</span>
            </button>
          </th>
          <th class="px-3 py-2 text-right select-none">
            <button class="inline-flex items-center gap-1" @click="$emit('sort','city')">
              <span>شهر</span>
              <span v-if="sortKey==='city'" class="text-xs">{{ sortOrder==='asc' ? '▲' : '▼' }}</span>
            </button>
          </th>
          <th class="px-3 py-2 text-right">آدرس</th>
          <th class="px-3 py-2 text-right select-none">
            <button class="inline-flex items-center gap-1" @click="$emit('sort','priority')">
              <span>اولویت</span>
              <span v-if="sortKey==='priority'" class="text-xs">{{ sortOrder==='asc' ? '▲' : '▼' }}</span>
            </button>
          </th>
          <th class="px-3 py-2 text-right">پیش‌فرض</th>
          <th class="px-3 py-2 text-right">فعال</th>
          <th class="px-3 py-2 text-right">اقدامات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="w in items" :key="w.id" class="border-b hover:bg-gray-50">
          <td class="px-3 py-2"><input type="checkbox" :checked="selectedIds.includes(w.id)" @change="$emit('toggleSelect', w.id)"/></td>
          <td class="px-3 py-2">{{ w.id }}</td>
          <td class="px-3 py-2 font-mono">{{ w.code }}</td>
          <td class="px-3 py-2 font-medium">{{ w.name }}</td>
          <td class="px-3 py-2">{{ w.city || '-' }}</td>
          <td class="px-3 py-2 truncate max-w-[300px]" :title="w.address">{{ w.address || '-' }}</td>
          <td class="px-3 py-2">{{ w.priority ?? 100 }}</td>
          <td class="px-3 py-2">
            <input type="radio" name="defaultWarehouse" :checked="w.is_default" @change="$emit('setDefault', w)"/>
          </td>
          <td class="px-3 py-2">
            <input type="checkbox" :checked="w.is_active" @change="$emit('toggleActive', w)"/>
          </td>
          <td class="px-3 py-2 flex gap-3">
            <button @click="$emit('edit', w)" class="text-blue-600 hover:underline">ویرایش</button>
            <button @click="$emit('remove', w)" class="text-red-600 hover:underline">حذف</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  items: { type: Array, default: () => [] },
  selectedIds: { type: Array, default: () => [] },
  sortKey: { type: String, default: 'id' },
  sortOrder: { type: String, default: 'desc' },
})

const allCurrentSelected = computed(() => props.items.length > 0 && props.items.every((i) => props.selectedIds.includes(i.id)))
</script>


