<template>
  <div class="w-full max-w-md">
    <div class="flex flex-row items-center justify-between mb-4">
      <h3 class="font-bold">مدیریت سطوح کاربران</h3>
      <button @click="addLevelDialog = true" class="px-2 py-1 text-xs rounded bg-blue-600 text-white">افزودن سطح</button>
    </div>
    <table class="min-w-full text-sm text-right border border-gray-200 rounded-lg overflow-hidden">
      <thead class="bg-gray-50">
        <tr>
          <th class="px-4 py-2 border-b">عنوان</th>
          <th class="px-4 py-2 border-b">امتیاز لازم</th>
          <th class="px-4 py-2 border-b">مزایا</th>
          <th class="px-4 py-2 border-b"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="level,idx in levels" :key="level.key">
          <td class="px-4 py-2 border-b font-bold">{{ level.name }}</td>
          <td class="px-4 py-2 border-b">{{ level.minScore }}+</td>
          <td class="px-4 py-2 border-b">{{ level.benefits }}</td>
          <td class="px-4 py-2 border-b">
            <button class="text-blue-600 text-xs mr-2" @click="startEdit(idx)">ویرایش</button>
            <button class="text-red-600 text-xs" @click="$emit('removeLevel',level.key)">حذف</button>
          </td>
        </tr>
      </tbody>
    </table>
    <!-- افزودن/ویرایش -->
    <div v-if="addLevelDialog || editIndex !== null" class="p-6 bg-gray-50 mt-4 rounded">
      <label>عنوان:</label>
      <input v-model="levelForm.name" class="input" />
      <label>امتیاز لازم:</label>
      <input v-model.number="levelForm.minScore" type="number" min="0" class="input" />
      <label>مزایا:</label>
      <input v-model="levelForm.benefits" class="input" />
      <div class="mt-2 flex gap-2">
        <button @click="saveLevel" class="bg-blue-500 text-white px-2 rounded">ذخیره</button>
        <button @click="resetForm" class="bg-gray-200 px-2 rounded">انصراف</button>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue'

interface Level {
  key: string
  name: string
  minScore: number
  benefits: string
}

const props = defineProps<{ levels: Level[] }>()
const emits = defineEmits<{
  addLevel: [level: Level]
  editLevel: [data: { idx: number } & Level]
  removeLevel: [key: string]
}>()

const addLevelDialog = ref(false)
const editIndex = ref<number | null>(null)
const levelForm = ref({ name:'', minScore:0, benefits:'' })

function startEdit(idx: number) {
  editIndex.value = idx
  const level = props.levels[idx]
  if (level) {
    Object.assign(levelForm.value, level)
  }
}
function resetForm() {
  editIndex.value = null
  addLevelDialog.value = false
  levelForm.value = { name:'', minScore:0, benefits:'' }
}
function saveLevel() {
  if(addLevelDialog.value) {
    emits('addLevel', {...levelForm.value, key:Math.random().toString(36).slice(2)} as Level)
  } else if(editIndex.value !== null) {
    emits('editLevel', {idx:editIndex.value, ...levelForm.value} as { idx: number } & Level)
  }
  resetForm()
}
watch(addLevelDialog, val=>{ if(val) levelForm.value = { name:'', minScore:0, benefits:'' } })
</script>
<style scoped>
.input { border:1px solid #ddd; border-radius:4px; padding:2px 8px; margin:3px 0;}
</style>
