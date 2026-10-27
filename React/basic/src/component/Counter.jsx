import React, { useState, useCallback } from "react";

function Counter() {
  const [count, setCount] = useState(0);
  const [text, setText] = useState("");

  const increment = useCallback(() => {
    console.log("Increment function created");
    setCount((prevCount) => prevCount + 1);
    console.log("count update");
  }, []);

  const handChange = useCallback((e) => {
    console.log("handChange function created");
    setText(e.target.value);
    console.log("input change");
  }, []);

  return (
    <div>
      <input type="text" value={text} onChange={handChange} />
      <button onClick={increment}>Increment</button>
      <p>Count: {count}</p>
    </div>
  );
}

export default Counter;
