<template>
  <div class="product-settings">
    <div class="settings-card">
      <div class="card-header">
        <h3>تنظیمات محصول</h3>
        <div class="toggle-switch">
          <input id="productToggle" v-model="enabled" type="checkbox" />
          <label for="productToggle"></label>
        </div>
      </div>
      
      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>محدودیت وزن</h4>
          <div class="weight-limitations">
            <div class="form-row">
              <div class="form-group">
                <label>حداقل وزن (گرم)</label>
                <input
                    v-model.number="weightLimits.min"
                    type="number"
                    min="0"
                    step="1"
                    class="form-input"
                    placeholder="0"
                />
              </div>
              <div class="form-group">
                <label>حداکثر وزن (گرم)</label>
                <input
                    v-model.number="weightLimits.max"
                    type="number"
                    min="0"
                    step="1"
                    class="form-input"
                    placeholder="50000"
                />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>وزن اضافی مجاز (گرم)</label>
                <input
                    v-model.number="weightLimits.extraAllowed"
                    type="number"
                    min="0"
                    step="1"
                    class="form-input"
                    placeholder="1000"
                />
              </div>
              <div class="form-group">
                <label>هزینه اضافی برای وزن اضافی (تومان)</label>
                <input
                    v-model.number="weightLimits.extraCost"
                    type="number"
                    min="0"
                    step="1000"
                    class="form-input"
                    placeholder="0"
                />
              </div>
            </div>
            <div class="form-group">
              <label>واحد وزن</label>
              <select v-model="weightLimits.unit" class="form-input">
                <option value="gram">گرم</option>
                <option value="kilogram">کیلوگرم</option>
                <option value="pound">پوند</option>
              </select>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>محدودیت ابعاد</h4>
          <div class="dimension-limitations">
            <div class="form-row">
              <div class="form-group">
                <label>حداکثر طول (سانتی‌متر)</label>
                <input
                    v-model.number="dimensionLimits.maxLength"
                    type="number"
                    min="0"
                    step="0.1"
                    class="form-input"
                    placeholder="100"
                />
              </div>
              <div class="form-group">
                <label>حداکثر عرض (سانتی‌متر)</label>
                <input
                    v-model.number="dimensionLimits.maxWidth"
                    type="number"
                    min="0"
                    step="0.1"
                    class="form-input"
                    placeholder="100"
                />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>حداکثر ارتفاع (سانتی‌متر)</label>
                <input
                    v-model.number="dimensionLimits.maxHeight"
                    type="number"
                    min="0"
                    step="0.1"
                    class="form-input"
                    placeholder="100"
                />
              </div>
              <div class="form-group">
                <label>حداکثر محیط (سانتی‌متر)</label>
                <input
                    v-model.number="dimensionLimits.maxPerimeter"
                    type="number"
                    min="0"
                    step="0.1"
                    class="form-input"
                    placeholder="300"
                />
              </div>
            </div>
            <div class="form-group">
              <label>نوع محدودیت ابعاد</label>
              <select v-model="dimensionLimits.type" class="form-input">
                <option value="strict">سختگیرانه</option>
                <option value="flexible">انعطاف‌پذیر</option>
                <option value="none">بدون محدودیت</option>
              </select>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>محصولات ممنوعه</h4>
          <div class="prohibited-products">
            <div class="form-group">
              <label>دسته‌بندی‌های ممنوعه</label>
              <div class="category-list">
                <div v-for="(category, index) in prohibitedProducts.categories" :key="index" class="category-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>نام دسته‌بندی</label>
                      <input
                          v-model="category.name"
                          type="text"
                          class="form-input"
                          placeholder="نام دسته‌بندی"
                      />
                    </div>
                    <div class="form-group">
                      <label>دلیل ممنوعیت</label>
                      <select v-model="category.reason" class="form-input">
                        <option value="legal">ممنوعیت قانونی</option>
                        <option value="safety">مشکل ایمنی</option>
                        <option value="logistics">مشکل لجستیکی</option>
                        <option value="policy">سیاست شرکت</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removeProhibitedCategory(index)">حذف</button>
                  </div>
                  <div class="form-group">
                    <label>توضیحات</label>
                    <textarea
                        v-model="category.description"
                        class="form-input"
                        placeholder="توضیحات بیشتر"
                        rows="2"
                    ></textarea>
                  </div>
                </div>
                <button class="add-category-btn" @click="addProhibitedCategory">+ افزودن دسته‌بندی</button>
              </div>
            </div>

            <div class="form-group">
              <label>کلمات کلیدی ممنوعه</label>
              <div class="keywords-list">
                <div v-for="(keyword, index) in prohibitedProducts.keywords" :key="index" class="keyword-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>کلمه کلیدی</label>
                      <input
                          v-model="keyword.value"
                          type="text"
                          class="form-input"
                          placeholder="کلمه کلیدی"
                      />
                    </div>
                    <div class="form-group">
                      <label>شدت ممنوعیت</label>
                      <select v-model="keyword.severity" class="form-input">
                        <option value="high">بالا</option>
                        <option value="medium">متوسط</option>
                        <option value="low">کم</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removeProhibitedKeyword(index)">حذف</button>
                  </div>
                </div>
                <button class="add-keyword-btn" @click="addProhibitedKeyword">+ افزودن کلمه کلیدی</button>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>محصولات حساس</h4>
          <div class="sensitive-products">
            <div class="form-group">
              <label>دسته‌بندی‌های حساس</label>
              <div class="sensitive-categories">
                <div v-for="(category, index) in sensitiveProducts.categories" :key="index" class="sensitive-category">
                  <div class="form-row">
                    <div class="form-group">
                      <label>نام دسته‌بندی</label>
                      <input
                          v-model="category.name"
                          type="text"
                          class="form-input"
                          placeholder="نام دسته‌بندی"
                      />
                    </div>
                    <div class="form-group">
                      <label>نوع حساسیت</label>
                      <select v-model="category.sensitivityType" class="form-input">
                        <option value="fragile">شکننده</option>
                        <option value="perishable">فاسدشدنی</option>
                        <option value="valuable">گران‌قیمت</option>
                        <option value="dangerous">خطرناک</option>
                        <option value="temperature">حساس به دما</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removeSensitiveCategory(index)">حذف</button>
                  </div>
                  <div class="form-row">
                    <div class="form-group">
                      <label>هزینه اضافی (تومان)</label>
                      <input
                          v-model.number="category.additionalCost"
                          type="number"
                          min="0"
                          step="1000"
                          class="form-input"
                          placeholder="0"
                      />
                    </div>
                    <div class="form-group">
                      <label>شرایط خاص</label>
                      <select v-model="category.specialConditions" class="form-input">
                        <option value="none">بدون شرایط خاص</option>
                        <option value="refrigerated">سردخانه</option>
                        <option value="fragile_handling">حمل شکننده</option>
                        <option value="insurance">بیمه اجباری</option>
                        <option value="signature">امضا اجباری</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label>توضیحات</label>
                    <textarea
                        v-model="category.description"
                        class="form-input"
                        placeholder="توضیحات شرایط خاص"
                        rows="2"
                    ></textarea>
                  </div>
                </div>
                <button class="add-sensitive-category-btn" @click="addSensitiveCategory">+ افزودن دسته‌بندی حساس</button>
              </div>
            </div>
          </div>
        </div>
        
        <div class="form-section">
          <h4>نیازمندی‌های خاص بسته‌بندی</h4>
          <div class="packaging-requirements">
            <div class="form-group">
              <label>نوع بسته‌بندی اجباری</label>
              <div class="packaging-types">
                <label v-for="type in packagingRequirements.requiredTypes" :key="type.value" class="packaging-checkbox">
                  <input
                      v-model="packagingRequirements.selectedTypes"
                      type="checkbox"
                      :value="type.value"
                  />
                  {{ type.label }}
                </label>
              </div>
            </div>

            <div class="form-group">
              <label>مواد بسته‌بندی ممنوعه</label>
              <div class="prohibited-materials">
                <div v-for="(material, index) in packagingRequirements.prohibitedMaterials" :key="index" class="material-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>نام ماده</label>
                      <input
                          v-model="material.name"
                          type="text"
                          class="form-input"
                          placeholder="نام ماده"
                      />
                    </div>
                    <div class="form-group">
                      <label>دلیل ممنوعیت</label>
                      <select v-model="material.reason" class="form-input">
                        <option value="environmental">زیست محیطی</option>
                        <option value="safety">ایمنی</option>
                        <option value="regulatory">مقرراتی</option>
                        <option value="policy">سیاست شرکت</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removeProhibitedMaterial(index)">حذف</button>
                  </div>
                </div>
                <button class="add-material-btn" @click="addProhibitedMaterial">+ افزودن ماده ممنوعه</button>
              </div>
            </div>

            <div class="form-group">
              <label>استانداردهای بسته‌بندی</label>
              <div class="packaging-standards">
                <div v-for="(standard, index) in packagingRequirements.standards" :key="index" class="standard-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>نام استاندارد</label>
                      <input
                          v-model="standard.name"
                          type="text"
                          class="form-input"
                          placeholder="نام استاندارد"
                      />
                    </div>
                    <div class="form-group">
                      <label>سطح اجباری</label>
                      <select v-model="standard.level" class="form-input">
                        <option value="required">اجباری</option>
                        <option value="recommended">توصیه شده</option>
                        <option value="optional">اختیاری</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removePackagingStandard(index)">حذف</button>
                  </div>
                  <div class="form-group">
                    <label>توضیحات</label>
                    <textarea
                        v-model="standard.description"
                        class="form-input"
                        placeholder="توضیحات استاندارد"
                        rows="2"
                    ></textarea>
                  </div>
                </div>
                <button class="add-standard-btn" @click="addPackagingStandard">+ افزودن استاندارد</button>
              </div>
            </div>

            <div class="form-group">
              <label>هزینه بسته‌بندی اضافی</label>
              <div class="form-row">
                <div class="form-group">
                  <label>هزینه پایه (تومان)</label>
                  <input
                      v-model.number="packagingRequirements.baseCost"
                      type="number"
                      min="0"
                      step="1000"
                      class="form-input"
                      placeholder="0"
                  />
                </div>
                <div class="form-group">
                  <label>هزینه اضافی برای بسته‌بندی خاص (تومان)</label>
                  <input
                      v-model.number="packagingRequirements.specialCost"
                      type="number"
                      min="0"
                      step="1000"
                      class="form-input"
                      placeholder="0"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled || false)

