/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./internal/**/*.{go, templ}", "./pkg/**/*.{go, templ}", "./dashboard/*.{go, templ}"],
    theme: {
        extend: {},
    },
    daisyui: {
        themes: ["light", "dark"],
    },
    plugins: [require("daisyui")],
}

