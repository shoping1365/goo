import { defineEventHandler } from 'h3'

export default defineEventHandler((event) => {
  const id = event.context.params?.id || '1'
  // داده تستی
  return {
    id,
    comment: 'نظر تستی',
    rating: 5,
    status: 'approved',
    createdAt: new Date().toISOString(),
    customer: {
      name: 'کاربر تست',
      email: 'test@example.com',
      joinDate: new Date().toISOString(),
      reviewCount: 1
    },
    product: {
      id: '100',
      name: 'محصول تستی',
      image: '/uploads/media/products/images/test.jpg',
      price: 100000,
      thumbnail: '/uploads/media/products/images/test_thumbnail.jpg'
    }
  }
}) 