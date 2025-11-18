import eslintPluginVue from 'eslint-plugin-vue'
import eslintConfigPrettier from 'eslint-config-prettier'
import tseslint from '@typescript-eslint/eslint-plugin'
import tsparser from '@typescript-eslint/parser'

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
        files: ['**/*.{js,ts,vue}'],
        languageOptions: {
            parser: tsparser,
            parserOptions: {
                ecmaVersion: 'latest',
                sourceType: 'module',
                parser: '@typescript-eslint/parser'
            }
        },
        plugins: {
            '@typescript-eslint': tseslint
        },
        rules: {
            // Vue rules
            'vue/multi-word-component-names': 'off',
            'vue/no-v-html': 'warn',
            'vue/require-default-prop': 'off',
            'vue/require-prop-types': 'off',
            'vue/no-setup-props-destructure': 'off',
            'vue/no-v-text-v-html-on-component': 'warn',
            'vue/html-self-closing': ['error', {
                'html': {
                    'void': 'always',
                    'normal': 'never',
                    'component': 'always'
                }
            }],

            // TypeScript rules
            '@typescript-eslint/no-unused-vars': ['warn', {
                'argsIgnorePattern': '^_',
                'varsIgnorePattern': '^_'
            }],
            '@typescript-eslint/no-explicit-any': 'warn',

            // General rules
            'no-console': ['warn', { allow: ['warn', 'error'] }],
            'no-debugger': 'warn',
            'no-unused-vars': 'off', // Use TypeScript version instead
            'prefer-const': 'warn',
            'no-var': 'error'
        }
    },
    eslintConfigPrettier
]
