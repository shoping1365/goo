import { defineEventHandler, readBody, createError } from 'h3'
import { getHeader } from '../_utils/getHeader'
import { getCookieValue } from '../_utils/cookieHelper'

// API ���� ��� ����� ����
export default defineEventHandler(async (event) => {
  try {
    // ������ ������� �������
    const body = await readBody(event)
    
    // ����� ����� ���� �����
    const userId = event.context?.user?.id
    if (!userId) {
      throw createError({
        statusCode: 401,
        message: '����� ����� ���� ���� ���'
      })
    }

    // ������ ����� �������
    const { getDatabase } = await import('../_utils/database')

    // ���������� ������� �����
    const { 
      notes, 
      paymentMethod,
      items,
      total
    } = body

    console.log('������� �������:', { notes, paymentMethod, items, total })

    if (!items || !Array.isArray(items) || items.length === 0) {
      throw createError({
        statusCode: 400,
        message: '��� ���� ���� ���'
      })
    }

    // ����� ������ ������� � ������ total
    const db = await getDatabase()
    let calculatedTotal = 0
    
    for (const item of items) {
      const product = await db.query(
        'SELECT id, name, price, stock_quantity FROM products WHERE id = $1',
        [item.product_id]
      )
      
      if (!product || product.length === 0) {
        throw createError({
          statusCode: 400,
          message: `����� ${item.product?.name || '������'} ���� ���`
        })
      }
      
      if ((product[0] as any).stock_quantity < item.quantity) {
        throw createError({
          statusCode: 400,
          message: `������ ����� ${(product[0] as any).name} ���� ����`
        })
      }
      
      // ������ total
      calculatedTotal += (product[0] as any).price * item.quantity
    }
    
    // ������� �� total ������ ��� �� ��� total ������ �� ����ʝ���
    const finalTotal = calculatedTotal

    // ������ ���� �� profile_data �����
    console.log('������ ���� �� profile_data ���� �����:', userId)
    const userResult = await db.query(
      'SELECT profile_data FROM users WHERE id = $1',
      [userId]
    )
    
    console.log('����� ������ �����:', userResult)
    
    if (!userResult || userResult.length === 0) {
      throw createError({
        statusCode: 400,
        message: '����� ���� ���'
      })
    }
    
    const userData = userResult[0] as any
    const profileData = userData.profile_data
    
    if (!profileData || !profileData.selected_address) {
      throw createError({
        statusCode: 400,
        message: '���� ������ ��� ���� ���. ����� ����� ���� ��� �� ��� ����'
      })
    }
    
    const address = {
      id: 1, // ���� ����� �� profile_data
      full_address: profileData.selected_address,
      recipient_name: profileData.first_name + ' ' + profileData.last_name,
      recipient_mobile: event.context?.user?.mobile || '',
      postal_code: profileData.postal_code || ''
    }

    // ����� ����� ����� ����� �� ���
    const random = Math.floor(Math.random() * 10000).toString().padStart(4, '0')
    const orderNumber = `ORD-${random}`

    // ������ IP � User Agent
    const clientIP = getClientIP(event) || '������'
    const userAgent = getHeader(event, 'user-agent') || '������'

    // ����� ����� �� �������
    const orderResult = await db.query(
      `INSERT INTO orders (
        user_id, 
        shipping_address_id, 
        payment_method, 
        total_amount, 
        final_amount,
        status, 
        payment_status,
        notes,
        order_number,
        created_at,
        customer_ip,
        user_agent
      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`,
      [
        userId,
        address.id,
        paymentMethod,
        finalTotal,
        finalTotal, // final_amount
        'pending',
        'pending',
        notes || null,
        orderNumber,
        new Date().toISOString(), // created_at
        clientIP, // customer_ip
        userAgent // user_agent
      ]
    )

    const orderId = (orderResult[0] as any).id

    // ������ ������� �����
    for (const item of items) {
      // ������ ������� ���� �����
      const productResult = await db.query(
        'SELECT id, name, image, sku, price FROM products WHERE id = $1',
        [item.product_id]
      )
      
      const product = productResult[0] as any
      
      await db.query(
        `INSERT INTO order_items (
          order_id, 
          product_id, 
          product_name,
          product_image,
          product_sku,
          quantity, 
          unit_price, 
          total_price,
          final_price
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
        [
          orderId,
          item.product_id,
          product.name,
          product.image,
          product.sku,
          item.quantity,
          product.price,
          product.price * item.quantity,
          product.price * item.quantity
        ]
      )

      // ���� ������ �����
      await db.query(
        'UPDATE products SET stock_quantity = stock_quantity - $1 WHERE id = $2',
        [item.quantity, item.product_id]
      )
    }

    // �ǘ ���� ��� ���� ����� (�� �� ������� � �� �� session)
    try {
      await db.query('DELETE FROM cart_items WHERE cart_id IN (SELECT id FROM carts WHERE user_id = $1)', [userId])
      
      // حذف سبد خرید از session storage
      const { clearCart } = await import('../cart/data')
      const sessionId = getCookieValue(event, 'session_id')
      if (sessionId) {
        clearCart(sessionId)
      }
    } catch (cartError) {
      console.error('��� �� �ǘ ���� ��� ����:', cartError)
      // ����� ������ ��� ǐ� �ǘ ���� ��� ���� �� ��� ����� ���
    }

    // ǐ� ������ ������ ���ϡ �ѐ�� ������ ����� ���
    if (paymentMethod === 'online') {
      // ���� ���� �� ����� �ѐ�� ������ ��ԝ���
      let gatewayId = 1
      
      try {
        // ����� ���� �ѐ�� ������
        const gatewayResult = await db.query(
          'SELECT id FROM payment_gateways WHERE id = $1',
          [gatewayId]
        )
        
        if (!gatewayResult || gatewayResult.length === 0) {
          // ����� �ѐ�� ������ ��ԝ���
          const newGatewayResult = await db.query(
            `INSERT INTO payment_gateways (
              name, 
              english_name, 
              type, 
              status, 
              color, 
              fee, 
              min_amount, 
              max_amount, 
              priority, 
              is_test_mode, 
              description, 
              created_at
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW()) RETURNING id`,
            [
              '�ѐ�� ��ԝ���',
              'default_gateway',
              'online',
              'active',
              'bg-blue-500',
              0,
              0,
              999999999,
              1,
              true,
              '�ѐ�� ������ ��ԝ��� �����'
            ]
          )
          gatewayId = (newGatewayResult[0] as any).id
        }
      } catch (gatewayError) {
        console.error('��� �� �����/����� �ѐ�� ������:', gatewayError)
        // ����� �� gateway_id = 1
      }

      // ����� ��ǘ�� ������
      const paymentResult = await db.query(
        `INSERT INTO payment_transactions (
          gateway_id, 
          order_id, 
          amount, 
          payment_method, 
          status, 
          created_at
        ) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`,
        [gatewayId, orderId, finalTotal, 'online', 'pending']
      )

      return {
        success: true,
        message: '����� �� ������ ��� ��',
        data: {
          orderId,
          orderNumber,
          paymentTransactionId: (paymentResult[0] as any).id,
          redirectToPayment: true
        }
      }
    }

    // ǐ� ������ �� ��� ����
    return {
      success: true,
      message: '����� �� ������ ��� �� � �� ������ ����� ���',
      data: {
        orderId,
        orderNumber,
        redirectToPayment: false
      }
    }

  } catch (error: any) {
    console.error('��� �� ��� �����:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: '��� �� ��� �����'
    })
  }
})

// ���� ��� ���� ������ ����� �� session
async function getUserFromSession(event: any) {
  try {
    // ������� �� requireAuth �� _utils/auth.ts
    // Auth checks removed
    
    return event.context?.user || null
  } catch (error: any) {
    console.error('��� �� ������ ������� �����:', error)
    return null
  }
}

// ���� ��� ���� ������ IP ������
function getClientIP(event: any): string {
  return event.node.req.headers['x-forwarded-for'] || 
         event.node.req.headers['x-real-ip'] || 
         event.node.req.connection?.remoteAddress || 
         'unknown'
} 
