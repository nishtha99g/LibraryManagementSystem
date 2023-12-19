import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  List,
  ListItem,
  ListItemText,
  IconButton,
  Typography,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Button,
} from "@mui/material";
import "../css/BookList.css";
import RemoveIcon from "@mui/icons-material/Remove";
import EditIcon from "@mui/icons-material/Edit";

import "../css/BookList.css";
const BookList = () => {
  const [books, setBooks] = useState([]);
  const [editBook, setEditBook] = useState(null);
  const [openEditDialog, setOpenEditDialog] = useState(false);
  const [updatedTitle, setUpdatedTitle] = useState("");
  const [updatedAuthor, setUpdatedAuthor] = useState("");
  const [updatedGenre, setUpdatedGenre] = useState("");

  const fetchBooks = () => {
    axios
      .get("http://localhost:8080/books")
      .then((response) => setBooks(response.data))
      .catch((error) => console.error("Error fetching books:", error));
  };

  const handleDeleteBook = (id) => {
    axios
      .delete(`http://localhost:8080/books/${id}`)
      .then(() => {
        console.log("Book deleted successfully");
        fetchBooks(); // Refresh the book list after deletion
      })
      .catch((error) => console.error("Error deleting book:", error));
  };

  const handleOpenEditDialog = (book) => {
    setEditBook(book);
    setUpdatedTitle(book.title);
    setUpdatedAuthor(book.author);
    setOpenEditDialog(true);
  };

  const handleUpdateBook = (id) => {
    const updatedBook = {
      title: updatedTitle,
      author: updatedAuthor,
      genre: updatedGenre,
    };

    axios
      .put(`http://localhost:8080/books/${id}`, updatedBook)
      .then(() => {
        console.log("Book updated successfully");
        handleCloseEditDialog(); // Close the edit dialog after updating
        fetchBooks(); // Refresh the book list after updating
      })
      .catch((error) => console.error("Error updating book:", error));
  };

  const handleCloseEditDialog = () => {
    setOpenEditDialog(false);
  };

  const handleTitleChange = (event) => {
    setUpdatedTitle(event.target.value);
  };

  const handleAuthorChange = (event) => {
    setUpdatedAuthor(event.target.value);
  };

  const handleGenreChange = (event) => {
    setUpdatedGenre(event.target.value);
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  return (
    <div className="book-list-container">
      <Typography variant="h4" gutterBottom>
        Book List
      </Typography>
      <List>
        {books.map((book) => (
          <ListItem key={book.ID}>
            <ListItemText
              primary={book.title}
              secondary={`Author: ${book.author} | Genre: ${book.genre}`}
            />
            <IconButton
              color="secondary"
              onClick={() => handleOpenEditDialog(book)}
              aria-label="edit"
            >
              <EditIcon />
            </IconButton>
            <IconButton
              color="secondary"
              onClick={() => handleDeleteBook(book.ID)}
              aria-label="delete"
            >
              <RemoveIcon />
            </IconButton>
          </ListItem>
        ))}
      </List>
      <Dialog open={openEditDialog} onClose={handleCloseEditDialog}>
        <DialogTitle>Edit Book</DialogTitle>
        <DialogContent>
          <TextField
            label="Title"
            value={updatedTitle}
            onChange={handleTitleChange}
            fullWidth
            margin="normal"
          />
          <TextField
            label="Author"
            value={updatedAuthor}
            onChange={handleAuthorChange}
            fullWidth
            margin="normal"
          />
          <TextField
            label="Genre"
            value={updatedGenre}
            onChange={handleGenreChange}
            fullWidth
            margin="normal"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseEditDialog}>Cancel</Button>
          <Button onClick={() => handleUpdateBook(editBook.ID)}>Update</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default BookList;
