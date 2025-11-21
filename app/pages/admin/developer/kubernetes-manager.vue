<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">مدیریت کلاستر Kubernetes</h1>
        <p class="text-gray-600">ابزارهای مدیریت و مانیتورینگ کلاستر Kubernetes</p>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Node ها</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.nodes }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Pod های فعال</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.activePods }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">استفاده از CPU</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.cpuUsage }}%</p>
            </div>
          </div>
        </div>

        <div class="bg-white rounded-lg shadow p-6">
          <div class="flex items-center">
            <div class="p-2 bg-yellow-100 rounded-lg">
              <svg class="w-6 h-6 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">استفاده از حافظه</p>
              <p class="text-2xl font-semibold text-gray-900">{{ stats.memoryUsage }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Cluster Overview -->
      <div class="mb-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">نمای کلی کلاستر</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">وضعیت Node ها</h3>
              <div class="space-y-2">
                <div v-for="node in clusterNodes" :key="node.name" class="flex items-center justify-between p-3 border rounded-lg">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ node.name }}</span>
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      node.status === 'Ready' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                    ]">
                      {{ node.status }}
                    </span>
                  </div>
                  <div class="text-sm text-gray-500">
                    {{ node.cpu }} CPU / {{ node.memory }}
                  </div>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">Namespace ها</h3>
              <div class="space-y-2">
                <div v-for="namespace in namespaces" :key="namespace.name" class="flex items-center justify-between p-3 border rounded-lg">
                  <span class="font-medium text-gray-900">{{ namespace.name }}</span>
                  <span class="text-sm text-gray-500">{{ namespace.pods }} Pod</span>
                </div>
              </div>
            </div>

            <div class="space-y-4">
              <h3 class="font-medium text-gray-900">سرویس‌ها</h3>
              <div class="space-y-2">
                <div v-for="service in services" :key="service.name" class="flex items-center justify-between p-3 border rounded-lg">
                  <div>
                    <span class="font-medium text-gray-900">{{ service.name }}</span>
                    <p class="text-sm text-gray-500">{{ service.type }}</p>
                  </div>
                  <span class="text-sm text-gray-500">{{ service.ports }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Pod Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h2 class="text-xl font-semibold text-gray-900">مدیریت Pod ها</h2>
              <button 
                class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="refreshPods"
              >
                بروزرسانی
              </button>
            </div>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="pod in pods" :key="pod.name" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ pod.name }}</span>
                    <span
