<template>
  <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
    <div class="w-full max-w-3xl bg-white rounded-2xl shadow-2xl px-4 py-4 overflow-y-auto max-h-[90vh] relative animate-fadein">
      <button class="absolute left-4 topx-4 py-4 btn btn-sm btn-error z-10" @click="$emit('close')">بستن</button>
      <div v-if="user">
        <div class="flex flex-col md:flex-row gapx-4 py-4 mb-6">
          <div class="flex-1">
            <UserProfileTab :user="user" />
          </div>
          <div class="flex-1 flex flex-col gapx-4 py-4">
            <UserAdminActionsTab :user="user" @action="handleAction" />
          </div>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gapx-4 py-4">
          <UserActivityTab :user="user" />
          <UserOrdersTab :user="user" />
          <UserBehaviorTab :user="user" />
        </div>
      </div>
      <div v-else>
        <span>کاربری انتخاب نشده است.</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ user: any, open: boolean }>();
// Tabs are removed for now, but keep imports for future use
import UserProfileTab from './UserProfileTab.vue';
import UserActivityTab from './UserActivityTab.vue';
import UserOrdersTab from './UserOrdersTab.vue';
import UserBehaviorTab from './UserBehaviorTab.vue';
import UserAdminActionsTab from './UserAdminActionsTab.vue';

function handleAction(action: string) {
  // Placeholder for admin actions (block, reset password, etc.)
}
</script>
<style scoped>
.animate-fadein {
  animation: fadein 0.3s;
}
@keyframes fadein {
  from { opacity: 0; transform: translateY(40px); }
  to { opacity: 1; transform: none; }
}
</style> 
