import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const q = getQuery(event)
  const url = new URL(`${config.public.goApiBase}/api/admin/fraud/cases`)
  if (q.risk) url.searchParams.set('risk', String(q.risk))
  if (q.status) url.searchParams.set('status', String(q.status))
  if (q.page) url.searchParams.set('page', String(q.page))
  if (q.pageSize) url.searchParams.set('pageSize', String(q.pageSize))
  if (q.search) url.searchParams.set('search', String(q.search))
  if (q.paymentMethod) url.searchParams.set('paymentMethod', String(q.paymentMethod))
  if (q.minAmount) url.searchParams.set('minAmount', String(q.minAmount))
  if (q.maxAmount) url.searchParams.set('maxAmount', String(q.maxAmount))
  if (q.from) url.searchParams.set('from', String(q.from))
  if (q.to) url.searchParams.set('to', String(q.to))

  const res = await $fetch(url.toString())
  return res
})