:class="[
                      'px-2 py-1 rounded text-xs font-medium',
                      pod.status === 'Running' ? 'bg-green-100 text-green-800' :
                      pod.status === 'Pending' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-red-100 text-red-800'
                    ]">
                      {{ pod.status }}
                    </span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="viewPodLogs(pod)"
                    >
                      لاگ‌ها
                    </button>
                    <button 
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="deletePod(pod)"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500 space-y-1">
                  <p>Namespace: {{ pod.namespace }}</p>
                  <p>Node: {{ pod.node }}</p>
                  <p>IP: {{ pod.ip }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Deployment Management -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">مدیریت Deployment ها</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="deployment in deployments" :key="deployment.name" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ deployment.name }}</span>
                    <span class="text-sm text-gray-500">{{ deployment.replicas }} replicas</span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-green-600 hover:text-green-800 text-sm font-medium"
                      @click="scaleDeployment(deployment)"
                    >
                      مقیاس
                    </button>
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="rolloutDeployment(deployment)"
                    >
                      Rollout
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500 space-y-1">
                  <p>Image: {{ deployment.image }}</p>
                  <p>Available: {{ deployment.available }}/{{ deployment.replicas }}</p>
                  <p>Updated: {{ deployment.updated }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- ConfigMap & Secrets -->
        <div class="bg-white rounded-lg shadow">
          <div class="p-6 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-900">ConfigMap & Secrets</h2>
          </div>
          <div class="p-6">
            <div class="space-y-4">
              <div v-for="config in configMaps" :key="config.name" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ config.name }}</span>
                    <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">ConfigMap</span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="editConfigMap(config)"
                    >
                      ویرایش
                    </button>
                    <button 
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="deleteConfigMap(config)"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500">
                  <p>{{ config.dataCount }} data entries</p>
                </div>
              </div>

              <div v-for="secret in secrets" :key="secret.name" class="border rounded-lg p-6">
                <div class="flex items-center justify-between mb-2">
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <span class="font-medium text-gray-900">{{ secret.name }}</span>
                    <span class="text-xs bg-red-100 text-red-800 px-2 py-1 rounded">Secret</span>
                  </div>
                  <div class="flex items-center space-x-2 space-x-reverse">
                    <button 
                      class="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      @click="editSecret(secret)"
                    >
                      ویرایش
                    </button>
                    <button 
                      class="text-red-600 hover:text-red-800 text-sm font-medium"
                      @click="deleteSecret(secret)"
                    >
                      حذف
                    </button>
                  </div>
                </div>
                <div class="text-sm text-gray-500">
                  <p>{{ secret.dataCount }} secret entries</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Resource Monitoring -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-900">مانیتورینگ منابع</h2>
        </div>
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از CPU</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-blue-600 h-2 rounded-full" :style="{ width: stats.cpuUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.cpuUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از حافظه</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-green-600 h-2 rounded-full" :style="{ width: stats.memoryUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.memoryUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">استفاده از دیسک</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-yellow-600 h-2 rounded-full" :style="{ width: stats.diskUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.diskUsage }}%</p>
            </div>
            <div class="space-y-2">
              <h3 class="font-medium text-gray-900">شبکه</h3>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div class="bg-purple-600 h-2 rounded-full" :style="{ width: stats.networkUsage + '%' }"></div>
              </div>
              <p class="text-sm text-gray-500">{{ stats.networkUsage }}%</p>
            </div>
          </div>
        </div>
      </div>

      <!-- YAML Editor -->
      <div class="mt-8 bg-white rounded-lg shadow">
        <div class="p-6 border-b border-gray-200">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-semibold text-gray-900">ویرایشگر YAML</h2>
            <div class="flex items-center space-x-2 space-x-reverse">
              <button 
                class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="applyYaml"
              >
                اعمال
              </button>
              <button 
                class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                @click="validateYaml"
              >
                اعتبارسنجی
              </button>
            </div>
          </div>
        </div>
        <div class="p-6">
          <textarea
            v-model="yamlContent"
            rows="15"
            class="w-full border border-gray-300 rounded-lg px-3 py-2 font-mono text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="apiVersion: v1&#10;kind: Pod&#10;metadata:&#10;  name: my-pod&#10;spec:&#10;  containers:&#10;  - name: my-container&#10;    image: nginx:latest"
          ></textarea>
        </div>
      </div>

      <!-- Pod Logs Modal -->
      <div v-if="logsModal.show" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white rounded-lg max-w-4xl max-h-4xl overflow-auto">
          <div class="p-6 border-b border-gray-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-semibold text-gray-900">لاگ‌های {{ logsModal.pod?.name }}</h3>
              <button class="text-gray-400 hover:text-gray-600" @click="logsModal.show = false">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>
          </div>
          <div class="p-6">
            <pre class="bg-gray-900 text-gray-100 p-6 rounded-lg overflow-x-auto text-sm">{{ logsModal.content }}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';

definePageMeta({
  layout: 'admin-main',
  middleware: ['developer-only']
});

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()

// Reactive data
const yamlContent = ref(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80`)

const logsModal = reactive({
  show: false,
  pod: null,
  content: ''
})

const stats = reactive({
  nodes: 5,
  activePods: 24,
  cpuUsage: 65,
  memoryUsage: 78,
  diskUsage: 45,
  networkUsage: 32
})

const clusterNodes = ref([
  { name: 'node-1', status: 'Ready', cpu: '4', memory: '8Gi' },
  { name: 'node-2', status: 'Ready', cpu: '4', memory: '8Gi' },
  { name: 'node-3', status: 'Ready', cpu: '2', memory: '4Gi' },
  { name: 'node-4', status: 'NotReady', cpu: '4', memory: '8Gi' }
])

const namespaces = ref([
  { name: 'default', pods: 8 },
  { name: 'kube-system', pods: 12 },
  { name: 'monitoring', pods: 4 },
  { name: 'ingress-nginx', pods: 2 }
])

const services = ref([
  { name: 'kubernetes', type: 'ClusterIP', ports: '443/TCP' },
  { name: 'nginx-service', type: 'LoadBalancer', ports: '80:30000/TCP' },
  { name: 'redis-service', type: 'ClusterIP', ports: '6379/TCP' }
])

const pods = ref([
  {
    name: 'nginx-deployment-6b474476c4-abc12',
    status: 'Running',
    namespace: 'default',
    node: 'node-1',
    ip: '10.244.0.5'
  },
  {
    name: 'redis-deployment-8f5d4c2a1b-def34',
    status: 'Running',
    namespace: 'default',
    node: 'node-2',
    ip: '10.244.1.3'
  },
  {
    name: 'app-deployment-7e3c1b9a8f-ghi56',
    status: 'Pending',
    namespace: 'default',
    node: 'node-3',
    ip: '10.244.2.1'
  }
])

const deployments = ref([
  {
    name: 'nginx-deployment',
    replicas: 3,
    image: 'nginx:latest',
    available: 3,
    updated: 3
  },
  {
    name: 'redis-deployment',
    replicas: 1,
    image: 'redis:alpine',
    available: 1,
    updated: 1
  },
  {
    name: 'app-deployment',
    replicas: 2,
    image: 'myapp:v1.0.0',
    available: 1,
    updated: 2
  }
])

const configMaps = ref([
  { name: 'app-config', dataCount: 5 },
  { name: 'nginx-config', dataCount: 3 },
  { name: 'redis-config', dataCount: 2 }
])

const secrets = ref([
  { name: 'db-credentials', dataCount: 2 },
  { name: 'api-keys', dataCount: 4 },
  { name: 'tls-cert', dataCount: 1 }
])

// Methods
function refreshPods() {
  alert('Pod ها بروزرسانی شدند')
}

function viewPodLogs(pod) {
  logsModal.pod = pod
  logsModal.content = `[2024-01-15 10:30:15] INFO: Pod started
[2024-01-15 10:30:16] INFO: Container initialized
[2024-01-15 10:30:17] INFO: Application ready
[2024-01-15 10:30:18] INFO: Health check passed
[2024-01-15 10:30:19] INFO: Service listening on port 80`
  logsModal.show = true
}

function deletePod(pod) {
  if (confirm(`آیا می‌خواهید Pod ${pod.name} را حذف کنید؟`)) {
    alert(`Pod ${pod.name} حذف شد`)
  }
}

function scaleDeployment(deployment) {
  const replicas = prompt(`تعداد replicas برای ${deployment.name} را وارد کنید:`, deployment.replicas)
  if (replicas) {
    alert(`Deployment ${deployment.name} به ${replicas} replicas مقیاس شد`)
  }
}

function rolloutDeployment(deployment) {
  alert(`Rollout برای ${deployment.name} شروع شد`)
}

function editConfigMap(config) {
  alert(`ویرایش ConfigMap ${config.name}`)
}

function deleteConfigMap(config) {
  if (confirm(`آیا می‌خواهید ConfigMap ${config.name} را حذف کنید؟`)) {
    alert(`ConfigMap ${config.name} حذف شد`)
  }
}

function editSecret(secret) {
  alert(`ویرایش Secret ${secret.name}`)
}

function deleteSecret(secret) {
  if (confirm(`آیا می‌خواهید Secret ${secret.name} را حذف کنید؟`)) {
    alert(`Secret ${secret.name} حذف شد`)
  }
}

function applyYaml() {
  if (yamlContent.value.trim()) {
    alert('YAML اعمال شد')
  }
}

function validateYaml() {
  if (yamlContent.value.trim()) {
    alert('YAML معتبر است')
  } else {
    alert('YAML نامعتبر است')
  }
}
</script> 
