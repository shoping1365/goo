<template>
  <div class="group-visibility">
    <div class="section-header">
      <h3>نمایش به گروه‌های خاص</h3>
      <p>مدیریت نمایش روش‌های ارسال برای گروه‌های مختلف مشتریان</p>
    </div>

    <!-- Group Selection -->
    <div class="group-selection">
      <h4>انتخاب گروه مشتریان</h4>
      <div class="groups-list">
        <label v-for="group in customerGroups" :key="group.id" class="group-checkbox">
          <input v-model="selectedGroups" type="checkbox" :value="group.id">
          <span>{{ group.name }}</span>
        </label>
      </div>
    </div>

    <!-- Shipping Method Visibility Table -->
    <div class="visibility-table">
      <h4>دسترسی روش‌های ارسال</h4>
      <table>
        <thead>
        <tr>
          <th>روش ارسال</th>
          <th v-for="group in customerGroups" :key="group.id">{{ group.name }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="method in shippingMethods" :key="method.id">
          <td>{{ method.name }}</td>
          <td v-for="group in customerGroups" :key="group.id">
            <input
type="checkbox"
                   :checked="isMethodVisibleForGroup(method.id, group.id)"
                   @change="toggleVisibility(method.id, group.id)">
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Save Button -->
    <div class="actions">
      <button class="btn btn-success" @click="saveVisibility">
        <i class="fas fa-save"></i>
        ذخیره تغییرات
      </button>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, reactive } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

const customerGroups = ref([
  { id: 'all', name: 'همه مشتریان' },
  { id: 'vip', name: 'مشتریان VIP' },
  { id: 'wholesale', name: 'عمده‌فروشان' },
  { id: 'retail', name: 'خرده‌فروشان' },
  { id: 'guest', name: 'مهمان' }
])

const shippingMethods = ref([
  { id: 'standard', name: 'ارسال استاندارد' },
  { id: 'express', name: 'ارسال سریع' },
  { id: 'free', name: 'ارسال رایگان' }
])

const selectedGroups = ref(['all'])

const visibilityMatrix = reactive({
  standard: { all: true, vip: true, wholesale: true, retail: true, guest: true },
  express: { all: true, vip: true, wholesale: false, retail: true, guest: false },
  free:    { all: true, vip: false, wholesale: false, retail: true, guest: false }
})

const isMethodVisibleForGroup = (methodId: string, groupId: string) => {
  return !!visibilityMatrix[methodId][groupId]
}

const toggleVisibility = (methodId: string, groupId: string) => {
  visibilityMatrix[methodId][groupId] = !visibilityMatrix[methodId][groupId]
}

const saveVisibility = () => {
  // Implementation for saving visibility settings
  alert('تغییرات ذخیره شد!')
}
</script>

<style scoped>
.group-visibility {
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

.group-selection {
  margin-bottom: 30px;
}

.groups-list {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.group-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 15px;
}

.visibility-table {
  margin-bottom: 30px;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

th, td {
  padding: 12px;
  text-align: center;
  border-bottom: 1px solid #eee;
}

th {
  background: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
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
