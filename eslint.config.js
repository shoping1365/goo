import tseslint from '@typescript-eslint/eslint-plugin'
import tsparser from '@typescript-eslint/parser'
import eslintConfigPrettier from 'eslint-config-prettier'
import vueParser from 'vue-eslint-parser'
import eslintPluginVue from 'eslint-plugin-vue'

export default [
    {
        ignores: [
            'node_modules/**',
            '.nuxt/**',
            '.output/**',
            'dist/**',
            '.git/**',
            'coverage/**',
            'my-go-backend/**',
            'android-app/**',
            'codeql-dbs/**',
            // Ignore TinyMCE third-party library files
            'public/tinymce-assets/**',
            '**/tinymce.js',
            '**/tinymce.min.js',
            '**/tinymce-assets/**'
        ]
    },
    ...eslintPluginVue.configs['flat/recommended'],
    {
        files: ['**/*.{js,ts}'],
        languageOptions: {
            parser: tsparser,
            parserOptions: {
                ecmaVersion: 'latest',
                sourceType: 'module'
            }
        },
        plugins: {
            '@typescript-eslint': tseslint
        },
        rules: {
            // ============================================
            // قوانین TypeScript - صریح و محکم
            // ============================================

            // ممنوعیت استفاده از any - خطا
            '@typescript-eslint/no-explicit-any': 'error',

            // متغیرهای استفاده نشده - خطا (با استثنا برای متغیرهایی که با _ شروع می‌شوند)
            '@typescript-eslint/no-unused-vars': ['error', {
                'argsIgnorePattern': '^_',
                'varsIgnorePattern': '^_',
                'caughtErrorsIgnorePattern': '^_',
                'destructuredArrayIgnorePattern': '^_'
            }],

            // ============================================
            // قوانین عمومی - صریح و محکم
            // ============================================

            // ممنوعیت استفاده از console.log - خطا (فقط console.warn و console.error مجاز)
            'no-console': ['error', { allow: ['warn', 'error'] }],

            // ممنوعیت استفاده از debugger - خطا
            'no-debugger': 'error',

            // استفاده از نسخه TypeScript برای no-unused-vars
            'no-unused-vars': 'off',
        }
    },
    {
        files: ['**/*.vue'],
        languageOptions: {
            parser: vueParser,
            parserOptions: {
                parser: tsparser,
                ecmaVersion: 'latest',
                sourceType: 'module',
                extraFileExtensions: ['.vue']
            }
        },
        plugins: {
            '@typescript-eslint': tseslint
        },
        rules: {
            // ============================================
            // قوانین Vue - صریح و محکم
            // ============================================

            // نام کامپوننت باید چند کلمه‌ای باشد
            'vue/multi-word-component-names': 'off',

            // هشدار برای استفاده از v-html
            'vue/no-v-html': 'warn',

            // غیرفعال کردن require-default-prop
            'vue/require-default-prop': 'off',

            // غیرفعال کردن require-prop-types
            'vue/require-prop-types': 'off',

            // غیرفعال کردن no-setup-props-destructure
            'vue/no-setup-props-destructure': 'off',

            // هشدار برای v-text و v-html روی کامپوننت
            'vue/no-v-text-v-html-on-component': 'warn',

            // ============================================
            // قوانین Vue - Mutation Props (خطا)
            // ============================================

            // ممنوعیت تغییر مستقیم props - خطا
            'vue/no-mutating-props': 'error',

            // ============================================
            // قوانین Vue - DefineEmits (خطا)
            // ============================================

            // ممنوعیت فراخوانی چندباره defineEmits - خطا
            'vue/valid-define-emits': 'error',

            // ============================================
            // قوانین Vue - Transition (خطا)
            // ============================================

            // الزام استفاده از v-if یا v-show در transition - خطا
            'vue/require-toggle-inside-transition': 'error',

            // ============================================
            // قوانین Vue - Duplicate Keys (خطا)
            // ============================================

            // ممنوعیت کلیدهای تکراری - خطا
            'vue/no-dupe-keys': 'error',

            // ============================================
            // قوانین Vue - Unused Variables (خطا)
            // ============================================

            // متغیرهای استفاده نشده در Vue - خطا
            'vue/no-unused-vars': ['error', {
                'ignorePattern': '^_'
            }],

            // ============================================
            // قوانین TypeScript - صریح و محکم
            // ============================================

            // ممنوعیت استفاده از any - خطا
            '@typescript-eslint/no-explicit-any': 'error',

            // متغیرهای استفاده نشده - خطا (با استثنا برای متغیرهایی که با _ شروع می‌شوند)
            '@typescript-eslint/no-unused-vars': ['error', {
                'argsIgnorePattern': '^_',
                'varsIgnorePattern': '^_',
                'caughtErrorsIgnorePattern': '^_',
                'destructuredArrayIgnorePattern': '^_'
            }],

            // ============================================
            // قوانین عمومی - صریح و محکم
            // ============================================

            // ممنوعیت استفاده از console.log - خطا (فقط console.warn و console.error مجاز)
            'no-console': ['error', { allow: ['warn', 'error'] }],

            // ممنوعیت استفاده از debugger - خطا
            'no-debugger': 'error',

            // استفاده از نسخه TypeScript برای no-unused-vars
            'no-unused-vars': 'off',
        }
    },
    eslintConfigPrettier
]
