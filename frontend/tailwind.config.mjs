/** @type {import('tailwindcss').Config} */

export default {
  content: [
    './node_modules/preline/dist/*.{ts,js}',
    './src/**/*.{mjs,js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {}
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('preline/plugin'),
  ]
}
