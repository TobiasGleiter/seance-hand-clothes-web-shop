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