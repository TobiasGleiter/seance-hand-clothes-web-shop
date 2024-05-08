Below is the implementation of the article details page with the specified functionalities:

```jsx
import { useState, useEffect } from "react";
import axios from "axios";

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
        `/articleDetails?articleId=${articleId}`,
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
```

This component fetches the details of a specific article from the backend when the page loads and displays them. It provides an input field to specify the quantity of the article to add to the cart and a button to add it to the cart. Once the article is successfully added to the cart, a success message is displayed.
