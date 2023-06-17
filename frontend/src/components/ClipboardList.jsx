import React from 'react';
import { List, ListItem, ListItemText, Button, IconButton, ListItemAvatar, Avatar } from '@mui/material'
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import TextSnippetIcon from '@mui/icons-material/TextSnippet';

const ClipboardList = ({ entries }) => {
  const handleCopy = (data) => {
    navigator.clipboard.writeText(data)
      .then(() => {
        console.log('Text copied to clipboard');
      })
      .catch((error) => {
        console.error('Error copying text to clipboard:', error);
      });
  };

  return (
    <List>
      {entries.map((entry) => (
        <ListItem key={entry.ID} divider secondaryAction={
          <IconButton edge='end' aria-label='copy' onClick={() => handleCopy(entry.Data)}>
            <ContentCopyIcon />
          </IconButton>
        }>
          <ListItemAvatar>
            <Avatar>
              <TextSnippetIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary={entry.Data} secondary={entry.Timestamp} />
          {/* <Button variant="contained" color="primary" onClick={() => handleCopy(entry.Data)}>
            Copy
          </Button> */}
        </ListItem>
      ))}
    </List>
  );
};

export default ClipboardList;