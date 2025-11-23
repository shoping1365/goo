<template>
  <div v-if="showItemsModal && hasAccess" class="modal-overlay" @click="closeItemsModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h3>انتخاب ایتم‌های لایه</h3>
        <button class="modal-close" @click="closeItemsModal">×</button>
      </div>
      
      <div class="modal-body">
        <div class="items-grid">
          <div 
            v-for="item in (availableItems as { id: string, name: string }[])" 
            :key="item.id"
            class="item-card"
            :class="{ 'selected': isItemSelected(item.id) }"
            @click="toggleItem(item.id)"
          >
            <div class="item-icon">
              <div v-if="item.id === 'logo'" class="logo-preview">
                <span class="text-lg font-bold text-blue-600">🏢</span>
              </div>
              <div v-else-if="item.id === 'search'" class="search-preview">
                <div class="search-box-preview">
                  <input type="text" placeholder="جستجو..." class="search-input-preview" readonly />
                  <span class="search-icon">🔍</span>
                </div>
              </div>
              <div v-else-if="item.id === 'auth'" class="auth-preview">
                <div class="auth-buttons-preview">
                  <button class="auth-btn-preview login-btn">ورود</button>
                  <button class="auth-btn-preview register-btn">ثبت‌نام</button>
                </div>
              </div>
              <div v-else-if="item.id === 'cart'" class="cart-preview">
                <div class="cart-icon-preview">
                  <span class="cart-icon">🛒</span>
                  <span class="cart-badge-preview">0</span>
                </div>
              </div>
              <div v-else-if="item.id === 'notification'" class="notification-preview">
                <div class="notification-icon-preview">
                  <span class="notification-icon">🔔</span>
                  <span class="notification-badge-preview">3</span>
                </div>
              </div>
              <div v-else-if="item.id === 'custom'" class="custom-preview">
                <div class="custom-box-preview">
                  <span class="custom-icon">📦</span>
                </div>
              </div>
              <div v-else-if="item.id === 'menu'" class="menu-preview">
                <div class="menu-icon-preview">
                  <span class="menu-icon">☰</span>
                </div>
              </div>
              <div v-else-if="item.id === 'language'" class="language-preview">
                <div class="language-selector-preview">
                  <span class="language-icon">🌐</span>
                  <span class="language-text">فارسی</span>
                </div>
              </div>
              <div v-else-if="item.id === 'currency'" class="currency-preview">
                <div class="currency-selector-preview">
                  <span class="currency-icon">💰</span>
                  <span class="currency-text">تومان</span>
                </div>
              </div>
              <div v-else-if="item.id === 'image'" class="image-preview">
                <div class="image-upload-preview">
                  <span class="image-icon">📷</span>
                  <span class="image-text">عکس</span>
                </div>
              </div>
              <div v-else-if="item.id === 'social'" class="social-preview">
                <div class="social-icons-preview">
                  <span class="social-icon">📘</span>
                  <span class="social-icon">📷</span>
                  <span class="social-icon">📱</span>
                </div>
              </div>
              <div v-else class="default-icon">
                <span class="text-lg text-gray-600">📱</span>
              </div>
            </div>
            <div class="item-name">{{ item.name }}</div>
            <div v-if="isItemSelected(item.id)" class="item-check">✓</div>
          </div>
        </div>
      </div>
      
      <div class="modal-footer">
        <button class="btn btn-primary" @click="closeItemsModal">تایید</button>
        <button class="btn btn-secondary" @click="closeItemsModal">انصراف</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const navigateTo: (to: string, options?: { redirectCode?: number; external?: boolean }) => Promise<void>
</script>

<script setup lang="ts">
import type { Ref } from 'vue'
import { computed, inject, onMounted, ref, watch } from 'vue';
import { useAuth } from '~/composables/useAuth';

// احراز هویت
const { user, isAuthenticated } = useAuth();

// بررسی دسترسی admin
const hasAccess = computed(() => {
  if (!isAuthenticated.value) {
    return false;
  }

  const userRole = user.value?.role?.toLowerCase() || '';
  const adminRoles = ['admin', 'developer'];
  return adminRoles.includes(userRole);
});

// بررسی احراز هویت و دسترسی admin - نمایش 404 در صورت عدم دسترسی
const checkAuth = async (): Promise<void> => {
  if (!hasAccess.value) {
    await navigateTo('/404', { external: false });
  }
};

// بررسی احراز هویت در هنگام mount
onMounted(async () => {
  await checkAuth();
});

// بررسی احراز هویت هنگام تغییر وضعیت احراز هویت
watch([isAuthenticated, hasAccess], async () => {
  if (!hasAccess.value) {
    await checkAuth();
  }
});

