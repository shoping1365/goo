<template>
  <div class="bg-gradient-to-tr from-green-50 to-white rounded-2xl shadow-lg p-6 mt-2">
    <div class="font-bold text-green-700 mb-4 flex items-center gap-2">
      <svg class="w-5 h-5 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a5 5 0 00-10 0v2M5 13h14M12 17v.01"/></svg>
      کیف پول
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="flex items-center gap-2 bg-white rounded-lg p-3 shadow-sm">
        <span class="bg-green-100 text-green-600 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 1.343-3 3s1.343 3 3 3 3-1.343 3-3-1.343-3-3-3zm0 13c-4.418 0-8-3.582-8-8s3.582-8 8-8 8 3.582 8 8-3.582 8-8 8z"/></svg>
        </span>
        <span class="font-semibold">موجودی کیف پول:</span>
        <span class="text-gray-700">{{ walletBalance }}</span>
      </div>
      <div class="flex items-center gap-2 bg-white rounded-lg p-3 shadow-sm">
        <span class="bg-blue-100 text-blue-600 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3"/></svg>
        </span>
        <span class="font-semibold">آخرین تراکنش:</span>
        <span class="text-gray-700">{{ lastTransactionAmount }}</span>
        <span class="text-xs text-gray-400">({{ lastTransactionDate }})</span>
      </div>
      <div class="flex items-center gap-2 bg-white rounded-lg p-3 shadow-sm">
        <span class="bg-yellow-100 text-yellow-600 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/></svg>
        </span>
        <span class="font-semibold">وضعیت کیف پول:</span>
        <span :class="walletStatus === 'فعال' ? 'text-green-600' : 'text-red-600'">{{ walletStatus }}</span>
      </div>
      <button
        class="flex items-center justify-center gap-2 bg-white rounded-lg p-3 shadow-sm hover:bg-green-100 text-green-700 font-semibold transition h-full w-full"
        @click="$emit('view-all')"
        type="button"
      >
        <span class="bg-green-100 text-green-600 rounded-lg p-2">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
        </span>
        مشاهده تمام تراکنش‌ها
      </button>
    </div>
  </div>
</template>
<script lang="ts">
declare const $fetch: <T = unknown>(url: string, options?: { credentials?: string }) => Promise<T>
</script>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{ user?: any }>()

const walletBalance = ref<string>('-')
const lastTransactionAmount = ref<string>('-')
const lastTransactionDate = ref<string>('-')
const walletStatus = ref<string>('نامشخص')

function formatAmount(n: number) {
  try { return n.toLocaleString('fa-IR') + ' تومان' } catch { return String(n) }
}

function toFaDate(d?: string | Date | null) {
  if (!d) return '-'
  try { return new Date(d as any).toLocaleDateString('fa-IR') } catch { return '-' }
}

async function loadWallet(userId?: number | string) {
  if (!userId) return
  const summary: any = await $fetch(`/api/users/${userId}/wallet/summary`, { credentials: 'include' })
  walletBalance.value = formatAmount(Number(summary?.balance || 0))
  walletStatus.value = summary?.status === 'active' ? 'فعال' : 'غیرفعال'

  const txs: any[] = await $fetch(`/api/users/${userId}/wallet/transactions?limit=1`, { credentials: 'include' })
  if (Array.isArray(txs) && txs.length) {
    const t = txs[0]
    const sign = t.type === 'credit' ? '+' : '-'
    lastTransactionAmount.value = `${sign}${formatAmount(Number(t.amount || 0))}`
    lastTransactionDate.value = toFaDate(t.created_at || t.createdAt)
  } else {
    lastTransactionAmount.value = '-'
    lastTransactionDate.value = '-'
  }
}

watch(() => props.user?.id, (id) => loadWallet(id), { immediate: true })
</script> 
