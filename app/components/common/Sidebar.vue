<template>
  <div class="sidebar bg-white shadow-lg border-l border-gray-200 h-full">
    <!-- Header Bar -->
    <div class="header-bar">
      <div class="header-left">
        <div class="header-icons">
          <div class="icon-wrapper">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 19.128a9.38 9.38 0 002.625.372 9.337 9.337 0 004.121-.952 4.125 4.125 0 00-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 018.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0111.964-3.07M12 6.375a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zm8.25 2.25a2.625 2.625 0 11-5.25 0 2.625 2.625 0 015.25 0z" />
            </svg>
          </div>
          <div class="icon-wrapper">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
            </svg>
          </div>
        </div>
      </div>
      
      <div class="header-right">
        <div class="user-info">
          <div class="user-name">آ . آذری</div>
          <div class="user-status">
            <div class="online-dot"></div>
            آنلاین
          </div>
        </div>
        <div class="user-avatar">
          <div class="avatar-content">
            <div class="avatar-character">😊</div>
          </div>
        </div>
      </div>
      
      <div class="menu-icon">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z" />
        </svg>
      </div>
    </div>

    <!-- Search and Online Users Section -->
    <div class="search-section">
      <div class="online-users-info">
        <div class="online-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 17.25v1.007a3 3 0 01-.879 2.122L7.5 21h9l-.621-.621A3 3 0 0115 18.257V17.25m6-12V15a2.25 2.25 0 01-2.25 2.25H5.25A2.25 2.25 0 013 15V5.25m18 0A2.25 2.25 0 0020.75 3H3.75A2.25 2.25 0 001.5 5.25m18 0V12a2.25 2.25 0 01-2.25 2.25H5.25A2.25 2.25 0 013 12V5.25" />
          </svg>
        </div>
        <span class="online-text">کاربران آنلاین</span>
        <div class="online-count">۲۴۲</div>
      </div>
      
      <div class="search-bar">
        <input 
          v-model="searchQuery" 
          type="text"
          placeholder="جست و جو در بین مکالمات"
          class="search-input"
        />
        <div class="search-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
          </svg>
        </div>
      </div>
    </div>

    <!-- Chat Categories -->
    <div class="chat-categories">
      <div 
        class="category-item"
        :class="{ active: activeCategory === 'other' }"
        @click="setActiveCategory('other')"
      >
        <div class="category-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />
          </svg>
        </div>
        <span class="category-text">مکالمه های دیگر</span>
      </div>
      
      <div 
        class="category-item"
        :class="{ active: activeCategory === 'waiting' }"
        @click="setActiveCategory('waiting')"
      >
        <div class="category-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <span class="category-text">در انتظار پاسخ</span>
      </div>
      
      <div 
        class="category-item active"
        :class="{ active: activeCategory === 'conversation' }"
        @click="setActiveCategory('conversation')"
      >
        <div class="category-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
          </svg>
        </div>
        <span class="category-text">در حال مکالمه</span>
        <div class="category-count">۳۰</div>
      </div>
    </div>

    <!-- Toolbar -->
    <div class="toolbar">
      <div class="toolbar-left">
        <button class="toolbar-btn">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
        <span class="remaining-text">باقی مانده : ۱۵۰۰۰</span>
        <button class="toolbar-btn">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
          </svg>
        </button>
      </div>
      
      <div class="toolbar-right">
        <button class="toolbar-btn">
          <span>جدید ترین</span>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
          </svg>
        </button>
        <button class="toolbar-btn">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Chat List -->
    <div class="chat-list">
      <div 
        v-for="chat in filteredChats" 
        :key="chat.id"
        class="chat-item"
        :class="{ 'has-new-message': chat.hasNewMessage }"
        @click="selectChat(chat)"
      >
        <div class="chat-icon">
          <div class="chat-icon-inner">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12c0 4.556-4.03 8.25-9 8.25a9.764 9.764 0 01-2.555-.337A5.972 5.972 0 015.41 20.97a5.969 5.969 0 01-.474-.065 4.48 4.48 0 00.978-2.025c.09-.457-.133-.901-.467-1.226C3.93 16.178 3 14.189 3 12c0-4.556 4.03-8.25 9-8.25s9 3.694 9 8.25z" />
            </svg>
          </div>
          <div v-if="chat.hasNewMessage" class="new-message-indicator">۱</div>
        </div>
        
        <div class="chat-content">
          <div class="chat-header">
            <span class="user-label">کاربر :</span>
            <span class="user-id">{{ chat.userId }}</span>
          </div>
          <div class="chat-link">{{ chat.productLink }}</div>
          <div class="chat-message">{{ chat.message }}</div>
        </div>
      </div>
    </div>

    <!-- Load More Button -->
    <div class="load-more">
      <button class="load-more-btn">بیشتر</button>
    </div>
  </div>
