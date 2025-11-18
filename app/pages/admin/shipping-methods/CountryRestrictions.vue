<template>
  <div class="country-restrictions">
    <div class="section-header">
      <h3>محدودیت‌های کشورها</h3>
      <p>مدیریت محدودیت‌های ارسال به کشورهای مختلف</p>
    </div>

    <!-- Country Restrictions List -->
    <div class="restrictions-list">
      <h4>لیست محدودیت‌های کشورها</h4>
      <table>
        <thead>
        <tr>
          <th>کشور</th>
          <th>نوع محدودیت</th>
          <th>توضیحات</th>
          <th>تاریخ اعمال</th>
          <th>وضعیت</th>
          <th>عملیات</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="restriction in countryRestrictions" :key="restriction.id">
          <td>{{ restriction.country }}</td>
          <td>{{ restriction.restrictionType }}</td>
          <td>{{ restriction.description }}</td>
          <td>{{ restriction.effectiveDate }}</td>
          <td>
            <span :class="['status', restriction.status]">{{ restriction.statusText }}</span>
          </td>
          <td>
            <button @click="editRestriction(restriction)" class="btn-edit">ویرایش</button>
            <button @click="deleteRestriction(restriction.id)" class="btn-delete">حذف</button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>

    <!-- Add New Restriction -->
    <div class="add-restriction-section">
      <h4>افزودن محدودیت جدید</h4>
      <form @submit.prevent="addNewRestriction">
        <div class="form-row">
          <div class="form-group">
            <label>کشور:</label>
            <select v-model="newRestriction.country" required>
              <option value="">انتخاب کنید</option>
              <option value="usa">آمریکا</option>
              <option value="uk">انگلیس</option>
              <option value="germany">آلمان</option>
              <option value="france">فرانسه</option>
              <option value="australia">استرالیا</option>
            </select>
          </div>
          <div class="form-group">
            <label>نوع محدودیت:</label>
            <select v-model="newRestriction.restrictionType" required>
              <option value="">انتخاب کنید</option>
              <option value="complete_ban">ممنوعیت کامل</option>
              <option value="partial_ban">ممنوعیت جزئی</option>
              <option value="weight_limit">محدودیت وزن</option>
              <option value="value_limit">محدودیت ارزش</option>
              <option value="documentation">نیاز به مستندات خاص</option>
            </select>
          </div>
        </div>
        <div class="form-group">
          <label>توضیحات:</label>
          <textarea v-model="newRestriction.description" rows="3" placeholder="توضیحات کامل محدودیت"></textarea>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>تاریخ اعمال:</label>
            <input v-model="newRestriction.effectiveDate" type="date" required>
          </div>
          <div class="form-group">
            <label>وضعیت:</label>
            <select v-model="newRestriction.status" required>
              <option value="active">فعال</option>
              <option value="inactive">غیرفعال</option>
            </select>
          </div>
        </div>
        <button type="submit" class="btn-primary">افزودن محدودیت</button>
      </form>
    </div>

    <!-- World Map Visualization -->
    <div class="world-map-section">
      <h4>نقشه محدودیت‌های جهانی</h4>
      <div class="map-container">
        <div class="map-legend">
          <div class="legend-item">
            <span class="legend-color no-restriction"></span>
            <span>بدون محدودیت</span>
          </div>
          <div class="legend-item">
            <span class="legend-color partial-restriction"></span>
            <span>محدودیت جزئی</span>
          </div>
          <div class="legend-item">
            <span class="legend-color complete-ban"></span>
            <span>ممنوعیت کامل</span>
          </div>
        </div>
        <div class="world-map">
          <!-- World map visualization would go here -->
          <div class="map-placeholder">
            نقشه تعاملی محدودیت‌های کشورها
          </div>
        </div>
      </div>
    </div>

    <!-- Restriction Statistics -->
    <div class="restriction-statistics">
      <h4>آمار محدودیت‌ها</h4>
      <div class="stats-grid">
        <div class="stat-card">
          <div class="stat-number">{{ stats.totalCountries }}</div>
          <div class="stat-label">کل کشورها</div>
        </div>
        <div class="stat-card">
          <div class="stat-number">{{ stats.restrictedCountries }}</div>
          <div class="stat-label">کشورهای محدود</div>
        </div>
        <div class="stat-card">
          <div class="stat-number">{{ stats.completeBans }}</div>
          <div class="stat-label">ممنوعیت کامل</div>
        </div>
        <div class="stat-card">
          <div class="stat-number">{{ stats.partialRestrictions }}</div>
          <div class="stat-label">محدودیت جزئی</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const countryRestrictions = ref([
  {
    id: 1,
    country: 'آمریکا',
    restrictionType: 'محدودیت ارزش',
    description: 'ارسال کالا با ارزش بیش از 800 دلار نیاز به مجوز خاص دارد',
    effectiveDate: '1402/01/01',
    status: 'active',
    statusText: 'فعال'
  },
  {
    id: 2,
    country: 'استرالیا',
    restrictionType: 'ممنوعیت کامل',
    description: 'ارسال مواد غذایی و گیاهان به استرالیا ممنوع است',
    effectiveDate: '1402/03/15',
    status: 'active',
    statusText: 'فعال'
  }
])

