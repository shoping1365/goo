-- اسکریپت برای به‌روزرسانی order_items که product_name، product_image و product_sku ندارند
-- این اسکریپت را در PostgreSQL اجرا کنید

-- به‌روزرسانی تمام order_items که اطلاعات محصول ندارند
UPDATE order_items oi
SET 
    product_name = p.name,
    product_image = p.image,
    product_sku = p.sku
FROM products p
WHERE 
    oi.product_id = p.id 
    AND (
        oi.product_name IS NULL 
        OR oi.product_name = '' 
        OR oi.product_image IS NULL 
        OR oi.product_sku IS NULL 
        OR oi.product_sku = ''
    );

-- بررسی نتیجه
SELECT 
    oi.id,
    oi.order_id,
    oi.product_id,
    oi.product_name,
    oi.product_image,
    oi.product_sku,
    oi.quantity,
    oi.unit_price,
    oi.final_price
FROM order_items oi
WHERE order_id IN (
    SELECT id FROM orders ORDER BY created_at DESC LIMIT 10
)
ORDER BY oi.order_id DESC, oi.id ASC;

-- آمار order_items که اطلاعات کامل دارند
SELECT 
    COUNT(*) as total_items,
    COUNT(CASE WHEN product_name IS NOT NULL AND product_name != '' THEN 1 END) as items_with_name,
    COUNT(CASE WHEN product_image IS NOT NULL THEN 1 END) as items_with_image,
    COUNT(CASE WHEN product_sku IS NOT NULL AND product_sku != '' THEN 1 END) as items_with_sku
FROM order_items;



















