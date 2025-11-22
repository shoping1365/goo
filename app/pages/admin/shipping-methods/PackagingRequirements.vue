<template>
  <div class="packaging-requirements">
       <div class="settings-card">
         <div class="card-header">
           <h3>نیازمندی‌های خاص بسته‌بندی</h3>
           <div class="toggle-switch">
             <input id="packagingToggle" v-model="enabled" type="checkbox" />
             <label for="packagingToggle"></label>
           </div>
         </div>
   
         <div v-if="enabled" class="card-body">
           <div class="form-section">
             <h4>نوع بسته‌بندی اجباری</h4>
             <div class="packaging-types">
               <label v-for="type in packaging.requiredTypes" :key="type.value" class="packaging-checkbox">
                 <input
                   v-model="packaging.selectedTypes"
                   type="checkbox"
                   :value="type.value"
                 />
                 {{ type.label }}
               </label>
             </div>
           </div>
   
           <div class="form-section">
             <h4>مواد بسته‌بندی ممنوعه</h4>
             <div class="prohibited-materials">
               <div v-for="(material, index) in packaging.prohibitedMaterials" :key="index" class="material-item">
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
                   <button class="remove-btn" @click="removeMaterial(index)">حذف</button>
                 </div>
               </div>
               <button class="add-btn" @click="addMaterial">+ افزودن ماده ممنوعه</button>
             </div>
           </div>
   
           <div class="form-section">
             <h4>استانداردهای بسته‌بندی</h4>
             <div class="packaging-standards">
               <div v-for="(standard, index) in packaging.standards" :key="index" class="standard-item">
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
                   <button class="remove-btn" @click="removeStandard(index)">حذف</button>
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
               <button class="add-btn" @click="addStandard">+ افزودن استاندارد</button>
             </div>
           </div>
   
           <div class="form-section">
             <h4>هزینه بسته‌بندی اضافی</h4>
             <div class="form-row">
               <div class="form-group">
                 <label>هزینه پایه (تومان)</label>
                 <input
                   v-model.number="packaging.baseCost"
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
                   v-model.number="packaging.specialCost"
                   type="number"
                   min="0"
                   step="1000"
                   class="form-input"
                   placeholder="0"
                 />
               </div>
             </div>
           </div>
   
           <div class="form-section">
             <h4>تنظیمات اضافی</h4>
             <div class="additional-settings">
               <div class="form-group">
                 <label>نحوه اعمال بسته‌بندی</label>
                 <div class="application-methods">
                   <label class="radio-label">
                     <input v-model="packaging.applicationMethod" type="radio" value="automatic" />
                     اعمال خودکار
                   </label>
                   <label class="radio-label">
                     <input v-model="packaging.applicationMethod" type="radio" value="manual" />
                     اعمال دستی
                   </label>
                   <label class="radio-label">
                     <input v-model="packaging.applicationMethod" type="radio" value="conditional" />
                     اعمال شرطی
                   </label>
                 </div>
               </div>
   
               <div class="form-group">
                 <label>اقدام در صورت عدم رعایت</label>
                 <select v-model="packaging.nonComplianceAction" class="form-input">
                   <option value="warning">هشدار</option>
                   <option value="reject">رد سفارش</option>
                   <option value="apply_automatically">اعمال خودکار</option>
                   <option value="require_approval">نیاز به تایید</option>
                 </select>
               </div>
   
               <div class="form-group">
                 <label>پیام اطلاع‌رسانی</label>
                 <textarea
                   v-model="packaging.notificationMessage"
                   class="form-input"
                   placeholder="پیام نمایشی برای نیازمندی‌های بسته‌بندی"
                   rows="3"
                 ></textarea>
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
  import { ref, watch } from 'vue'

  definePageMeta({ layout: 'admin-main', middleware: 'admin' })
   
   const props = defineProps({
     modelValue: {
       type: Object,
       default: () => ({})
     }
   })
   
   const emit = defineEmits(['update:modelValue'])
   
   const enabled = ref(props.modelValue.enabled ?? true)
   
   const packaging = ref({
     requiredTypes: [
       { value: 'bubble_wrap', label: 'حباب‌دار' },
       { value: 'cardboard_box', label: 'جعبه مقوایی' },
       { value: 'plastic_wrap', label: 'پلاستیک' },
       { value: 'wooden_box', label: 'جعبه چوبی' },
       { value: 'special_container', label: 'ظرف مخصوص' },
       { value: 'vacuum_seal', label: 'مهر و موم خلاء' },
       { value: 'foam_padding', label: 'فوم محافظ' },
       { value: 'corner_protectors', label: 'محافظ گوشه' }
     ],
     selectedTypes: props.modelValue.selectedTypes || ['bubble_wrap', 'cardboard_box'],
     prohibitedMaterials: props.modelValue.prohibitedMaterials || [
       { name: 'پلاستیک یکبار مصرف', reason: 'environmental' },
       { name: 'فوم پلی‌استایرن', reason: 'environmental' }
     ],
     standards: props.modelValue.standards || [
       { name: 'استاندارد ISO 11683', level: 'required', description: 'استاندارد بین‌المللی بسته‌بندی' },
       { name: 'استاندارد ملی ایران', level: 'recommended', description: 'استاندارد ملی بسته‌بندی' }
     ],
     baseCost: props.modelValue.baseCost || 0,
     specialCost: props.modelValue.specialCost || 0,
     applicationMethod: props.modelValue.applicationMethod || 'automatic',
     nonComplianceAction: props.modelValue.nonComplianceAction || 'warning',
     notificationMessage: props.modelValue.notificationMessage || 'این محصول نیاز به بسته‌بندی خاص دارد.'
   })
   
   function addMaterial() {
     packaging.value.prohibitedMaterials.push({ name: '', reason: 'environmental' })
   }
   
   function removeMaterial(index) {
     packaging.value.prohibitedMaterials.splice(index, 1)
   }
   
   function addStandard() {
     packaging.value.standards.push({ name: '', level: 'recommended', description: '' })
   }
   
   function removeStandard(index) {
     packaging.value.standards.splice(index, 1)
   }
   
   // Watch for changes and emit updates
   watch([enabled, packaging], () => {
     emit('update:modelValue', {
       enabled: enabled.value,
       ...packaging.value
     })
   }, { deep: true })
   </script>
   
   <style scoped>
   .packaging-requirements {
     background: #fff;
     border-radius: 12px;
     box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
     overflow: hidden;
   }
   
   .settings-card {
     background: #fff;
   }
   
   .card-header {
     background: linear-gradient(135deg, #17a2b8 0%, #138496 100%);
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
     border-bottom: 2px solid #17a2b8;
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
     border-color: #17a2b8;
     box-shadow: 0 0 0 2px rgba(23, 162, 184, 0.25);
   }
   
   .packaging-types {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
     gap: 1rem;
     margin-top: 1rem;
   }
   
   .packaging-checkbox {
     display: flex;
     align-items: center;
     gap: 0.5rem;
     cursor: pointer;
     padding: 0.75rem;
     border-radius: 6px;
     transition: background-color 0.3s;
     border: 1px solid #e9ecef;
     background: white;
   }
   
   .packaging-checkbox:hover {
     background: #e9ecef;
   }
   
   .packaging-checkbox input[type="checkbox"] {
     width: 18px;
     height: 18px;
   }
   
   .material-item,
   .standard-item {
     background: white;
     border: 1px solid #e9ecef;
     border-radius: 8px;
     padding: 1rem;
     margin-bottom: 1rem;
     border-left: 4px solid #17a2b8;
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
   
   .application-methods {
     display: flex;
     flex-direction: column;
     gap: 0.5rem;
     margin-top: 0.5rem;
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
