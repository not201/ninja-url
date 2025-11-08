import { ThemeProvider } from './components/theme-provider'

function App({ children }: { children: React.ReactNode }) {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme" >
      <div id="app" className="h-screen p-4 grid grid-rows-[auto_1fr_auto]">
        {children}
      </div>
    </ThemeProvider>
  )
}

export default App