// محدودیت وزن
const weightLimits = ref({
  min: props.modelValue.weightLimits?.min || 0,
  max: props.modelValue.weightLimits?.max || 50000,
  extraAllowed: props.modelValue.weightLimits?.extraAllowed || 1000,
  extraCost: props.modelValue.weightLimits?.extraCost || 0,
  unit: props.modelValue.weightLimits?.unit || 'gram'
})

// محدودیت ابعاد
const dimensionLimits = ref({
  maxLength: props.modelValue.dimensionLimits?.maxLength || 100,
  maxWidth: props.modelValue.dimensionLimits?.maxWidth || 100,
  maxHeight: props.modelValue.dimensionLimits?.maxHeight || 100,
  maxPerimeter: props.modelValue.dimensionLimits?.maxPerimeter || 300,
  type: props.modelValue.dimensionLimits?.type || 'flexible'
})

// محصولات ممنوعه
const prohibitedProducts = ref({
  categories: props.modelValue.prohibitedProducts?.categories || [
    { name: 'مواد مخدر', reason: 'legal', description: 'ممنوعیت قانونی' },
    { name: 'سلاح گرم', reason: 'legal', description: 'ممنوعیت قانونی' }
  ],
  keywords: props.modelValue.prohibitedProducts?.keywords || [
    { value: 'مواد مخدر', severity: 'high' },
    { value: 'سلاح', severity: 'high' }
  ]
})

