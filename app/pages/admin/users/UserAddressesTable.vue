<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-orange-700 mb-2 flex justify-between items-center">
      <span>آدرس‌های کاربر</span>
      <button class="bg-orange-100 hover:bg-orange-200 text-orange-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-addresses')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نوع</th>
          <th class="p-2 text-right font-bold">آدرس</th>
          <th class="p-2 text-right font-bold">کد پستی</th>
          <th class="p-2 text-right font-bold">شماره تماس</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="address in addresses" :key="address.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ address.type }}</td>
          <td class="p-2 text-right">{{ address.address }}</td>
          <td class="p-2 text-right">{{ address.postalCode }}</td>
          <td class="p-2 text-right">{{ address.phone }}</td>
          <td class="p-2 text-right">
            <span :class="address.isDefault ? 'text-green-600' : 'text-gray-600'">{{ address.isDefault ? 'پیش‌فرض' : 'عادی' }}</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup lang="ts">
import { ref, watch } from 'vue';
import type { User } from '~/types/user';

const props = defineProps<{ user: User }>()
defineEmits(['view-all-addresses']);

type AddressRow = {
  id: number
  type: string
  address: string
  postalCode?: string
  phone?: string
  isDefault: boolean
}

interface AddressResponse {
  id: number;
  label?: string;
  province?: string;
  city?: string;
  street?: string;
  postal_code?: string;
  postalCode?: string;
  phone?: string;
  recipient_mobile?: string;
  recipientMobile?: string;
  is_default?: boolean;
  isDefault?: boolean;
  [key: string]: unknown;
}

const addresses = ref<AddressRow[]>([])

async function loadAddresses(userId?: number | string) {
  if (!userId) return
  try {
    const response = await fetch(`/api/users/${userId}/addresses`, {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    
    const data = (await response.json()) as AddressResponse[]

    if (Array.isArray(data)) {
      addresses.value = data.map(a => ({
        id: a.id,
        type: a.label || 'عادی',
        address: [a.province, a.city, a.street].filter(Boolean).join('، '),
        postalCode: a.postal_code || a.postalCode,
        phone: a.phone || a.recipient_mobile || a.recipientMobile,
        isDefault: !!a.is_default || !!a.isDefault
      }))
    } else {
      addresses.value = []
    }
  } catch {
    addresses.value = []
  }
}

watch(() => props.user?.id, (id) => loadAddresses(id), { immediate: true })
</script> 
