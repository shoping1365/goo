 <template>
  <div class="bg-white rounded-xl shadow-lg p-6">
    <h3 class="text-lg font-bold text-gray-900 mb-4">صفحات پربازدید</h3>
    
    <div class="space-y-3">
      <div v-for="(page, index) in pages" :key="page.page" class="flex items-center justify-between p-3 rounded-lg hover:bg-gray-50 transition-colors duration-200">
        <div class="flex items-center gap-3">
          <div
class="flex items-center justify-center w-8 h-8 rounded-full text-sm font-bold"
               :class="getRankClass(index)">
            {{ index + 1 }}
          </div>
          <div>
            <div class="font-semibold text-gray-900">{{ page.page }}</div>
            <div class="text-sm text-gray-500">{{ page.visits.toLocaleString('fa-IR') }} بازدید</div>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="w-20">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div 
                class="h-2 rounded-full transition-all duration-500 bg-gradient-to-r from-blue-400 to-blue-600"
                :style="{ width: page.percentage + '%' }"
              ></div>
            </div>
          </div>
          <div class="text-sm font-semibold text-gray-700 w-12 text-left">
            {{ page.percentage }}%
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Page {
  page: string
  visits: number
  percentage: number
}

defineProps<{
  pages: Page[]
}>()

const getRankClass = (index: number) => {
  if (index === 0) return 'bg-gradient-to-br from-yellow-400 to-yellow-500 text-white'
  if (index === 1) return 'bg-gradient-to-br from-gray-400 to-gray-500 text-white'
  if (index === 2) return 'bg-gradient-to-br from-orange-400 to-orange-500 text-white'
  return 'bg-gray-100 text-gray-600'
}
</script> 
