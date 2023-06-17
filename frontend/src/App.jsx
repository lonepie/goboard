import { useEffect, useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
// import './App.css'
// import { styled } from '@mui/material/styles'
import { Box, TextField, Typography } from '@mui/material'
import ClipboardList from './components/ClipboardList'


function App() {
  // const [count, setCount] = useState(0)
  const [clipboardEntries, setClipboardEntries] = useState([]);
  const [filterText, setFilterText] = useState('');

  useEffect(() => {
    fetch("http://localhost:3000/api/clipboard")
      .then((response) => response.json())
      .then((data) => setClipboardEntries(data))
      .catch((error) => console.log(error));
  }, []);

  const handleFilterChange = (event) => {
    setFilterText(event.target.value);
  };

  const filteredEntries = clipboardEntries.filter((entry) =>
    entry.Data.toLowerCase().includes(filterText.toLowerCase())
  );

  return (
    <>
      {/* <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p> */}
      <Typography variant='h1' gutterBottom>
        goBoard Clipboard Entries
      </Typography>
      <Box sx={{ bgcolor: 'background.paper'}}>
        <TextField label="Filter" value={filterText} onChange={handleFilterChange} fullWidth margin='normal' />
        <ClipboardList entries={filteredEntries} />
      </Box>
    </>
  )
}

export default App
