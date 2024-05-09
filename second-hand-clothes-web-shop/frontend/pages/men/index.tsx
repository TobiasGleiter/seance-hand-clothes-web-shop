import { useState, useEffect } from "react";
import axios from "axios";

const MenPage = () => {
  const [menArticles, setMenArticles] = useState([]);
  const [filteredArticles, setFilteredArticles] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState("");

  useEffect(() => {
    fetchMenArticles();
  }, []);

  const fetchMenArticles = async () => {
    try {
      const response = await axios.get("http://localhost:8080/articles/men"); //?category=men
      setMenArticles(response.data);
      setFilteredArticles(response.data);
    } catch (error) {
      console.error("Error fetching men articles:", error);
    }
  };

  const handleCategoryFilter = (category) => {
    setSelectedCategory(category);
    if (category === "") {
      setFilteredArticles(menArticles);
    } else {
      const filtered = menArticles.filter(
        (article) => article.category === category,
      );
      setFilteredArticles(filtered);
    }
  };


  return (
    <div>
      <h1>Men&apos;s Articles</h1>
      <div>
        <button onClick={() => handleCategoryFilter("")}>All</button>
        <button onClick={() => handleCategoryFilter("Hosen & Jeans")}>
          Hosen & Jeans
        </button>
        <button onClick={() => handleCategoryFilter("Pullover & Strickjacken")}>
          Pullover & Strickjacken
        </button>
        <button onClick={() => handleCategoryFilter("T-Shirts & Poloshirts")}>
          T-Shirts & Poloshirts
        </button>
        <button onClick={() => handleCategoryFilter("Hemden")}>Hemden</button>
        <button onClick={() => handleCategoryFilter("Unterwäsche")}>
          Unterwäsche
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

export default MenPage;