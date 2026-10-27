import { useState, useEffect } from "react";

// This is the custom hook
function useFetchData(url) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Fetch data when the hook is used
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      });
  }, [url]);

  // Return the fetched data and loading state
  return { data, loading };
}

export default useFetchData;
