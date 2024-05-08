import { useState, useEffect } from "react";
import axios from "axios";

const CheckoutPage = ({ cartItems }) => {
  const [totalPrice, setTotalPrice] = useState(0);
  const [orderSubmitted, setOrderSubmitted] = useState(false);

  // Calculate total price
  const calculateTotalPrice = () => {
    let total = 0;
    cartItems.forEach((item) => {
      total += item.price * item.quantity;
    });
    setTotalPrice(total);
  };

  // Function to send order to backend
  const sendOrderToBackend = async () => {
    try {
      const orderData = {
        name: "John Doe", // You can replace this with user's name
        totalPrice: totalPrice,
        userId: "user123", // You can replace this with the actual user ID
        articles: cartItems.map((item) => ({
          articleId: item.id,
          quantity: item.quantity,
        })),
      };

      const response = await axios.post("/orders", orderData);
      console.log("Order submitted successfully:", response.data);
      setOrderSubmitted(true);
    } catch (error) {
      console.error("Error submitting order:", error);
    }
  };

  // Calculate total price when component mounts
  useEffect(() => {
    calculateTotalPrice();
  }, []);

  return (
    <div>
      <h1>Checkout</h1>
      {orderSubmitted ? (
        <p>Order submitted successfully!</p>
      ) : (
        <div>
          <p>Total Price: {totalPrice}</p>
          <button onClick={sendOrderToBackend}>Submit Order</button>
        </div>
      )}
    </div>
  );
};

export default CheckoutPage;
