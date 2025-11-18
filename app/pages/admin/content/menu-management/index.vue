<template>
  <div class="page-surface">
    <!-- Header -->
    <div class="header-bg">
      <div class="page-header-flex">
        <div>
          <h1 class="page-title">مدیریت منوها</h1>
          <p class="page-subtitle">لیست منوهای ثبت شده در سامانه</p>
        </div>
        <div class="header-actions">
          <button class="btn btn-secondary" @click="refreshMenus" :disabled="isLoading">
            <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v6h6M20 20v-6h-6M5 19A9 9 0 1120 9" />
            </svg>
            بروزرسانی
          </button>
          <button class="btn btn-primary" @click="addNewMenu">افزودن منو</button>
        </div>
      </div>
    </div>

    <div class="content-wrapper px-4 sm:px-6 lg:px-8 py-6">
      <div class="content-card bg-white rounded-lg shadow-soft">
        <div class="table-header">
          <h2 class="table-title">لیست منوها</h2>
          <span class="table-count" v-if="menusList.length">{{ menusList.length }} منو</span>
        </div>

        <div v-if="isLoading" class="table-loading">
          <div class="loader" />
          <p class="loading-text">در حال بارگذاری منوها...</p>
        </div>

        <div v-else-if="!menusList.length" class="empty-state">
          <svg class="mx-auto h-12 w-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7V3m8 4V3M4 11h16M5 21h14a1 1 0 001-1v-9H4v9a1 1 0 001 1z" />
          </svg>
          <h3 class="empty-title">هنوز منویی ایجاد نشده است</h3>
          <p class="empty-text">برای شروع روی دکمه افزودن منو کلیک کنید.</p>
          <button class="btn btn-primary mt-4" @click="addNewMenu">ایجاد منو جدید</button>
        </div>

        <div v-else class="menu-cards">
          <div v-for="menu in menusList" :key="menu.id ?? menu.slug" class="menu-card">
            <div class="card-header">
              <div class="card-info">
                <h3 class="card-title">{{ menu.name || 'بدون عنوان' }}</h3>
                <div class="card-meta">
                  <span class="meta-item">
                    <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 20l4-16m2 16l4-16M6 9h14M4 15h14"/>
                    </svg>
                    {{ menu.id ?? '—' }}
                  </span>
                  <span class="meta-item">
                    <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
                    </svg>
                    {{ menu.items?.length ?? 0 }} آیتم
                  </span>
                </div>
              </div>
              <span 
                :class="['status-badge', menu.enabled ? 'status-active' : 'status-inactive']"
                @click.stop="toggleMenuStatus(menu)"
                style="cursor: pointer;"
                title="کلیک کنید برای تغییر وضعیت"
              >
                {{ menu.enabled ? 'فعال' : 'غیرفعال' }}
              </span>
            </div>
            
            <div class="card-body">
              <div class="info-row">
                <span class="info-label">شناسه:</span>
                <code class="info-value">{{ truncateSlug(menu.slug) }}</code>
              </div>
              <div class="info-row" v-if="menu.updatedAt || menu.updated_at || menu.createdAt || menu.created_at">
                <span class="info-label">آخرین تغییر:</span>
                <span class="info-value">{{ formatDate(menu.updatedAt || menu.updated_at || menu.createdAt || menu.created_at) }}</span>
              </div>
            </div>

            <div class="card-actions">
              <button class="btn-card btn-edit" @click="editMenu(menu)">
                <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                </svg>
                ویرایش
              </button>
              <button class="btn-card btn-delete" @click="handleDeleteClick(menu)" :disabled="isDeletingId === menu.id">
                <svg v-if="isDeletingId !== menu.id" class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                </svg>
                <span v-if="isDeletingId === menu.id" class="inline-loader" />
                حذف
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <ClientOnly>
      <DeleteConfirmModal
        ref="deleteModal"
        single-delete-title="حذف منو"
        :single-delete-message="`آیا از حذف منوی «${selectedMenuName}» مطمئن هستید؟ این عملیات غیرقابل بازگشت است.`"
        @single-delete="confirmDelete"
        @close-single="selectedMenuId = null"
      />
    </ClientOnly>
  </div>
</template>

<script setup>
definePageMeta({
  layout: 'admin-main'
})