</template>

<script setup>
const searchQuery = ref('')
const activeCategory = ref('conversation')

// Sample chat data
const chats = ref([
  {
    id: 1,
    userId: '۴۳۰۸۹۸',
    productLink: 'https://iranxia.net/product/xiao...',
    message: 'سلام قیمت سیم اصل ۶۷ وات چنده ....',
    hasNewMessage: false
  },
  {
    id: 2,
    userId: '۴۳۲۲۷۶',
    productLink: 'https://iranxia.net/product/airfryer...',
    message: 'لمسی هست یا دکمه ای ؟',
    hasNewMessage: true
  },
  {
    id: 3,
    userId: '۴۳۱۸۰۰',
    productLink: 'https://iranxia.net/cart',
    message: 'سلام',
    hasNewMessage: false
  },
  {
    id: 4,
    userId: '۴۳۱۵۴۱',
    productLink: 'https://iranxia.net/product/roborock...',
    message: 'سلام برای دریافت مشاوره تخصصی ....',
    hasNewMessage: false
  },
  {
    id: 5,
    userId: '۴۳۱۱۳۸',
    productLink: 'https://iranxia.net/product/Mi-۵۰۰۰۰...',
    message: 'سلام احتمالا اورجینال نبوده',
    hasNewMessage: false
  },
  {
    id: 6,
    userId: '۴۳۱۵۸۲',
    productLink: 'https://iranxia.net/product/xiaomi...',
    message: 'قیمت دوچرخه چنده',
    hasNewMessage: false
  }
])

const setActiveCategory = (category) => {
  activeCategory.value = category
}

const selectChat = (chat) => {
  // Emit the chat-selected event to parent component
  emit('chatSelected', chat)
}

const filteredChats = computed(() => {
  if (!searchQuery.value) return chats.value
  
  return chats.value.filter(chat => 
    chat.message.includes(searchQuery.value) ||
    chat.userId.includes(searchQuery.value) ||
    chat.productLink.includes(searchQuery.value)
  )
})

// Emit events for parent component
const emit = defineEmits(['chatSelected', 'categoryChanged', 'searchChanged'])

watch(activeCategory, (newCategory) => {
  emit('categoryChanged', newCategory)
})

watch(searchQuery, (newQuery) => {
  emit('searchChanged', newQuery)
})
</script>

<style scoped>
.sidebar {
  width: 350px;
  background: white;
  border-right: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: hidden;
}

/* Header Bar */
.header-bar {
  background: #8b5cf6;
  padding: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white;
  position: relative;
}

.header-left {
  display: flex;
  align-items: center;
}

.header-icons {
  display: flex;
  gap: 12px;
}

