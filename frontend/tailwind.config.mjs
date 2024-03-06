/** @type {import('tailwindcss').Config} */

export default {
  content: [
    './src/**/*.{mjs,js,ts,jsx,tsx}',
    'node_modules/preline/dist/*.js',
  ],
  theme: {
    extend: {}
  },
  plugins: [
    require('preline/plugin'),
  ]
}