const newRestriction = reactive({
  country: '',
  restrictionType: '',
  description: '',
  effectiveDate: '',
  status: 'active'
})

const stats = ref({
  totalCountries: 195,
  restrictedCountries: 45,
  completeBans: 12,
  partialRestrictions: 33
})

const editRestriction = (restriction) => {
  // Edit restriction logic
}

const deleteRestriction = (id) => {
  // Delete restriction logic
}

const addNewRestriction = () => {
  // Add new restriction logic
}
</script>

<style scoped>
.country-restrictions {
  padding: 20px;
}

.section-header {
  margin-bottom: 30px;
}

.section-header h3 {
  color: #2c3e50;
  margin-bottom: 10px;
}

.section-header p {
  color: #7f8c8d;
}

.restrictions-list,
.add-restriction-section,
.world-map-section,
.restriction-statistics {
  background: white;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 15px;
}

th, td {
  padding: 12px;
  text-align: right;
  border-bottom: 1px solid #ecf0f1;
}

th {
  background-color: #f8f9fa;
  font-weight: 600;
  color: #2c3e50;
}

.status {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.status.active {
  background-color: #d4edda;
  color: #155724;
}

.status.inactive {
  background-color: #f8d7da;
  color: #721c24;
}

.btn-edit,
.btn-delete {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  margin-left: 5px;
  cursor: pointer;
  font-size: 12px;
}

.btn-edit {
  background-color: #007bff;
  color: white;
}

.btn-delete {
  background-color: #dc3545;
  color: white;
}

.form-row {
  display: flex;
  gap: 20px;
  margin-bottom: 15px;
}

.form-group {
  flex: 1;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: 500;
  color: #2c3e50;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.btn-primary {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover {
  background-color: #218838;
}

.map-container {
  display: flex;
  gap: 20px;
  margin-top: 15px;
}

.map-legend {
  flex: 0 0 200px;
}

.legend-item {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.legend-color {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  margin-left: 10px;
}

.legend-color.no-restriction {
  background-color: #28a745;
}

.legend-color.partial-restriction {
  background-color: #ffc107;
}

.legend-color.complete-ban {
  background-color: #dc3545;
}

.world-map {
  flex: 1;
  height: 400px;
  border: 1px solid #ddd;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f8f9fa;
}

.map-placeholder {
  color: #6c757d;
  font-size: 18px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-top: 15px;
}

.stat-card {
  background-color: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  text-align: center;
  border: 1px solid #dee2e6;
}

.stat-number {
  font-size: 32px;
  font-weight: bold;
  color: #007bff;
  margin-bottom: 5px;
}

.stat-label {
  color: #6c757d;
  font-size: 14px;
}
</style>
