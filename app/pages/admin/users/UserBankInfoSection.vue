<template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <div class="font-bold text-blue-700 mb-2 flex justify-between items-center">
      <span class="flex items-center gap-2">
        <svg class="w-5 h-5 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/>
        </svg>
        اطلاعات بانکی
      </span>
      <button 
        v-if="bankingInfos.length > 0" 
        class="bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg px-3 py-1 text-xs flex items-center transition font-bold" 
        @click="showAll = !showAll"
      >
        <svg class="w-4 h-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
        {{ showAll ? 'نمایش کمتر' : 'مشاهده همه' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="text-center py-4">
      <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
      <p class="mt-2 text-sm text-gray-500">در حال بارگذاری...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="text-center py-4">
      <div class="text-red-500 mb-2">
        <i class="fas fa-exclamation-triangle"></i>
      </div>
      <p class="text-red-600 text-sm">{{ error }}</p>
      <button
        @click="fetchBankingInfo"
        class="mt-2 px-3 py-1 bg-red-100 text-red-600 rounded text-sm hover:bg-red-200"
      >
        تلاش مجدد
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="bankingInfos.length === 0" class="text-center py-4">
      <div class="text-gray-500 mb-2">
        <svg class="w-12 h-12 mx-auto text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/>
        </svg>
      </div>
      <p class="text-gray-600 text-sm">اطلاعات بانکی وجود ندارد</p>
      <p class="text-xs text-gray-500 mt-1">این کاربر هنوز اطلاعات بانکی ثبت نکرده است</p>
    </div>

    <!-- Data Display -->
    <div v-else class="space-y-3">
      <div
        v-for="info in displayedBankingInfos"
        :key="info.id"
        class="bg-gradient-to-r from-blue-50 to-white rounded-lg p-6 border border-blue-100"
      >
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-2">
            <button
              v-if="!info.isVerified"
              @click="verifyBankingInfo(info.id)"
              class="bg-green-100 hover:bg-green-200 text-green-700 px-3 py-1 rounded text-xs font-bold transition flex items-center gap-1"
              title="تایید اطلاعات بانکی"
            >
              <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
              </svg>
              تایید
            </button>
            <button
              v-else
              @click="unverifyBankingInfo(info.id)"
              class="bg-red-100 hover:bg-red-200 text-red-700 px-3 py-1 rounded text-xs font-bold transition flex items-center gap-1"
              title="لغو تایید اطلاعات بانکی"
            >
              <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
              لغو تایید
            </button>
          </div>
          <div class="flex items-center gap-2">
            <span v-if="info.isDefault" class="bg-green-100 text-green-700 px-2 py-1 rounded-full text-xs font-bold">
              پیشفرض
            </span>
            <span v-if="info.isVerified" class="bg-blue-100 text-blue-700 px-2 py-1 rounded-full text-xs font-bold">
              تایید شده
            </span>
            <span v-else class="bg-yellow-100 text-yellow-700 px-2 py-1 rounded-full text-xs font-bold">
              تایید نشده
            </span>
          </div>
        </div>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-3 text-sm">
          <div class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">نام بانک:</span>
            <span class="text-gray-900">{{ info.bankName }}</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">شماره کارت:</span>
            <span class="ltr text-gray-900 font-mono">{{ formatCardNumber(info.cardNumber) }}</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">شماره حساب:</span>
            <span class="ltr text-gray-900 font-mono">{{ info.accountNumber }}</span>
          </div>
          <div v-if="info.shebaNumber" class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">شماره شبا:</span>
            <span class="ltr text-gray-900 font-mono text-xs">{{ formatShebaNumber(info.shebaNumber) }}</span>
          </div>
          <div v-if="info.accountHolderName" class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">صاحب حساب:</span>
            <span class="text-gray-900">{{ info.accountHolderName }}</span>
          </div>
          <div v-if="info.accountType" class="flex items-center gap-2">
            <span class="text-gray-600 font-medium">نوع حساب:</span>
            <span class="text-gray-900">{{ info.accountType }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <ViewAllModal v-model="showAll" title="اطلاعات بانکی کاربر">
      <div class="space-y-3">
        <div
          v-for="info in bankingInfos"
          :key="info.id"
          class="bg-gradient-to-r from-blue-50 to-white rounded-lg p-6 border border-blue-100"
        >
          <div class="flex items-center justify-between mb-3">
            <div class="flex items-center gap-2">
              <button
                v-if="!info.isVerified"
                @click="verifyBankingInfo(info.id)"
                class="bg-green-100 hover:bg-green-200 text-green-700 px-3 py-1 rounded text-xs font-bold transition flex items-center gap-1"
                title="تایید اطلاعات بانکی"
              >
                <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
                تایید
              </button>
              <button
                v-else
                @click="unverifyBankingInfo(info.id)"
                class="bg-red-100 hover:bg-red-200 text-red-700 px-3 py-1 rounded text-xs font-bold transition flex items-center gap-1"
                title="لغو تایید اطلاعات بانکی"
              >
                <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                </svg>
                لغو تایید
              </button>
            </div>
            <div class="flex items-center gap-2">
              <span v-if="info.isDefault" class="bg-green-100 text-green-700 px-2 py-1 rounded-full text-xs font-bold">
                پیشفرض
              </span>
              <span v-if="info.isVerified" class="bg-blue-100 text-blue-700 px-2 py-1 rounded-full text-xs font-bold">
                تایید شده
              </span>
              <span v-else class="bg-yellow-100 text-yellow-700 px-2 py-1 rounded-full text-xs font-bold">
                تایید نشده
              </span>
            </div>
          </div>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-3 text-sm">
            <div class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">نام بانک:</span>
              <span class="text-gray-900">{{ info.bankName }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">شماره کارت:</span>
              <span class="ltr text-gray-900 font-mono">{{ formatCardNumber(info.cardNumber) }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">شماره حساب:</span>
              <span class="ltr text-gray-900 font-mono">{{ info.accountNumber }}</span>
            </div>
            <div v-if="info.shebaNumber" class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">شماره شبا:</span>
              <span class="ltr text-gray-900 font-mono text-xs">{{ formatShebaNumber(info.shebaNumber) }}</span>
            </div>
            <div v-if="info.accountHolderName" class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">صاحب حساب:</span>
              <span class="text-gray-900">{{ info.accountHolderName }}</span>
            </div>
            <div v-if="info.accountType" class="flex items-center gap-2">
              <span class="text-gray-600 font-medium">نوع حساب:</span>
              <span class="text-gray-900">{{ info.accountType }}</span>
            </div>
          </div>
        </div>
      </div>
    </ViewAllModal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import ViewAllModal from '~/components/admin/modals/ViewAllModal.vue';

const props = defineProps<{ user: any }>();
const showAll = ref(false);

// Real data for banking info
const bankingInfos = ref([]);
const loading = ref(false);
const error = ref('');

const displayedBankingInfos = computed(() => {
  return showAll.value ? bankingInfos.value : bankingInfos.value.slice(0, 2);
});

const fetchBankingInfo = async () => {
  if (!props.user?.id) return;

  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch(`/api/admin/users/${props.user.id}/banking-info`) as any;

    if (response.success) {
      bankingInfos.value = response.data.bankingInfos.map((item: any) => ({
        id: item.id,
        bankName: item.bank_name,
        cardNumber: item.card_number,
        accountNumber: item.account_number,
        shebaNumber: item.sheba_number,
        accountHolderName: item.account_holder_name,
        accountType: item.account_type,
        isDefault: item.is_default,
        isVerified: item.is_verified
      }));
    } else {
      error.value = 'خطا در دریافت داده‌ها';
    }
  } catch (err: any) {
    console.error('خطا در دریافت اطلاعات بانکی:', err);
    error.value = err.data?.message || 'خطا در دریافت اطلاعات بانکی';
  } finally {
    loading.value = false;
  }
};

const formatCardNumber = (cardNumber: string) => {
  if (!cardNumber) return '';
  return cardNumber.replace(/(\d{4})(?=\d)/g, '$1-');
};

const formatShebaNumber = (shebaNumber: string) => {
  if (!shebaNumber) return '';
  return shebaNumber.replace(/(.{4})/g, '$1 ').trim();
};

const verifyBankingInfo = async (bankingInfoId: number) => {
  const note = prompt('یادداشت تایید (اختیاری):');
  if (note === null) return; // کاربر کنسل کرد

  try {
    const response = await $fetch(`/api/admin/banking-info/${bankingInfoId}/verify`, {
      method: 'POST',
      body: { note: note || '' }
    }) as any;

    if (response.success) {
      alert('اطلاعات بانکی با موفقیت تایید شد');
      await fetchBankingInfo(); // رفرش لیست
    }
  } catch (err: any) {
    console.error('خطا در تایید اطلاعات بانکی:', err);
    alert(err.data?.message || 'خطا در تایید اطلاعات بانکی');
  }
};

const unverifyBankingInfo = async (bankingInfoId: number) => {
  const confirmed = confirm('آیا از لغو تایید این اطلاعات بانکی مطمئن هستید؟');
  if (!confirmed) return;

  const note = prompt('دلیل لغو تایید (اختیاری):');
  if (note === null) return; // کاربر کنسل کرد

  try {
    const response = await $fetch(`/api/admin/banking-info/${bankingInfoId}/unverify`, {
      method: 'POST',
      body: { note: note || '' }
    }) as any;

    if (response.success) {
      alert('تایید اطلاعات بانکی با موفقیت لغو شد');
      await fetchBankingInfo(); // رفرش لیست
    }
  } catch (err: any) {
    console.error('خطا در لغو تایید اطلاعات بانکی:', err);
    alert(err.data?.message || 'خطا در لغو تایید اطلاعات بانکی');
  }
};

onMounted(() => {
  fetchBankingInfo();
});
</script> 
