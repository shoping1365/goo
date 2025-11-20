<template>
  <div class="flex flex-row items-end h-36 gap-6">
    <div v-for="(r,idx) in rankCounts" :key="idx" class="flex flex-col items-center w-20">
      <div
:class="[r.color, 'w-10 rounded-t', 'transition-all']"
        :style="{height:(r.value*20)+'px'}"></div>
      <span class="mt-2 text-xs font-bold">{{ r.name }}</span>
      <span class="text-[10px]">({{ r.value }})</span>
    </div>
  </div>
</template>
<script setup lang="ts">
interface RankCount {
  name: string;
  value: number;
  color: string;
}

defineProps<{rankCounts: RankCount[], ranks?: unknown}>()
</script>
