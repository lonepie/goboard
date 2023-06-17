import { useEffect, useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
// import './App.css'
// import { styled } from '@mui/material/styles'
import { AppBar, Avatar, Box, Grid, Paper, TextField, Typography } from '@mui/material'
import ContentPasteIcon from '@mui/icons-material/ContentPaste';
import ClipboardList from './components/ClipboardList'
import { blue, blueGrey } from '@mui/material/colors';


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
    <AppBar position='static'>
      <Grid container>
        <Grid item>
          <Avatar sx={{ bgcolor: blue[500], display: 'flex', mr: 1 }}>
            <ContentPasteIcon />
          </Avatar>
        </Grid>
        <Grid item>
          <Typography variant='h6' noWrap component='a' href='/' sx={{ display: 'flex', mr: 2, textDecoration: 'none', color: 'inherit', lineHeight: '2em' }}>
            goBoard Entries
          </Typography>
        </Grid>
      </Grid>
    </AppBar>
    <Paper>
      <Box>
        <TextField label="Filter" value={filterText} onChange={handleFilterChange} fullWidth margin='normal' />
        <ClipboardList entries={filteredEntries} />
      </Box>
    </Paper>
    </>
  )
}

export default App
