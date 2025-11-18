export default defineEventHandler((event) => {
  const rows = [
    ['id','name','email','role'],
    ['1','Ali','ali@example.com','user'],
    ['2','Sara','sara@example.com','admin']
  ]
  const csv = rows.map(r=>r.join(',')).join('\n')
  setResponseHeader(event, 'Content-Type', 'text/csv; charset=utf-8')
  return `\uFEFF${csv}`
})