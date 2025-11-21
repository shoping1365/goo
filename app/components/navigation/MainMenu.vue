<template>
  <nav class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex">
          <div class="flex-shrink-0 flex items-center">
            <!-- لوگو -->
            <div class="w-10 h-10 bg-indigo-600 rounded-full flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
            </div>
          </div>
          <div class="hidden sm:mr-6 sm:flex sm:space-x-8 sm:space-x-reverse">
            <NuxtLink
              to="/"
              class="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
            >
              صفحه اصلی
            </NuxtLink>
            <NuxtLink
              to="/settings/security"
              class="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
            >
              تنظیمات امنیت
            </NuxtLink>
          </div>
        </div>
        <div class="hidden sm:mr-6 sm:flex sm:items-center">
          <button
            class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            @click="logout"
          >
            خروج
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
// احراز هویت - غیرفعال شده
const logout = async () => {

}
</script> 

