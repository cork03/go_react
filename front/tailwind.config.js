/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "white": "#ffffff",
        "gray": "#e5e7eb"
      }
    },
  },
  plugins: [],
}
