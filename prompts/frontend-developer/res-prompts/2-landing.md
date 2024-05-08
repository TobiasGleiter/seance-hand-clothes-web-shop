Here's a basic implementation of the landing page functionality using Next.js and React. We'll create a component called `LandingPage` to handle fetching articles from the backend, displaying them, implementing search by article name, and filtering by size.

```jsx
import { useState, useEffect } from "react";
import axios from "axios";

const LandingPage = () => {
  const [articles, setArticles] = useState([]);
  const [filteredArticles, setFilteredArticles] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [sizeFilter, setSizeFilter] = useState("");

  useEffect(() => {
    fetchArticles();
  }, []);

  const fetchArticles = async () => {
    try {
      const response = await axios.get("/articles");
      setArticles(response.data);
      setFilteredArticles(response.data);
    } catch (error) {
      console.error("Error fetching articles:", error);
    }
  };

  const handleSearch = (event) => {
    setSearchTerm(event.target.value);
    filterArticles(event.target.value, sizeFilter);
  };

  const handleSizeFilter = (event) => {
    setSizeFilter(event.target.value);
    filterArticles(searchTerm, event.target.value);
  };

  const filterArticles = (searchTerm, sizeFilter) => {
    let filtered = articles.filter((article) =>
      article.name.toLowerCase().includes(searchTerm.toLowerCase()),
    );

    if (sizeFilter) {
      filtered = filtered.filter((article) => article.size === sizeFilter);
    }

    setFilteredArticles(filtered);
  };

  return (
    <div>
      <h1>Second-Hand Clothes Web Shop</h1>
      <input
        type="text"
        placeholder="Search by article name"
        value={searchTerm}
        onChange={handleSearch}
      />
      <select onChange={handleSizeFilter} value={sizeFilter}>
        <option value="">Filter by size</option>
        <option value="S">Small</option>
        <option value="M">Medium</option>
        <option value="L">Large</option>
      </select>
      <div>
        {filteredArticles.map((article) => (
          <div key={article.id}>
            <h2>{article.name}</h2>
            <p>Size: {article.size}</p>
            <p>Category: {article.category}</p>
            <p>Price: {article.price}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default LandingPage;
```

Make sure you have Axios installed (`npm install axios`) for making HTTP requests. This component fetches articles from the backend when the page loads, allows searching articles by name, and filtering articles by size. It displays the article name, size, category, and price for each article.

You can include this `LandingPage` component in your Next.js application wherever you want the landing page to appear.
