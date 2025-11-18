package com.shop.android.fragments

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.recyclerview.widget.LinearLayoutManager
import com.shop.android.R
import com.shop.android.adapters.ProductsAdapter
import com.shop.android.databinding.FragmentHomeBinding
import com.shop.android.viewmodels.HomeViewModel
import dagger.hilt.android.AndroidEntryPoint

/**
 * فرگمنت صفحه اصلی
 * نمایش محصولات ویژه و جدید
 */
@AndroidEntryPoint
class HomeFragment : Fragment() {

    private var _binding: FragmentHomeBinding? = null
    private val binding get() = _binding!!
    
    private val viewModel: HomeViewModel by viewModels()
    private lateinit var productsAdapter: ProductsAdapter

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        _binding = FragmentHomeBinding.inflate(inflater, container, false)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        
        setupRecyclerView()
        setupObservers()
        viewModel.loadProducts()
    }

    private fun setupRecyclerView() {
        productsAdapter = ProductsAdapter { product ->
            // انتقال به صفحه جزئیات محصول
            // TODO: Navigate to product detail
        }
        
        binding.recyclerViewProducts.apply {
            layoutManager = LinearLayoutManager(context)
            adapter = productsAdapter
        }
    }

    private fun setupObservers() {
        viewModel.products.observe(viewLifecycleOwner) { products ->
            productsAdapter.submitList(products)
        }
        
        viewModel.isLoading.observe(viewLifecycleOwner) { isLoading ->
            binding.progressBar.visibility = if (isLoading) View.VISIBLE else View.GONE
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
} 