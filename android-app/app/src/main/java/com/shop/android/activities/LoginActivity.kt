package com.shop.android.activities

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.shop.android.R
import dagger.hilt.android.AndroidEntryPoint

/**
 * اکتیویتی ورود
 */
@AndroidEntryPoint
class LoginActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)
        
        // TODO: پیاده‌سازی ورود کاربر
    }
} 