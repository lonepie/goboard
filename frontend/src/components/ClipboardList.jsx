import React, { useState } from 'react';
import { List, ListItem, ListItemText, Button, IconButton, ListItemAvatar, Avatar, Snackbar } from '@mui/material'
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import TextSnippetIcon from '@mui/icons-material/TextSnippet';

const ClipboardList = ({ entries }) => {
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState('');
  const [copiedEntryId, setCopiedEntryId] = useState(null);

  const handleCopy = (entry) => {
    navigator.clipboard.writeText(entry.Data)
      .then(() => {
        setSnackbarMessage('Text copied to clipboard');
        setOpenSnackbar(true);
        setCopiedEntryId(entry.ID);
        setTimeout(() => {
          setCopiedEntryId(null);
        }, 3000);
        console.log('Text copied to clipboard');
      })
      .catch((error) => {
        console.error('Error copying text to clipboard:', error);
      });
  };

  const handleSnackbarClose = () => {
    setOpenSnackbar(false);
  };

  return (
    <>
      <List>
        {entries.map((entry) => (
          <ListItem key={entry.ID} divider secondaryAction={
            <Button 
              variant='contained' 
              color={copiedEntryId === entry.ID ? 'secondary' : 'primary'} 
              disabled={copiedEntryId === entry.ID}
              startIcon={<ContentCopyIcon />} 
              onClick={() => handleCopy(entry)}
            >
              {copiedEntryId === entry.ID ? 'Copied' : 'Copy'}
            </Button>
            // <IconButton edge='end' aria-label='copy' onClick={() => handleCopy(entry.Data)}>
            //   <ContentCopyIcon />
            // </IconButton>
          }>
            <ListItemAvatar>
              <Avatar>
                <TextSnippetIcon />
              </Avatar>
            </ListItemAvatar>
            <ListItemText primary={entry.Data} secondary={entry.Timestamp} />
          </ListItem>
        ))}
      </List>
      <Snackbar
          open={openSnackbar}
          autoHideDuration={3000}
          onClose={handleSnackbarClose}
          message={snackbarMessage}
        />
    </>
  );
};

export default ClipboardList;