package com.shop.android.activities

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.shop.android.R
import dagger.hilt.android.AndroidEntryPoint

/**
 * اکتیویتی سبد خرید
 */
@AndroidEntryPoint
class CartActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_cart)
        
        // TODO: پیاده‌سازی نمایش سبد خرید
    }
} 