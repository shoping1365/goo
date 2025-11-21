<template>
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">نقش‌های کاربری</h1>
            <p class="text-sm text-gray-500 mt-1">مدیریت نقش‌های مختلف کاربران در سیستم</p>
          </div>
          
          <div class="flex items-center gap-3">
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
              @click="showAddRoleModal = true"
            >
              <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
              </svg>
              افزودن نقش جدید
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="p-6">
      <!-- آمار کارت‌ها -->
      <div class="mx-auto px-4 sm:px-6 lg:px-8 py-8 mb-8">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
            <TemplateCard 
              title="کل نقش‌ها" 
              :value="statistics.total_roles || 0" 
              variant="blue"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="کل کاربران" 
              :value="statistics.total_users || 0" 
              variant="green"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="مجوزهای فعال" 
              :value="statistics.total_permissions || 0" 
              variant="orange"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
                </svg>
              </template>
            </TemplateCard>

            <TemplateCard 
              title="نقش‌های فعال" 
              :value="statistics.active_roles || 0" 
              variant="cyan"
            >
              <template #icon>
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                </svg>
              </template>
            </TemplateCard>
        </div>
      </div>

      <!-- بخش مدیریت دسترسی‌ها -->
      <SectionCard>
        <div class="px-6 py-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">نقش‌های کاربری</h2>
          <p class="text-sm text-gray-500 mt-1">مدیریت نقش‌های مختلف کاربران در سیستم</p>
        </div>
        
        <div class="p-6">
          <div v-if="loading" class="text-center py-8">
            <div class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-blue-500 hover:bg-blue-400 transition ease-in-out duration-150 cursor-not-allowed">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال بارگذاری...
            </div>
          </div>
          
          <div v-else-if="roles.length === 0" class="text-center py-8">
            <svg class="w-12 h-12 text-gray-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
            <h3 class="text-lg font-medium text-gray-900 mb-2">هیچ نقشی وجود ندارد</h3>
            <p class="text-gray-500 mb-4">برای شروع، اولین نقش را ایجاد کنید</p>
            <button 
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700"
              @click="showAddRoleModal = true"
            >
              ایجاد نقش جدید
            </button>
          </div>
          
          <div v-else class="space-y-4">
            <div 
              v-for="role in roles" 
              :key="role.id"
              class="border border-gray-200 rounded-lg p-6 hover:shadow-md transition-shadow"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <div class="w-10 h-10 bg-gradient-to-r from-blue-400 to-blue-600 rounded-lg flex items-center justify-center">
                    <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                    </svg>
                  </div>
                  <div>
                    <h3 class="text-lg font-medium text-gray-900">{{ role.display_name }}</h3>
                    <p class="text-sm text-gray-500">{{ role.description }}</p>
                    <div class="flex items-center mt-1">
                      <span class="text-xs text-gray-400">{{ role.users_count }} کاربر</span>
                      <span class="mx-2 text-gray-300">•</span>
                      <span class="text-xs text-gray-400">{{ role.permissions_count }} مجوز</span>
                      <span class="mx-2 text-gray-300">•</span>
                      <span class="text-xs text-gray-400">نام: {{ role.name }}</span>
                    </div>
                  </div>
                </div>
                
                <div class="flex items-center space-x-2">
                  <button 
                    class="p-2 text-gray-400 hover:text-blue-600 transition-colors"
                    @click="editRole(role)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    class="p-2 text-gray-400 hover:text-red-600 transition-colors"
                    @click="deleteRole(role.id)"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </SectionCard>
    </div>

    <!-- Modal افزودن/ویرایش نقش -->
    <div v-if="showAddRoleModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
        <div class="px-6 py-4 border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ editingRole ? 'ویرایش نقش' : 'افزودن نقش جدید' }}
          </h3>
        </div>
        
        <div class="p-6">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نام نقش <span class="text-red-500">*</span>
                <span class="text-xs text-gray-500">(فقط حروف انگلیسی، اعداد و underscore)</span>
              </label>
              <input 
                v-model="roleForm.name"
                type="text"
                placeholder="مثال: admin, manager, operator"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                dir="ltr"
                :class="{ 'border-red-500': nameError }"
              >
              <p v-if="nameError" class="text-red-500 text-xs mt-1">{{ nameError }}</p>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                نام نمایشی <span class="text-red-500">*</span>
              </label>
              <input 
                v-model="roleForm.display_name"
                type="text"
                placeholder="مثال: مدیر سیستم"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                dir="rtl"
              >
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">توضیحات</label>
              <textarea 
                v-model="roleForm.description"
                rows="3"
                placeholder="توضیحات مربوط به این نقش..."
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                dir="rtl"
              ></textarea>
            </div>
          </div>
        </div>
        
        <div class="px-6 py-4 border-t border-gray-200 flex justify-end space-x-3 space-x-reverse">
          <button 
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-lg transition-colors"
            @click="closeModal"
          >
            انصراف
          </button>
          <button 
            :disabled="saving"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            @click="saveRole"
          >
            <span v-if="saving" class="inline-flex items-center">
              <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              در حال ذخیره...
            </span>
            <span v-else>{{ editingRole ? 'ویرایش' : 'افزودن' }}</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue';
