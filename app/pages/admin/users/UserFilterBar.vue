<template>
  <div class="flex flex-wrap gap-3 mb-4 items-center bg-white rounded-xl px-4 py-3 shadow-sm font-sans">
    <div class="flex-1">
      <input v-model="filters.search" @input="onFilterChange" type="text" placeholder="جستجو (نام، ایمیل، شماره)"
        class="input input-bordered input-lg w-full max-w-xl" />
    </div>
    <div class="flex items-center gap-2">
      <select v-model="filters.role" @change="onFilterChange" class="input input-sm border border-gray-300 rounded-lg px-3 py-2 bg-white shadow-none focus:shadow-none font-sans text-sm font-bold">
        <option value="">همه نقش‌ها</option>
        <option value="admin">ادمین</option>
        <option value="user">کاربر</option>
        <option value="developer">دولوپر</option>
        <option value="customer">مشتری</option>
        <option value="editor">ویرایشگر</option>
      </select>
    </div>
  </div>
</template>
<script setup lang="ts">


const props = defineProps({
  filters: Object
})

function onFilterChange() {
  // userStore.fetchUsers();
}
</script>
<style scoped>
.btn-success,
.btn-export-excel {
  background-color: #22c55e !important;
  color: #fff !important;
  border: none !important;
  box-shadow: 0 2px 8px 0 #22c55e22;
  font-weight: bold;
  padding: 0.5rem 1.2rem;
  border-radius: 0.5rem;
  transition: background 0.2s;
}
.btn-success:hover,
.btn-export-excel:hover {
  background-color: #16a34a !important;
}
input:focus, select:focus {
  outline: none;
  border-color: #e5e7eb; /* tailwind gray-200 */
  box-shadow: none;
}
select, select option {
  font-family: inherit;
  font-size: 0.875rem; /* text-sm */
  font-weight: bold;
}
</style> 
