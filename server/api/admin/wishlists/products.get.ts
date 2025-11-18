import type { H3Event } from 'h3'
import { defineEventHandler, getQuery } from 'h3'

export default defineEventHandler((event: H3Event) => {
  const query = getQuery(event)
  const page = Math.max(1, Number(query.page || 1))
  const pageSize = Math.min(100, Math.max(1, Number(query.pageSize || 25)))
  const sort = String(query.sort || 'wishlist_count')
  const order = String(query.order || 'desc').toLowerCase() === 'asc' ? 'asc' : 'desc'
  const q = String(query.q || '').toLowerCase()
  const category = String(query.category || '')

  // Mock data source (aggregate by product)
  const allItems = [
    {
      product: {
        id: 101,
        name: 'گوشی سامسونگ گلکسی A54',
        sku: 'SAM-A54-128',
        price: 8500000,
        discount: 15,
        category: 'الکترونیکی',
        image: 'https://via.placeholder.com/100x100?text=Phone',
        in_stock: true
      },
      users: [
        { id: 1, user: { id: 1001, name: 'علی احمدی', email: 'ali@example.com' }, created_at: '2024-01-15T10:30:00Z' },
        { id: 2, user: { id: 1006, name: 'فاطمه رضایی', email: 'fatemeh@example.com' }, created_at: '2024-01-16T14:20:00Z' }
      ],
      latestDate: '2024-01-16T14:20:00Z'
    },
    {
      product: {
        id: 102,
        name: 'لپ‌تاپ ایسوس VivoBook',
        sku: 'ASUS-VB15-512',
        price: 25000000,
        category: 'الکترونیکی',
        image: 'https://via.placeholder.com/100x100?text=Laptop',
        in_stock: false
      },
      users: [
        { id: 3, user: { id: 1002, name: 'مریم محمدی', email: 'maryam@example.com' }, created_at: '2024-01-14T14:20:00Z' }
      ],
      latestDate: '2024-01-14T14:20:00Z'
    }
  ]

  // Filters
  let filtered = allItems
  if (q) {
    filtered = filtered.filter(i =>
      i.product.name.toLowerCase().includes(q) ||
      i.product.sku.toLowerCase().includes(q)
    )
  }
  if (category) {
    filtered = filtered.filter(i => i.product.category === category)
  }

  // Sort by requested column
  interface WishlistItem {
    product: {
      name: string
      price?: number
      in_stock?: boolean
    }
    users?: unknown[]
    latestDate: string
  }
  const sorters: Record<string, (a: WishlistItem, b: WishlistItem) => number> = {
    name: (a, b) => a.product.name.localeCompare(b.product.name, 'fa'),
    price: (a, b) => (a.product.price || 0) - (b.product.price || 0),
    wishlist_count: (a, b) => (a.users?.length || 0) - (b.users?.length || 0),
    latest_date: (a, b) => new Date(a.latestDate).getTime() - new Date(b.latestDate).getTime(),
    stock_status: (a, b) => (a.product.in_stock ? 1 : 0) - (b.product.in_stock ? 1 : 0)
  }
  if (sorters[sort]) {
    filtered = filtered.sort((a, b) => sorters[sort](a, b))
    if (order === 'desc') filtered.reverse()
  }

  const total = filtered.length
  const start = (page - 1) * pageSize
  const end = start + pageSize
  const items = filtered.slice(start, end)

  return {
    items,
    page,
    pageSize,
    total,
    totalPages: Math.ceil(total / pageSize)
  }
})


