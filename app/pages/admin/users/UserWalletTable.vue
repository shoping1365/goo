<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-emerald-700 mb-2 flex justify-between items-center">
      <span>کیف پول</span>
      <button class="bg-emerald-100 hover:bg-emerald-200 text-emerald-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" @click="$emit('view-all-transactions')">
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        مشاهده همه
      </button>
    </div>
    <table class="min-w-full text-sm">
      <thead>
        <tr class="bg-gray-100">
          <th class="p-2 text-right font-bold">نوع تراکنش</th>
          <th class="p-2 text-right font-bold">تاریخ</th>
          <th class="p-2 text-right font-bold">مبلغ</th>
          <th class="p-2 text-right font-bold">توضیحات</th>
          <th class="p-2 text-right font-bold">وضعیت</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="transaction in transactions" :key="transaction.id" class="border-b hover:bg-gray-50">
          <td class="p-2 text-right">{{ transaction.type }}</td>
          <td class="p-2 text-right">{{ transaction.date }}</td>
          <td class="p-2 text-right" :class="transaction.amount > 0 ? 'text-green-600' : 'text-red-600'">
            {{ transaction.amount > 0 ? '+' : '' }}{{ transaction.amount }} تومان
          </td>
          <td class="p-2 text-right">{{ transaction.description }}</td>
          <td class="p-2 text-right">
            <span :class="transaction.status === 'موفق' ? 'text-green-600' : 'text-red-600'">{{ transaction.status }}</span>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="mt-4 p-3 bg-emerald-50 rounded-lg">
      <div class="text-emerald-800 font-bold">موجودی فعلی: <span class="text-2xl">{{ balance }} تومان</span></div>
    </div>
  </div>
</template>
<script lang="ts">
declare const $fetch: <T = unknown>(url: string, options?: { credentials?: string }) => Promise<T>
</script>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{ user: any }>()

const balance = ref<number>(0)
const transactions = ref<any[]>([])

function toDateString(d?: string | Date | null) {
  if (!d) return '-'
  try { return new Date(d as any).toLocaleDateString('fa-IR') } catch { return '-' }
}

async function loadWallet(userId?: number | string) {
  if (!userId) return
  const summary: any = await $fetch(`/api/users/${userId}/wallet/summary`, { credentials: 'include' })
  balance.value = Number(summary?.balance || 0)
  const txs: any[] = await $fetch(`/api/users/${userId}/wallet/transactions?limit=10`, { credentials: 'include' })
  transactions.value = (txs || []).map(t => ({
    id: t.id,
    type: t.type === 'credit' ? 'شارژ کیف پول' : 'برداشت/مصرف',
    date: toDateString(t.created_at || t.createdAt),
    amount: Number(t.amount || 0),
    description: t.gateway || t.method || '-',
    status: (t.status === 'success') ? 'موفق' : (t.status === 'pending' ? 'در انتظار' : 'ناموفق')
  }))
}

watch(() => props.user?.id, (id) => loadWallet(id), { immediate: true })
</script> 