// محصولات حساس
const sensitiveProducts = ref({
  categories: props.modelValue.sensitiveProducts?.categories || [
    {
      name: 'شیشه و سرامیک',
      sensitivityType: 'fragile',
      additionalCost: 50000,
      specialConditions: 'fragile_handling',
      description: 'نیاز به حمل و نقل ویژه'
    },
    {
      name: 'مواد غذایی فاسدشدنی',
      sensitivityType: 'perishable',
      additionalCost: 30000,
      specialConditions: 'refrigerated',
      description: 'نیاز به سردخانه'
    }
  ]
})

// نیازمندی‌های بسته‌بندی
const packagingRequirements = ref({
  requiredTypes: [
    { value: 'bubble_wrap', label: 'حباب‌دار' },
    { value: 'cardboard_box', label: 'جعبه مقوایی' },
    { value: 'plastic_wrap', label: 'پلاستیک' },
    { value: 'wooden_box', label: 'جعبه چوبی' },
    { value: 'special_container', label: 'ظرف مخصوص' }
  ],
  selectedTypes: props.modelValue.packagingRequirements?.selectedTypes || ['bubble_wrap', 'cardboard_box'],
  prohibitedMaterials: props.modelValue.packagingRequirements?.prohibitedMaterials || [
    { name: 'پلاستیک یکبار مصرف', reason: 'environmental' }
  ],
  standards: props.modelValue.packagingRequirements?.standards || [
    {
      name: 'استاندارد ISO 11683',
      level: 'required',
      description: 'استاندارد بین‌المللی بسته‌بندی'
    }
  ],
  baseCost: props.modelValue.packagingRequirements?.baseCost || 0,
  specialCost: props.modelValue.packagingRequirements?.specialCost || 0
})