import TemplateCard from '~/components/common/TemplateCard.vue';
import SectionCard from '~/pages/admin/shipping-methods/SectionCard.vue';

// تعریف definePageMeta و useHead برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string }) => void

// تعریف interface برای خطا
interface ApiError {
  data?: {
    code?: string
    message?: string
    details?: Record<string, unknown>
  }
}

// تعریف interface برای نقش
interface Role {
  id: number
  name: string
  display_name: string
  description: string
  users_count: number
  permissions_count: number
  is_active: boolean
  is_system: boolean
  priority: number
  created_at: string
  updated_at: string
}

// تعریف interface برای آمار
interface Statistics {
  total_roles: number
  total_users: number
  total_permissions: number
  active_roles: number
}

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'نقش‌های کاربری - پنل مدیریت'
})

// Reactive data
const roles = ref<Role[]>([])
const statistics = ref<Statistics>({
  total_roles: 0,
  total_users: 0,
  total_permissions: 0,
  active_roles: 0
})
const loading = ref(false)
const saving = ref(false)
const showAddRoleModal = ref(false)
const editingRole = ref<Role | null>(null)
const nameError = ref('')

// فرم نقش
const roleForm = reactive({
  name: '',
  display_name: '',
  description: ''
})

// Computed - kept for potential future use
const _isFormValid = computed(() => {
  return roleForm.name.trim() && roleForm.display_name.trim() && !nameError.value
})

// Methods
const validateRoleName = (name: string): string => {
  if (!name.trim()) return 'نام نقش الزامی است'
  
  // بررسی اینکه نام فقط شامل حروف انگلیسی، اعداد و underscore باشد
  const englishRegex = /^[a-zA-Z0-9_]+$/
  if (!englishRegex.test(name)) {
    return 'نام نقش باید فقط شامل حروف انگلیسی، اعداد و underscore باشد'
  }
  
  return ''
}

// تعریف interface برای پاسخ API
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

const fetchRoles = async () => {
  loading.value = true
  try {
    const response: ApiResponse<Role[]> = await $fetch('/api/roles')
    // فیلتر کردن نقش customer از لیست نمایشی
    roles.value = (response.data || []).filter(role => role.name !== 'customer')
  } catch (error) {
    console.error('خطا در دریافت نقش‌ها:', error)
    alert('خطا در دریافت نقش‌ها')
  } finally {
    loading.value = false
  }
}

