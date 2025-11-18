<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-indigo-700 mb-2 flex justify-between items-center">
      <span>اطلاعات بانکی</span>
      <button class="bg-indigo-100 hover:bg-indigo-200 text-indigo-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-bank-info')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نام بانک</th>
          <th class="p-2 text-right font-bold">شماره کارت</th>
          <th class="p-2 text-right font-bold">شماره شبا</th>
          <th class="p-2 text-right font-bold">صاحب حساب</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="bank in bankAccounts" :key="bank.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ bank.bankName }}</td>
          <td class="p-2 text-right">{{ bank.cardNumber }}</td>
          <td class="p-2 text-right">{{ bank.shebaNumber }}</td>
          <td class="p-2 text-right">{{ bank.accountHolder }}</td>
          <td class="p-2 text-right">
            <span :class="bank.isVerified ? 'text-green-600' : 'text-yellow-600'">{{ bank.isVerified ? 'تایید شده' : 'در انتظار تایید' }}</span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script lang="ts">
declare const $fetch: <T = unknown>(url: string, options?: { credentials?: string }) => Promise<T>
</script>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{ user: any }>()

type BankRow = {
  id: number
  bankName?: string
  cardNumber?: string
  shebaNumber?: string
  accountHolder?: string
  isVerified?: boolean
}

const bankAccounts = ref<BankRow[]>([])

async function loadBankAccounts(userId?: number | string) {
  if (!userId) return
  const res = await $fetch(`/api/users/${userId}/bank-accounts`, { credentials: 'include' }) as any[]
  bankAccounts.value = (res || []).map(a => ({
    id: a.id,
    bankName: a.bank_name || a.bankName,
    cardNumber: a.card_number_last4 ? `****-****-****-${a.card_number_last4}` : '',
    shebaNumber: a.iban || a.IBAN,
    accountHolder: props.user?.name || props.user?.username,
    isVerified: !!a.verified_at
  }))
}

watch(() => props.user?.id, (id) => loadBankAccounts(id), { immediate: true })
</script> 
