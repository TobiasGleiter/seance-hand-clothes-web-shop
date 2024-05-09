import { useState, useEffect } from "react";
import axios from "axios";
import { useRouter } from "next/router";

const ArticleDetailsPage = ({ articleId }) => {
  const [articleDetails, setArticleDetails] = useState(null);
  const [quantity, setQuantity] = useState(1);
  const [addToCartSuccess, setAddToCartSuccess] = useState(false);

  useEffect(() => {
    fetchArticleDetails();
  }, []);

  const fetchArticleDetails = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/articleDetails?articleId=${articleId}`,
      );
      setArticleDetails(response.data);
    } catch (error) {
      console.error("Error fetching article details:", error);
    }
  };

  const handleAddToCart = () => {
    const cartItem = {
      id: articleId,
      name: articleDetails.name,
      price: articleDetails.price,
      quantity: quantity,
    };
    // Implement your own logic to add the item to cart
    console.log("Added to cart:", cartItem);
    setAddToCartSuccess(true);
  };

  return (
    <div>
      {articleDetails ? (
        <div>
          <h1>{articleDetails.name}</h1>
          <p>Price: {articleDetails.price}</p>
          <p>Description: {articleDetails.description}</p>
          <input
            type="number"
            min="1"
            value={quantity}
            onChange={(e) => setQuantity(parseInt(e.target.value))}
          />
          <button onClick={handleAddToCart}>Add to Cart</button>
          {addToCartSuccess && <p>Added to cart successfully!</p>}
        </div>
      ) : (
        <p>Loading article details...</p>
      )}
    </div>
  );
};

export default ArticleDetailsPage;