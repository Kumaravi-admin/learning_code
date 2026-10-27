import { useRef, useState } from "react";

export default function Hook() {
  const countRef = useRef(0);
  const [count, setCount] = useState(0);

  function handleClick() {
    countRef.current = countRef.current + 1;
    console.log(countRef.current);
    setCount((count) => count + 1);
  }

  return (
    <>
      <h1>useRef count : {countRef.current}</h1>
      <button onClick={handleClick}>Incriment</button>
    </>
  );
}
