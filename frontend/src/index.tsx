import { createRoot } from 'react-dom/client'
import 'tailwindcss/tailwind.css'
import App from 'components/App'
import Configuration from 'config/config'

const container = document.getElementById('root') as HTMLDivElement
const root = createRoot(container)

console.log("Configuration has been loaded.", Configuration);

root.render(<App />)