const fetchStatistics = async () => {
  try {
    const response: ApiResponse<Statistics> = await $fetch('/api/roles/statistics')
    statistics.value = response.data || {
      total_roles: 0,
      total_users: 0,
      total_permissions: 0,
      active_roles: 0
    }
  } catch (error) {
    console.error('خطا در دریافت آمار:', error)
  }
}

const editRole = (role: Role) => {
  editingRole.value = role
  roleForm.name = role.name
  roleForm.display_name = role.display_name
  roleForm.description = role.description
  nameError.value = ''
  showAddRoleModal.value = true
}

const closeModal = () => {
  showAddRoleModal.value = false
  editingRole.value = null
  roleForm.name = ''
  roleForm.display_name = ''
  roleForm.description = ''
  nameError.value = ''
}

const saveRole = async () => {
  // بررسی validation
  nameError.value = validateRoleName(roleForm.name)
  if (nameError.value) return
  
  if (!roleForm.display_name.trim()) {
    alert('لطفاً نام نمایشی را وارد کنید')
    return
  }

  saving.value = true
  try {
    if (editingRole.value) {
      // ویرایش نقش موجود
      await $fetch(`/api/roles/${editingRole.value.id}`, {
        method: 'PUT',
        body: roleForm
      })
      
      alert('نقش با موفقیت ویرایش شد')
    } else {
      // افزودن نقش جدید
      await $fetch('/api/roles', {
        method: 'POST',
        body: roleForm
      })
      
      alert('نقش با موفقیت ایجاد شد')
    }
    
    closeModal()
    await fetchRoles()
    await fetchStatistics()
  } catch (error: unknown) {
    console.error('خطا در ذخیره نقش:', error)
    
    const apiError = error as ApiError
    if (apiError.data?.code === 'INVALID_ROLE_NAME') {
      nameError.value = apiError.data.message || ''
    } else if (apiError.data?.code === 'DUPLICATE_ROLE') {
      alert('نقش با این نام قبلاً وجود دارد')
    } else {
      alert('خطا در ذخیره نقش')
    }
  } finally {
    saving.value = false
  }
}

interface DeleteRoleResponse {
  users_without_role?: number
  message?: string
}

const deleteRole = async (roleId: number) => {
  const role = roles.value.find(r => r.id === roleId)
  if (!role) return
  
  let message = 'آیا از حذف این نقش اطمینان دارید؟'
  if (role.users_count > 0) {
    message = `این نقش ${role.users_count} کاربر دارد. با حذف آن، تمام کاربران بدون نقش خواهند شد. آیا ادامه می‌دهید؟`
  }
  
  if (!confirm(message)) {
    return
  }

  try {
    const response: ApiResponse<DeleteRoleResponse> = await $fetch(`/api/roles/${roleId}`, {
      method: 'DELETE'
    })
    
    if (response.data?.users_without_role && response.data.users_without_role > 0) {
      alert(`نقش حذف شد و ${response.data.users_without_role} کاربر بدون نقش شدند`)
    } else {
      alert('نقش با موفقیت حذف شد')
    }
    
    await fetchRoles()
    await fetchStatistics()
  } catch (error: unknown) {
    console.error('خطا در حذف نقش:', error)
    
    // نمایش پیام سفارشی بک‌اند یا پیام پیش‌فرض
    const apiError = error as ApiError
    if (apiError.data?.message) {
      alert(apiError.data.message)
    } else if (apiError.data?.code === 'SYSTEM_ROLE_DELETE_ERROR') {
      alert('نقش‌های سیستمی قابل حذف نیستند')
    } else if (apiError.data?.code === 'ROLE_IN_USE') {
      alert('این نقش در حال استفاده است و نمی‌توان آن را حذف کرد')
    } else {
      alert('خطا در حذف نقش')
    }
  }
}

// Watch for name changes to validate
watch(() => roleForm.name, (newName) => {
  nameError.value = validateRoleName(newName)
})

// Lifecycle
onMounted(async () => {
  await fetchRoles()
  await fetchStatistics()
})
</script> 

