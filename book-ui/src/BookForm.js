import React, { useState } from "react";
import axios from "axios";
import { TextField, Button, Typography } from "@mui/material";
import "./css/BookForm.css";

const BookForm = () => {
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [genre, setGenre] = useState("");

  const handleCreateBook = () => {
    axios
      .post("http://localhost:8080/books", { title, author, genre })
      .then((response) => console.log("Book created:", response.data))
      .catch((error) => console.error("Error creating book:", error));
  };

  return (
    <div className="book-form-container">
      <Typography variant="h4" gutterBottom>
        Add your Book to Library !!
      </Typography>
      <form>
        <TextField
          label="Title"
          variant="outlined"
          fullWidth
          margin="normal"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <TextField
          label="Author"
          variant="outlined"
          fullWidth
          margin="normal"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
        />
        <TextField
          label="Genre"
          variant="outlined"
          fullWidth
          margin="normal"
          value={genre}
          onChange={(e) => setGenre(e.target.value)}
        />
        <Button
          variant="contained"
          color="secondary"
          onClick={handleCreateBook}
        >
          Add Book
        </Button>
      </form>
    </div>
  );
};

export default BookForm;
