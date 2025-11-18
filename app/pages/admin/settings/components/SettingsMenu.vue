<template>
  <div class="settings-menu">
    <div v-for="item in menu" :key="item.key" class="w-full">
      <!-- Parent item (always has arrow) -->
      <div
        :class="['settings-menu-item', { active: activeKey.startsWith(item.key) }]"
        @click="handleMenuClick(item)"
      >
        <span :class="item.icon + ' ml-2'" v-if="item.icon"></span>
        {{ item.title }}
        <span class="mr-auto text-xs">
          {{ expanded[item.key] ? '▾' : '▸' }}
        </span>
      </div>

      <!-- Children (render if any) -->
      <transition name="fade">
        <div v-if="expanded[item.key] && item.children.length" class="pr-4 border-r border-gray-600">
          <div v-for="child in item.children" :key="child.key"
            :class="['settings-menu-item child', { active: activeKey === child.key }]"
            @click.stop="activeKey = child.key"
          >
            <span :class="child.icon + ' ml-2'" v-if="child.icon"></span>
            {{ child.title }}
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  menu: {
    type: Array,
    required: true
  },
  activeKey: {
    type: String,
    required: true
  },
  expanded: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update:activeKey', 'menuClick'])

function handleMenuClick(item) {
  if (item.children && item.children.length > 0) {
    // اگر آیتم فرزند دارد، فقط expand/collapse کن
    toggleGroup(item.key)
  } else {
    // اگر آیتم فرزند ندارد، activeKey را تغییر بده
    emit('update:activeKey', item.key)
    emit('menuClick', item)
  }
}

function toggleGroup(k) {
  emit('menuClick', { key: k, action: 'toggle' })
}
</script>

<style scoped>
.settings-menu {
  background: #222d32;
  padding: 16px 0 16px 16px;
  border-radius: 0 8px 8px 0;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 180px;
  max-width: 200px;
  width: 200px;
  margin-left: 0;
  margin-right: 0;
  height: fit-content;
  text-align: right;
  box-shadow: 0 2px 8px #0001;
  position: absolute;
  right: 0;
  top: 0;
  z-index: 2;
}
.settings-menu-item {
  color: #fff;
  padding: 8px 0;
  border-radius: 6px;
  font-size: 1em;
  display: flex;
  flex-direction: row;
  align-items: center;
  cursor: pointer;
  transition: background 0.2s;
  width: 100%;
  justify-content: flex-end;
  margin-bottom: 2px;
  text-align: right;
}
.settings-menu-item.active {
  background: #e74c3c;
  color: #fff;
}
.settings-menu-item:hover {
  background: #2d3842;
}
.settings-menu-item .ml-2 {
  margin-left: 8px;
  margin-right: 0;
  font-size: 1.1em;
}
.settings-menu-item.child {
  padding-right: 20px;
  font-size: 0.9em;
}
</style>
