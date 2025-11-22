<template>
  <div class="time-settings">
    <div class="settings-card">
      <div class="card-header">
        <h3>تنظیمات زمانی</h3>
        <div class="toggle-switch">
          <input id="timeToggle" v-model="enabled" type="checkbox" />
          <label for="timeToggle"></label>
        </div>
      </div>

      <div v-if="enabled" class="card-body">
        <!-- زمان تحویل تخمینی -->
        <div class="form-section">
          <h4>زمان تحویل تخمینی</h4>
          <div class="delivery-time-settings">
            <div class="form-row">
              <div class="form-group">
                <label>حداقل زمان تحویل (روز)</label>
                <input
                    v-model.number="estimatedDelivery.minDays"
                    type="number"
                    min="1"
                    step="1"
                    class="form-input"
                    placeholder="1"
                />
              </div>
              <div class="form-group">
                <label>حداکثر زمان تحویل (روز)</label>
                <input
                    v-model.number="estimatedDelivery.maxDays"
                    type="number"
                    min="1"
                    step="1"
                    class="form-input"
                    placeholder="7"
                />
              </div>
            </div>
            <div class="form-group">
              <label>پیام نمایشی</label>
              <input
                  v-model="estimatedDelivery.message"
                  type="text"
                  class="form-input"
                  placeholder="تحویل بین 2 تا 5 روز کاری"
              />
            </div>
          </div>
        </div>

        <!-- ساعات کاری -->
        <div class="form-section">
          <h4>ساعات کاری</h4>
          <div class="working-hours">
            <div class="form-row">
              <div class="form-group">
                <label>ساعت شروع</label>
                <input
                    v-model="workingHours.start"
                    type="time"
                    class="form-input"
                />
              </div>
              <div class="form-group">
                <label>ساعت پایان</label>
                <input
                    v-model="workingHours.end"
                    type="time"
                    class="form-input"
                />
              </div>
            </div>
            <div class="form-group">
              <label>منطقه زمانی</label>
              <select v-model="workingHours.timezone" class="form-input">
                <option value="Asia/Tehran">تهران (UTC+3:30)</option>
                <option value="Asia/Dubai">دبی (UTC+4)</option>
                <option value="UTC">UTC</option>
              </select>
            </div>
          </div>
        </div>

        <!-- روزهای کاری -->
        <div class="form-section">
          <h4>روزهای کاری</h4>
          <div class="working-days">
            <div class="weekdays-grid">
              <label v-for="day in weekDays" :key="day.value" class="weekday-checkbox">
                <input
                    v-model="workingDays"
                    type="checkbox"
                    :value="day.value"
                />
                <span class="weekday-label">{{ day.label }}</span>
              </label>
            </div>
            <div class="form-group">
              <label>استثناهای روزانه</label>
              <div class="exceptions-list">
                <div v-for="(exception, index) in dayExceptions" :key="index" class="exception-item">
                  <div class="form-row">
                    <div class="form-group">
                      <label>تاریخ</label>
                      <input
                          v-model="exception.date"
                          type="date"
                          class="form-input"
                      />
                    </div>
                    <div class="form-group">
                      <label>نوع</label>
                      <select v-model="exception.type" class="form-input">
                        <option value="closed">تعطیل</option>
                        <option value="limited">ساعات محدود</option>
                      </select>
                    </div>
                    <button class="remove-btn" @click="removeDayException(index)">حذف</button>
                  </div>
                  <div v-if="exception.type === 'limited'" class="form-row">
                    <div class="form-group">
                      <label>ساعت شروع</label>
                      <input v-model="exception.startTime" type="time" class="form-input" />
                    </div>
                    <div class="form-group">
                      <label>ساعت پایان</label>
                      <input v-model="exception.endTime" type="time" class="form-input" />
                    </div>
                  </div>
                </div>
                <button class="add-exception-btn" @click="addDayException">+ افزودن استثنا</button>
              </div>
            </div>
          </div>
        </div>

        <!-- محدودیت‌های زمانی -->
        <div class="form-section">
          <h4>محدودیت‌های زمانی</h4>
          <div class="time-limitations">
            <div class="form-group">
              <label>سفارش قبل از ارسال</label>
              <div class="order-before-shipping">
                <div class="form-row">
                  <div class="form-group">
                    <label>حداقل زمان (ساعت)</label>
                    <input
                        v-model.number="timeLimitations.minOrderBeforeShipping"
                        type="number"
                        min="0"
                        step="0.5"
                        class="form-input"
                        placeholder="2"
                    />
                  </div>
                  <div class="form-group">
                    <label>حداکثر زمان (ساعت)</label>
                    <input
                        v-model.number="timeLimitations.maxOrderBeforeShipping"
                        type="number"
                        min="0"
                        step="0.5"
                        class="form-input"
                        placeholder="48"
                    />
                  </div>
                </div>
              </div>
            </div>

            <div class="form-group">
              <label>محدودیت‌های روزانه</label>
              <div class="daily-limitations">
                <div class="form-row">
                  <div class="form-group">
                    <label>حداکثر سفارش در روز</label>
                    <input
                        v-model.number="timeLimitations.maxOrdersPerDay"
                        type="number"
                        min="0"
                        step="1"
                        class="form-input"
                        placeholder="0 (بدون محدودیت)"
                    />
                  </div>
                  <div class="form-group">
                    <label>ساعت آخرین سفارش</label>
                    <input
                        v-model="timeLimitations.lastOrderTime"
                        type="time"
                        class="form-input"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- گزینه‌های تحویل سریع -->
        <div class="form-section">
          <h4>گزینه‌های تحویل سریع</h4>
          <div class="express-delivery-options">
            <div class="form-group">
              <label>تحویل فوری</label>
              <div class="express-options">
                <label class="checkbox-label">
                  <input v-model="expressDelivery.enabled" type="checkbox" />
                  فعال‌سازی تحویل فوری
                </label>
                <div v-if="expressDelivery.enabled" class="express-settings">
                  <div class="form-row">
                    <div class="form-group">
                      <label>زمان تحویل (ساعت)</label>
                      <input
                          v-model.number="expressDelivery.deliveryHours"
                          type="number"
                          min="1"
                          step="1"
                          class="form-input"
                          placeholder="24"
                      />
                    </div>
                    <div class="form-group">
                      <label>هزینه اضافی</label>
                      <div class="price-input-group">
                        <input
                            v-model.number="expressDelivery.additionalCost"
                            type="number"
                            min="0"
                            step="1000"
                            class="form-input"
                            placeholder="0"
                        />
                        <span class="currency-label">تومان</span>
                      </div>
                    </div>
                  </div>
                  <div class="form-group">
                    <label>شرایط تحویل فوری</label>
                    <textarea
                        v-model="expressDelivery.conditions"
                        class="form-input"
                        placeholder="شرایط و محدودیت‌های تحویل فوری"
                        rows="3"
                    ></textarea>
                  </div>
                </div>
              </div>
            </div>

            <div class="form-group">
              <label>تحویل در همان روز</label>
              <div class="same-day-delivery">
                <label class="checkbox-label">
                  <input v-model="sameDayDelivery.enabled" type="checkbox" />
                  فعال‌سازی تحویل در همان روز
                </label>
                <div v-if="sameDayDelivery.enabled" class="same-day-settings">
                  <div class="form-row">
                    <div class="form-group">
                      <label>ساعت آخرین سفارش</label>
                      <input
                          v-model="sameDayDelivery.lastOrderTime"
                          type="time"
                          class="form-input"
                      />
                    </div>
                    <div class="form-group">
                      <label>هزینه اضافی</label>
                      <div class="price-input-group">
                        <input
                            v-model.number="sameDayDelivery.additionalCost"
                            type="number"
                            min="0"
                            step="1000"
                            class="form-input"
                            placeholder="0"
                        />
                        <span class="currency-label">تومان</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- تحویل در ساعات خاص -->
        <div class="form-section">
          <h4>تحویل در ساعات خاص</h4>
          <div class="specific-time-delivery">
            <div class="form-group">
              <label>ساعات تحویل</label>
              <div class="delivery-time-slots">
                <div v-for="(slot, index) in specificTimeDelivery.slots" :key="index" class="time-slot">
                  <div class="form-row">
                    <div class="form-group">
                      <label>از ساعت</label>
                      <input v-model="slot.startTime" type="time" class="form-input" />
                    </div>
                    <div class="form-group">
                      <label>تا ساعت</label>
                      <input v-model="slot.endTime" type="time" class="form-input" />
                    </div>
                    <div class="form-group">
                      <label>هزینه اضافی</label>
                      <div class="price-input-group">
                        <input
                            v-model.number="slot.additionalCost"
                            type="number"
                            min="0"
                            step="1000"
                            class="form-input"
                            placeholder="0"
                        />
                        <span class="currency-label">تومان</span>
                      </div>
                    </div>
                    <button class="remove-btn" @click="removeTimeSlot(index)">حذف</button>
                  </div>
                </div>
                <button class="add-slot-btn" @click="addTimeSlot">+ افزودن بازه زمانی</button>
              </div>
            </div>

            <div class="form-group">
              <label>روزهای هفته</label>
              <div class="weekdays-grid">
                <label v-for="day in weekDays" :key="day.value" class="weekday-checkbox">
                  <input
                      v-model="specificTimeDelivery.weekdays"
                      type="checkbox"
                      :value="day.value"
                  />
                  <span class="weekday-label">{{ day.label }}</span>
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref, watch } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const enabled = ref(props.modelValue.enabled || false)

