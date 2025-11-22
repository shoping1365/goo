<template>
      <div class="py-6 px-2 rtl">
        <div class="flex items-center justify-between mb-4">
          <button class="btn btn-sm btn-outline" @click="goBack">بازگشت به لیست کاربران</button>
          <button v-if="user && !isEditingInline" class="px-3 py-2 rounded bg-blue-600 text-white text-sm" @click="startInlineEdit">ویرایش کاربر</button>
        </div>
        <div v-if="user">
          <div class="mb-6">
            <template v-if="!isEditingInline">
              <UserProfileTab :user="user" @edit="startInlineEdit" />
            </template>
            <template v-else>
              <div class="bg-white rounded-xl shadow-lg p-6">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                  <div>
                    <label class="block text-sm mb-1">نام</label>
                    <input v-model="editForm.firstName" class="w-full border rounded p-2" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">نام خانوادگی</label>
                    <input v-model="editForm.lastName" class="w-full border rounded p-2" />
                  </div>
                  <div class="md:col-span-2">
                    <label class="block text-sm mb-1">نام نمایشی/کامل</label>
                    <input v-model="editForm.name" class="w-full border rounded p-2" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">ایمیل</label>
                    <input v-model="editForm.email" class="w-full border rounded p-2 ltr" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">نام کاربری</label>
                    <input v-model="editForm.username" class="w-full border rounded p-2 ltr" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">شماره همراه</label>
                    <input v-model="editForm.mobile" class="w-full border rounded p-2 ltr" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">شماره ثابت</label>
                    <input v-model="editForm.landline" class="w-full border rounded p-2 ltr" />
                  </div>
                  <div>
                    <label class="block text-sm mb-1">کد ملی</label>
                    <input v-model="editForm.nationalCode" class="w-full border rounded p-2 ltr" />
                  </div>
                  <div class="md:col-span-2">
                    <label class="block text-sm mb-1">آدرس منتخب</label>
                    <input v-model="editForm.selectedAddress" class="w-full border rounded p-2" />
                  </div>
                </div>
                <div class="mt-4 flex justify-end gap-2">
                  <button class="px-3 py-1 rounded bg-gray-100" @click="isEditingInline=false">انصراف</button>
                  <button class="px-3 py-1 rounded bg-blue-600 text-white" @click="saveUserEdits">ذخیره</button>
                </div>
              </div>
            </template>
          </div>
          <div class="mb-6">
            <UserAdminActionsTab :user="user" @action="handleAction" />
          </div>
          <div class="mb-6">
            <UserAddressesTable :user="user" />
          </div>
          <div class="mb-6">
            <UserBankInfoSection :user="user" />
          </div>
          <div class="mb-6">
            <UserWalletSection :user="user" />
          </div>
          <div class="mb-6">
            <UserMessagesSection :user="user" @new-message="handleNewMessage" />
          </div>
          <div class="mb-6">
            <UserCartSection :user="user" />
          </div>
          <div class="mb-6">
            <UserNextCartSection :user="user" />
          </div>
          <div class="mb-6">
            <UserWishlistSection :user="user" />
          </div>
          <div class="mb-6">
            <UserStockAlertsSection :user="user" />
          </div>
          <div class="mb-6">
            <UserDiscountAlertsSection :user="user" />
          </div>
          <div class="mb-6">
            <UserActivityTable :user="user" />
          </div>
          <div class="mb-6">
            <UserOrdersTable :user="user" />
          </div>
          <div class="mb-6">
            <UserBehaviorTable :user="user" />
          </div>
          <div class="mb-6">
            <UserRecentViewsSection :user="user" />
          </div>
          <div class="mb-6">
            <UserActiveSessionsSection :user="user" />
          </div>
          <div class="mb-6">
            <UserSuccessfulLoginsSection :user="user" />
          </div>
          <div class="mb-6">
            <UserFailedAttemptsSection :user="user" />
          </div>
        </div>
        <div v-else class="text-center text-gray-500 py-10">کاربر پیدا نشد.</div>
      </div>
    <!-- حذف مودال قدیمی؛ ویرایش به صورت inline پیاده شد -->
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const useRoute: () => { params: { id: string } }
declare const useRouter: () => { push: (path: string) => void }
declare const $fetch: <T = unknown>(url: string, options?: { method?: string; body?: unknown; credentials?: string }) => Promise<T>
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import UserActiveSessionsSection from './UserActiveSessionsSection.vue';
import UserActivityTable from './UserActivityTable.vue';
import UserAddressesTable from './UserAddressesTable.vue';
import UserAdminActionsTab from './UserAdminActionsTab.vue';
import UserBankInfoSection from './UserBankInfoSection.vue';
import UserBehaviorTable from './UserBehaviorTable.vue';
import UserCartSection from './UserCartSection.vue';
import UserDiscountAlertsSection from './UserDiscountAlertsSection.vue';
import UserFailedAttemptsSection from './UserFailedAttemptsSection.vue';
import UserMessagesSection from './UserMessagesSection.vue';
import UserNextCartSection from './UserNextCartSection.vue';
import UserOrdersTable from './UserOrdersTable.vue';
import UserProfileTab from './UserProfileTab.vue';
import UserRecentViewsSection from './UserRecentViewsSection.vue';
import UserStockAlertsSection from './UserStockAlertsSection.vue';
import UserSuccessfulLoginsSection from './UserSuccessfulLoginsSection.vue';
import UserWalletSection from './UserWalletSection.vue';
import UserWishlistSection from './UserWishlistSection.vue';

