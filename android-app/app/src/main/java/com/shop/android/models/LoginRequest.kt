package com.shop.android.models

import com.google.gson.annotations.SerializedName

/**
 * مدل درخواست ورود
 */
data class LoginRequest(
    @SerializedName("email")
    val email: String,
    
    @SerializedName("password")
    val password: String
) 