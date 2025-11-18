package com.shop.android.activities

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.shop.android.R
import dagger.hilt.android.AndroidEntryPoint

/**
 * اکتیویتی تکمیل خرید
 */
@AndroidEntryPoint
class CheckoutActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_checkout)
        
        // TODO: پیاده‌سازی تکمیل خرید
    }
} 