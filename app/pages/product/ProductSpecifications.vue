<template>
  <div class="space-y-6" dir="rtl">
    <!-- مشخصات گروه‌بندی شده -->
    <div v-if="groupedSpecs.length" class="space-y-4">
      <div
        v-for="group in groupedSpecs"
        :key="group.name"
        class="bg-white rounded-lg border border-gray-200 overflow-hidden"
      >
        <div class="overflow-x-auto">
          <table class="w-full">
            <tbody class="divide-y divide-gray-200">
              <tr
                v-for="spec in group.specs"
                :key="spec.id"
                class="hover:bg-gray-50 transition-colors"
              >
                <td class="px-6 py-4 w-1/3 text-sm font-medium text-gray-900 bg-gray-50">
                  {{ spec.name }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-700">
                  <span v-if="spec.unit" class="inline-flex items-center gap-1">
                    {{ spec.value }}
                    <span class="text-xs text-gray-500">{{ spec.unit }}</span>
                  </span>
                  <span v-else>{{ spec.value }}</span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- مشخصات ساده (بدون گروه‌بندی) -->
    <div v-else-if="specifications?.length" class="bg-white rounded-lg border border-gray-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مشخصه
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                مقدار
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr
              v-for="spec in filteredSpecs"
              :key="spec.id || spec.name"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 text-sm font-medium text-gray-900 bg-gray-50">
                {{ spec.name }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-700">
                <span v-if="spec.unit" class="inline-flex items-center gap-1">
                  {{ spec.value }}
                  <span class="text-xs text-gray-500">{{ spec.unit }}</span>
                </span>
                <span v-else>{{ spec.value }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- حالت خالی -->
    <div v-else class="bg-gray-50 rounded-lg p-12 text-center">
      <svg class="w-16 h-16 text-gray-300 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
      </svg>
      <h3 class="text-lg font-medium text-gray-900 mb-2">مشخصات فنی موجود نیست</h3>
      <p class="text-gray-500">مشخصات فنی این محصول هنوز ثبت نشده است.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';

interface Props {
  product: Record<string, unknown>
}

const props = defineProps<Props>()

// State
const searchQuery = ref('')
const expandedGroups = reactive<Record<string, boolean>>({})

// Computed
const specifications = computed(() => {
  return props.product?.specifications || []
})

const groupedSpecs = computed(() => {
  const groups: { name: string; specs: Record<string, unknown>[] }[] = []
  const groupMap = new Map()

  specifications.value.forEach((spec: Record<string, unknown>) => {
    const groupName = (spec.group as string) || 'عمومی'
    if (!groupMap.has(groupName)) {
      const group = { name: groupName, specs: [] }
      groups.push(group)
      groupMap.set(groupName, group)
    }
    groupMap.get(groupName).specs.push(spec)
  })

  return groups
})

const filteredSpecs = computed(() => {
  if (!searchQuery.value) return specifications.value
  
  const query = searchQuery.value.toLowerCase()
  return specifications.value.filter((spec: Record<string, unknown>) =>
    (spec.name as string)?.toLowerCase().includes(query) ||
    (spec.value as string)?.toLowerCase().includes(query)
  )
})

// const filteredGroups = computed(() => {
//   if (!searchQuery.value) return groupedSpecs.value

//   const query = searchQuery.value.toLowerCase()
//   return groupedSpecs.value
//     .map(group => ({
//       ...group,
//       specs: group.specs.filter(spec =>
//         (spec.name as string)?.toLowerCase().includes(query) ||
//         (spec.value as string)?.toLowerCase().includes(query)
//       )
//     }))
//     .filter(group => group.specs.length > 0)
// })

// Methods
// const toggleGroup = (groupName: string) => {
//   expandedGroups[groupName] = !expandedGroups[groupName]
// }

// const expandAll = () => {
//   groupedSpecs.value.forEach(group => {
//     expandedGroups[group.name] = true
//   })
// }

// const collapseAll = () => {
//   groupedSpecs.value.forEach(group => {
//     expandedGroups[group.name] = false
//   })
// }

// Initialize expanded state
groupedSpecs.value.forEach(group => {
  if (!(group.name in expandedGroups)) {
    expandedGroups[group.name] = true
  }
})
</script> 
