<template>
  <div class="personalization-by-history">
    <div class="section-header">
      <h3>شخصی‌سازی بر اساس سابقه</h3>
      <p>پیشنهاد و نمایش روش‌های ارسال بر اساس رفتار و سابقه خرید مشتری</p>
    </div>

    <!-- Personalization Settings -->
    <div class="personalization-settings">
      <h4>تنظیمات شخصی‌سازی</h4>
      <div class="settings-list">
        <div class="setting-item">
          <label>
            <input v-model="settings.useOrderHistory" type="checkbox">
            استفاده از تاریخچه سفارشات
          </label>
        </div>
        <div class="setting-item">
          <label>
            <input v-model="settings.useLocation" type="checkbox">
            استفاده از موقعیت جغرافیایی
          </label>
        </div>
        <div class="setting-item">
          <label>
            <input v-model="settings.usePreferredMethod" type="checkbox">
            نمایش روش ترجیحی مشتری
          </label>
        </div>
        <div class="setting-item">
          <label>
            <input v-model="settings.useSeasonalOffers" type="checkbox">
            پیشنهاد ویژه بر اساس فصل و مناسبت
          </label>
        </div>
      </div>
    </div>

    <!-- Personalization Preview -->
    <div class="personalization-preview">
      <h4>پیش‌نمایش شخصی‌سازی</h4>
      <div class="preview-list">
        <div v-for="item in previewData" :key="item.id" class="preview-item">
          <div class="customer-info">
            <span class="customer-name">{{ item.customer }}</span>
            <span class="customer-type">{{ item.type }}</span>
          </div>
          <div class="suggested-method">
            <span>روش پیشنهادی:</span>
            <strong>{{ item.suggestedMethod }}</strong>
          </div>
          <div class="history">
            <span>سابقه:</span>
            <span>{{ item.history }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Save Button -->
    <div class="actions">
      <button class="btn btn-success" @click="saveSettings">
        <i class="fas fa-save"></i>
        ذخیره تنظیمات
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const settings = ref({
  useOrderHistory: true,
  useLocation: false,
  usePreferredMethod: true,
  useSeasonalOffers: false
})

const previewData = ref([
  {
    id: 1,
    customer: 'علی احمدی',
    type: 'VIP',
    suggestedMethod: 'ارسال سریع',
    history: 'اغلب ارسال سریع، آخرین سفارش: ۲ روز پیش'
  },
  {
    id: 2,
    customer: 'فاطمه محمدی',
    type: 'مهمان',
    suggestedMethod: 'ارسال استاندارد',
    history: '۲ سفارش، هر دو با ارسال استاندارد'
  }
])

const saveSettings = () => {
  // Implementation for saving settings
  alert('تنظیمات شخصی‌سازی ذخیره شد!')
}
</script>

<style scoped>
.personalization-by-history {
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.section-header {
  margin-bottom: 30px;
  text-align: center;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
  font-size: 24px;
}

.section-header p {
  color: #7f8c8d;
  font-size: 14px;
}

.personalization-settings {
  margin-bottom: 30px;
}

.settings-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.setting-item {
  font-size: 16px;
  color: #2c3e50;
}

.personalization-preview {
  margin-bottom: 30px;
}

.preview-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.preview-item {
  background: #f8f9fa;
  padding: 15px;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.customer-info {
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 5px;
}

.suggested-method {
  color: #27ae60;
  font-size: 15px;
  margin-bottom: 5px;
}

.history {
  color: #6c757d;
  font-size: 13px;
}

.actions {
  text-align: center;
  margin-top: 20px;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
  background: #27ae60;
  color: white;
  transition: all 0.3s ease;
}

.btn:hover {
  background: #229954;
}
</style>
