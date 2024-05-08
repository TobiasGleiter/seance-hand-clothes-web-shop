# Occured erros in the frontend:

# 1. Axios

Axios was not installed :)

# 2. Expected 0-1 arguments, but got 2.ts(2554)

The error message "Expected 0-1 arguments, but got 2" typically occurs when there's a mismatch in the number of arguments passed to a function. In your code, the issue might be related to the `useState` hook usage inside the `useEffect` hook.

To resolve this, you can replace `useState` with `useEffect` to calculate the total price when the component mounts. Here's the corrected code:

```javascript
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
```

By using `useEffect`, the `calculateTotalPrice` function will be called only once when the component mounts, avoiding the error caused by passing extra arguments to `useState`. Let me know if this resolves the issue!
