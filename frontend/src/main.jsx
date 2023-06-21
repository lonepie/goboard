import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
// import './index.css'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { CssBaseline } from '@mui/material'
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import '@fontsource/roboto-mono';

const themeOptions = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#cba6f7',
    },
    secondary: {
      main: '#fab387',
    },
    background: {
      default: '#11111b',
      paper: '#181825',
    },
    text: {
      primary: '#cdd6f4',
      secondary: '#bac2de',
      disabled: '#6c7086',
      hint: '#a6adc8',
    },
    error: {
      main: '#f38ba8',
    },
    warning: {
      main: '#f9e2af',
    },
    info: {
      main: '#89b4fa',
    },
    success: {
      main: '#a6e3a1',
    },
    divider: '#7f849c',
  },
});

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#674ab7',
    },
    secondary: {
      main: '#ff5722',
    },
  },
});

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <ThemeProvider theme={themeOptions}>
      <CssBaseline />
      <App />
    </ThemeProvider>
  </React.StrictMode>,
)
