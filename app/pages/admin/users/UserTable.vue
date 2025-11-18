<template>
  <div class="overflow-x-auto bg-white rounded-xl shadow-lg">
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-center align-middle">
            <input type="checkbox" v-model="allSelected" />
          </th>
          <th class="p-2 text-center align-middle">ردیف</th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-purple-700"><i class="fa fa-id-card"></i></span>
              <span>ID</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-purple-700"><i class="fa fa-user"></i></span>
              <span>نام کامل</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-purple-700"><i class="fa fa-user"></i></span>
              <span>نام کاربری</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-blue-700"><i class="fa fa-envelope"></i></span>
              <span>ایمیل</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-blue-700"><i class="fa fa-calculator"></i></span>
              <span>شماره همراه</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-pink-700"><i class="fa fa-phone"></i></span>
              <span>شماره ثابت</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-yellow-700"><i class="fa fa-tag"></i></span>
              <span>نقش</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-blue-700"><i class="fa fa-calendar"></i></span>
              <span>تاریخ ثبت نام</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-pink-700"><i class="fa fa-clock"></i></span>
              <span>آخرین بازدید</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-green-700"><i class="fa fa-circle"></i></span>
              <span>وضعیت</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">
            <div class="flex items-center justify-center gap-1">
              <span class="text-yellow-700"><i class="fa fa-key"></i></span>
              <span>۲مرحله‌ای</span>
            </div>
          </th>
          <th class="p-2 text-center align-middle">جزئیات</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(user, idx) in displayedUsers" :key="user.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-center align-middle">
            <input type="checkbox" v-model="selected" :value="user.id" />
          </td>
          <td class="p-2 text-center align-middle">{{ idx + 1 }}</td>
          <td class="p-2 text-center align-middle">{{ user.id }}</td>
          <td class="p-2 text-center align-middle">{{ user.name }}</td>
          <td class="p-2 text-center align-middle">{{ user.username || '-' }}</td>
          <td class="p-2 text-center align-middle">{{ user.email || '-' }}</td>
          <td class="p-2 text-center align-middle">{{ user.mobile || user.phone || '-' }}</td>
          <td class="p-2 text-center align-middle">{{ user.landline || '-' }}</td>
          <td class="p-2 text-center align-middle">{{ user.roleName || user.role || '-' }}</td>
          <td class="p-2 text-center align-middle">{{ toPersianDate(user.registered) }}</td>
          <td class="p-2 text-center align-middle">{{ toPersianDateTime(user.lastSeen) }}</td>
          <td class="p-2 text-center align-middle">
            <span v-if="user.online" class="inline-flex items-center gap-1 text-green-600 font-bold">
              <span class="w-2 h-2 rounded-full bg-green-500 inline-block"></span>
              آنلاین
            </span>
            <span v-else class="text-gray-400">آفلاین</span>
          </td>
          <td class="p-2 text-center align-middle">
            <span :class="user.twofa ? 'text-green-600' : 'text-gray-400'">
              {{ user.twofa ? 'فعال' : 'غیرفعال' }}
            </span>
          </td>
          <td class="p-2 text-center align-middle">
            <NuxtLink :to="`/admin/users/${user.id}`" class="btn btn-xs btn-info">جزئیات</NuxtLink>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import { computed, ref, watch, onMounted } from "vue";
import { toPersianDate, toPersianDateTime } from "~/utils/dateUtils";

interface User {
  id: number | string;
  name: string;
  username?: string;
  email?: string;
  phone?: string;
  mobile?: string;
  landline?: string;
  role?: string;
  roleName?: string;
  status?: string;
  registered?: string;
  lastSeen?: string;
  online?: boolean;
  twofa?: boolean;
}

interface Filters {
  search?: string;
  role?: string;
  status?: string;
  onlineOnly?: boolean;
}

const props = defineProps<{ 
  users: User[];
  filters?: Filters;
}>()

const selected = ref<(string | number)[]>([]);
const allSelected = computed({
  get: () => displayedUsers.value.length > 0 && displayedUsers.value.every(u => selected.value.includes(u.id)),
  set: val => toggleAll(val)
});

function toggleAll(val: boolean) {
  if (val) {
    selected.value = displayedUsers.value.map(u => u.id);
  } else {
    selected.value = [];
  }
}

const displayedUsers = computed<User[]>(() => {
  let list = props.users;
  
  // استفاده از فیلترهای ارسال شده از والد
  const filters = props.filters || {};
  
  // role filter
  if (filters.role) {
    list = list.filter(u => u.role === filters.role || u.roleName === filters.role);
  }
  
  // status filter
  if (filters.status) {
    list = list.filter(u => u.status === filters.status);
  }
  
  // online filter
  if (filters.onlineOnly) {
    list = list.filter(u => u.online);
  }
  
  // search filter
  if (filters.search) {
    const s = filters.search.toLowerCase();
    list = list.filter(u => 
      (u.name || '').toLowerCase().includes(s) || 
      (u.email || '').toLowerCase().includes(s) || 
      (u.phone || '').toLowerCase().includes(s) ||
      (u.mobile || '').toLowerCase().includes(s)
    );
  }
  
  return list;
});

const emit = defineEmits(['selectionChanged']);

watch(selected, (val: number[]) => {
  // Emit selected user ids to parent
  emit('selectionChanged', val);
});
</script> 
