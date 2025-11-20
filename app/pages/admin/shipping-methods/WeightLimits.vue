<template>
  <div class="weight-limits">
    <div class="settings-card">
      <div class="card-header">
        <h3>محدودیت وزن</h3>
        <div class="toggle-switch">
          <input id="weightToggle" v-model="enabled" type="checkbox" />
          <label for="weightToggle"></label>
        </div>
      </div>

      <div v-if="enabled" class="card-body">
        <div class="form-section">
          <h4>محدودیت‌های اصلی وزن</h4>
          <div class="weight-settings">
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
                <small class="form-help">حداکثر وزن قابل قبول برای ارسال</small>
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
                <small class="form-help">مقدار اضافی که بدون هزینه اضافی قابل قبول است</small>
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
                <small class="form-help">هزینه اضافی برای هر واحد وزن اضافی</small>
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
          <h4>محدودیت‌های پیشرفته</h4>
          <div class="advanced-weight-settings">
            <div class="form-group">
              <label>محدودیت وزن بر اساس نوع محصول</label>
              <div class="product-weight-limits">
                <div v-for="(limit, index) in weightLimits.productLimits" :key="index" class="product-limit-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>نوع محصول</label>
                      <select v-model="limit.productType" class="form-input">
                        <option value="electronics">الکترونیکی</option>
                        <option value="clothing">پوشاک</option>
                        <option value="books">کتاب</option>
                        <option value="fragile">شکننده</option>
                        <option value="heavy">سنگین</option>
                        <option value="custom">سفارشی</option>
                      </select>
                    </div>
                    <div class="form-group">
                      <label>حداکثر وزن (گرم)</label>
                      <input
                          v-model.number="limit.maxWeight"
                          type="number"
                          min="0"
                          step="1"
                          class="form-input"
                          placeholder="0"
                      />
                    </div>
                    <button class="remove-btn" @click="removeProductLimit(index)">حذف</button>
                  </div>
                  <div v-if="limit.productType === 'custom'" class="form-group">
                    <label>نام سفارشی</label>
                    <input
                        v-model="limit.customName"
                        type="text"
                        class="form-input"
                        placeholder="نام نوع محصول"
                    />
                  </div>
                </div>
                <button class="add-limit-btn" @click="addProductLimit">+ افزودن محدودیت محصول</button>
              </div>
            </div>

            <div class="form-group">
              <label>محدودیت وزن بر اساس منطقه</label>
              <div class="regional-weight-limits">
                <div v-for="(limit, index) in weightLimits.regionalLimits" :key="index" class="regional-limit-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>منطقه</label>
                      <select v-model="limit.region" class="form-input">
                        <option value="local">محلی</option>
                        <option value="provincial">استانی</option>
                        <option value="national">ملی</option>
                        <option value="international">بین‌المللی</option>
                      </select>
                    </div>
                    <div class="form-group">
                      <label>حداکثر وزن (گرم)</label>
                      <input
                          v-model.number="limit.maxWeight"
                          type="number"
                          min="0"
                          step="1"
                          class="form-input"
                          placeholder="0"
                      />
                    </div>
                    <button class="remove-btn" @click="removeRegionalLimit(index)">حذف</button>
                  </div>
                </div>
                <button class="add-limit-btn" @click="addRegionalLimit">+ افزودن محدودیت منطقه</button>
              </div>
            </div>
          </div>
        </div>

        <div class="form-section">
          <h4>تنظیمات اضافی</h4>
          <div class="additional-settings">
            <div class="form-group">
              <label>محاسبه وزن</label>
              <div class="weight-calculation">
                <label class="radio-label">
                  <input v-model="weightLimits.calculationMethod" type="radio" value="actual" />
                  وزن واقعی
                </label>
                <label class="radio-label">
                  <input v-model="weightLimits.calculationMethod" type="radio" value="volumetric" />
                  وزن حجمی
                </label>
                <label class="radio-label">
                  <input v-model="weightLimits.calculationMethod" type="radio" value="higher" />
                  بالاترین مقدار
                </label>
              </div>
            </div>

            <div class="form-group">
              <label>گرد کردن وزن</label>
              <div class="weight-rounding">
                <label class="radio-label">
                  <input v-model="weightLimits.roundingMethod" type="radio" value="none" />
                  بدون گرد کردن
                </label>
                <label class="radio-label">
                  <input v-model="weightLimits.roundingMethod" type="radio" value="up" />
                  گرد کردن به بالا
                </label>
                <label class="radio-label">
                  <input v-model="weightLimits.roundingMethod" type="radio" value="nearest" />
                  گرد کردن به نزدیک‌ترین
                </label>
              </div>
            </div>

            <div class="form-group">
              <label>دقت محاسبه (اعشار)</label>
              <input
                  v-model.number="weightLimits.precision"
                  type="number"
                  min="0"
                  max="3"
                  step="1"
                  class="form-input"
                  placeholder="2"
              />
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

      const weightLimits = ref({
        min: props.modelValue.min || 0,
        max: props.modelValue.max || 50000,
        extraAllowed: props.modelValue.extraAllowed || 1000,
        extraCost: props.modelValue.extraCost || 0,
        unit: props.modelValue.unit || 'gram',
        productLimits: props.modelValue.productLimits || [
          { productType: 'fragile', maxWeight: 2000, customName: '' },
          { productType: 'heavy', maxWeight: 10000, customName: '' }
        ],
        regionalLimits: props.modelValue.regionalLimits || [
          { region: 'local', maxWeight: 5000 },
          { region: 'provincial', maxWeight: 10000 },
          { region: 'national', maxWeight: 20000 }
        ],
        calculationMethod: props.modelValue.calculationMethod || 'actual',
        roundingMethod: props.modelValue.roundingMethod || 'up',
        precision: props.modelValue.precision || 2
      })

      function addProductLimit() {
        weightLimits.value.productLimits.push({
          productType: 'electronics',
          maxWeight: 0,
          customName: ''
        })
      }

      function removeProductLimit(index) {
        weightLimits.value.productLimits.splice(index, 1)
      }

      function addRegionalLimit() {
        weightLimits.value.regionalLimits.push({
          region: 'local',
          maxWeight: 0
        })
      }

      function removeRegionalLimit(index) {
        weightLimits.value.regionalLimits.splice(index, 1)
      }

      // Watch for changes and emit updates
      watch([enabled, weightLimits], () => {
        emit('update:modelValue', {
          enabled: enabled.value,
          ...weightLimits.value
        })
      }, { deep: true })
    </script>

    <style scoped>
      .weight-limits {
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

      .form-help {
        display: block;
        margin-top: 0.25rem;
        color: #666;
        font-size: 0.8rem;
      }

      .radio-label {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        cursor: pointer;
        margin-bottom: 0.5rem;
        font-size: 0.9rem;
      }

      .radio-label input[type="radio"] {
        width: 16px;
        height: 16px;
      }

      .product-limit-item,
      .regional-limit-item {
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

      .add-limit-btn {
        background: #28a745;
        color: white;
        border: none;
        padding: 0.75rem 1.5rem;
        border-radius: 6px;
        cursor: pointer;
        font-size: 0.9rem;
        margin-top: 1rem;
      }

      .add-limit-btn:hover {
        background: #218838;
      }

      .weight-calculation,
      .weight-rounding {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        margin-top: 0.5rem;
      }

      @media (max-width: 768px) {
        .form-row {
          grid-template-columns: 1fr;
        }

        .card-body {
          padding: 1rem;
        }

        .form-section {
          padding: 1rem;
        }
      }
    </style>