interface AvailableItem {
  id: string
  name: string
  path: string
  icon: string
}

// Inject data and functions from parent
const showItemsModal = inject<Ref<boolean>>('showItemsModal', ref(false))
const availableItems = inject<AvailableItem[]>('availableItems', [])
const closeItemsModal = inject<() => void>('closeItemsModal', () => {})
const toggleItem = inject<(id: string) => void>('toggleItem', () => {})
const isItemSelected = inject<(id: string) => boolean>('isItemSelected', () => false)
</script>

<style scoped>
/* مودال استایل‌ها */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 12px;
  width: 90%;
  max-width: 800px;
  max-height: 80vh;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-bottom: 1px solid #e9ecef;
}

.modal-header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #2c3e50;
}

.modal-close {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #6c757d;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s;
}

.modal-close:hover {
  background: #f8f9fa;
  color: #495057;
}

.modal-body {
  padding: 24px;
  max-height: 65vh;
  overflow-y: auto;
}
@media (min-width: 768px) {
  .modal-body { max-height: 75vh; }
}

.items-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.item-card {
  border: 2px solid #e9ecef;
  border-radius: 8px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
  position: relative;
  text-align: center;
}

.item-card:hover {
  border-color: #667eea;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.1);
}

.item-card:hover .item-icon > div {
  background-color: #f5f5f5;
  transition: background-color 0.2s ease;
}

.item-card.selected {
  border-color: #667eea;
  background: #f8f9ff;
}

.item-icon {
  font-size: 32px;
  margin-bottom: 8px;
}

.item-name {
  font-weight: 500;
  color: #2c3e50;
  font-size: 14px;
}

.item-check {
  position: absolute;
  top: 8px;
  right: 8px;
  background: #667eea;
  color: white;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: bold;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid #e9ecef;
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.btn-secondary {
  background: #e9ecef;
  color: #6c757d;
  border: 1px solid #ced4da;
}

.btn-secondary:hover {
  background: #f8f9fa;
  color: #495057;
}

/* استایل‌های Preview آیتم‌ها */
.logo-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #e3f2fd;
  border-radius: 8px;
  min-height: 60px;
}

.search-preview {
  width: 100%;
}

.search-box-preview {
  position: relative;
  width: 100%;
}

.search-input-preview {
  width: 100%;
  padding: 8px 32px 8px 12px;
  border: 1px solid #e0e0e0;
  border-radius: 20px;
  background: #f5f5f5;
  font-size: 12px;
  color: #666;
}

.search-icon {
  position: absolute;
  right: 8px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 14px;
  color: #999;
}

.auth-preview {
  width: 100%;
}

.auth-buttons-preview {
  display: flex;
  gap: 4px;
  flex-direction: column;
}

.auth-btn-preview {
  padding: 6px 12px;
  border: none;
  border-radius: 6px;
  font-size: 11px;
  cursor: default;
  font-weight: 500;
}

.login-btn {
  background: #e3f2fd;
  color: #1976d2;
}

.register-btn {
  background: #f3e5f5;
  color: #7b1fa2;
}

.cart-preview {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #fff3e0;
  border-radius: 8px;
  min-height: 60px;
}

.cart-icon {
  font-size: 24px;
}

.cart-badge-preview {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #f44336;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 16px;
  text-align: center;
}

.notification-preview {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #fff8e1;
  border-radius: 8px;
  min-height: 60px;
}

.notification-icon {
  font-size: 24px;
}

.notification-badge-preview {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #ff9800;
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 16px;
  text-align: center;
}

.custom-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #f1f8e9;
  border-radius: 8px;
  min-height: 60px;
}

.custom-icon {
  font-size: 24px;
}

.menu-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #e0f2f1;
  border-radius: 8px;
  min-height: 60px;
}

.menu-icon {
  font-size: 24px;
}

.language-preview, .currency-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #fce4ec;
  border-radius: 8px;
  min-height: 60px;
  gap: 8px;
}

.language-icon, .currency-icon {
  font-size: 20px;
}

.language-text, .currency-text {
  font-size: 11px;
  color: #666;
  font-weight: 500;
}

.social-preview {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #e8eaf6;
  border-radius: 8px;
  min-height: 60px;
  gap: 4px;
}

.social-icon {
  font-size: 18px;
  padding: 4px;
  background: white;
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
}

.image-preview {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #e3f2fd;
  border-radius: 8px;
  min-height: 60px;
  gap: 4px;
}

.image-icon {
  font-size: 24px;
}

.image-text {
  font-size: 11px;
  color: #666;
  font-weight: 500;
}

.default-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  background: #f5f5f5;
  border-radius: 8px;
  min-height: 60px;
}
</style>
