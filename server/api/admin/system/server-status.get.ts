export default defineEventHandler(() => ({
  servers: [
    { id: 'nuxt', name: 'Nuxt Server', status: 'online' },
    { id: 'go', name: 'Go API', status: 'online' }
  ]
}))