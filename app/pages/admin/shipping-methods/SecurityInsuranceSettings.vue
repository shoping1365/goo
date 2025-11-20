<template>
  <div class="security-insurance-settings">
    <div class="settings-card">
      <div class="card-header">
        <h3>تنظیمات امنیتی و بیمه</h3>
        <div class="toggle-switch">
          <input id="securityToggle" v-model="enabled" type="checkbox" />
          <label for="securityToggle"></label>
        </div>
      </div>

      <div v-if="enabled" class="card-body">
        <!-- بیمه کالا -->
        <div class="form-section">
          <h4>بیمه کالا</h4>
          <div class="insurance-settings">
            <label class="checkbox-label">
              <input v-model="insurance.enabled" type="checkbox" />
              فعال‌سازی بیمه کالا
            </label>
            <div v-if="insurance.enabled" class="insurance-options">
              <div class="form-row">
                <div class="form-group">
                  <label>نوع بیمه</label>
                  <select v-model="insurance.type" class="form-input">
                    <option value="basic">پایه</option>
                    <option value="comprehensive">جامع</option>
                    <option value="premium">پریمیوم</option>
                    <option value="custom">سفارشی</option>
                  </select>
                </div>
                <div class="form-group">
                  <label>مبلغ بیمه (تومان)</label>
                  <input
                      v-model.number="insurance.amount"
                      type="number"
                      min="0"
                      step="1000"
                      class="form-input"
                      placeholder="0"
                  />
                </div>
              </div>
              <div class="form-row">
                <div class="form-group">
                  <label>هزینه بیمه (تومان)</label>
                  <input
                      v-model.number="insurance.cost"
                      type="number"
                      min="0"
                      step="1000"
                      class="form-input"
                      placeholder="0"
                  />
                </div>
                <div class="form-group">
                  <label>درصد پوشش بیمه</label>
                  <input
                      v-model.number="insurance.coveragePercentage"
                      type="number"
                      min="0"
                      max="100"
                      step="1"
                      class="form-input"
                      placeholder="100"
                  />
                </div>
              </div>
              <div class="form-group">
                <label>شرایط بیمه</label>
                <textarea
                    v-model="insurance.conditions"
                    class="form-input"
                    placeholder="شرایط و قوانین بیمه"
                    rows="3"
                ></textarea>
              </div>
            </div>
          </div>
        </div>

        <!-- امضای گیرنده -->
        <div class="form-section">
          <h4>امضای گیرنده</h4>
          <div class="signature-settings">
            <label class="checkbox-label">
              <input v-model="signature.enabled" type="checkbox" />
              الزام امضای گیرنده
            </label>
            <div v-if="signature.enabled" class="signature-options">
              <div class="form-row">
                <div class="form-group">
                  <label>نوع امضا</label>
                  <select v-model="signature.type" class="form-input">
                    <option value="digital">دیجیتال</option>
                    <option value="physical">فیزیکی</option>
                    <option value="both">هر دو</option>
                  </select>
                </div>
                <div class="form-group">
                  <label>هویت‌سنجی</label>
                  <select v-model="signature.verification" class="form-input">
                    <option value="id_card">کارت شناسایی</option>
                    <option value="phone">تلفن همراه</option>
                    <option value="code">کد تایید</option>
                    <option value="none">بدون تایید</option>
                  </select>
                </div>
              </div>
              <div class="form-group">
                <label>پیام امضا</label>
                <textarea
                    v-model="signature.message"
                    class="form-input"
                    placeholder="پیام نمایشی برای امضا"
                    rows="2"
                ></textarea>
              </div>
            </div>
          </div>
        </div>

        <!-- عکس از تحویل -->
        <div class="form-section">
          <h4>عکس از تحویل</h4>
          <div class="photo-settings">
            <label class="checkbox-label">
              <input v-model="photo.enabled" type="checkbox" />
              الزام عکس از تحویل
            </label>
            <div v-if="photo.enabled" class="photo-options">
              <div class="form-row">
                <div class="form-group">
                  <label>تعداد عکس‌های الزامی</label>
                  <input
                      v-model.number="photo.requiredCount"
                      type="number"
                      min="1"
                      max="10"
                      step="1"
                      class="form-input"
                      placeholder="3"
                  />
                </div>
                <div class="form-group">
                  <label>کیفیت عکس</label>
                  <select v-model="photo.quality" class="form-input">
                    <option value="low">کم</option>
                    <option value="medium">متوسط</option>
                    <option value="high">بالا</option>
                    <option value="ultra">خیلی بالا</option>
                  </select>
                </div>
              </div>
              <div class="form-group">
                <label>نوع عکس‌های الزامی</label>
                <div class="photo-types">
                  <label class="checkbox-label">
                    <input v-model="photo.types" type="checkbox" value="package" />
                    بسته بسته‌بندی شده
                  </label>
                  <label class="checkbox-label">
                    <input v-model="photo.types" type="checkbox" value="delivery" />
                    لحظه تحویل
                  </label>
                  <label class="checkbox-label">
                    <input v-model="photo.types" type="checkbox" value="signature" />
                    امضای گیرنده
                  </label>
                  <label class="checkbox-label">
                    <input v-model="photo.types" type="checkbox" value="location" />
                    موقعیت تحویل
                  </label>
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