const estimatedDelivery = ref({
  minDays: props.modelValue.estimatedDelivery?.minDays || 2,
  maxDays: props.modelValue.estimatedDelivery?.maxDays || 5,
  message: props.modelValue.estimatedDelivery?.message || 'تحویل بین 2 تا 5 روز کاری'
})

const workingHours = ref({
  start: props.modelValue.workingHours?.start || '08:00',
  end: props.modelValue.workingHours?.end || '18:00',
  timezone: props.modelValue.workingHours?.timezone || 'Asia/Tehran'
})

const weekDays = [
  { value: 'monday', label: 'دوشنبه' },
  { value: 'tuesday', label: 'سه‌شنبه' },
  { value: 'wednesday', label: 'چهارشنبه' },
  { value: 'thursday', label: 'پنج‌شنبه' },
  { value: 'friday', label: 'جمعه' },
  { value: 'saturday', label: 'شنبه' },
  { value: 'sunday', label: 'یکشنبه' }
]

const workingDays = ref(props.modelValue.workingDays || ['monday', 'tuesday', 'wednesday', 'thursday', 'friday'])

const dayExceptions = ref(props.modelValue.dayExceptions || [])

const timeLimitations = ref({
  minOrderBeforeShipping: props.modelValue.timeLimitations?.minOrderBeforeShipping || 2,
  maxOrderBeforeShipping: props.modelValue.timeLimitations?.maxOrderBeforeShipping || 48,
  maxOrdersPerDay: props.modelValue.timeLimitations?.maxOrdersPerDay || 0,
  lastOrderTime: props.modelValue.timeLimitations?.lastOrderTime || '16:00'
})

