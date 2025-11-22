<template>
  <div class="bg-white rounded-lg shadow-md border border-gray-200 overflow-hidden">
    <div class="bg-gradient-to-r from-slate-50 to-gray-50 px-4 py-3 border-b border-gray-200">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="w-6 h-6 bg-slate-600 rounded-md flex items-center justify-center mr-2 ml-2">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
            </svg>
          </div>
          <h3 class="text-sm font-semibold text-gray-900">مدیریت مشخصات فنی</h3>
        </div>
        
        <div class="flex items-center space-x-2 space-x-reverse">
          <div v-if="selectedAttributes && selectedAttributes.length > 0" class="flex items-center space-x-2 space-x-reverse bg-red-50 rounded-md px-2 py-1.5 border border-red-200">
            <span class="text-xs text-red-700 font-medium">{{ selectedAttributes.length }} مورد انتخاب شده</span>
            <button
              v-if="hasPermission('attribute.delete')"
              class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-100 hover:bg-red-200 transition-colors"
              @click="$emit('bulk-delete')"
            >
              🗑️ حذف انتخاب شده‌ها
            </button>
          </div>
          
          <!-- Search -->
          <div class="relative">
            <input 
              :value="searchQuery"
              type="text"
              class="block w-56 pl-8 pr-3 py-1.5 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-xs shadow-sm" 
              placeholder="جستجو نام مشخصه فنی" 
              dir="rtl"
              @input="(event) => $emit('update:searchQuery', (event.target as HTMLInputElement).value)"
            />
            <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
              <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Table Content -->
    <div class="overflow-hidden">
      <div class="overflow-x-auto">
        <div dir="rtl">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                  <input 
                    type="checkbox" 
                    :checked="isAllSelected"
                    class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                    @change="$emit('toggle-select-all')"
                  />
                </th>
                <th scope="col" class="py-2 text-right text-xs font-medium text-gray-600 uppercase tracking-wider pr-4">
                  ردیف
                </th>
                <th scope="col" class="px-4 py-2 text-right text-xs font-medium text-gray-600 uppercase tracking-wider">
                  نام مشخصه فنی
                </th>
                <th scope="col" class="px-4 py-2 text-center text-xs font-medium text-gray-600 uppercase tracking-wider">
                  عملیات
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="(attribute, index) in paginatedAttributes" :key="attribute.id" class="hover:bg-gray-50 transition-colors">
                                       <!-- Checkbox -->
                       <td class="px-4 py-3 whitespace-nowrap text-center">
                         <input 
                           type="checkbox" 
                           :value="attribute.id"
                           :checked="selectedAttributes.includes(attribute.id)"
                           class="h-3 w-3 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                           @change="(event) => handleCheckboxChange(attribute.id, (event.target as HTMLInputElement).checked)"
                         />
                       </td>

                <!-- Row Number -->
                <td class="py-3 whitespace-nowrap text-right pr-4">
                  <span class="inline-flex items-center w-6 h-6 bg-gray-100 text-gray-700 text-xs font-medium rounded-full text-right pr-4">{{ (currentPage - 1) * itemsPerPage + index + 1 }}</span>
                </td>

                <!-- Name -->
                <td class="py-3 whitespace-nowrap text-xs font-medium text-gray-900 text-right pr-4">
                  {{ attribute.name }}
                </td>

                <!-- Actions -->
                <td class="px-4 py-3 whitespace-nowrap text-center">
                  <div class="flex flex-row items-center justify-center gap-x-4">
                    <button 
                      v-if="hasPermission('attribute.update')"
                      class="inline-flex items-center px-2 py-1 border border-blue-300 text-xs font-medium rounded-sm text-blue-700 bg-blue-50 hover:bg-blue-100 transition-colors"
                      @click="$emit('edit', attribute)"
                    >
                      ✏️ ویرایش
                    </button>
                    <button 
                      v-if="hasPermission('attribute.delete')"
                      class="inline-flex items-center px-2 py-1 border border-red-300 text-xs font-medium rounded-sm text-red-700 bg-red-50 hover:bg-red-100 transition-colors mr-2"
                      @click="$emit('delete', attribute)"
                    >
                      🗑️ حذف
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
</script>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})
// تعریف interface برای مشخصات فنی
interface Attribute {
  id: string
  name: string
  order: number
  description?: string
  isActive: boolean
  createdAt: string
}

// تعریف props
interface Props {
  attributes: Attribute[]
  searchQuery: string
  selectedAttributes: string[]
  currentPage: number
  itemsPerPage: number
  paginatedAttributes: Attribute[]
  isAllSelected: boolean
}

// تعریف emits
interface Emits {
  (e: 'update:searchQuery', value: string): void
  (e: 'toggle-select-all'): void
  (e: 'bulk-delete'): void
  (e: 'edit', attribute: Attribute): void
  (e: 'delete', attribute: Attribute): void
  (e: 'update:selected-attributes', value: string[]): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

// استفاده از useAuth برای چک کردن پرمیژن
const { hasPermission } = useAuth()

// تابع مدیریت تغییر checkbox
const handleCheckboxChange = (attributeId: string, isChecked: boolean) => {
  const newSelectedAttributes = [...props.selectedAttributes]
  
  if (isChecked) {
    if (!newSelectedAttributes.includes(attributeId)) {
      newSelectedAttributes.push(attributeId)
    }
  } else {
    const index = newSelectedAttributes.indexOf(attributeId)
    if (index > -1) {
      newSelectedAttributes.splice(index, 1)
    }
  }
  
  emit('update:selected-attributes', newSelectedAttributes)
}
</script>
