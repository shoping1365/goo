package com.shop.android.utils

import java.text.NumberFormat
import java.util.Locale

/**
 * ابزارهای مربوط به قیمت
 */
object PriceUtils {
    
    /**
     * فرمت کردن قیمت به صورت فارسی
     */
    fun formatPrice(price: Double): String {
        val formatter = NumberFormat.getNumberInstance(Locale("fa", "IR"))
        return "${formatter.format(price)} تومان"
    }
    
    /**
     * فرمت کردن قیمت بدون واحد
     */
    fun formatPriceOnly(price: Double): String {
        val formatter = NumberFormat.getNumberInstance(Locale("fa", "IR"))
        return formatter.format(price)
    }
}

// Extension function برای راحتی استفاده
fun Double.formatPrice(): String = PriceUtils.formatPrice(this)
fun Double.formatPriceOnly(): String = PriceUtils.formatPriceOnly(this) 