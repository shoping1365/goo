<template>
  <div class="shipping-method-pricing-settings">
    <h2 class="section-title">تنظیمات قیمت‌گذاری</h2>
    
    <div class="pricing-grid">
      <FlatRatePricing v-model="pricingData.flatRate" />
      <WeightBasedPricing v-model="pricingData.weightBased" />
      <DistanceBasedPricing v-model="pricingData.distanceBased" />
      <ItemCountPricing v-model="pricingData.itemCount" />
      <OrderValuePricing v-model="pricingData.orderValue" />
      <SpecialDiscountPricing v-model="pricingData.specialDiscount" />
      <RegionalPricing v-model="pricingData.regional" />
      <TimeBasedPricing v-model="pricingData.timeBased" />
    </div>
    
    <div class="summary-section">
      <h3>خلاصه تنظیمات قیمت‌گذاری</h3>
      <div class="summary-grid">
        <div class="summary-item" v-for="(setting, key) in pricingData" :key="key">
          <span class="setting-name">{{ getSettingName(key) }}</span>
          <span class="setting-status" :class="{ active: setting.enabled }">
            {{ setting.enabled ? 'فعال' : 'غیرفعال' }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import FlatRatePricing from './FlatRatePricing.vue'
import WeightBasedPricing from './WeightBasedPricing.vue'
import DistanceBasedPricing from './DistanceBasedPricing.vue'
import ItemCountPricing from './ItemCountPricing.vue'
import OrderValuePricing from './OrderValuePricing.vue'
import SpecialDiscountPricing from './SpecialDiscountPricing.vue'
import RegionalPricing from './RegionalPricing.vue'
import TimeBasedPricing from './TimeBasedPricing.vue'

const pricingData = reactive({
  flatRate: { enabled: false, price: 0, currency: 'IRR' },
  weightBased: { enabled: false, weightUnit: 'kg', weightRanges: [] },
  distanceBased: { enabled: false, distanceUnit: 'km', distanceRanges: [] },
  itemCount: { enabled: false, itemRanges: [] },
  orderValue: { enabled: false, currency: 'IRR', valueRanges: [] },
  specialDiscount: { enabled: false, discounts: [] },
  regional: { enabled: false, regions: [] },
  timeBased: { enabled: false, timeRanges: [] }
})

function getSettingName(key) {
  const names = {
    flatRate: 'قیمت ثابت',
    weightBased: 'قیمت بر اساس وزن',
    distanceBased: 'قیمت بر اساس فاصله',
    itemCount: 'قیمت بر اساس تعداد آیتم',
    orderValue: 'قیمت بر اساس ارزش سفارش',
    specialDiscount: 'تخفیف‌های خاص',
    regional: 'قیمت‌گذاری منطقه‌ای',
    timeBased: 'قیمت‌گذاری در ساعات خاص'
  }
  return names[key] || key
}
</script>

<style scoped>
.shipping-method-pricing-settings {
  direction: rtl;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px #0001;
  padding: 2rem;
  max-width: 1200px;
  margin: 2rem auto;
}

.section-title {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 2rem;
  text-align: right;
  color: #333;
}

.pricing-grid {
  display: grid;
  gap: 2rem;
  margin-bottom: 2rem;
}

.summary-section {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #e9ecef;
}

.summary-section h3 {
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: #555;
}

.summary-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}

.summary-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.setting-name {
  font-weight: 500;
  color: #555;
}

.setting-status {
  padding: 0.3rem 0.8rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
  background: #dc3545;
  color: white;
}

.setting-status.active {
  background: #28a745;
}
</style> 
