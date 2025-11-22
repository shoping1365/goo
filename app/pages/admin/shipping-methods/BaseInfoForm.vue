<template>
  <div class="shipping-method-base-info-form">
    <h2 class="form-title">اطلاعات پایه روش ارسال</h2>
    <form @submit.prevent>
      <div class="form-row">
        <label>نام روش ارسال</label>
        <input v-model="form.name" type="text" class="input" placeholder="مثلاً پست پیشتاز" />
      </div>
      <div class="form-row">
        <label>توضیحات</label>
        <textarea v-model="form.description" class="input" placeholder="توضیح مختصر درباره این روش ارسال"></textarea>
      </div>
      <div class="form-row">
        <label>آیکون/لوگو</label>
        <input type="file" accept="image/*" @change="onIconChange" />
        <img v-if="form.iconUrl" :src="form.iconUrl" class="icon-preview" />
      </div>
      <div class="form-row">
        <label>وضعیت</label>
        <select v-model="form.active" class="input">
          <option :value="true">فعال</option>
          <option :value="false">غیرفعال</option>
        </select>
      </div>
      <div class="form-row">
        <label>اولویت نمایش</label>
        <input v-model.number="form.priority" type="number" min="1" class="input" placeholder="مثلاً 1" />
      </div>
      <div class="form-row">
        <label>دسته‌بندی</label>
        <select v-model="form.category" class="input">
          <option value="post">پست</option>
          <option value="courier">پیک</option>
          <option value="freight">حمل و نقل سنگین</option>
        </select>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });
const form = ref({
  name: '',
  description: '',
  iconUrl: '',
  active: true,
  priority: 1,
  category: 'post',
})
function onIconChange(e) {
  const file = e.target.files[0]
  if (file) {
    form.value.iconUrl = URL.createObjectURL(file)
  }
}
</script>

<style scoped>
.shipping-method-base-info-form {
  direction: rtl;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px #0001;
  padding: 2rem;
  max-width: 500px;
  margin: 2rem auto;
}
.form-title {
  font-size: 1.3rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
  text-align: right;
}
.form-row {
  margin-bottom: 1.2rem;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}
label {
  font-size: 1rem;
  margin-bottom: 0.3rem;
  color: #333;
}
.input {
  width: 100%;
  padding: 0.5rem 0.8rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 1rem;
  background: #fafbfc;
  transition: border 0.2s;
}
.input:focus {
  border-color: #1976d2;
  outline: none;
}
.icon-preview {
  margin-top: 0.5rem;
  max-width: 60px;
  max-height: 60px;
  border-radius: 8px;
  border: 1px solid #eee;
}
</style> 