.icon-wrapper {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-info {
  text-align: right;
}

.user-name {
  font-size: 14px;
  font-weight: 600;
  line-height: 1.2;
}

.user-status {
  font-size: 12px;
  opacity: 0.9;
  line-height: 1.2;
  display: flex;
  align-items: center;
  gap: 4px;
}

.online-dot {
  width: 8px;
  height: 8px;
  background: #10b981;
  border-radius: 50%;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.avatar-content {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.avatar-character {
  font-size: 18px;
}

.menu-icon {
  position: absolute;
  top: 16px;
  left: 16px;
  color: white;
  cursor: pointer;
}

/* Search Section */
.search-section {
  padding: 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  border-bottom: 1px solid #f3f4f6;
}

.online-users-info {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #374151;
}

.online-icon {
  color: #6b7280;
}

.online-text {
  font-size: 12px;
  color: #374151;
}

.online-count {
  background: #374151;
  color: white;
  border-radius: 12px;
  padding: 4px 8px;
  font-size: 11px;
  font-weight: 600;
  min-width: 32px;
  text-align: center;
}

.search-bar {
  position: relative;
  flex: 1;
  max-width: 200px;
  margin-right: 12px;
}

.search-input {
  width: 100%;
  padding: 8px 32px 8px 12px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 12px;
  background: #f9fafb;
  color: #374151;
}

.search-input::placeholder {
  color: #9ca3af;
}

.search-icon {
  position: absolute;
  top: 50%;
  right: 8px;
  transform: translateY(-50%);
  color: #9ca3af;
}

/* Chat Categories */
.chat-categories {
  display: flex;
  background: white;
  border-bottom: 1px solid #e5e7eb;
}

.category-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  border-bottom: 2px solid transparent;
}

.category-item:hover {
  background: #f9fafb;
}

.category-item.active {
  border-bottom-color: #8b5cf6;
  background: #f3f4f6;
}

.category-icon {
  color: #6b7280;
  margin-bottom: 4px;
}

.category-item.active .category-icon {
  color: #8b5cf6;
}

.category-text {
  font-size: 11px;
  color: #374151;
  text-align: center;
  line-height: 1.2;
}

.category-item.active .category-text {
  color: #8b5cf6;
  font-weight: 500;
}

.category-count {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #ef4444;
  color: white;
  border-radius: 50%;
  width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: 500;
}

/* Toolbar */
.toolbar {
  padding: 12px 16px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
  border-bottom: 1px solid #f3f4f6;
}

.toolbar-left, .toolbar-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.toolbar-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 6px 8px;
  border: none;
  background: transparent;
  color: #6b7280;
  font-size: 12px;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.toolbar-btn:hover {
  background: #f3f4f6;
  color: #374151;
}

.remaining-text {
  font-size: 12px;
  color: #6b7280;
}

/* Chat List */
.chat-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.chat-item {
  display: flex;
  gap: 12px;
  padding: 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  margin-bottom: 8px;
  background: white;
  border: 1px solid #f3f4f6;
}

.chat-item:hover {
  background: #f9fafb;
  border-color: #e5e7eb;
}

.chat-item.has-new-message {
  background: #fef3c7;
  border-color: #fbbf24;
}

.chat-icon {
  position: relative;
  flex-shrink: 0;
}

.chat-icon-inner {
  width: 40px;
  height: 40px;
  background: #8b5cf6;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.new-message-indicator {
  position: absolute;
  top: -4px;
  right: -4px;
  width: 16px;
  height: 16px;
  background: #374151;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: 600;
}

.chat-content {
  flex: 1;
  min-width: 0;
}

.chat-header {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-bottom: 4px;
}

.user-label {
  font-size: 12px;
  color: #6b7280;
}

.user-id {
  font-size: 12px;
  color: #374151;
  font-weight: 500;
}

.chat-link {
  font-size: 11px;
  color: #8b5cf6;
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.chat-message {
  font-size: 12px;
  color: #374151;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* Load More */
.load-more {
  padding: 16px;
  border-top: 1px solid #f3f4f6;
}

.load-more-btn {
  width: 100%;
  padding: 10px;
  background: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  color: #374151;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.load-more-btn:hover {
  background: #e5e7eb;
}
</style> 