import React from "react";
import useFetchData from "./useFetchData";

function MyComponent() {
  // Using the custom hook to fetch data
  const { data, loading } = useFetchData(
    "https://jsonplaceholder.typicode.com/users"
  );

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Data:</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default MyComponent;
