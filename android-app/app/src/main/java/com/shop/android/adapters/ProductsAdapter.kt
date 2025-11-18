package com.shop.android.adapters

import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import com.bumptech.glide.Glide
import com.shop.android.databinding.ItemProductBinding
import com.shop.android.models.Product
import com.shop.android.utils.formatPrice

/**
 * آداپتر برای نمایش محصولات در RecyclerView
 */
class ProductsAdapter(
    private val onProductClick: (Product) -> Unit
) : ListAdapter<Product, ProductsAdapter.ProductViewHolder>(ProductDiffCallback()) {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ProductViewHolder {
        val binding = ItemProductBinding.inflate(
            LayoutInflater.from(parent.context),
            parent,
            false
        )
        return ProductViewHolder(binding, onProductClick)
    }

    override fun onBindViewHolder(holder: ProductViewHolder, position: Int) {
        holder.bind(getItem(position))
    }

    class ProductViewHolder(
        private val binding: ItemProductBinding,
        private val onProductClick: (Product) -> Unit
    ) : RecyclerView.ViewHolder(binding.root) {

        fun bind(product: Product) {
            binding.apply {
                textProductName.text = product.name
                textProductPrice.text = formatPrice(product.price)
                textProductBrand.text = product.brand
                
                // نمایش تصویر محصول
                if (product.images.isNotEmpty()) {
                    Glide.with(imageProduct)
                        .load(product.images.first())
                        .centerCrop()
                        .into(imageProduct)
                }
                
                // نمایش وضعیت موجودی
                if (product.stock > 0) {
                    textStockStatus.text = "موجود"
                    textStockStatus.setTextColor(root.context.getColor(com.shop.android.R.color.stock_green))
                } else {
                    textStockStatus.text = "ناموجود"
                    textStockStatus.setTextColor(root.context.getColor(com.shop.android.R.color.stock_red))
                }
                
                // نمایش امتیاز
                product.rating?.let { rating ->
                    textRating.text = rating.toString()
                }
                
                // کلیک روی محصول
                root.setOnClickListener {
                    onProductClick(product)
                }
            }
        }
    }

    private class ProductDiffCallback : DiffUtil.ItemCallback<Product>() {
        override fun areItemsTheSame(oldItem: Product, newItem: Product): Boolean {
            return oldItem.id == newItem.id
        }

        override fun areContentsTheSame(oldItem: Product, newItem: Product): Boolean {
            return oldItem == newItem
        }
    }
} 