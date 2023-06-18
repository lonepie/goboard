import { useEffect, useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
// import './App.css'
// import { styled } from '@mui/material/styles'
import { AppBar, Avatar, Box, Button, Grid, IconButton, Paper, TextField, Toolbar, Typography } from '@mui/material'
import ContentPasteIcon from '@mui/icons-material/ContentPaste';
import RefreshIcon from '@mui/icons-material/Refresh';
import ClipboardList from './components/ClipboardList';


function App() {
  // const [count, setCount] = useState(0)
  const [clipboardEntries, setClipboardEntries] = useState([]);
  const [filterText, setFilterText] = useState('');

  useEffect(() => {
    fetchClipboardEntries();
  }, []);

  const fetchClipboardEntries = () => {
    fetch("http://localhost:3000/api/clipboard")
      .then((response) => response.json())
      .then((data) => setClipboardEntries(data))
      .catch((error) => console.log(error));
  };

  const handleFilterChange = (event) => {
    setFilterText(event.target.value);
  };
  
  const handleRefresh = () => {
    fetchClipboardEntries();
  };

  const filteredEntries = clipboardEntries.filter((entry) =>
    entry.Data.toLowerCase().includes(filterText.toLowerCase())
  );

  return (
    <>
    <AppBar color='primary' position='static'>
      <Toolbar>
          <Avatar sx={{ color: 'text.primary', bgcolor: 'primary.main', display: 'flex', mr: 1 }}>
            <ContentPasteIcon />
          </Avatar>
          <Typography variant='h6' noWrap component='div' sx={{ display: 'flex', mr: 2, textDecoration: 'none', color: 'inherit', flexGrow: 1 }}>
            goBoard Entries
          </Typography>
          <Box>
            <Button variant='outlined' startIcon={<RefreshIcon />} onClick={handleRefresh}>Refresh</Button>
          </Box>
      </Toolbar>
    </AppBar>
    <Paper>
      <Box>
        <TextField label="Filter" value={filterText} onChange={handleFilterChange} fullWidth margin='normal' />
        <ClipboardList entries={filteredEntries} filterText={filterText} fetchClipboardEntries={fetchClipboardEntries} />
      </Box>
    </Paper>
    </>
  )
}

export default App
