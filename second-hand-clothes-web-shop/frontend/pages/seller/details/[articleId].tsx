import { useState, useEffect } from "react";
import axios from "axios";

const SellerArticleDetailsPage = ({ userId, articleId }) => {
  const [articleDetails, setArticleDetails] = useState(null);
  const [deleteSuccess, setDeleteSuccess] = useState(false);
  const [updateSuccess, setUpdateSuccess] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    category: "",
    price: 0,
    size: "",
    rating: 0,
    subcategory: "",
  });

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

  const handleDeleteArticle = async () => {
    try {
      await axios.post("/deleteArticleForUser", {
        userId: userId,
        articleId: articleId,
      });
      setDeleteSuccess(true);
      // Redirect to seller page or perform other actions after deletion
    } catch (error) {
      console.error("Error deleting article:", error);
    }
  };
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSaveArticle = async () => {
    try {
      await axios.post("/updateArticleForUser", {
        userId: userId,
        article: { _id: articleId, ...formData },
      });
      setUpdateSuccess(true);
      // Redirect to seller page or perform other actions after update
    } catch (error) {
      console.error("Error updating article:", error);
    }
  };

  return (
    <div>
      {articleDetails ? (
        <div>
          <h1>Article Details</h1>
          <input
            type="text"
            name="name"
            placeholder="Name"
            value={formData.name}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="category"
            placeholder="Category"
            value={formData.category}
            onChange={handleChange}
            required
          />
          <input
            type="number"
            name="price"
            placeholder="Price"
            value={formData.price}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="size"
            placeholder="Size"
            value={formData.size}
            onChange={handleChange}
            required
          />
          <input
            type="number"
            name="rating"
            placeholder="Rating"
            value={formData.rating}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="subcategory"
            placeholder="Subcategory"
            value={formData.subcategory}
            onChange={handleChange}
            required
          />
          <button onClick={handleSaveArticle}>Save Article</button>
          {updateSuccess && <p>Article updated successfully!</p>}
        </div>
      ) : (
        <p>Loading article details...</p>
      )}
    </div>
  );
};

export default SellerArticleDetailsPage;