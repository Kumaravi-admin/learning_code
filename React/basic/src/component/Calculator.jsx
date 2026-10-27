import React, { useState, useMemo } from "react";

function Calculator() {
  const [count, setCount] = useState(0);
  const [otherCount, setOtherCount] = useState(0);

  // A simple expensive function
  function calculateDouble(num) {
    console.log("Calculating...");
    for (let i = 0; i < 1000000000; i++) {} // Simulate a heavy task
    return num * 2;
  }

  // useMemo to cache the result
  const doubledValue = useMemo(() => calculateDouble(count), [count]);

  return (
    <div>
      <p>Count: {count}</p>
      <p>Doubled Value: {doubledValue}</p>
      <button onClick={() => setCount(count + 1)}>Increment Count</button>
      <p>Other Count: {otherCount}</p>
      <button onClick={() => setOtherCount(otherCount + 1)}>
        Increment Other Count
      </button>
    </div>
  );
}

export default Calculator;
