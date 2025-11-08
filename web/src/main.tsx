import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { Footer } from './components/footer.tsx'
import { Header } from './components/header.tsx'
import { Index } from './components/index.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App>
      <Header />
      <Index />
      <Footer />
    </App>
  </StrictMode>,
)
