Here's the implementation of the checkout page with the specified functionalities:

```jsx
import { useState } from "react";
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
  useState(() => {
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
```

This component calculates the total price of the articles in the cart and provides a button to submit the order. When the "Submit Order" button is clicked, it sends a POST request to the backend with the order details. Once the order is successfully submitted, a message is displayed indicating that the order has been submitted successfully.
