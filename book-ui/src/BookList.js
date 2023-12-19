import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  List,
  ListItem,
  ListItemText,
  IconButton,
  Typography,
} from "@mui/material";
import "./css/BookList.css";
import RemoveIcon from "@mui/icons-material/Remove";
const BookList = () => {
  const [books, setBooks] = useState([]);

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
              secondary={`Author: ${book.author}`}
            />
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
    </div>
  );
};

export default BookList;