const enabled = ref(props.modelValue.enabled ?? true)

const insurance = ref({
  enabled: props.modelValue.insurance?.enabled ?? false,
  type: props.modelValue.insurance?.type ?? 'basic',
  amount: props.modelValue.insurance?.amount ?? 0,
  cost: props.modelValue.insurance?.cost ?? 0,
  coveragePercentage: props.modelValue.insurance?.coveragePercentage ?? 100,
  conditions: props.modelValue.insurance?.conditions ?? ''
})

const signature = ref({
  enabled: props.modelValue.signature?.enabled ?? false,
  type: props.modelValue.signature?.type ?? 'digital',
  verification: props.modelValue.signature?.verification ?? 'phone',
  message: props.modelValue.signature?.message ?? ''
})

const photo = ref({
  enabled: props.modelValue.photo?.enabled ?? false,
  requiredCount: props.modelValue.photo?.requiredCount ?? 3,
  quality: props.modelValue.photo?.quality ?? 'medium',
  types: props.modelValue.photo?.types ?? ['package', 'delivery']
})

watch([enabled, insurance, signature, photo], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    insurance: insurance.value,
    signature: signature.value,
    photo: photo.value
  })
}, { deep: true })
</script>

<style scoped>
.security-insurance-settings {
  direction: rtl;
}

.settings-card {
  background: #f8f9fa;
  border-radius: 12px;
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  background: #fff;
  border-bottom: 1px solid #e9ecef;
}

.card-header h3 {
  margin: 0;
  font-size: 1.1rem;
  font-weight: 600;
  color: #333;
}

.toggle-switch {
  position: relative;
  width: 50px;
  height: 24px;
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
  background-color: #ccc;
  transition: .4s;
  border-radius: 24px;
}

.toggle-switch label:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}

.toggle-switch input:checked + label {
  background-color: #2196F3;
}

.toggle-switch input:checked + label:before {
  transform: translateX(26px);
}

.card-body {
  padding: 1.5rem;
}

.form-section {
  margin-bottom: 2rem;
}

.form-section h4 {
  margin: 0 0 1rem 0;
  font-size: 1rem;
  color: #555;
  border-bottom: 1px solid #e9ecef;
  padding-bottom: 0.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #555;
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 0.9rem;
  background: white;
}

.form-input:focus {
  border-color: #2196F3;
  outline: none;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
  cursor: pointer;
}

.checkbox-label input[type="checkbox"] {
  width: 18px;
  height: 18px;
}

.insurance-options,
.signature-options,
.photo-options {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.photo-types {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.5rem;
  margin-top: 0.5rem;
}

.radio-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
  cursor: pointer;
}

.radio-label input[type="radio"] {
  width: 16px;
  height: 16px;
}
</style>
