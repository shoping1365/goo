package com.shop.android.models

import com.google.gson.annotations.SerializedName

/**
 * مدل پاسخ ورود
 */
data class LoginResponse(
    @SerializedName("token")
    val token: String,
    
    @SerializedName("user")
    val user: User
)

/**
 * مدل کاربر
 */
data class User(
    @SerializedName("id")
    val id: Int,
    
    @SerializedName("name")
    val name: String,
    
    @SerializedName("email")
    val email: String,
    
    @SerializedName("phone")
    val phone: String? = null
) 