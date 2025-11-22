<template>
  <button :class="buttonClass" @click="exportToCSV">
    <slot>
      <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
      </svg>
      خروجی اکسل
    </slot>
  </button>
</template>

<script setup lang="ts">
interface Props {
  data?: Record<string, unknown>[]
  filename?: string
  buttonClass?: string
}

const props = withDefaults(defineProps<Props>(), {
  data: () => [],
  filename: 'data.csv',
  buttonClass: 'inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-emerald-400 to-green-600 hover:from-emerald-500 hover:to-green-700 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105'
})

// تابع ساده برای تبدیل آرایه به CSV
function arrayToCSV(data: Record<string, unknown>[]): string {
  if (!data || !data.length) return '';
  
  // دریافت کلیدهای ستون‌ها
  const headers = Object.keys(data[0]);
  
  // ایجاد header row
  const headerRow = headers.map(header => `"${header}"`).join(',');
  
  // ایجاد data rows
  const dataRows = (data as Record<string, unknown>[]).map(row =>
    headers.map(header => {
      const value = row[header];
      // اگر مقدار شامل کاما، نقل قول یا خط جدید باشد، آن را در نقل قول قرار بده
      if (typeof value === 'string' && (value.includes(',') || value.includes('"') || value.includes('\n'))) {
        return `"${value.replace(/"/g, '""')}"`;
      }
      return `"${value || ''}"`;
    }).join(',')
  );
  
  return [headerRow, ...dataRows].join('\n');
}

function exportToCSV() {
  if (!props.data || !props.data.length) return;
  
  const csv = '\uFEFF' + arrayToCSV(props.data as Record<string, unknown>[]); // BOM برای اکسل و فارسی
  const blob = new Blob([csv], { type: 'text/csv;charset=utf-8;' });
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.setAttribute('download', props.filename);
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
}
</script> 
