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
            // Vue rules
            'vue/multi-word-component-names': 'off',
            'vue/no-v-html': 'warn',
            'vue/require-default-prop': 'off',
            'vue/require-prop-types': 'off',
            'vue/no-setup-props-destructure': 'off',
            'vue/no-v-text-v-html-on-component': 'warn',
            
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
        }
    },
    eslintConfigPrettier
]
