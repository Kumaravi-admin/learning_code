import { useState } from "react";

export default function Form() {
  const [city, setCity] = useState("Delhi");

  function handleClick(event) {
    event.preventDefault();
    alert(city);
  }

  return (
    <form>
      <select value={city} onChange={(event) => setCity(event.target.value)}>
        <option value="Delhi">Delhi</option>
        <option value="Mumbai">Mumbai</option>
        <option value="Kolkata">Kolkata</option>
      </select>
      <button onClick={handleClick}>Click</button>
    </form>
  );
}
