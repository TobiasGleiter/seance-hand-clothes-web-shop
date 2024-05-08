Here's how you can implement the cart page with the specified functionalities:

```jsx
import { useState } from "react";

const CartPage = () => {
  const [cartItems, setCartItems] = useState([]);

  const handleRemoveFromCart = (itemId) => {
    const updatedCartItems = cartItems.filter((item) => item.id !== itemId);
    setCartItems(updatedCartItems);
  };

  return (
    <div>
      <h1>Cart</h1>
      <button onClick={() => setCartItems([])}>Clear Cart</button>
      <div>
        {cartItems.length === 0 ? (
          <p>Your cart is empty.</p>
        ) : (
          <ul>
            {cartItems.map((item) => (
              <li key={item.id}>
                <div>
                  <h3>{item.name}</h3>
                  <p>Price: {item.price}</p>
                  <button onClick={() => handleRemoveFromCart(item.id)}>
                    Remove
                  </button>
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export default CartPage;
```

This component displays the items in the cart, along with a "Remove" button for each item. Clicking the "Remove" button removes the corresponding item from the cart. Additionally, there's a "Clear Cart" button to remove all items from the cart. If the cart is empty, a message is displayed indicating that the cart is empty.