const expressDelivery = ref({
  enabled: props.modelValue.expressDelivery?.enabled || false,
  deliveryHours: props.modelValue.expressDelivery?.deliveryHours || 24,
  additionalCost: props.modelValue.expressDelivery?.additionalCost || 0,
  conditions: props.modelValue.expressDelivery?.conditions || ''
})

const sameDayDelivery = ref({
  enabled: props.modelValue.sameDayDelivery?.enabled || false,
  lastOrderTime: props.modelValue.sameDayDelivery?.lastOrderTime || '12:00',
  additionalCost: props.modelValue.sameDayDelivery?.additionalCost || 0
})

const specificTimeDelivery = ref({
  slots: props.modelValue.specificTimeDelivery?.slots || [
    { startTime: '09:00', endTime: '12:00', additionalCost: 0 },
    { startTime: '14:00', endTime: '18:00', additionalCost: 0 }
  ],
  weekdays: props.modelValue.specificTimeDelivery?.weekdays || ['monday', 'tuesday', 'wednesday', 'thursday', 'friday']
})

function addDayException() {
  dayExceptions.value.push({
    date: '',
    type: 'closed',
    startTime: '09:00',
    endTime: '17:00'
  })
}

function removeDayException(index) {
  dayExceptions.value.splice(index, 1)
}

function addTimeSlot() {
  specificTimeDelivery.value.slots.push({
    startTime: '09:00',
    endTime: '12:00',
    additionalCost: 0
  })
}

function removeTimeSlot(index) {
  specificTimeDelivery.value.slots.splice(index, 1)
}

// Watch for changes and emit updates
watch([enabled, estimatedDelivery, workingHours, workingDays, dayExceptions, timeLimitations, expressDelivery, sameDayDelivery, specificTimeDelivery], () => {
  emit('update:modelValue', {
    enabled: enabled.value,
    estimatedDelivery: estimatedDelivery.value,
    workingHours: workingHours.value,
    workingDays: workingDays.value,
    dayExceptions: dayExceptions.value,
    timeLimitations: timeLimitations.value,
    expressDelivery: expressDelivery.value,
    sameDayDelivery: sameDayDelivery.value,
    specificTimeDelivery: specificTimeDelivery.value
  })
}, { deep: true })
</script>

<style scoped>
.time-settings {
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

.price-input-group {
  position: relative;
  display: flex;
  align-items: center;
}

.price-input-group .form-input {
  padding-right: 60px;
}

.currency-label {
  position: absolute;
  right: 10px;
  color: #666;
  font-size: 0.8rem;
}

.weekdays-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.weekday-checkbox {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 6px;
  transition: background-color 0.3s;
}

.weekday-checkbox:hover {
  background: #e9ecef;
}

.weekday-checkbox input[type="checkbox"] {
  width: 18px;
  height: 18px;
}

.weekday-label {
  font-size: 0.9rem;
  color: #333;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
  margin-bottom: 1rem;
  font-weight: 500;
}

.checkbox-label input[type="checkbox"] {
  width: 18px;
  height: 18px;
}

.exceptions-list {
  margin-top: 1rem;
}

.exception-item {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
}

.time-slot {
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

.add-exception-btn,
.add-slot-btn {
  background: #28a745;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  margin-top: 1rem;
}

.add-exception-btn:hover,
.add-slot-btn:hover {
  background: #218838;
}

.express-settings,
.same-day-settings {
  background: white;
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  margin-top: 1rem;
}

.delivery-time-slots {
  margin-top: 1rem;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .weekdays-grid {
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
