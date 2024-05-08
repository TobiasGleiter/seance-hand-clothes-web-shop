import { useState, useEffect } from "react";
import axios from "axios";

const KidsPage = () => {
  const [kidsArticles, setKidsArticles] = useState([]);
  const [filteredArticles, setFilteredArticles] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState("");

  useEffect(() => {
    fetchKidsArticles();
  }, []);

  const fetchKidsArticles = async () => {
    try {
      const response = await axios.get("/articles/kids");
      setKidsArticles(response.data);
      setFilteredArticles(response.data);
    } catch (error) {
      console.error("Error fetching kids articles:", error);
    }
  };

  const handleCategoryFilter = (category) => {
    setSelectedCategory(category);
    if (category === "") {
      setFilteredArticles(kidsArticles);
    } else {
      const filtered = kidsArticles.filter(
        (article) => article.category === category,
      );
      setFilteredArticles(filtered);
    }
  };

  return (
    <div>
      <h1>Kids' Articles</h1>
      <div>
        <button onClick={() => handleCategoryFilter("")}>All</button>
        <button onClick={() => handleCategoryFilter("Hosen & Jeans")}>
          Hosen & Jeans
        </button>
        <button onClick={() => handleCategoryFilter("Hemden")}>Hemden</button>
        <button onClick={() => handleCategoryFilter("Jumpsuits & Sets")}>
          Jumpsuits & Sets
        </button>
        <button onClick={() => handleCategoryFilter("Unterwäsche")}>
          Unterwäsche
        </button>
        <button onClick={() => handleCategoryFilter("Oberteile")}>
          Oberteile
        </button>
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

export default KidsPage;