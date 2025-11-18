<template>
  <div class="min-h-screen bg-gray-50 p-8">
    <div class="max-w-4xl mx-auto">
      <h1 class="text-3xl font-bold text-gray-900 mb-8">تست باکس جستجو</h1>
      
      <div class="bg-white rounded-lg shadow-lg p-8">
        <SearchBox 
          placeholder="جستجو کنید..."
          search-type="all"
          @search="handleSearch"
        />
      </div>

      <div v-if="lastSearch" class="mt-8 p-4 bg-blue-50 rounded-lg">
        <p class="text-sm text-gray-600">آخرین جستجو:</p>
        <p class="text-lg font-semibold text-gray-900">{{ lastSearch }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import SearchBox from '~/components/common/SearchBox.vue'

const lastSearch = ref('')

const handleSearch = (query: string) => {
  console.log('جستجو شد:', query)
  lastSearch.value = query
}
</script>
