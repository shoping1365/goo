<template>
  <div class="w-full max-w-md">
    <div class="flex flex-row items-center justify-between mb-4">
      <h3 class="font-bold">مدیریت معیارهای امتیازدهی</h3>
      <button @click="addCritDialog = true" class="px-2 py-1 text-xs rounded bg-blue-600 text-white">افزودن معیار</button>
    </div>
    <table class="min-w-full text-sm text-right border border-gray-200 rounded-lg overflow-hidden">
      <thead class="bg-gray-50">
        <tr>
          <th class="px-2 py-2 border-b">عنوان</th>
          <th class="px-2 py-2 border-b">امتیاز</th>
          <th class="px-2 py-2 border-b">فعال است؟</th>
          <th class="px-2 py-2 border-b"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="crit,idx in criteria" :key="crit.key">
          <td class="px-2 py-2 border-b">{{ crit.label }}</td>
          <td class="px-2 py-2 border-b">{{ crit.point }}</td>
          <td class="px-2 py-2 border-b">
            <input type="checkbox" v-model="crit.active" @change="$emit('editCriterion',{idx, ...crit})"/>
          </td>
          <td class="px-2 py-2 border-b">
            <button class="text-blue-600 text-xs mr-2" @click="startEdit(idx)">ویرایش</button>
            <button class="text-red-600 text-xs" @click="$emit('removeCriterion',crit.key)">حذف</button>
          </td>
        </tr>
      </tbody>
    </table>
    <!-- افزودن/ویرایش -->
    <div v-if="addCritDialog || editIndex !== null" class="p-6 bg-gray-50 mt-4 rounded">
      <label>عنوان:</label>
      <input v-model="critForm.label" class="input" />
      <label>امتیاز:</label>
      <input v-model.number="critForm.point" type="number" class="input" />
      <div class="mt-2 flex gap-2">
        <button @click="saveCriterion" class="bg-blue-500 text-white px-2 rounded">ذخیره</button>
        <button @click="resetForm" class="bg-gray-200 px-2 rounded">انصراف</button>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue'

interface Criterion {
  key: string
  label: string
  point: number
  active: boolean
}

const props = defineProps<{ criteria: Criterion[] }>()
const emits = defineEmits<{
  addCriterion: [criterion: Criterion]
  editCriterion: [data: { idx: number } & Criterion]
  removeCriterion: [key: string]
}>()

const addCritDialog = ref(false)
const editIndex = ref<number | null>(null)
const critForm = ref({ label:'', point:0, active:true })

function startEdit(idx: number) {
  editIndex.value = idx
  const crit = props.criteria[idx]
  if (crit) {
    Object.assign(critForm.value, crit)
  }
}
function resetForm() {
  editIndex.value = null
  addCritDialog.value = false
  critForm.value = { label:'', point:0, active:true }
}
function saveCriterion() {
  if(addCritDialog.value) {
    emits('addCriterion', {...critForm.value, key:Math.random().toString(36).slice(2)} as Criterion)
  } else if(editIndex.value !== null) {
    emits('editCriterion', {idx:editIndex.value, ...critForm.value} as { idx: number } & Criterion)
  }
  resetForm()
}
watch(addCritDialog, val=>{ if(val) critForm.value = { label:'', point:0, active:true } })
</script>
<style scoped>
.input { border:1px solid #ddd; border-radius:4px; padding:2px 8px; margin:3px 0;}
</style>
