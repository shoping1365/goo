<template>
     <div class="sensitive-products">
       <div class="settings-card">
         <div class="card-header">
           <h3>محصولات حساس</h3>
           <div class="toggle-switch">
             <input type="checkbox" v-model="enabled" id="sensitiveToggle" />
             <label for="sensitiveToggle"></label>
           </div>
         </div>
   
         <div v-if="enabled" class="card-body">
           <div class="form-section">
             <h4>دسته‌بندی‌های حساس</h4>
             <div class="sensitive-categories">
               <div v-for="(category, index) in sensitive.categories" :key="index" class="sensitive-category">
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
                       <option value="humidity">حساس به رطوبت</option>
                       <option value="light">حساس به نور</option>
                       <option value="custom">سفارشی</option>
                     </select>
                   </div>
                   <button @click="removeCategory(index)" class="remove-btn">حذف</button>
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
                       <option value="temperature_controlled">کنترل دما</option>
                       <option value="humidity_controlled">کنترل رطوبت</option>
                     </select>
                   </div>
                 </div>
   
                 <div v-if="category.sensitivityType === 'temperature' || category.sensitivityType === 'humidity'" class="form-row">
                   <div class="form-group">
                     <label>حداقل دما (°C)</label>
                     <input
                       v-model.number="category.minTemperature"
                       type="number"
                       class="form-input"
                       placeholder="0"
                     />
                   </div>
                   <div class="form-group">
                     <label>حداکثر دما (°C)</label>
                     <input
                       v-model.number="category.maxTemperature"
                       type="number"
                       class="form-input"
                       placeholder="25"
                     />
                   </div>
                 </div>
   
                 <div v-if="category.sensitivityType === 'humidity'" class="form-row">
                   <div class="form-group">
                     <label>حداقل رطوبت (%)</label>
                     <input
                       v-model.number="category.minHumidity"
                       type="number"
                       min="0"
                       max="100"
                       class="form-input"
                       placeholder="30"
                     />
                   </div>
                   <div class="form-group">
                     <label>حداکثر رطوبت (%)</label>
                     <input
                       v-model.number="category.maxHumidity"
                       type="number"
                       min="0"
                       max="100"
                       class="form-input"
                       placeholder="70"
                     />
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
               <button @click="addCategory" class="add-btn">+ افزودن دسته‌بندی حساس</button>
             </div>
           </div>
   
           <div class="form-section">
             <h4>تنظیمات اضافی</h4>
             <div class="additional-settings">
               <div class="form-group">
                 <label>نحوه تشخیص محصولات حساس</label>
                 <div class="detection-methods">
                   <label class="radio-label">
                     <input type="radio" v-model="sensitive.detectionMethod" value="automatic" />
                     تشخیص خودکار
                   </label>
                   <label class="radio-label">
                     <input type="radio" v-model="sensitive.detectionMethod" value="manual" />
                     تشخیص دستی
                   </label>
                   <label class="radio-label">
                     <input type="radio" v-model="sensitive.detectionMethod" value="both" />
                     هر دو روش
                   </label>
                 </div>
               </div>
   
               <div class="form-group">
                 <label>اقدام در صورت تشخیص محصول حساس</label>
                 <select v-model="sensitive.action" class="form-input">
                   <option value="apply_conditions">اعمال شرایط خاص</option>
                   <option value="require_approval">نیاز به تایید</option>
                   <option value="warning">هشدار</option>
                   <option value="custom">اقدام سفارشی</option>
                 </select>
               </div>
   
               <div class="form-group">
                 <label>پیام اطلاع‌رسانی</label>
                 <textarea
                   v-model="sensitive.notificationMessage"
                   class="form-input"
                   placeholder="پیام نمایشی برای محصولات حساس"
                   rows="3"
                 ></textarea>
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
   
   const sensitive = ref({
     categories: props.modelValue.categories || [
       {
         name: 'شیشه و سرامیک',
         sensitivityType: 'fragile',
         additionalCost: 50000,
         specialConditions: 'fragile_handling',
         description: 'نیاز به حمل و نقل ویژه',
         minTemperature: null,
         maxTemperature: null,
         minHumidity: null,
         maxHumidity: null
       },
       {
         name: 'مواد غذایی فاسدشدنی',
         sensitivityType: 'perishable',
         additionalCost: 30000,
         specialConditions: 'refrigerated',
         description: 'نیاز به سردخانه',
         minTemperature: 2,
         maxTemperature: 8,
         minHumidity: null,
         maxHumidity: null
       },
       {
         name: 'لوازم الکترونیکی',
         sensitivityType: 'valuable',
         additionalCost: 75000,
         specialConditions: 'insurance',
         description: 'نیاز به بیمه و حمل ویژه',
         minTemperature: null,
         maxTemperature: null,
         minHumidity: null,
         maxHumidity: null
       }
     ],
     detectionMethod: props.modelValue.detectionMethod || 'automatic',
     action: props.modelValue.action || 'apply_conditions',
     notificationMessage: props.modelValue.notificationMessage || 'این محصول در دسته محصولات حساس قرار دارد و شرایط خاصی برای حمل و نقل اعمال خواهد شد.'
   })
   
   function addCategory() {
     sensitive.value.categories.push({
       name: '',
       sensitivityType: 'fragile',
       additionalCost: 0,
       specialConditions: 'none',
       description: '',
       minTemperature: null,
       maxTemperature: null,
       minHumidity: null,
       maxHumidity: null
     })
   }
   
   function removeCategory(index) {
     sensitive.value.categories.splice(index, 1)
   }
   
   // Watch for changes and emit updates
   watch([enabled, sensitive], () => {
     emit('update:modelValue', {
       enabled: enabled.value,
       ...sensitive.value
     })
   }, { deep: true })
   </script>
   
   <style scoped>
   .sensitive-products {
     background: #fff;
     border-radius: 12px;
     box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
     overflow: hidden;
   }
   
   .settings-card {
     background: #fff;
   }
   
   .card-header {
     background: linear-gradient(135deg, #ffc107 0%, #e0a800 100%);
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
     border-bottom: 2px solid #ffc107;
     padding-bottom: 0.5rem;
   }
   
   .form-row {
     display: grid;
     grid-template-columns: 1fr 1fr auto;
     gap: 1.5rem;
     margin-bottom: 1rem;
     align-items: end;
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
     border-color: #ffc107;
     box-shadow: 0 0 0 2px rgba(255, 193, 7, 0.25);
   }
   
   .sensitive-category {
     background: white;
     border: 1px solid #e9ecef;
     border-radius: 8px;
     padding: 1rem;
     margin-bottom: 1rem;
     border-left: 4px solid #ffc107;
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
     transition: background-color 0.3s;
   }
   
   .remove-btn:hover {
     background: #c82333;
   }
   
   .add-btn {
     background: #28a745;
     color: white;
     border: none;
     padding: 0.75rem 1.5rem;
     border-radius: 6px;
     cursor: pointer;
     font-size: 0.9rem;
     margin-top: 1rem;
     transition: background-color 0.3s;
   }
   
   .add-btn:hover {
     background: #218838;
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
   
   .detection-methods {
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
