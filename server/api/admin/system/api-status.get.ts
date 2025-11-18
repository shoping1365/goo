export default defineEventHandler(() => ({
  connections: [
    { id: 'go', name: 'Go Backend', status: 'online' },
    { id: 'postgres', name: 'PostgreSQL', status: 'online' }
  ]
}))