import { computed, onMounted, ref } from 'vue'
import { navigateTo } from '#app'
import { useMenuManagement } from '~/composables/useMenuManagement'
import { useToast } from '~/composables/useToast'
import DeleteConfirmModal from '~/components/common/DeleteConfirmModal.vue'

const { menus, fetchMenus, deleteMenu, isLoading } = useMenuManagement()
const { showError, showSuccess } = useToast()

const isDeletingId = ref(null)
const deleteModal = ref(null)
const selectedMenuId = ref(null)
const selectedMenuName = ref('')

const menusList = computed(() => Array.isArray(menus.value) ? menus.value : [])

const loadMenus = async () => {
  try {
    await fetchMenus()
  } catch (error) {
    console.error('Error loading menus:', error)
    const message = error?.message || error?.data?.message || 'خطا در دریافت لیست منوها'
    showError(message)
  }
}

const refreshMenus = async () => {
  await loadMenus()
}

const addNewMenu = () => {
  navigateTo('/admin/content/menu-management/create')
}

const editMenu = (menu) => {
  if (!menu?.id) {
    showError('شناسه منو نامعتبر است')
    return
  }
  navigateTo(`/admin/content/menu-management/edit/${menu.id}`)
}

const handleDeleteClick = (menu) => {
  if (!menu?.id) {
    showError('شناسه منو نامعتبر است')
    return
  }
  
  selectedMenuId.value = menu.id
  selectedMenuName.value = menu.name || menu.slug || menu.id
  deleteModal.value?.openDeleteConfirm(menu.id)
}

const confirmDelete = async (menuId) => {
  isDeletingId.value = menuId
  try {
    await deleteMenu(menuId)
    showSuccess('منو با موفقیت حذف شد')
    await loadMenus()
  } catch (error) {
    console.error('Error deleting menu:', error)
    const message = error?.message || error?.data?.message || 'خطا در حذف منو'
    showError(message)
  } finally {
    isDeletingId.value = null
    selectedMenuId.value = null
  }
}

const toggleMenuStatus = async (menu) => {
  const originalStatus = menu.enabled
  try {
    menu.enabled = !menu.enabled
    
    // فقط enabled رو تغییر بده، items رو نفرست!
    const updateData = {
      name: menu.name,
      slug: menu.slug,
      enabled: menu.enabled
    }
    
    await $fetch(`/api/admin/menus/${menu.id}`, {
      method: 'PUT',
      body: updateData
    })
    
    showSuccess(menu.enabled ? 'منو فعال شد' : 'منو غیرفعال شد')
    await loadMenus() // رفرش کن تا از backend آیتم‌های کامل بیاد
  } catch (error) {
    menu.enabled = originalStatus // revert on error
    console.error('Error toggling menu status:', error)
    showError('خطا در تغییر وضعیت منو')
  }
}