// توابع افزودن
function addProhibitedCategory() {
  prohibitedProducts.value.categories.push({
    name: '',
    reason: 'legal',
    description: ''
  })
}

function addProhibitedKeyword() {
  prohibitedProducts.value.keywords.push({
    value: '',
    severity: 'medium'
  })
}

function addSensitiveCategory() {
  sensitiveProducts.value.categories.push({
    name: '',
    sensitivityType: 'fragile',
    additionalCost: 0,
    specialConditions: 'none',
    description: ''
  })
}

function addProhibitedMaterial() {
  packagingRequirements.value.prohibitedMaterials.push({
    name: '',
    reason: 'environmental'
  })
}

function addPackagingStandard() {
  packagingRequirements.value.standards.push({
    name: '',
    level: 'recommended',
    description: ''
  })
}

// توابع حذف
function removeProhibitedCategory(index) {
  prohibitedProducts.value.categories.splice(index, 1)
}

function removeProhibitedKeyword(index) {
  prohibitedProducts.value.keywords.splice(index, 1)
}

function removeSensitiveCategory(index) {
  sensitiveProducts.value.categories.splice(index, 1)
}

function removeProhibitedMaterial(index) {
  packagingRequirements.value.prohibitedMaterials.splice(index, 1)
}

function removePackagingStandard(index) {
  packagingRequirements.value.standards.splice(index, 1)
}

// Watch for changes and emit updates
watch([enabled, weightLimits, dimensionLimits, prohibitedProducts, sensitiveProducts, packagingRequirements], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    weightLimits: weightLimits.value,
    dimensionLimits: dimensionLimits.value,
    prohibitedProducts: prohibitedProducts.value,
    sensitiveProducts: sensitiveProducts.value,
    packagingRequirements: packagingRequirements.value
  })
}, { deep: true })
</script>

<style scoped>
.product-settings {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.settings-card {
  background: #fff;
}

.card-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h3 {
  margin: 0;
  font-size: 1.3rem;
}

.toggle-switch {
  position: relative;
  width: 60px;
  height: 30px;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-switch label {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(255, 255, 255, 0.3);
  transition: 0.4s;
  border-radius: 30px;
}

.toggle-switch label:before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

.toggle-switch input:checked + label {
  background-color: #28a745;
}

.toggle-switch input:checked + label:before {
  transform: translateX(30px);
}

.card-body {
  padding: 2rem;
}

.form-section {
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.form-section h4 {
  margin: 0 0 1.5rem 0;
  color: #333;
  font-size: 1.1rem;
  border-bottom: 2px solid #007bff;
  padding-bottom: 0.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
  margin-bottom: 1rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
  font-weight: 500;
  font-size: 0.9rem;
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  font-family: inherit;
  transition: border-color 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.packaging-types {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.packaging-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.packaging-checkbox:hover {
  background: #e9ecef;
}

.packaging-checkbox input[type="checkbox"] {
  width: 18px;
  height: 18px;
}

.category-item,
.keyword-item,
.sensitive-category,
.material-item,
.standard-item {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
}

.remove-btn {
  background: #dc3545;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.8rem;
  align-self: end;
}

.remove-btn:hover {
  background: #c82333;
}

.add-category-btn,
.add-keyword-btn,
.add-sensitive-category-btn,
.add-material-btn,
.add-standard-btn {
  background: #28a745;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  margin-top: 1rem;
}

.add-category-btn:hover,
.add-keyword-btn:hover,
.add-sensitive-category-btn:hover,
.add-material-btn:hover,
.add-standard-btn:hover {
  background: #218838;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .packaging-types {
    grid-template-columns: repeat(2, 1fr);
  }

  .card-body {
    padding: 1rem;
  }

  .form-section {
    padding: 1rem;
  }
}
</style> 
