import { useState, useEffect } from "react";
import axios from "axios";
import Navigation from "../components/Navigation";
import Link from "next/link";

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
      const response = await axios.get("http://localhost:8080/articles");
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
      <Navigation />
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
          <Link href={`/articles/${article._id}`} key={article._id}>
            <h2>{article.name}</h2>
            <p>Size: {article.size}</p>
            <p>Category: {article.category}</p>
            <p>Price: {article.price}</p>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default LandingPage;