definePageMeta({ layout: 'admin-main', middleware: 'admin' });

interface UserDetail {
  id: number | string
  ID?: number | string
  name: string
  fullName?: string
  username: string
  email: string
  mobile: string
  landline?: string
  landlineNumber?: string
  nationalCode?: string
  national_code?: string
  selectedAddress?: string
  selected_address?: string
  registered_at?: string
  registeredAt?: string
  registered?: string
  last_seen?: string
  lastSeenAt?: string
  lastSeen?: string
  online?: boolean
  role_id?: number
  status?: string
  first_name?: string
  last_name?: string
  profile_data?: {
    first_name?: string
    last_name?: string
  }
}

interface EditForm {
  name: string
  firstName: string
  lastName: string
  username: string
  email: string
  mobile: string
  landline: string
  nationalCode: string
  selectedAddress: string
}

const route = useRoute();
const router = useRouter();
const user = ref<UserDetail | null>(null);
// const editOpen = ref(false) // deprecated (modal removed)
const isEditingInline = ref(false)
const editForm = ref<EditForm>({ name: '', firstName: '', lastName: '', username: '', email: '', mobile: '', landline: '', nationalCode: '', selectedAddress: '' })

onMounted(async () => {
  const id = route.params.id;
  try {
    const detail = await $fetch<UserDetail>(`/api/users/${id}`)
    user.value = {
      id: detail.id || detail.ID || '',
      name: detail.name || detail.fullName || detail.username || '-',
      username: detail.username,
      email: detail.email,
      mobile: detail.mobile,
      landline: detail.landline || detail.landlineNumber || '',
      nationalCode: detail.nationalCode || detail.national_code || '',
      selectedAddress: detail.selectedAddress || detail.selected_address || '',
      registered: detail.registered_at || detail.registeredAt || detail.registered || '-',
      last_seen: detail.last_seen || detail.lastSeenAt || detail.lastSeen || '-',
      online: !!detail.online,
      role_id: detail.role_id,
      status: detail.status,
    }
    editForm.value = {
      name: user.value.name,
      firstName: detail.first_name || detail.profile_data?.first_name || '',
      lastName: detail.last_name || detail.profile_data?.last_name || '',
      username: user.value.username || '',
      email: user.value.email || '',
      mobile: user.value.mobile || '',
      landline: user.value.landline || '',
      nationalCode: user.value.nationalCode || '',
      selectedAddress: user.value.selectedAddress || ''
    }
  } catch (err) {
    user.value = null
  }
});

function goBack() {
  router.push('/admin/users');
}
function handleAction(_action: string) {
  // Placeholder for admin actions
}
function handleNewMessage() {
  // Placeholder for sending a new message
}

function openEditModal() {
  if (!user.value) return
  editForm.value = {
    name: user.value.name,
    firstName: editForm.value.firstName,
    lastName: editForm.value.lastName,
    username: user.value.username || '',
    email: user.value.email || '',
    mobile: user.value.mobile || '',
    landline: user.value.landline || '',
    nationalCode: user.value.nationalCode || '',
    selectedAddress: user.value.selectedAddress || ''
  }
  isEditingInline.value = true
}

function startInlineEdit(){
  openEditModal()
}

async function saveUserEdits() {
  if (!user.value) return
  const payload: Record<string, unknown> = {
    name: editForm.value.name,
    email: editForm.value.email,
    mobile: editForm.value.mobile,
    username: editForm.value.username,
    profile_data: {
      first_name: editForm.value.firstName,
      last_name: editForm.value.lastName,
      landline: editForm.value.landline,
      national_code: editForm.value.nationalCode,
      selected_address: editForm.value.selectedAddress
    }
  }
  await $fetch(`/api/users/${user.value.id}`, { method: 'PUT', body: payload, credentials: 'include' })
  Object.assign(user.value, {
    name: editForm.value.name,
    username: editForm.value.username,
    email: editForm.value.email,
    mobile: editForm.value.mobile,
    landline: editForm.value.landline,
    nationalCode: editForm.value.nationalCode,
    selectedAddress: editForm.value.selectedAddress
  })
  isEditingInline.value = false
}
</script> 