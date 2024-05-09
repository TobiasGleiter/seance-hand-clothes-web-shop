import { useState, useEffect } from "react";
import axios from "axios";

const WomenPage = () => {
  const [womenArticles, setWomenArticles] = useState([]);
  const [filteredArticles, setFilteredArticles] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState("");

  useEffect(() => {
    fetchWomenArticles();
  }, []);

  const fetchWomenArticles = async () => {
    try {
      const response = await axios.get("http://localhost:8080/articles/women"); //?category=women
      setWomenArticles(response.data);
      setFilteredArticles(response.data);
    } catch (error) {
      console.error("Error fetching women articles:", error);
    }
  };

  const handleCategoryFilter = (category) => {
    setSelectedCategory(category);
    if (category === "") {
      setFilteredArticles(womenArticles);
    } else {
      const filtered = womenArticles.filter(
        (article) => article.category === category,
      );
      setFilteredArticles(filtered);
    }
  };

  return (
    <div>
      <h1>Women&apos;s Articles</h1>
      <div>
        <button onClick={() => handleCategoryFilter("")}>All</button>
        <button onClick={() => handleCategoryFilter("Hosen & Jeans")}>
          Hosen & Jeans
        </button>
        <button onClick={() => handleCategoryFilter("Pullover & Strickjacken")}>
          Pullover & Strickjacken
        </button>
        <button onClick={() => handleCategoryFilter("Hemden")}>Hemden</button>
        <button onClick={() => handleCategoryFilter("Unterwäsche")}>
          Unterwäsche
        </button>
        <button onClick={() => handleCategoryFilter("Blusen")}>Blusen</button>
        <button onClick={() => handleCategoryFilter("Tops")}>Tops</button>
      </div>
      <div>
        {filteredArticles.map((article) => (
          <div key={article.id}>
            <h2>{article.name}</h2>
            <p>Category: {article.category}</p>
            <p>Price: {article.price}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default WomenPage;