const formatDate = (value) => {
  if (!value) {
    return '—'
  }
  try {
    const date = typeof value === 'string' ? new Date(value) : value
    if (Number.isNaN(date.getTime())) {
      return '—'
    }
    return date.toLocaleString('fa-IR', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch (error) {
    return '—'
  }
}

const truncateSlug = (slug) => {
  if (!slug) return '—'
  if (slug.length <= 30) return slug
  return slug.substring(0, 30) + '...'
}

onMounted(async () => {
  await loadMenus()
})

</script>

<style scoped>
.page-surface {
  min-height: 100vh;
  padding: 2rem 0 3rem;
}

.content-wrapper {
  max-width: 1120px;
  margin: 0 auto;
}

.content-card {
  width: 100%;
  overflow: hidden;
}

.shadow-soft {
  box-shadow: 0 20px 45px -25px rgba(15, 23, 42, 0.35);
}

.header-bg {
  max-width: 1120px;
  margin: 0 auto 30px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #ffffff;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 25px 60px -30px rgba(76, 29, 149, 0.45);
}

.page-header-flex {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 2rem;
  font-weight: 700;
  margin: 0;
}

.page-subtitle {
  margin-top: 0.5rem;
  color: rgba(255, 255, 255, 0.8);
  font-size: 0.95rem;
}

.header-actions {
  display: flex;
  gap: 0.75rem;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.65rem 1.4rem;
  border-radius: 0.75rem;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-primary {
  background: rgba(255, 255, 255, 0.2);
  border: 2px solid rgba(255, 255, 255, 0.3);
  color: #ffffff;
}

.btn-primary:hover {
  background: rgba(255, 255, 255, 0.28);
  border-color: rgba(255, 255, 255, 0.45);
}

.btn-secondary {
  background: rgba(255, 255, 255, 0.15);
  border: 2px solid rgba(255, 255, 255, 0.25);
  color: #ffffff;
}

.btn-secondary:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 1.75rem;
  border-bottom: 1px solid #eef2ff;
}

.table-title {
  font-size: 1.15rem;
  font-weight: 700;
  color: #1e293b;
}

.table-count {
  background: #eef2ff;
  color: #4338ca;
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-size: 0.85rem;
  font-weight: 600;
}

.menu-cards {
  padding: 1.5rem;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 1.25rem;
}

.menu-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 0.75rem;
  overflow: hidden;
  transition: all 0.2s ease;
}

.menu-card:hover {
  border-color: #cbd5e1;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.card-header {
  padding: 1.25rem;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 1rem;
}

.card-info {
  flex: 1;
  min-width: 0;
}

.card-title {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 0.5rem 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.card-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.meta-item {
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  font-size: 0.85rem;
  color: #64748b;
}

.meta-item .icon {
  width: 1rem;
  height: 1rem;
  flex-shrink: 0;
}

.card-body {
  padding: 1.25rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.info-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
}

.info-label {
  color: #64748b;
  font-weight: 600;
  white-space: nowrap;
}

.info-value {
  color: #334155;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

code.info-value {
  background: #f1f5f9;
  padding: 0.25rem 0.5rem;
  border-radius: 0.375rem;
  font-family: 'Courier New', monospace;
  font-size: 0.85rem;
  color: #4338ca;
}

.card-actions {
  padding: 1rem 1.25rem;
  background: #f8fafc;
  border-top: 1px solid #e2e8f0;
  display: flex;
  gap: 0.75rem;
}

.btn-card {
  flex: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.65rem 1rem;
  border-radius: 0.5rem;
  font-weight: 600;
  font-size: 0.9rem;
  transition: all 0.2s ease;
  border: 1px solid transparent;
}

.btn-card .icon {
  width: 1.1rem;
  height: 1.1rem;
  flex-shrink: 0;
}

.btn-edit {
  background: #eff6ff;
  color: #1d4ed8;
  border-color: #dbeafe;
}

.btn-edit:hover {
  background: #dbeafe;
  border-color: #bfdbfe;
}

.btn-delete {
  background: #fef2f2;
  color: #dc2626;
  border-color: #fee2e2;
}

.btn-delete:hover:not(:disabled) {
  background: #fee2e2;
  border-color: #fecaca;
}

.btn-delete:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.status-badge {
  display: inline-flex;
  align-items: center;
  padding: 0.4rem 0.85rem;
  border-radius: 999px;
  font-size: 0.8rem;
  font-weight: 600;
  white-space: nowrap;
  flex-shrink: 0;
}

.status-active {
  background: #dcfce7;
  color: #166534;
}

.status-inactive {
  background: #fee2e2;
  color: #991b1b;
}

.table-loading {
  padding: 3rem 1.75rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 1rem;
}

.loader {
  width: 2.75rem;
  height: 2.75rem;
  border: 4px solid #e0e7ff;
  border-top-color: #4338ca;
  border-radius: 50%;
  animation: spin 0.9s linear infinite;
}

.loading-text {
  color: #475569;
  font-weight: 500;
}

.empty-state {
  padding: 4rem 1.75rem;
  text-align: center;
}

.empty-title {
  font-size: 1.1rem;
  font-weight: 700;
  color: #1f2937;
}

.empty-text {
  margin-top: 0.75rem;
  color: #6b7280;
}

.inline-loader {
  display: inline-block;
  margin-left: 0.5rem;
  width: 1rem;
  height: 1rem;
  border: 2px solid rgba(239, 68, 68, 0.2);
  border-top-color: #dc2626;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

@media (max-width: 768px) {
  .page-header-flex {
    flex-direction: column;
    align-items: flex-start;
    gap: 1.5rem;
  }

  .header-actions {
    width: 100%;
    justify-content: flex-start;
  }

  .btn {
    width: 100%;
  }

  .menu-cards {
    grid-template-columns: 1fr;
    padding: 1rem;
  }
}
</style>
