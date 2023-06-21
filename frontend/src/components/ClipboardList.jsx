import React, { useState } from 'react';
import { 
  List, 
  ListItem, 
  ListItemText, 
  Button, 
  IconButton, 
  ListItemAvatar, 
  Avatar, 
  Snackbar, 
  Menu, 
  MenuItem,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  DialogContentText,
  TextField,
} from '@mui/material'
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import TextSnippetIcon from '@mui/icons-material/TextSnippet';

const ClipboardList = ({ entries, filterText, fetchClipboardEntries }) => {
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState('');
  const [copiedEntryId, setCopiedEntryId] = useState(null);
  const [anchorEl, setAnchorEl] = useState(null);
  const [selectedEntry, setSelectedEntry] = useState(null);
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [openEditDialog, setOpenEditDialog] = useState(false);
  const [editedData, setEditedData] = useState('');

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

  const handleEdit = (entry) => {
    setSelectedEntry(entry);
    setAnchorEl(null);
    setEditedData(entry.Data);
    setOpenEditDialog(true);
    //TODO: implement edit logic
  };

  const handleDelete = (entry) => {
    setSelectedEntry(entry);
    setAnchorEl(null);
    setOpenDeleteDialog(true);
  };

  const handleEditConfirm = () => {
    fetch(`http://localhost:3000/api/clipboard/${selectedEntry.ID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ Data: editedData }),
    })
      .then((response) => {
        if (response.ok) {
          setSnackbarMessage('Clipboard entry updated');
          setOpenSnackbar(true);
          setSelectedEntry(null);
          setOpenEditDialog(false);
          fetchClipboardEntries();
        } else {
          throw new Error('Failed to update clipboard entry');
        }
      })
      .catch((error) => {
        console.error(error);
        setSnackbarMessage(error);
        setOpenSnackbar(true);
      });
  };

  const handleDeleteConfirm = () => {
    
    fetch(`http://localhost:3000/api/clipboard/${selectedEntry.ID}`, {
      method: 'DELETE',
    }).then((response) => {
      response.json().then((response) => {
        setSnackbarMessage("Clipboard entry deleted");
        setOpenSnackbar(true);
        setSelectedEntry(null);
        setOpenDeleteDialog(false);
        fetchClipboardEntries();
        console.log(response);
      })
    }).catch(err => {
      console.error(err);
      setSnackbarMessage("Error deleting clipboard entry");
      setOpenSnackbar(true);
    });
  };

  const handleEditCancel = () => {
    setOpenEditDialog(false);
  };

  const handleDeleteCancel = () => {
    setOpenDeleteDialog(false);
  };

  const handleMenuOpen = (event, entry) => {
    setAnchorEl(event.currentTarget);
    setSelectedEntry(entry);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
    setSelectedEntry(null);
  };

  const handleSnackbarClose = () => {
    setOpenSnackbar(false);
  };

  const highlightText = (text) => {
    if (!filterText || filterText.trim() === '') {
      return text;
    }

    const regex = new RegExp(`(${filterText})`, 'gi');
    const parts = text.split(regex);
    return parts.map((part, index) => (
      regex.test(part) ? <mark key={index}>{part}</mark> : <span key={index}>{part}</span>
    ));
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
          }>
            <ListItemAvatar>
              <IconButton
                onClick={(event) => handleMenuOpen(event, entry)}
              >
                <Avatar sx={{ bgcolor: 'secondary.main'}}>
                  <TextSnippetIcon />
                </Avatar>
              </IconButton>
              <Menu
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl) && selectedEntry?.ID === entry.ID}
                onClose={handleMenuClose}
              >
                <MenuItem onClick={() => handleEdit(entry)}>Edit</MenuItem>
                <MenuItem onClick={() => handleDelete(entry)}>Delete</MenuItem>
              </Menu>
            </ListItemAvatar>
            <ListItemText 
              primary={highlightText(entry.Data)}
              primaryTypographyProps={{ component: 'div', noWrap: true, maxWidth: 'md'}}
              secondary={entry.Timestamp}
            />
          </ListItem>
        ))}
      </List>
      <Dialog open={openEditDialog} onClose={handleEditCancel} fullWidth>
        <DialogTitle>Edit Clipboard Entry</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label="Data"
            type="text"
            fullWidth
            multiline
            value={editedData}
            onChange={(e) => setEditedData(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleEditCancel} color="primary">
            Cancel
          </Button>
          <Button onClick={handleEditConfirm} color="primary">
            Update
          </Button>
        </DialogActions>
      </Dialog>
      <Dialog open={openDeleteDialog} onClose={handleDeleteCancel}>
        <DialogTitle>Confirm Delete</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Are you sure you want to delete this clipboard entry?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleDeleteCancel} color="primary">
            Cancel
          </Button>
          <Button onClick={handleDeleteConfirm} color="primary" autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>
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